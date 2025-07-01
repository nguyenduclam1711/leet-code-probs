#include <iostream>
#include <vector>

using namespace std;

class Solution {
private:
  void recurse(int i, int n, int k, vector<int> &arr,
               vector<vector<int>> &res) {
    if (arr.size() == k) {
      vector<int> newArr;
      newArr.assign(arr.begin(), arr.end());
      res.push_back(newArr);
      return;
    }
    for (int j = i + 1; j <= n; j++) {
      arr.push_back(j);
      this->recurse(j, n, k, arr, res);
      arr.assign(arr.begin(), arr.end() - 1);
    }
  }

public:
  vector<vector<int>> combine(int n, int k) {
    vector<vector<int>> res;
    vector<int> arr;
    this->recurse(0, n, k, arr, res);
    return res;
  }
};

int main() {
  Solution s;
  auto res1 = s.combine(4, 2);
  for (auto vec : res1) {
    for (auto n : vec) {
      cout << n << " ";
    }
    cout << endl;
  }
  return 0;
}
