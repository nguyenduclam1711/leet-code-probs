function possibleStringCount(word: string): number {
  if (word.length === 1) {
    return 1;
  }
  let res = 1;
  for (let i = 1; i < word.length; i++) {
    if (word[i] === word[i - 1]) {
      res++;
    }
  }
  return res;
};
