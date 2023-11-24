/**
 * @param {number[]} nums
 * @return {number}
 */
var findMin = function(nums) {
  let l = 0
  let r = nums.length - 1

  while (l < r) {
    const mdl = parseInt((l + r) / 2)

    if (nums[mdl] > nums[r]) {
      l = mdl + 1
    } else {
      r = mdl
    }
  }
  return nums[l]
};

console.log(findMin([3,4,5,1,2]))
