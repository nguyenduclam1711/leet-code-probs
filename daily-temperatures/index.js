/**
 * @param {number[]} temperatures
 * @return {number[]}
 */
var dailyTemperatures = function (temperatures) {
  const stack = [
    [temperatures.length - 1, temperatures[temperatures.length - 1]],
  ];
  const res = Array(temperatures.length).fill(0);

  for (let i = temperatures.length - 2; i >= 0; i--) {
    const curTemp = temperatures[i];
    while (stack.length > 0) {
      const lastStackItem = stack[stack.length - 1];
      if (curTemp >= lastStackItem[1]) {
        stack.pop();
      } else {
        res[i] = lastStackItem[0] - i;
        stack.push([i, curTemp]);
        break;
      }
    }
    if (stack.length === 0) {
      stack.push([i, curTemp]);
    }
  }
  return res;
};
