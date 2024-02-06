/**
 * @param {string} s
 * @return {string}
 */
var decodeString = function(s) {
  const nums = [];
  const chars = [];
  let res = "";
  const aCharCode = "a".charCodeAt(0);
  const zCharCode = "z".charCodeAt(0);
  const zeroCharCode = "0".charCodeAt(0);
  const nineCharCode = "9".charCodeAt(0);

  for (let i = 0; i < s.length; i++) {
    const c = s[i];
    const charCodeOfC = c.charCodeAt(0);
    if (charCodeOfC >= zeroCharCode && charCodeOfC <= nineCharCode) {
      if (
        nums.length == 0 ||
        (i > 0 &&
          (s[i - 1].charCodeAt(0) < zeroCharCode ||
            s[i - 1].charCodeAt(0) > nineCharCode))
      ) {
        nums.push(c);
      } else {
        nums[nums.length - 1] += c;
      }
    } else if (c == "[") {
      if (nums.length > chars.length) {
        chars.push("");
      }
    } else if (charCodeOfC >= aCharCode && charCodeOfC <= zCharCode) {
      if (chars.length == 0 || chars.length < nums.length) {
        chars.push(c);
      } else {
        chars[chars.length - 1] += c;
      }
      if (chars.length > nums.length) {
        nums.push("1");
      }
    } else if (c == "]") {
      const numOfChars = Number(nums[nums.length - 1]);
      if (nums.length == 1) {
        for (let i = 0; i < numOfChars; i++) {
          res += chars[chars.length - 1];
        }
      } else {
        for (let i = 0; i < numOfChars; i++) {
          chars[chars.length - 2] += chars[chars.length - 1];
        }
      }
      chars.pop();
      nums.pop();
    }
  }
  if (chars.length > 0) {
    res += chars[0];
  }
  return res;
};
