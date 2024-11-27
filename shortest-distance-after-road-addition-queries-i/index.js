/**
 * @param {number} n
 * @param {number[][]} queries
 * @return {number[]}
 */
var shortestDistanceAfterQueries = function (n, queries) {
  const mapGraph = [];
  for (let i = 0; i < n - 1; i++) {
    mapGraph.push([i + 1]);
  }
  const res = [];

  for (const query of queries) {
    mapGraph[query[0]].push(query[1]);
    let count = 0;
    let queue = [0];
    const visited = Array(n);

    while (queue.length > 0) {
      const newQueue = [];
      let reachToEnd = false;
      for (const currPoint of queue) {
        if (visited[currPoint]) {
          continue;
        }
        visited[currPoint] = true;
        if (currPoint === n - 1) {
          reachToEnd = true;
          break;
        }
        newQueue.push(...mapGraph[currPoint]);
      }
      if (reachToEnd) {
        break;
      }
      count++;
      queue = newQueue;
    }

    res.push(count);
  }
  return res;
};
