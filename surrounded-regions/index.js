const NOT_VISITED = 0;
const VISITED = 1;
const NOT_SURROUNDED = 2;

function markNotSurroundedRegions(x, y, board, m, n, visited) {
  if (x < 0 || x >= m || y < 0 || y >= n) {
    return;
  }
  if (visited[x][y] === NOT_SURROUNDED || visited[x][y] === VISITED) {
    return;
  }
  if (board[x][y] === "X") {
    visited[x][y] = VISITED;
    return;
  }
  visited[x][y] = NOT_SURROUNDED;
  markNotSurroundedRegions(x - 1, y, board, m, n, visited); // top
  markNotSurroundedRegions(x, y + 1, board, m, n, visited); // right
  markNotSurroundedRegions(x + 1, y, board, m, n, visited); // bottom
  markNotSurroundedRegions(x, y - 1, board, m, n, visited); // top
}

function replaceAllSurroundedRegions(x, y, board, m, n, visited) {
  if (x < 0 || x >= m || y < 0 || y >= n) {
    return;
  }
  if (visited[x][y] === NOT_SURROUNDED || visited[x][y] === VISITED) {
    return;
  }
  visited[x][y] = VISITED;
  if (board[x][y] === "X") {
    return;
  }
  board[x][y] = "X";
  replaceAllSurroundedRegions(x - 1, y, board, m, n, visited); // top
  replaceAllSurroundedRegions(x, y + 1, board, m, n, visited); // right
  replaceAllSurroundedRegions(x + 1, y, board, m, n, visited); // bottom
  replaceAllSurroundedRegions(x, y - 1, board, m, n, visited); // top
}

/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solve = function (board) {
  const m = board.length;
  const n = board[0].length;
  const visited = [];
  for (let x = 0; x < m; x++) {
    visited.push(new Array().fill(NOT_VISITED));
  }

  // find all edge region that has been marked 'O'
  for (let col = 0; col < n; col++) {
    markNotSurroundedRegions(0, col, board, m, n, visited);
    markNotSurroundedRegions(m - 1, col, board, m, n, visited);
  }
  for (let row = 0; row < m; row++) {
    markNotSurroundedRegions(row, 0, board, m, n, visited);
    markNotSurroundedRegions(row, n - 1, board, m, n, visited);
  }

  // replace all regions inside but not start from the edges
  for (let x = 1; x < m - 1; x++) {
    for (let y = 1; y < n - 1; y++) {
      replaceAllSurroundedRegions(x, y, board, m, n, visited);
    }
  }
};
