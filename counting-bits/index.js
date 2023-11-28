var countBits = function(n) {
  const len = n + 1;
  const res = Array(len);
  let offset = 1;
  res[0] = 0;

  for (let i = 1; i < len; i++) {
    if (offset * 2 === i) {
      offset = i;
    }
    res[i] = 1 + res[i - offset];
  }
  return res;
};

console.log(countBits(31));
