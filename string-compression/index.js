function appendChars(insertPos, currChar, chars, count) {
  let nextInsertPos = insertPos;
  chars[nextInsertPos] = currChar;
  nextInsertPos++;
  if (count > 1) {
    const strCount = count.toString();
    for (const n of strCount) {
      chars[nextInsertPos] = n;
      nextInsertPos++;
    }
  }
  return nextInsertPos;
}

/**
 * @param {character[]} chars
 * @return {number}
 */
var compress = function (chars) {
  let insertPos = 0;
  let currChar = chars[0];
  let count = 1;

  for (let i = 1; i < chars.length; i++) {
    if (chars[i] === currChar) {
      count++;
    } else {
      insertPos = appendChars(insertPos, currChar, chars, count);
      currChar = chars[i];
      count = 1;
    }
  }
  insertPos = appendChars(insertPos, currChar, chars, count);
  return insertPos;
};
