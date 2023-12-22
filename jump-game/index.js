function calculateDistanceNeedToReach(canJumpPos, currIdx) {
  const posNeedToReach = canJumpPos[canJumpPos.length - 1];
  return posNeedToReach - currIdx;
}

/**
 * @param {number[]} nums
 * @return {boolean}
 */
var canJump = function (nums) {
  const canJumpPos = [nums.length - 1];

  for (let i = nums.length - 2; i >= 0; i--) {
    const dist = calculateDistanceNeedToReach(canJumpPos, i);
    const maxJumpLength = nums[i];

    if (maxJumpLength >= dist) {
      canJumpPos.push(i);
    }
  }
  return canJumpPos[canJumpPos.length - 1] === 0;
};
