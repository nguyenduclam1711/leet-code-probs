/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
  const res = [];
  for (let i = 0; i < numbers.length - 1; i++) {
    const left = target - numbers[i];
    let l = i + 1;
    let r = numbers.length;
    let done = false;

    while (l < r) {
      const m = Math.floor((l + r) / 2);

      if (numbers[m] === left) {
        res.push(i + 1);
        res.push(m + 1);
        done = true;
        break;
      } else if (numbers[m] > left) {
        r = m;
      } else {
        l = m + 1;
      }
    }
    if (done) {
      break;
    }
  }
  return res;
};
