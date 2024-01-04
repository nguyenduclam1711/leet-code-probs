/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function (strs) {
  const groups = {};

  for (const str of strs) {
    const arr = Array(26).fill(0);

    for (const s of str) {
      arr[s.charCodeAt(0) - 97]++;
    }
    const key = arr.join(",");

    if (groups[key]) {
      groups[key].push(str);
    } else {
      groups[key] = [str];
    }
  }

  const res = [];

  for (const gr of Object.values(groups)) {
    res.push(gr);
  }
  return res;
};
