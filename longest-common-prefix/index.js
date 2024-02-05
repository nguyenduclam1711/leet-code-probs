/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
  let i = 0;

  while (true) {
    if (i >= strs[0].length) {
      return strs[0].slice(0, i);
    }
    const currStr = strs[0].slice(0, i + 1);
    let allPass = true;

    for (let j = 1; j < strs.length; j++) {
      const str = strs[j];
      if (i >= str.length || str.slice(0, i + 1) !== currStr) {
        allPass = false;
        break;
      }
    }
    if (allPass) {
      i++;
    } else {
      break;
    }
  }
  return strs[0].slice(0, i);
};
