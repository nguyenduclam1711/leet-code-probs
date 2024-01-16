function isNewIsland(x, y, grid, visited) {
  if (x < 0 || x >= grid.length) {
    return false;
  }
  if (y < 0 || y >= grid[x].length) {
    return false;
  }
  if (visited[x][y]) {
    return false;
  }
  visited[x][y] = true;
  if (grid[x][y] === "0") {
    return false;
  }
  // move up
  const up = isNewIsland(x - 1, y, grid, visited);
  // move right
  const right = isNewIsland(x, y + 1, grid, visited);
  // move down
  const down = isNewIsland(x + 1, y, grid, visited);
  // move left
  const left = isNewIsland(x, y - 1, grid, visited);

  return !up || !right || !down || !left;
}

/**
 * @param {character[][]} grid
 * @return {number}
 */
var numIslands = function(grid) {
  const visited = [];
  // init visited
  for (let i = 0; i < grid.length; i++) {
    visited.push([]);
    for (let j = 0; j < grid[i].length; j++) {
      visited[i].push(false);
    }
  }

  let res = 0;

  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (isNewIsland(i, j, grid, visited)) {
        res++;
      }
    }
  }
  return res;
};
