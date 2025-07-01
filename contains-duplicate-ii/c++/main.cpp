#include <cstdlib>
#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
  bool containsNearbyDuplicate(vector<int> &nums, int k) {
    unordered_map<int, int> m;
    for (auto i = 0; i < nums.size(); i++) {
      auto n = nums[i];
      if (m.count(n) == 0) {
        m[n] = i;
      } else {
        auto diffAbs = abs(i - m[n]);
        if (diffAbs <= k) {
          return true;
        }
        m[n] = i;
      }
    }
    return false;
  }
};

int main() {
  Solution s;
  vector<int> nums = {1, 2, 3, 1};
  cout << s.containsNearbyDuplicate(nums, 3) << endl;
  return 0;
}
