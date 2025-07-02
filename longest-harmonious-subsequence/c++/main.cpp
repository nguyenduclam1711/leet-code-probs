#include <algorithm>
#include <iostream>
#include <set>
#include <string>
#include <unordered_map>
#include <vector>

using namespace std;

vector<string> split(string str, char del) {
  string temp = "";
  vector<string> res;
  for (int i = 0; i < (int)str.size(); i++) {
    if (str[i] != del) {
      temp += str[i];
    } else {
      res.push_back(temp);
      temp = "";
    }
  }
  res.push_back(temp);
  return res;
}

class Solution {
public:
  int findLHS(vector<int> &nums) {
    unordered_map<string, int> count;
    set<int> uniqueNums;

    for (auto n : nums) {
      uniqueNums.insert(n);
      string combination = to_string(n) + " " + to_string(n + 1);
      count[combination] = count[combination] + 1;
      string combination2 = to_string(n - 1) + " " + to_string(n);
      count[combination2] = count[combination2] + 1;
    }
    if (uniqueNums.size() == 1) {
      return 0;
    }
    int res = 1;
    for (auto i = count.begin(); i != count.end(); i++) {
      if (i->second > 1) {
        auto splitedKeys = split(i->first, ' ');
        auto n1 = stoi(splitedKeys[0]);
        auto n2 = stoi(splitedKeys[1]);
        if (uniqueNums.count(n1) && uniqueNums.count(n2)) {
          res = max(res, i->second);
        }
      }
    }
    if (res == 1) {
      return 0;
    }
    return res;
  }

  int findLHS2(vector<int> &nums) {
    sort(nums.begin(), nums.end());
    int res = 0;
    int j = 0;

    for (int i = 0; i < nums.size(); i++) {
      while (nums[i] - nums[j] > 1) {
        j++;
      }
      if (nums[i] - nums[j] == 1) {
        res = max(res, i - j + 1);
      }
    }
    return res;
  }
};

int main() {
  Solution s;
  vector<int> nums = {1, 3, 2, 2, 5, 2, 3, 7};
  cout << s.findLHS(nums) << endl;
  vector<int> nums2 = {1, 3, 2, 2, 5, 2, 3, 7};
  cout << s.findLHS2(nums2) << endl;
  return 0;
}
