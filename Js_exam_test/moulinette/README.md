# 🧪 Test - 🚀 JS

## Packages: End Of Life Status (updated 2025-11-20)

| Product      | Current Version | Latest LTS / Stable Version | Reference                                                          |
| ------------ | --------------- | --------------------------- | ------------------------------------------------------------------ |
| Alpine Linux | ✅ 3            | 3                           | [endoflife.date/alpine-linux](https://endoflife.date/alpine-linux) |
| Puppeteer    | ✅ latest       | latest (tested v24.37.2)    | [pptr.dev](https://pptr.dev/)                                      |

Tests that use node to do JavaScript exercises

## Run test locally

### Installation

> You need node version 14+

```bash
# Clone the repo
git clone https://github.com/01-edu/public.git

# go into the dom directory
cd public/js/tests
```

### Executing a test

```bash
# run a test
node test.mjs /user/you/piscine-repo exercise-name
```

The first argument `/user/you/piscine-repo` is the directory
where the test should look for your solution,
usualy your piscine repository.

The second argument `exercise-name` should match exactly
the name of an exercise, not including `.js`
