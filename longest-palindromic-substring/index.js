/**
 * @param {string} s
 * @return {string}
 */
function longestPalindrome(s) {
  const queue = [];
  let maxLen = 0;
  let res = "";

  // init queue
  for (let i = 0; i < s.length; i++) {
    queue.push({
      start: i,
      end: i,
    });
    if (i < s.length - 1) {
      const j = i + 1;
      if (s[i] === s[j]) {
        queue.push({
          start: i,
          end: j,
        });
      }
    }
  }
  while (queue.length > 0) {
    const pos = queue[0];
    const palindrome = s.slice(pos.start, pos.end + 1);

    if (palindrome.length > maxLen) {
      maxLen = palindrome.length;
      res = palindrome;
    }
    const i = pos.start - 1;
    const j = pos.end + 1;

    if (i >= 0 && j < s.length && s[i] === s[j]) {
      queue.push({
        start: i,
        end: j,
      });
    }
    queue.shift();
  }
  return res;
}
