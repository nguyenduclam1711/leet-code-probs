#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
  void dfs(int num, int k, int n, int total, vector<int> &nums,
           vector<vector<int>> &res) {
    if (total > n) {
      return;
    }
    if (nums.size() == k) {
      if (total == n) {
        vector<int> newNums;
        newNums.assign(nums.begin(), nums.end());
        res.push_back(newNums);
      }
      return;
    }
    for (int i = num + 1; i < 10; i++) {
      nums.push_back(i);
      dfs(i, k, n, total + i, nums, res);
      nums.assign(nums.begin(), nums.end() - 1);
    }
  }

public:
  vector<vector<int>> combinationSum3(int k, int n) {
    vector<vector<int>> res;
    vector<int> nums;
    this->dfs(0, k, n, 0, nums, res);
    return res;
  }
};

int main() {
  Solution s;
  auto res1 = s.combinationSum3(3, 7);
  cout << "Test 1\n";
  for (auto vec : res1) {
    for (auto n : vec) {
      cout << n << " ";
    }
    cout << endl;
  }

  auto res2 = s.combinationSum3(3, 9);
  cout << "Test 2\n";
  for (auto vec : res2) {
    for (auto n : vec) {
      cout << n << " ";
    }
    cout << endl;
  }

  return 0;
}
