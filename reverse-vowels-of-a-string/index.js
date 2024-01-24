/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function (s) {
  const vowels = {
    a: true,
    e: true,
    i: true,
    o: true,
    u: true,
    A: true,
    E: true,
    I: true,
    O: true,
    U: true,
  };
  const arr = Array(s.length);
  let i = 0;
  let j = s.length - 1;

  while (i <= j) {
    if (vowels[s[i]] && vowels[s[j]]) {
      arr[i] = s[j];
      arr[j] = s[i];
      i++;
      j--;
    } else {
      if (!vowels[s[i]]) {
        arr[i] = s[i];
        i++;
      }
      if (!vowels[s[j]]) {
        arr[j] = s[j];
        j--;
      }
    }
  }
  return arr.join("");
};

console.log(reverseVowels("hello"));
