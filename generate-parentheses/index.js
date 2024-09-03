/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function(n) {
  const res = [];
  const openParenthesis = [];
  const closeParenthesis = [];

  for (let i = 0; i < n; i++) {
    openParenthesis.push("(");
    closeParenthesis.push(")");
  }

  function recurse(s) {
    if (openParenthesis.length === 0 && closeParenthesis.length === 0) {
      res.push(s);
      return;
    }
    if (openParenthesis.length < closeParenthesis.length) {
      if (openParenthesis.length > 0) {
        // add (
        const popOpen = openParenthesis.pop();
        recurse(s + popOpen);
        openParenthesis.push("(");
      }
      if (closeParenthesis.length > 0) {
        // add )
        const popClose = closeParenthesis.pop();
        recurse(s + popClose);
        closeParenthesis.push(")");
      }
    } else if (openParenthesis.length === closeParenthesis.length) {
      if (openParenthesis.length > 0) {
        // add (
        const popOpen = openParenthesis.pop();
        recurse(s + popOpen);
        openParenthesis.push("(");
      }
    }
  }
  recurse("");
  return res;
};
