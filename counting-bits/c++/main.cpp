#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> countBits(int n) {
    auto len = n + 1, offset = 1;
    vector<int> res(len);

    for (int i = 1; i < len; i++) {
      if (offset * 2 == i) {
        offset = i;
      }
      res[i] = 1 + res[i - offset];
    }
    return res;
  }
};

int main() {
  Solution s;
  auto nums = s.countBits(5);
  for (auto n : nums) {
    cout << n << " ";
  }
  cout << endl;
  return 0;
}
