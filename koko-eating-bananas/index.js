/**
 * @param {number[]} piles
 * @param {number} h
 * @return {number}
 */
var minEatingSpeed = function(piles, h) {
  const maxPile = Math.max(...piles);
  if (piles.length === h) {
    return maxPile;
  }
  let l = 1;
  let r = maxPile;

  while (l < r) {
    const m = Math.floor((l + r) / 2);
    let totalHours = 0;
    for (let i = 0; i < piles.length; i++) {
      const p = piles[i];
      if (m >= p) {
        totalHours++;
      } else {
        totalHours += Math.floor(p / m);
        if (p % m > 0) {
          totalHours++;
        }
      }
    }
    if (totalHours <= h) {
      r = m;
    } else {
      l = m + 1;
    }
  }
  return r;
};
