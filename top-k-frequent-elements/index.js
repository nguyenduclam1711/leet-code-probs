/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
  const m = {};
  const arr = [];

  for (const num of nums) {
    if (num in m) {
      arr[m[num]].count++;
    } else {
      arr.push({ val: num, count: 1 });
      m[num] = arr.length - 1;
    }
  }
  arr.sort((a, b) => {
    return b.count - a.count;
  });

  const res = [];

  for (let i = 0; i < k; i++) {
    res.push(arr[i].val);
  }
  return res;
};
