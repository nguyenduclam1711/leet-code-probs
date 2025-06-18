#include <iostream>
#include <string>
class Solution {
public:
  bool isPalindrome(int x) {
    std::string str = std::to_string(x);
    int l = 0;
    int r = str.size() - 1;

    while (l < r) {
      if (str[l] != str[r]) {
        return false;
      }
      l++;
      r--;
    }
    return true;
  }
};

void test(Solution &s, int x, bool expected) {
  bool res = s.isPalindrome(x);

  std::cout << "Current test number: " << x << std::endl;
  if (res == expected) {
    std::cout << "Right\n";
  } else {
    std::cout << "Wrong\n";
  }
}

int main() {
  Solution s;
  test(s, 121, true);
  test(s, -121, false);
  test(s, 10, false);
  return 0;
}
