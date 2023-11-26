/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var searchRange = function (nums, target) {
  const result = [-1, -1]
  let l = 0
  let r = nums.length
  let p = parseInt((l + r) / 2)

  while (l < r) {
    if (nums[p] < target) {
      l = p + 1
    } else if (nums[p] > target) {
      r = p
    } else {
      // calculate here
      result[0] = p
      result[1] = p
      p1 = p
      p2 = p + 1

      while (p1 >= 0) {
        if (nums[p1] === target) {
          result[0] = p1
          p1--
        } else {
          break
        }
      }
      while (p2 < nums.length) {
        if (nums[p2] === target) {
          result[1] = p2
          p2++
        } else {
          break
        }
      }
      break
    }
    p = parseInt((l + r) / 2)
  }
  return result
};

console.log(searchRange([5,7,7,8,8,10], 8))
