const primeNumbers = [2, 3, 5, 7];

function checkIsPrimeNumber(n) {
  let isPrimeNumber = true;
  let largestDivisible = 0;

  for (let i = 0; i < primeNumbers.length; i++) {
    const primeNumber = primeNumbers[i];
    if (primeNumber === n) {
      return {
        isPrimeNumber,
        largestDivisible,
      };
    }
    if (n % primeNumber === 0) {
      isPrimeNumber = false;
      largestDivisible = n / primeNumber;
      break;
    }
  }
  if (isPrimeNumber) {
    primeNumbers.push(n);
  }
  return {
    isPrimeNumber,
    largestDivisible,
  };
}

/**
 * @param {number} n
 * @return {number}
 */
var minSteps = function(n) {
  const dp = Array.from(Array(n + 1));
  dp[0] = 0;
  dp[1] = 0;
  dp[2] = 2;

  for (let i = 3; i < dp.length; i++) {
    const { largestDivisible, isPrimeNumber } = checkIsPrimeNumber(i);

    if (isPrimeNumber) {
      dp[i] = i;
    } else {
      dp[i] = dp[largestDivisible] + i / largestDivisible;
    }
  }
  return dp[n];
};
