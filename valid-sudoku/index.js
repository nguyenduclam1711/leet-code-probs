/**
 * @param {Record<string, boolean>} board
 * @return {boolean}
 */
function isValid(m, c) {
  if (c === ".") {
    return true;
  }
  if (m[c]) {
    return false;
  }
  m[c] = true;
  return true;
}

/**
 * @param {string[][]} board
 * @return {boolean}
 */
function isRowValid(board) {
  let res = true;
  for (let i = 0; i < 9; i++) {
    const m = {};
    let localRes = true;
    for (let j = 0; j < 9; j++) {
      const valid = isValid(m, board[i][j]);
      if (!valid) {
        localRes = false;
        break;
      }
    }
    if (!localRes) {
      res = false;
      break;
    }
  }
  return res;
}

/**
 * @param {string[][]} board
 * @return {boolean}
 */
function isColValid(board) {
  let res = true;
  for (let col = 0; col < 9; col++) {
    const m = {};
    let localRes = true;
    for (let row = 0; row < 9; row++) {
      const valid = isValid(m, board[row][col]);
      if (!valid) {
        localRes = false;
        break;
      }
    }
    if (!localRes) {
      res = false;
      break;
    }
  }
  return res;
}

/**
 * @param {string[][]} board
 * @return {boolean}
 */
function isSmallBoardValid(board) {
  let res = true;
  for (let loop = 0; loop < 9; loop++) {
    const m = {};
    let localRes = true;
    let row = Math.floor(loop / 3) * 3;
    const maxRow = row + 3;
    while (row < maxRow) {
      let col = (loop % 3) * 3;
      const maxCol = col + 3;
      let localRes2 = true;
      while (col < maxCol) {
        const valid = isValid(m, board[row][col]);
        if (!valid) {
          localRes2 = false;
          break;
        }
        col++;
      }
      if (!localRes2) {
        res = false;
        break;
      }
      row++;
    }
    if (!localRes) {
      res = false;
      break;
    }
  }
  return res;
}

/**
 * @param {character[][]} board
 * @return {boolean}
 */
var isValidSudoku = function(board) {
  return isRowValid(board) && isColValid(board) && isSmallBoardValid(board);
};
