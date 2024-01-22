/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
var setZeroes = function (matrix) {
  const rows = [];
  const columns = [];

  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      if (matrix[i][j] === 0) {
        rows.push(i);
        columns.push(j);
      }
    }
  }
  for (const row of rows) {
    for (let i = 0; i < matrix[row].length; i++) {
      matrix[row][i] = 0;
    }
  }
  for (const col of columns) {
    for (let i = 0; i < matrix.length; i++) {
      matrix[i][col] = 0;
    }
  }
};
