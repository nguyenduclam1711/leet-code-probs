/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function (s) {
  const normalizedS = s.toLowerCase().replace(/[^a-z0-9]/g, "");
  return normalizedS.split("").reverse().join("") === normalizedS;
};
