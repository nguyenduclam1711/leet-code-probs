#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
  int possibleStringCount(string word) {
    if (word.size() == 1) {
      return 1;
    }
    int res = 1;
    for (int i = 1; i < word.size(); i++) {
      if (word[i] == word[i - 1]) {
        res++;
      }
    }
    return res;
  }
};

int main() {
  Solution s;
  auto test1 = "abbcccc";
  cout << "Test 1: " << test1 << endl << s.possibleStringCount(test1) << endl;
  auto test2 = "ere";
  cout << "Test 2: " << test2 << " " << s.possibleStringCount(test2) << endl;
  return 0;
}
