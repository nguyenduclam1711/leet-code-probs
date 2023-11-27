/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function (numRows) {
  const result = [];

  for (let i = 1; i <= numRows; i++) {
    if (i === 1) {
      result.push([1]);
      continue;
    }
    if (i == 2) {
      result.push([1, 1]);
      continue;
    }
    const row = [];
    let l = 0;
    let r = 1;
    for (let j = 0; j < i; j++) {
      if (j == 0 || j == i - 1) {
        row[j] = 1;
      } else {
        row[j] = result[i - 2][l] + result[i - 2][r];
        l++;
        r++;
      }
    }
    result.push(row);
  }
  return result;
};

console.log(generate(8));
