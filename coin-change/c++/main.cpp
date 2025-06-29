#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  int coinChange(vector<int> &coins, int amount) {
    auto maxLen = amount + 1;
    vector<int> s(maxLen);
    sort(coins.begin(), coins.end(),
         [](int first, int second) { return first < second; });
    for (int i = 0; i < s.size(); i++) {
      if (i == 0) {
        s[i] = 0;
        continue;
      }
      s[i] = maxLen;
      for (int coin : coins) {
        if (i == coin) {
          s[i] = 1;
          break;
        }
        if (i > coin) {
          s[i] = min(1 + s[i - coin], s[i]);
        } else {
          break;
        }
      }
    }
    if (s[amount] == maxLen) {
      return -1;
    }
    return s[amount];
  }
};

int main() {
  Solution s;
  vector<int> coins = {1, 2, 5};
  cout << s.coinChange(coins, 11) << endl;
  return 0;
}
