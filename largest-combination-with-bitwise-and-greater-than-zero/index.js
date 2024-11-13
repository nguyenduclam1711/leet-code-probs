function find(n, setBits) {
  let i = 31;

  while (n > 0) {
    const a = n & 1;
    setBits[i] += a;
    n >>= 1;
    i--;
  }
}

function largestCombination(candidates) {
  const setBits = Array.from(Array(32)).fill(0);
  for (const c of candidates) {
    find(c, setBits);
  }
  return Math.max(...setBits);
}
