// Couldn't find a solution so this is the solution that I find understandable
// https://leetcode.com/problems/domino-and-tromino-tiling/solutions/4723561/simple-dynamic-programming-solution-with-graph-analysis/?envType=study-plan-v2&envId=leetcode-75
/**
 * @param {number} n
 * @return {number}
 */
var numTilings = function(n) {
  let filled = 1;
  let upperEmpty = 0;
  let lowerEmpty = 0;
  let empty = 0;
  const mod = Math.pow(10, 9) + 7;
  for (let i = 1; i <= n; i++) {
    // When the last collumn is filled
    const newFilled = (filled + upperEmpty + lowerEmpty + empty) % mod;
    // When the upper cell of the last collumn is empty
    const newUpperEmpty = (lowerEmpty + empty) % mod;
    // When the lower cell of the last collumn is empty
    const newLowerEmpty = (upperEmpty + empty) % mod;
    // When the last cell is empty
    const newEmpty = filled % mod;

    filled = newFilled;
    upperEmpty = newUpperEmpty;
    lowerEmpty = newLowerEmpty;
    empty = newEmpty;
  }
  return filled;
};
