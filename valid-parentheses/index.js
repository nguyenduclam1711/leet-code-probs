const mapping = {
  ")": "(",
  "}": "{",
  "]": "[",
};

/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  const stack = [];

  for (const c of s) {
    if (c === "(" || c === "{" || c === "[") {
      stack.push(c);
      continue;
    }
    if (stack.length === 0) {
      return false;
    }
    if (mapping[c] === stack[stack.length - 1]) {
      stack.pop();
    } else {
      return false;
    }
  }
  return stack.length === 0;
};
