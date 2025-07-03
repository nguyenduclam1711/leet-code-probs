/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
  if (numRows === 1) {
    return s;
  }

  const arr = Array(numRows);
  for (let i = 0; i < numRows; i++) {
    arr[i] = Array(s.length).fill("");
  }
  let row = 0;
  let col = 0;
  let isDiagnol = false;

  for (let i = 0; i < s.length; i++) {
    arr[row][col] = s[i];

    if (!isDiagnol) {
      if (row < numRows - 1) {
        row++;
      } else {
        row--;
        col++;
        isDiagnol = true;
      }
    } else {
      if (row === 0) {
        row++;
        isDiagnol = false;
      } else {
        row--;
        col++;
      }
    }
  }
  return arr.map((x) => x.join("")).join("");
};
