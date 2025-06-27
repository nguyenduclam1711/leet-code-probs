#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> asteroidCollision(vector<int> &asteroids) {
    vector<int> stack;

    for (int i = 0; i < asteroids.size(); i++) {
      stack.push_back(asteroids[i]);
      while (stack.size() > 1) {
        int r = stack.size() - 1;
        int l = r - 1;
        if (stack[l] > 0 && stack[r] < 0) {
          int rAbs = abs(stack[r]);
          int lAbs = abs(stack[l]);

          if (rAbs == lAbs) {
            stack.pop_back();
            stack.pop_back();
          } else {
            if (rAbs > lAbs) {
              stack[l] = stack[r];
            }
            stack.pop_back();
          }
        } else {
          break;
        }
      }
    }
    return stack;
  }
};

int main() {
  Solution s;
  vector<int> arr1 = {5, 10, -5};
  vector<int> res1 = s.asteroidCollision(arr1);
  for (int i = 0; i < res1.size(); i++) {
    cout << arr1[i] << " ";
  }
  cout << endl;
  return 0;
}
