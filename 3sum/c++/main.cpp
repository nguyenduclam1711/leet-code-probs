#include <algorithm>
#include <vector>

using namespace std;

class Solution {
public:
  vector<vector<int>> threeSum(vector<int> &nums) {
    sort(nums.begin(), nums.end(), [](int a, int b) { return a < b; });
    vector<vector<int>> result;
    int prevA;

    for (int i = 0; i < nums.size(); i++) {
      int num = nums[i];
      if (num == prevA && i != 0) {
        continue;
      }
      int l = i + 1;
      int r = nums.size() - 1;
      int prevB = nums[0] - 1;

      while (l < r) {
        if (nums[l] == prevB) {
          l++;
          continue;
        }
        int tmp = nums[l] + nums[r];
        if (tmp == -num) {
          result.push_back({
              num,
              nums[l],
              nums[r],
          });
          prevB = nums[l];
          l++;
          r--;
        } else if (tmp < -num) {
          l++;
        } else {
          r--;
        }
      }
      prevA = num;
    }
    return result;
  }
};
