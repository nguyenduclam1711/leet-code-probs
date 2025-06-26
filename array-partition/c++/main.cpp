#include <algorithm>
#include <iostream>
#include <vector>

class Solution {
public:
  int arrayPairSum(std::vector<int> &nums) {
    std::sort(nums.begin(), nums.end());
    int res = 0;

    for (int i = 0; i < nums.size(); i += 2) {
      res += nums[i];
    }
    return res;
  }
};

int main(int argc, char *argv[]) {
  Solution s;
  std::vector<int> nums1 = {1, 4, 3, 2};
  std::cout << s.arrayPairSum(nums1) << "\n";
  std::vector<int> nums2 = {6, 2, 6, 5, 1, 2};
  std::cout << s.arrayPairSum(nums2) << "\n";
  return 0;
}
