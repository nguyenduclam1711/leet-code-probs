#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> twoSum(vector<int> &nums, int target) {
    vector<int> result;
    unordered_map<int, int> mapNum;

    for (int i = 0; i < nums.size(); i++) {
      int currNum = nums[i];
      int leftNum = target - currNum;

      if (mapNum.count(leftNum)) {
        result.push_back(i);
        result.push_back(mapNum[leftNum]);
        return result;
      } else {
        mapNum[currNum] = i;
      }
    }
    return result;
  }
};

void test(Solution &solution, vector<int> &nums, int target) {
  vector<int> res = solution.twoSum(nums, target);
  if (!res.empty()) {
    cout << "Dit nhau: " << res[0] << " " << res[1] << endl;
  } else {
    cout << "This shit is wrong" << endl;
  }
}

int main() {
  Solution solution;
  vector<int> nums1 = {2, 7, 11, 15};
  test(solution, nums1, 9);
  vector<int> nums2 = {3, 2, 4};
  test(solution, nums2, 6);
  vector<int> nums3 = {3, 3};
  test(solution, nums3, 6);

  return 0;
}
