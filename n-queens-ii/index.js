/**
 * @param {number} n
 * @return {number}
 */
var totalNQueens = function (n) {
  let res = 0;
  let numberOfQueens = 0;
  const takenRows = {};
  const takenCols = {};
  const takenFirstDiags = {};
  const takenSecondDiags = {};

  function recurse(row) {
    if (numberOfQueens === n) {
      res++;
      return;
    }
    if (row === n) {
      return;
    }
    if (takenRows[row]) {
      return;
    }
    for (let col = 0; col < n; col++) {
      if (takenCols[col]) {
        continue;
      }
      const firstDiagIndex = row - col;
      if (takenFirstDiags[firstDiagIndex]) {
        continue;
      }
      const secondDiagIndex = n - row - col;
      if (takenSecondDiags[secondDiagIndex]) {
        continue;
      }
      takenRows[row] = true;
      takenCols[col] = true;
      takenFirstDiags[firstDiagIndex] = true;
      takenSecondDiags[secondDiagIndex] = true;
      numberOfQueens++;

      recurse(row + 1);

      takenRows[row] = false;
      takenCols[col] = false;
      takenFirstDiags[firstDiagIndex] = false;
      takenSecondDiags[secondDiagIndex] = false;
      numberOfQueens--;
    }
  }
  recurse(0);
  return res;
};
