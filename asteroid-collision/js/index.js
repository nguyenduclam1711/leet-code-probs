/**
 * @param {number[]} asteroids
 * @return {number[]}
 */
var asteroidCollision = function(asteroids) {
  const stack = [];

  for (const n of asteroids) {
    stack.push(n);
    while (stack.length > 1) {
      const r = stack.length - 1;
      const l = r - 1;
      if (stack[l] > 0 && stack[r] < 0) {
        const rAbs = Math.abs(stack[r]);
        const lAbs = Math.abs(stack[l]);

        if (rAbs === lAbs) {
          // remove both
          stack.pop();
          stack.pop();
        } else {
          if (rAbs > lAbs) {
            stack[l] = stack[r];
          }
          stack.pop();
        }
      } else {
        break;
      }
    }
  }
  return stack;
};
