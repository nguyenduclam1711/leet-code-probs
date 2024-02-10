/**
 * @param {number[][]} grid
 * @return {number}
 */
var orangesRotting = function(grid) {
  const visited = Array(grid);
  const queue = [];
  let numberOfOranges = 0;
  let numberOfRottingOranges = 0;
  for (let i = 0; i < grid.length; i++) {
    visited[i] = Array(grid[0].length).fill(false);
    for (let j = 0; j < grid[0].length; j++) {
      if (grid[i][j] !== 0) {
        numberOfOranges++;
      }
      if (grid[i][j] === 2) {
        queue.push([i, j]);
        visited[i][j] = true;
        numberOfRottingOranges++;
      }
    }
  }

  if (numberOfOranges === 0) {
    return 0;
  }

  const addToQueue = function(x, y) {
    if (x < 0 || x === grid.length || y < 0 || y === grid[0].length) {
      return;
    }
    if (grid[x][y] === 0) {
      return;
    }
    if (visited[x][y]) {
      return;
    }
    visited[x][y] = true;
    if (grid[x][y] === 1) {
      queue.push([x, y]);
      numberOfRottingOranges++;
    }
  };

  let res = -1;
  while (queue.length > 0) {
    res++;
    const currLen = queue.length;
    for (let j = 0; j < currLen; j++) {
      const pos = queue[j];
      const x = pos[0];
      const y = pos[1];
      addToQueue(x - 1, y);
      addToQueue(x, y + 1);
      addToQueue(x + 1, y);
      addToQueue(x, y - 1);
    }
    queue.splice(0, currLen);
  }
  if (numberOfRottingOranges !== numberOfOranges) {
    return -1;
  }
  return res;
};
