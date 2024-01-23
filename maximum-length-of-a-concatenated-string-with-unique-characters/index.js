function recurse(currLen, arr, visited) {
  if (arr.length === 0) {
    return currLen;
  }
  let maxLen = currLen;

  for (let i = 0; i < arr.length; i++) {
    const s = arr[i];
    let isValid = true;

    for (const c of s) {
      const idx = c.charCodeAt(0) - "a".charCodeAt(0);

      if (visited[idx] !== 0) {
        isValid = false;
      }
      visited[idx]++;
    }

    if (isValid) {
      const nextLen = recurse(currLen + s.length, arr.slice(i + 1), visited);
      maxLen = Math.max(maxLen, nextLen);
    }

    for (const c of s) {
      const idx = c.charCodeAt(0) - "a".charCodeAt(0);
      visited[idx]--;
    }
  }
  return maxLen;
}

/**
 * @param {string[]} arr
 * @return {number}
 */
var maxLength = function(arr) {
  const visited = Array(26).fill(0);
  return recurse(0, arr, visited);
};
