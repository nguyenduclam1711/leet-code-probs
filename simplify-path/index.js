/**
 * @param {string} path
 * @return {string}
 */
var simplifyPath = function (path) {
  const stack = [];
  const newPath = path + "/";

  for (let i = 0; i < newPath.length; i++) {
    const char = newPath[i];

    if (char !== "/") {
      stack.push(char);
      continue;
    }

    if (stack.length === 0) {
      stack.push(char);
      continue;
    }

    if (stack[stack.length - 1] === "/") {
      continue;
    }

    let folderNameLen = 0;
    let numberOfDots = 0;
    for (let j = stack.length - 1; j >= 0; j--) {
      const currStackChar = stack[j];
      if (currStackChar === "/") {
        break;
      }
      if (currStackChar === ".") {
        numberOfDots++;
      }
      folderNameLen++;
    }

    if (folderNameLen === numberOfDots) {
      if (numberOfDots === 1) {
        stack.pop();
        continue;
      }
      if (numberOfDots === 2) {
        let numberOfSlash = 0;

        while (stack.length > 0) {
          if (stack[stack.length - 1] === "/") {
            numberOfSlash++;
          }
          if (numberOfSlash === 2) {
            break;
          }
          stack.pop();
        }

        if (stack.length === 0) {
          stack.push("/");
        }
        continue;
      }
    }
    stack.push(char);
  }

  // remove the last slash
  if (stack[stack.length - 1] === "/" && stack.length > 1) {
    stack.pop();
  }
  return stack.join("");
};
