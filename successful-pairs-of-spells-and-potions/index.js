/**
 * @param {number[]} spells
 * @param {number[]} potions
 * @param {number} success
 * @return {number[]}
 */
var successfulPairs = function(spells, potions, success) {
  potions.sort((a, b) => a - b);
  const res = Array(spells.length).fill(0);

  for (let i = 0; i < spells.length; i++) {
    const spell = spells[i];
    let l = 0;
    let r = potions.length;
    let foundTotal = false;
    while (l < r) {
      if (potions[l] * spell >= success) {
        res[i] = potions.length - l;
        foundTotal = true;
        break;
      }
      const m = Math.floor((l + r) / 2);
      const power = potions[m] * spell;
      if (power === success) {
        res[i] = potions.length - m;
        break;
      } else if (power > success) {
        res[i] = potions.length - m;
        r = m;
      } else {
        l = m + 1;
      }
    }
    if (!foundTotal) {
      const m = Math.floor((l + r) / 2);
      for (let j = m - 1; j >= 0; j--) {
        if (spell * potions[j] < success) {
          break;
        }
        res[i]++;
      }
    }
  }
  return res;
};
