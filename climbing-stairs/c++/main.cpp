#include <iostream>

using namespace std;

class Solution {
public:
  int climbStairs(int n) {
    int a = 0, b = 1, i = 0;
    while (i < n) {
      auto tmp = a;
      a = b;
      b += tmp;
      i++;
    }
    return b;
  }
};

int main() {
  Solution s;
  cout << s.climbStairs(3) << endl;
  return 0;
}
