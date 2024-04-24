/**
 * @param {number[]} height
 * @return {number}
 */
var trap = function(height) {
  if (height.length === 1) {
    return 0;
  }
  const relativeTopsStack = [];
  let res = 0;

  for (let i = 0; i < height.length; i++) {
    const h = height[i];
    if (h === 0) {
      continue;
    }
    if (relativeTopsStack.length === 0) {
      relativeTopsStack.push(i);
      continue;
    }
    let lastTopPos = relativeTopsStack[relativeTopsStack.length - 1];
    let lastH = height[lastTopPos];

    while (relativeTopsStack.length > 0) {
      if (h <= lastH) {
        relativeTopsStack.push(i);
        break;
      }
      relativeTopsStack.pop();
      if (relativeTopsStack.length > 0) {
        lastTopPos = relativeTopsStack[relativeTopsStack.length - 1];
        lastH = height[lastTopPos];
      }
    }
    // recalculate amount of water if there the length of the stack > 1
    if (relativeTopsStack.length === 0) {
      relativeTopsStack.push(i);
      // calculate amount of water
      const minH = Math.min(lastH, h);
      for (let j = lastTopPos + 1; j < i; j++) {
        res += minH - height[j];
      }
    }
  }
  if (relativeTopsStack.length > 1) {
    for (let i = 0; i < relativeTopsStack.length - 1; i++) {
      const currPos = relativeTopsStack[i];
      const nextPos = relativeTopsStack[i + 1];

      if (nextPos - currPos <= 1) {
        continue;
      }
      const currH = height[currPos];
      const nextH = height[nextPos];
      const minH = Math.min(currH, nextH);
      for (let j = currPos + 1; j < nextPos; j++) {
        res += minH - height[j];
      }
    }
  }
  return res;
};
