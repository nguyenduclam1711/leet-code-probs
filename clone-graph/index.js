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
var cloneGraph = function(node) {
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
