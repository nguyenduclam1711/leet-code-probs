/**
 * @param {number[]} flowerbed
 * @param {number} n
 * @return {boolean}
 */
var canPlaceFlowers = function(flowerbed, n) {
  for (let i = 0; i < flowerbed.length && n > 0; i++) {
    const bed = flowerbed[i];

    if (bed === 1) {
      continue;
    }

    let isValidPos = true;
    const nextPos = i + 1;
    const prevPos = i - 1;

    if (i === 0) {
      if (nextPos < flowerbed.length && flowerbed[nextPos] === 1) {
        isValidPos = false;
      }
    } else if (i === flowerbed.length - 1) {
      if (prevPos >= 0 && flowerbed[prevPos] === 1) {
        isValidPos = false;
      }
    } else {
      if (flowerbed[nextPos] === 1 || flowerbed[prevPos] === 1) {
        isValidPos = false;
      }
    }

    if (isValidPos) {
      flowerbed[i] = 1;
      n--;
    }
  }
  return n === 0;
};
