#include <algorithm>
#include <map>
#include <sstream>
#include <vector>

using namespace std;

class Solution {
public:
  string makeKey(vector<int> &nums) {
    ostringstream oss;
    for (size_t i = 0; i < nums.size(); i++) {
      if (i != 0) {
        oss << ",";
      }
      oss << nums[i];
    }
    return oss.str();
  }

  vector<vector<int>> combinationSum(vector<int> &candidates, int target) {
    sort(candidates.begin(), candidates.end());
    vector<vector<vector<int>>> dp(target + 1);

    for (int i = 2; i < dp.size(); i++) {
      dp[i] = vector<vector<int>>{};
      map<string, bool> c;

      for (int num : candidates) {
        if (i < num) {
          break;
        } else if (i == num) {
          dp[i].push_back(vector<int>{num});
        } else {
          auto remain = i - num;
          for (auto nums : dp[remain]) {
            vector<int> newCombination = {num};
            newCombination.insert(newCombination.end(), nums.begin(),
                                  nums.end());
            sort(newCombination.begin(), newCombination.end());
            auto key = makeKey(newCombination);
            if (c[key]) {
              continue;
            } else {
              c[key] = true;
              dp[i].push_back(newCombination);
            }
          }
        }
      }
    }
    return dp[target];
  }
};
