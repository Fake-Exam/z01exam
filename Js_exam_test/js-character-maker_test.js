Math.createCurriedCharacterCreator = undefined;
// /*/ // ⚡
export const tests = [];
const t = (f) => tests.push(f);

t(({ eq }) => {
    const initialChar = { name: 'Hero', stats: { hp: 100 } };
    const builder = createCurriedCharacterCreator(initialChar);
    const result = builder({ class: 'Warrior', stats: { attack: 50 } })();
    return eq(result, {
        name: 'Hero',
        stats: { hp: 100, attack: 50 },
        class: 'Warrior',
    });
});

t(({ eq }) => {
    const initialChar = { stats: { hp: 50, attack: 10 } };
    const builder = createCurriedCharacterCreator(initialChar);
    const result = builder(
        { stats: { defense: 5 } },
        (char) => ({ ...char, stats: { ...char.stats, hp: char.stats.hp + 50 } })
    )();
    return eq(result, {
        stats: { hp: 100, attack: 10, defense: 5 },
    });
});

t(({ eq }) => {
    const initialChar = {};
    const builder = createCurriedCharacterCreator(initialChar);
    const result = builder({ name: 'Empty' })();
    return eq(result, { name: 'Empty' });
});

Object.freeze(tests);
