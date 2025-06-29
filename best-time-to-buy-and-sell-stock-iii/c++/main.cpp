#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  int maxProfit(vector<int> &prices) {
    int low = prices[0], high = prices[prices.size() - 1];
    vector<int> firstDp(prices.size(), 0);
    vector<int> secondDp(prices.size(), 0);

    for (int i = 1; i < prices.size(); i++) {
      firstDp[i] = max(prices[i] - low, firstDp[i - 1]);
      if (prices[i] < low) {
        low = prices[i];
      }
      int j = prices.size() - i - 1;
      secondDp[j] = max(high - prices[j], secondDp[j + 1]);
      if (prices[j] > high) {
        high = prices[j];
      }
    }
    int res = 0;
    for (int i = 0; i < firstDp.size(); i++) {
      int currRes;
      if (i < firstDp.size() - 1) {
        currRes = firstDp[i] + secondDp[i + 1];
      } else {
        currRes = firstDp[i];
      }
      if (currRes > res) {
        res = currRes;
      }
    }
    return res;
  }
};

int main() {
  Solution s;
  vector<int> prices = {3, 3, 5, 0, 0, 3, 1, 4};
  cout << s.maxProfit(prices) << endl;
  return 0;
}
