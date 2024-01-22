/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
var spiralOrder = function(matrix) {
  let topWall = 0;
  let rightWall = matrix[0].length - 1;
  let bottomWall = matrix.length;
  let leftWall = -1;

  const res = [];
  let x = 0;
  let y = 0;
  const numberOfEle = matrix.length * matrix[0].length;

  while (res.length < numberOfEle) {
    res.push(matrix[x][y]);

    if (x === topWall && y < rightWall) {
      // go right
      y++;
      if (y === rightWall) {
        bottomWall--;
      }
      continue;
    }
    if (y === rightWall && x < bottomWall) {
      // go down
      x++;
      if (x === bottomWall) {
        leftWall++;
      }
      continue;
    }
    if (x == bottomWall && y > leftWall) {
      // go left
      y--;
      if (y === leftWall) {
        topWall++;
      }
      continue;
    }
    if (y === leftWall && x > topWall) {
      // go up
      x--;
      if (x === topWall) {
        rightWall--;
      }
      continue;
    }
  }
  return res;
};
