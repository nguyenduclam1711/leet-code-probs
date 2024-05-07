/**
 * @param {string} ransomNote
 * @param {string} magazine
 * @return {boolean}
 */
var canConstruct = function(ransomNote, magazine) {
  if (magazine.length < ransomNote.length) {
    return false;
  }
  const hashMap = {};

  for (const c of magazine) {
    hashMap[c] = (hashMap[c] || 0) + 1;
  }
  for (const c of ransomNote) {
    if (hashMap[c] === undefined || hashMap[c] <= 0) {
      return false;
    }
    hashMap[c]--;
  }
  return true;
};
