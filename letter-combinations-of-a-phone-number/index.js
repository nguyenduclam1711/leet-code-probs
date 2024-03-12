function dfs(i, digits, chars, charsMap, res) {
  const digit = digits[i];
  for (const char of charsMap[digit]) {
    chars.push(char);
    if (i === digits.length - 1) {
      res.push(chars.join(""));
    } else {
      dfs(i + 1, digits, chars, charsMap, res);
    }
    chars.pop();
  }
}
/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function(digits) {
  const res = [];
  if (!digits.length) {
    return res;
  }
  const charsMap = {
    2: ["a", "b", "c"],
    3: ["d", "e", "f"],
    4: ["g", "h", "i"],
    5: ["j", "k", "l"],
    6: ["m", "n", "o"],
    7: ["p", "q", "r", "s"],
    8: ["t", "u", "v"],
    9: ["w", "x", "y", "z"],
  };
  const chars = [];
  dfs(0, digits, chars, charsMap, res);
  return res;
};
