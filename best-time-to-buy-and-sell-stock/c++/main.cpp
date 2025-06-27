#include <iostream>

using namespace std;

class Solution {
public:
  int maxProfit(vector<int> &prices) {
    int profit = 0, currentProfit = 0, buyDay = 0;

    for (int i = 1; i < prices.size(); i++) {
      currentProfit = prices[i] - prices[buyDay];
      if (currentProfit > profit) {
        profit = currentProfit;
      } else if (currentProfit < 0) {
        buyDay = i;
      }
    }
    return profit;
  }
};

int main() {
  Solution s;
  vector<int> prices = {7, 1, 5, 3, 6, 4};
  cout << "Profit: " << s.maxProfit(prices) << endl;
  return 0;
}
