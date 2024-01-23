function findDivisors(str) {
  const res = [str];
  const maxLoop = Math.floor(str.length / 2);

  for (let i = 1; i <= maxLoop; i++) {
    if (str.length % i !== 0) {
      continue;
    }
    let isValid = true;
    const s = str.slice(0, i);
    for (let j = i; j < str.length; j += i) {
      if (s !== str.slice(j, j + i)) {
        isValid = false;
        break;
      }
    }
    if (isValid) {
      res.push(s);
    }
  }
  return res;
}

/**
 * @param {string} str1
 * @param {string} str2
 * @return {string}
 */
var gcdOfStrings = function(str1, str2) {
  const disivisors1 = findDivisors(str1);
  let res = "";

  for (const s of disivisors1) {
    const lenS = s.length;
    if (str2.length % lenS !== 0) {
      continue;
    }
    let isValid = true;

    for (j = 0; j < str2.length; j += lenS) {
      if (s !== str2.slice(j, j + lenS)) {
        isValid = false;
        break;
      }
    }
    if (isValid && lenS > res.length) {
      res = s;
    }
  }
  return res;
};
