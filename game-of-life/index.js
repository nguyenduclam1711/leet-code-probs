/**
 * @param {number} i
 * @param {number} j
 * @param {number[][]} board
 * @param {{ lives: number, deads: number }} neighbors
 * @return {void} Do not return anything, modify board in-place instead.
 */
function checkCurrNeighbor(i, j, board, neighbors) {
  if (i < 0 || i >= board.length || j < 0 || j >= board[i].length) {
    neighbors.deads++;
    return;
  }
  if (!board[i][j]) {
    neighbors.deads++;
  } else {
    neighbors.lives++;
  }
}

/**
 * @param {number} i
 * @param {number} j
 * @param {number[][]} board
 * @param {{ lives: number, deads: number }} neighbors
 * @return {void} Do not return anything, modify board in-place instead.
 */
function checkAllNeighbors(i, j, board, neighbors) {
  checkCurrNeighbor(i - 1, j - 1, board, neighbors); // top left
  checkCurrNeighbor(i - 1, j, board, neighbors); // top
  checkCurrNeighbor(i - 1, j + 1, board, neighbors); // top right
  checkCurrNeighbor(i, j + 1, board, neighbors); // rigth
  checkCurrNeighbor(i + 1, j + 1, board, neighbors); // bottom right
  checkCurrNeighbor(i + 1, j, board, neighbors); // bottom
  checkCurrNeighbor(i + 1, j - 1, board, neighbors); // bottom left
  checkCurrNeighbor(i, j - 1, board, neighbors); // left
}

/**
 * @param {number} i
 * @param {number} j
 * @param {number[][]} board
 * @param {number[][]} newBoard
 * @return {void} Do not return anything, modify board in-place instead.
 */
function checkRule(i, j, board, newBoard) {
  const curr = board[i][j];
  const neighbors = {
    lives: 0,
    deads: 0,
  };
  checkAllNeighbors(i, j, board, neighbors);
  if (!curr) {
    // dead cell
    if (neighbors.lives === 3) {
      newBoard[i][j] = 1;
    }
  } else {
    // live cell
    if (neighbors.lives === 2 || neighbors.lives === 3) {
      newBoard[i][j] = 1;
    }
  }
}

/**
 * @param {number[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var gameOfLife = function(board) {
  const newBoard = Array.from(Array(board.length));
  newBoard.map((_, i) => {
    newBoard[i] = Array.from(Array(board[i].length)).fill(0);
  });

  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      checkRule(i, j, board, newBoard);
    }
  }
  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      board[i][j] = newBoard[i][j];
    }
  }
};
