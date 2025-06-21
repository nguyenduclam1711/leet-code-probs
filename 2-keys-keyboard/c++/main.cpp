#include <vector>

std::vector<int> primeNumbers = {2, 3, 5, 7};

struct CheckIsPrimeNumberRes {
  int largestDivisible;
  bool isPrimeNumber;
};

class Solution {
private:
  CheckIsPrimeNumberRes checkIsPrimeNumber(int n) {
    bool isPrimeNumber = true;
    int largestDivisible = 0;

    for (int i = 0; i < primeNumbers.size(); i++) {
      int primeNumber = primeNumbers[i];
      if (n == primeNumber) {
        return CheckIsPrimeNumberRes{
            .largestDivisible = largestDivisible,
            .isPrimeNumber = isPrimeNumber,
        };
      }
      if (n % primeNumber == 0) {
        isPrimeNumber = false;
        largestDivisible = n / primeNumber;
        break;
      }
    }

    if (isPrimeNumber) {
      primeNumbers.push_back(n);
    }
    return CheckIsPrimeNumberRes{
        .largestDivisible = largestDivisible,
        .isPrimeNumber = isPrimeNumber,
    };
  }

public:
  int minSteps(int n) {
    std::vector<int> dp(n + 1, 0);
    if (n > 1) {
      dp[2] = 2;
    }
    for (int i = 3; i < dp.size(); i++) {
      CheckIsPrimeNumberRes res = this->checkIsPrimeNumber(i);
      if (res.isPrimeNumber) {
        dp[i] = i;
      } else {
        dp[i] = dp[res.largestDivisible] + i / res.largestDivisible;
      }
    }
    return dp[n];
  }
};

int main() {
  Solution s;
  s.minSteps(6);
  return 0;
}
