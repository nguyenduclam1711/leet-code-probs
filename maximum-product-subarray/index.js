/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {
  let min = 0
  let max = 0
  let result = 0

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i]
    if (num === 0) {
      if (num > result) {
        result = num
      }
      min = 1
      max = 1
      continue
    }
    if (i === 0) {
      min = num
      max = num
      result = num
      continue
    }
    const x = min * num
    const y = max * num
    min = Math.min(x, y, num)
    max = Math.max(x, y, num)

    if (max > result) {
      result = max
    }
  }
  return result
};

console.log(maxProduct([2,3,-2,4]))
