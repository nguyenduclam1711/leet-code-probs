var hammingWeight = function (n) {
  let w = 0;
  while (n !== 0) {
    if (n & (1 == 1)) {
      w++;
    }
    n = n >> 1;
  }
  return w;
};

var countBits = function (n) {
  const res = Array(n + 1);

  for (let i = 0; i < n + 1; i++) {
    res[i] = hammingWeight(i);
  }
  return res;
};

console.log(countBits(10));
