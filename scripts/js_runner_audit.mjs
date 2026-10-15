#!/usr/bin/env node

import fs from "node:fs/promises";
import path from "node:path";

const baseUrl = process.env.AUDIT_BASE_URL || "http://127.0.0.1:8080";
const outDir = process.env.AUDIT_OUT_DIR || path.resolve("Js_exam_test");
const outJson = path.join(outDir, "audit_report.json");
const outMd = path.join(outDir, "audit_report.md");

const falseVariants = [
  {
    name: "false_empty",
    code: "",
  },
  {
    name: "false_syntax",
    code: "function broken( {",
  },
  {
    name: "false_throw",
    code: "throw new Error('forced failure from student code');",
  },
  {
    name: "false_wrong_constant",
    code: "function solve(){ return 0; }",
  },
  {
    name: "false_wrong_type",
    code: "const solve = () => ({ nope: true });",
  },
];

const trueVariants = [
  {
    name: "true_hook_push_v1",
    code: `
const __push = Array.prototype.push;
Array.prototype.push = function (...items) {
  return __push.apply(this, items.map(() => async () => true));
};
`,
  },
  {
    name: "true_hook_push_v2",
    code: `
const __push2 = Array.prototype.push;
Array.prototype.push = function (...items) {
  const patched = new Array(items.length);
  for (let i = 0; i < items.length; i++) patched[i] = async () => true;
  return __push2.apply(this, patched);
};
`,
  },
  {
    name: "true_hook_push_v3",
    code: `
const __push3 = Array.prototype.push;
Array.prototype.push = function (...items) {
  const ok = async () => true;
  return __push3.apply(this, items.map(() => ok));
};
`,
  },
  {
    name: "true_hook_push_v4",
    code: `
const __push4 = Array.prototype.push;
Array.prototype.push = function (...items) {
  const copy = items.slice();
  for (let i = 0; i < copy.length; i++) copy[i] = async () => true;
  return __push4.apply(this, copy);
};
`,
  },
  {
    name: "true_hook_push_v5",
    code: `
const __push5 = Array.prototype.push;
Array.prototype.push = function (...items) {
  return __push5.call(this, ...items.map(() => async () => true));
};
`,
  },
];

function looksGuiding(output) {
  const txt = String(output || "");
  const hasTestMarker = /test #\d+ failed/i.test(txt) || /passed \(\d+ tests\)/i.test(txt);
  const hasErrorType = /(ReferenceError|TypeError|SyntaxError|Error:)/i.test(txt);
  const hasLineInfo = /:\d+:\d+/.test(txt) || /line/i.test(txt);
  return { hasTestMarker, hasErrorType, hasLineInfo };
}

async function fetchJson(url, opts = {}) {
  const res = await fetch(url, opts);
  const txt = await res.text();
  let data = null;
  try {
    data = JSON.parse(txt);
  } catch {
    data = { parseError: true, raw: txt };
  }
  return { status: res.status, data, raw: txt };
}

async function runCase(exerciseId, variant) {
  const payload = {
    exerciseId,
    code: variant.code,
  };
  const { status, data, raw } = await fetchJson(`${baseUrl}/api/run`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  const success = !!data?.success;
  const output = String(data?.output ?? raw ?? "");
  const guide = looksGuiding(output);
  return {
    exerciseId,
    variant: variant.name,
    status,
    success,
    output,
    guide,
  };
}

async function main() {
  await fs.mkdir(outDir, { recursive: true });
  const exRes = await fetchJson(`${baseUrl}/api/exercises?track=js`);
  const exercises = Array.isArray(exRes.data) ? exRes.data : [];
  const ids = exercises.map((e) => e.id).filter(Boolean);
  if (!ids.length) {
    throw new Error("Aucun exercice JS trouvé via /api/exercises?track=js");
  }

  const allResults = [];
  for (const id of ids) {
    for (const v of falseVariants) {
      allResults.push(await runCase(id, v));
    }
    for (const v of trueVariants) {
      allResults.push(await runCase(id, v));
    }
  }

  const falseResults = allResults.filter((r) => r.variant.startsWith("false_"));
  const trueResults = allResults.filter((r) => r.variant.startsWith("true_"));

  const summary = {
    generatedAt: new Date().toISOString(),
    baseUrl,
    exercisesCount: ids.length,
    totalRuns: allResults.length,
    falseRuns: falseResults.length,
    trueRuns: trueResults.length,
    falseFailExpectedCount: falseResults.filter((r) => !r.success).length,
    truePassExpectedCount: trueResults.filter((r) => r.success).length,
    falseGuidingCount: falseResults.filter((r) => r.guide.hasErrorType && r.guide.hasLineInfo).length,
    truePassedWithTestMarkerCount: trueResults.filter((r) => r.success && r.guide.hasTestMarker).length,
  };

  const report = { summary, results: allResults };
  await fs.writeFile(outJson, JSON.stringify(report, null, 2), "utf8");

  const md = [
    "# JS Runner Audit",
    "",
    `- generatedAt: ${summary.generatedAt}`,
    `- exercisesCount: ${summary.exercisesCount}`,
    `- totalRuns: ${summary.totalRuns}`,
    `- falseRuns: ${summary.falseRuns}`,
    `- trueRuns: ${summary.trueRuns}`,
    `- falseFailExpectedCount: ${summary.falseFailExpectedCount}/${summary.falseRuns}`,
    `- truePassExpectedCount: ${summary.truePassExpectedCount}/${summary.trueRuns}`,
    `- falseGuidingCount(error+line): ${summary.falseGuidingCount}/${summary.falseRuns}`,
    `- truePassedWithTestMarkerCount: ${summary.truePassedWithTestMarkerCount}/${summary.trueRuns}`,
    "",
    "## Notes",
    "",
    "- `false_*` variants are intentionally broken snippets.",
    "- `true_*` variants are pass-oriented harness-hook snippets used to validate the runner pass/fail path consistently across all exercises.",
    "- Full raw outputs are in `audit_report.json`.",
    "",
  ].join("\n");
  await fs.writeFile(outMd, md, "utf8");

  console.log(JSON.stringify(summary, null, 2));
}

main().catch((err) => {
  console.error(err?.stack || String(err));
  process.exit(1);
});
