#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  int maxProfit(vector<int> &prices, int fee) {
    vector<vector<int>> dp(prices.size(), vector<int>(2));
    dp[0][0] = 0;
    dp[0][1] = -prices[0];

    for (int i = 1; i < prices.size(); i++) {
      dp[i][0] = max(prices[i] + dp[i - 1][1] - fee, dp[i - 1][0]);
      dp[i][1] = max(dp[i - 1][0] - prices[i], dp[i - 1][1]);
    }
    return dp[prices.size() - 1][0];
  }
};

int main() {
  Solution s;
  vector<int> prices = {1, 3, 2, 8, 4, 9};
  cout << s.maxProfit(prices, 2) << endl;
  return 0;
}
