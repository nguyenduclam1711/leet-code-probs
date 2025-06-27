#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  int maxProfit(vector<int> &prices) {
    // NOTE: to get all the profit by checking the last zero price (cover the
    // edge case: the last element is larger than the previous pointer price)
    prices.push_back(0);
    int res = 0, i = 0, j = 1, currMax = prices[1];

    while (i < prices.size() && j < prices.size()) {
      if (prices[i] >= prices[j]) {
        if (currMax > prices[i]) {
          res += currMax - prices[i];
        }
        i = j;
        j = i + 1;
        if (j < prices.size()) {
          currMax = prices[j];
        }
      } else {
        if (prices[j] > currMax) {
          currMax = prices[j];
          j++;
        } else {
          res += currMax - prices[i];
          i = j;
          j = i + 1;
          if (j < prices.size()) {
            currMax = prices[j];
          }
        }
      }
    }
    return res;
  }
};

int main() {
  Solution s;
  vector<int> prices = {7, 1, 5, 3, 6, 4};
  cout << s.maxProfit(prices) << endl;
  return 0;
}
