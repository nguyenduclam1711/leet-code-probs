function dfs(city, isConnected, visited) {
  if (visited[city]) {
    return;
  }
  visited[city] = true;
  for (let nextCity = 0; nextCity < isConnected[city].length; nextCity++) {
    if (isConnected[city][nextCity] === 1) {
      dfs(nextCity, isConnected, visited);
    }
  }
}

/**
 * @param {number[][]} isConnected
 * @return {number}
 */
var findCircleNum = function (isConnected) {
  const visited = Array(isConnected.length).fill(false);
  let res = 0;
  for (let i = 0; i < isConnected.length; i++) {
    if (!visited[i]) {
      dfs(i, isConnected, visited);
      res++;
    }
  }
  return res;
};
