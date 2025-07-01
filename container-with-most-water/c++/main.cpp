#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  int maxArea(vector<int> &height) {
    int area = 0, l = 0, r = height.size() - 1;
    while (l < r) {
      auto width = r - l;
      if (height[l] < height[r]) {
        area = max(area, width * height[l]);
        l++;
      } else {
        area = max(area, width * height[r]);
        r--;
      }
    }
    return area;
  }
};

int main() {
  Solution s;
  vector<int> height = {1, 8, 6, 2, 5, 4, 8, 3, 7};
  cout << s.maxArea(height) << endl;
  return 0;
}
