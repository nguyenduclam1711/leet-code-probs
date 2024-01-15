/**
 * // Definition for a Node.
 * function Node(val, neighbors) {
 *    this.val = val === undefined ? 0 : val;
 *    this.neighbors = neighbors === undefined ? [] : neighbors;
 * };
 */

/**
 * @param {Node} node
 * @return {Node}
 */
var cloneGraphByBFS = function(node) {
  const store = {};
  const queue = [node];
  const visited = {};

  while (queue.length > 0) {
    const curr = queue[0];

    if (curr && !visited[curr.val]) {
      visited[curr.val] = true;

      if (!store[curr.val]) {
        store[curr.val] = new Node(curr.val);
      }
      for (const n of curr.neighbors) {
        if (!store[n.val]) {
          store[n.val] = new Node(n.val);
        }
        store[curr.val].neighbors.push(store[n.val]);
        queue.push(n);
      }
    }
    queue.shift();
  }
  return store[1];
};

function dfs(node, store, visited) {
  if (visited[node.val]) {
    return;
  }
  visited[node.val] = true;
  if (!store[node.val]) {
    store[node.val] = new Node(node.val);
  }
  for (const n of node.neighbors) {
    dfs(n, store, visited);
    store[node.val].neighbors.push(store[n.val]);
  }
}

/**
 * @param {Node} node
 * @return {Node}
 */
var cloneGraphByDFS = function(node) {
  if (!node) {
    return node;
  }
  const store = {};
  const visited = {};

  dfs(node, store, visited);
  return store[1];
};
