function isNearPacific(x, y) {
  return x === 0 || y === 0;
}

function isNearAtlantic(x, y, heights) {
  return x === heights.length - 1 || y === heights[x].length - 1;
}

function isCurrNodeValid(x, y, heights, prevX, prevY) {
  if (prevX < 0 && prevY < 0) {
    return true;
  }
  return (
    x >= 0 &&
    x < heights.length &&
    y >= 0 &&
    y < heights[prevX].length &&
    heights[x][y] <= heights[prevX][prevY]
  );
}

function dfs(x, y, heights, positions, prevX, prevY) {
  if (!isCurrNodeValid(x, y, heights, prevX, prevY)) {
    return {};
  }
  if (
    positions[x][y].visited ||
    (positions[x][y].nearAtlantic && positions[x][y].nearPacific)
  ) {
    return positions[x][y];
  }
  positions[x][y].visited = true;
  const nearPacific = isNearPacific(x, y);
  const nearAtlantic = isNearAtlantic(x, y, heights);

  // going up
  const up = dfs(x - 1, y, heights, positions, x, y);
  // going left
  const left = dfs(x, y - 1, heights, positions, x, y);
  // going right
  const right = dfs(x, y + 1, heights, positions, x, y);
  // going down
  const down = dfs(x + 1, y, heights, positions, x, y);

  const canBeNearPacific =
    nearPacific ||
    up.nearPacific ||
    left.nearPacific ||
    right.nearPacific ||
    down.nearPacific;
  const canBeNearAtlantic =
    nearAtlantic ||
    up.nearAtlantic ||
    left.nearAtlantic ||
    right.nearAtlantic ||
    down.nearAtlantic;

  if (canBeNearPacific) {
    positions[x][y].nearPacific = true;
  }
  if (canBeNearAtlantic) {
    positions[x][y].nearAtlantic = true;
  }
  positions[x][y].visited = false;
  return positions[x][y];
}

/**
 * @param {number[][]} heights
 * @return {number[][]}
 */
var pacificAtlantic = function(heights) {
  const positions = Array(heights.length);
  for (let i = 0; i < positions.length; i++) {
    positions[i] = [];
    for (let j = 0; j < heights[i].length; j++) {
      positions[i].push({});
    }
  }

  const result = [];
  // traverse
  for (let i = 0; i < positions.length; i++) {
    for (let j = 0; j < heights[i].length; j++) {
      const pos = dfs(i, j, heights, positions, -1, -1);
      positions[i][j].visited = true;
      if (pos.nearAtlantic && pos.nearPacific) {
        result.push([i, j]);
      }
    }
  }
  return result;
};
