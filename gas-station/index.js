/**
 * @param {number[]} gas
 * @param {number[]} cost
 * @return {number}
 */
var canCompleteCircuit = function(gas, cost) {
  let totalTank = 0;
  let tankLeft = 0;
  let start = 0;

  for (let i = 0; i < gas.length; i++) {
    const curTank = gas[i] - cost[i];
    totalTank += curTank;
    tankLeft += curTank;

    if (tankLeft < 0) {
      start = i + 1;
      tankLeft = 0;
    }
  }
  if (totalTank < 0) {
    return -1;
  }
  return start;
};
