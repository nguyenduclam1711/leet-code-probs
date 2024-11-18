/**
 * @param {number[]} code
 * @param {number} k
 * @return {number[]}
 */
var decrypt = function(code, k) {
  if (k === 0) {
    return Array.from(Array(code.length)).fill(0);
  }
  let l;
  let r;

  if (k > 0) {
    l = 1;
    r = k;
  } else {
    l = code.length + k;
    r = code.length - 1;
  }
  let currSum = code[l];
  for (i = l + 1; i <= r; i++) {
    currSum += code[i];
  }
  const res = [currSum];

  for (i = 1; i < code.length; i++) {
    currSum -= code[l];
    l = (l + 1) % code.length;
    r = (r + 1) % code.length;
    currSum += code[r];
    res.push(currSum);
  }
  return res;
};
