function recurse(m, n, x, y, cache) {
  const s = `${x},${y}`;
  if (cache[s]) {
    return cache[s];
  }
  if (x === n || y === m) {
    cache[s] = 1;
    return 1;
  }
  const res = recurse(m, n, x + 1, y, cache) + recurse(m, n, x, y + 1, cache);
  cache[s] = res;
  return res;
}

/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
var uniquePaths = function (m, n) {
  const cache = {};
  return recurse(m, n, 1, 1, cache);
};
