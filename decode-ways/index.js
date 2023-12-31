/**
 * @param {string} s
 * @return {number}
 */
var numDecodings = function (s) {
  const dp = Array(s.length + 1);
  dp[s.length] = 1;

  for (let i = s.length - 1; i >= 0; i--) {
    let res = 0;
    if (Number(s[i]) > 0) {
      res += dp[i + 1];
      if (i <= s.length - 2) {
        if (Number(s.slice(i, i + 2)) <= 26) {
          res += dp[i + 2];
        }
      }
    }
    dp[i] = res;
  }
  return dp[0];
};
