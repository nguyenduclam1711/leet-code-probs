function checkStrings(s1, wordFromDict) {
  let i = 0;
  let j = 0;
  while (i < s1.length && j < wordFromDict.length) {
    if (s1[i] === wordFromDict[j]) {
      i++;
      j++;
    } else {
      return {
        pos: i,
        isValid: false,
      };
    }
  }
  if (j === wordFromDict.length) {
    return {
      pos: i,
      isValid: true,
    };
  }
  return {
    pos: i,
    isValid: false,
  };
}

function recurse(s, wordDict, caches) {
  if (s in caches) {
    return caches[s];
  }
  if (s.length == 0) {
    return true;
  }
  const validPos = [];
  for (const word of wordDict) {
    const { pos, isValid } = checkStrings(s, word);
    if (isValid) {
      validPos.push(pos);
    }
  }
  let result = false;
  for (const p of validPos) {
    const r = recurse(s.slice(p), wordDict, caches);
    if (r) {
      result = r;
      break;
    }
  }
  caches[s] = result;
  return result;
}

/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
var wordBreak = function (s, wordDict) {
  const caches = {};
  return recurse(s, wordDict, caches);
};
