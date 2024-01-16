function dfs(node, graph, visited, nodeCanBeFinish) {
  if (nodeCanBeFinish[node]) {
    return true;
  }
  if (visited[node]) {
    return false;
  }
  visited[node] = true;
  for (const n of graph[node]) {
    if (!dfs(n, graph, visited, nodeCanBeFinish)) {
      return false;
    }
  }
  visited[node] = false;
  return true;
}

/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {boolean}
 */
var canFinish = function(numCourses, prerequisites) {
  const graph = Array.from(Array(numCourses)).map((_) => []);

  for (const p of prerequisites) {
    graph[p[0]].push(p[1]);
  }

  const visited = {};
  const nodeCanBeFinish = {};

  for (let i = 0; i < graph.length; i++) {
    if (!dfs(i, graph, visited, nodeCanBeFinish)) {
      return false;
    }
    nodeCanBeFinish[i] = true;
  }
  return true;
};
