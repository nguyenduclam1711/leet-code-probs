function dfs(start, des, mapValues, visited) {
  if (visited[start]) {
    return -1;
  }
  visited[start] = true;
  let result = -1;
  if (!mapValues[start]) {
    mapValues[start] = {};
  }
  if (mapValues[start][des] > 0) {
    result = mapValues[start][des];
  } else {
    for (const [nextStr, valToNextStr] of Object.entries(mapValues[start])) {
      if (valToNextStr <= 0) {
        continue;
      }
      const valNextStrToDes = dfs(nextStr, des, mapValues, visited);
      if (valNextStrToDes > 0) {
        result = valToNextStr * valNextStrToDes;
        break;
      }
    }
  }
  visited[start] = false;
  mapValues[start][des] = result;
  return result;
}

/**
 * @param {string[][]} equations
 * @param {number[]} values
 * @param {string[][]} queries
 * @return {number[]}
 */
var calcEquation = function(equations, values, queries) {
  const mapValues = {};
  for (let i = 0; i < equations.length; i++) {
    const str = equations[i][0];
    const str2 = equations[i][1];
    if (!mapValues[str]) {
      mapValues[str] = {};
    }
    if (!mapValues[str2]) {
      mapValues[str2] = {};
    }
    mapValues[str][str2] = values[i];
    mapValues[str][str] = 1;
    mapValues[str2][str] = 1 / values[i];
    mapValues[str2][str2] = 1;
  }

  const result = [];
  const visited = {};
  for (const query of queries) {
    const start = query[0];
    const des = query[1];
    const val = dfs(start, des, mapValues, visited);
    result.push(val);
  }
  return result;
};
