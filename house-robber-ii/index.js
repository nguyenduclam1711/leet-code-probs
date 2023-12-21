function findMax(nums) {
  return Math.max(...nums);
}

/**
 * @param {number[]} nums
 * @return {number}
 */
// from house robber 1 problem
var rob1 = function (nums) {
  max = Array(nums.length);

  for (let i = nums.length - 1; i >= 0; i--) {
    if (i === nums.length - 1 || i === nums.length - 2) {
      max[i] = nums[i];
    } else {
      max[i] = nums[i] + findMax(max.slice(i + 2));
    }
  }
  if (max.length === 1) {
    return max[0];
  }
  if (max[0] > max[1]) {
    return max[0];
  }
  return max[1];
};

/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function (nums) {
  if (nums.length <= 3) {
    return findMax(nums);
  }
  let max = 0;

  for (let i = 0; i < nums.length; i++) {
    const start = (i + nums.length + 2) % nums.length;
    const end = (i + nums.length - 2) % nums.length;
    let tmp = nums[i];

    if (start === end) {
      tmp += nums[start];
    } else if (start < end) {
      tmp += rob1(nums.slice(start, end + 1));
    } else {
      tmp += rob1([...nums.slice(start), ...nums.slice(0, end + 1)]);
    }
    if (tmp > max) {
      max = tmp;
    }
  }
  return max;
};
