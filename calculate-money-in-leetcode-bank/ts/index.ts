function totalMoney(n: number): number {
  let total = 0;
  let week = 0;
  let pushMoney = 0;
  let i = 1;
  while (i <= n) {
    if (i % 7 === 1) { // monday
      week++;
      pushMoney = week;
      total += week;
    } else {
      pushMoney++;
      total += pushMoney;
    }
    i++;
  }
  return total;
};
