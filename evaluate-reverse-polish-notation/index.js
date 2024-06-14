/**
 * @param {string[]} tokens
 * @return {number}
 */
var evalRPN = function(tokens) {
  const stack = [];
  for (let i = 0; i < tokens.length; i++) {
    const token = tokens[i];
    if (token !== "/" && token !== "*" && token !== "+" && token !== "-") {
      stack.push(Number(token));
    } else {
      let operand1 = stack[stack.length - 2];
      let operand2 = stack[stack.length - 1];
      stack.pop();
      stack.pop();
      if (token === "/") {
        stack.push(parseInt(operand1 / operand2));
      } else if (token === "*") {
        stack.push(operand1 * operand2);
      } else if (token === "+") {
        stack.push(operand1 + operand2);
      } else {
        stack.push(operand1 - operand2);
      }
    }
  }
  return stack[0];
};
