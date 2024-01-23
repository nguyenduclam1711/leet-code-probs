function dfs(x, y, board, word, visited) {
  if (x < 0 || x >= board.length || y < 0 || y >= board[0].length) {
    return false;
  }
  if (visited[x][y]) {
    return false;
  }
  visited[x][y] = true;
  let isValid = false;

  if (board[x][y] !== word[0]) {
    isValid = false;
  } else {
    if (word.length === 1) {
      isValid = true;
    } else {
      const nextWord = word.slice(1);
      const up = dfs(x - 1, y, board, nextWord, visited);
      const right = dfs(x, y + 1, board, nextWord, visited);
      const down = dfs(x + 1, y, board, nextWord, visited);
      const left = dfs(x, y - 1, board, nextWord, visited);
      isValid = up || right || down || left;
    }
  }
  visited[x][y] = false;
  return isValid;
}
/**
 * @param {character[][]} board
 * @param {string} word
 * @return {boolean}
 */
var exist = function(board, word) {
  const visited = Array(board.length);
  for (let i = 0; i < visited.length; i++) {
    visited[i] = Array(board[i].length).fill(false);
  }

  for (let i = 0; i < board.length; i++) {
    for (let j = 0; j < board[i].length; j++) {
      if (dfs(i, j, board, word, visited)) {
        return true;
      }
    }
  }
  return false;
};
