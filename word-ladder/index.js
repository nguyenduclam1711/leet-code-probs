/**
 * @param {string} beginWord
 * @param {string} endWord
 * @param {string[]} wordList
 * @return {number}
 */
var ladderLength = function (beginWord, endWord, wordList) {
  const uniqueWords = {};
  for (const w of wordList) {
    uniqueWords[w] = true;
  }
  if (!uniqueWords[endWord]) {
    return 0;
  }
  const wLen = beginWord.length;
  let res = 0;
  let q = [beginWord];
  const checkedWords = {};
  checkedWords[beginWord] = true;
  let isValid = false;

  while (q.length > 0) {
    res++;
    const newQ = [];
    for (const w of q) {
      if (w === endWord) {
        isValid = true;
        break;
      }
      for (let i = 0; i < wLen; i++) {
        for (let j = 0; j < 26; j++) {
          const currNewChar = String.fromCharCode("a".charCodeAt(0) + j);
          let nextW = "";
          if (i === 0) {
            nextW = currNewChar + w.slice(1);
          } else if (i === wLen - 1) {
            nextW = w.slice(0, i) + currNewChar;
          } else {
            nextW = w.slice(0, i) + currNewChar + w.slice(i + 1);
          }
          if (uniqueWords[nextW] && !checkedWords[nextW]) {
            newQ.push(nextW);
            checkedWords[nextW] = true;
          }
        }
      }
    }
    if (isValid) {
      break;
    }
    q = newQ;
  }
  if (isValid) {
    return res;
  }
  return 0;
};
