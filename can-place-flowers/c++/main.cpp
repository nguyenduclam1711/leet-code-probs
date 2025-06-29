#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  bool canPlaceFlowers(vector<int> &flowerbed, int n) {
    auto localN = n;
    for (int i = 0; i < flowerbed.size() && localN > 0; i++) {
      int bed = flowerbed[i];

      if (bed == 1) {
        continue;
      }

      bool isValidPos = true;
      int nextBedPos = i + 1, prevBedPos = i - 1;

      if (i == 0) {
        if (nextBedPos < flowerbed.size() && flowerbed[nextBedPos] == 1) {
          isValidPos = false;
        }
      } else if (i == flowerbed.size() - 1) {
        if (prevBedPos >= 0 && flowerbed[prevBedPos] == 1) {
          isValidPos = false;
        }
      } else {
        if (flowerbed[prevBedPos] == 1 || flowerbed[nextBedPos] == 1) {
          isValidPos = false;
        }
      }
      if (isValidPos) {
        flowerbed[i] = 1;
        localN--;
      }
    }
    return localN == 0;
  }
};

int main() {
  Solution s;
  vector<int> flowerbed = {1, 0, 0, 0, 1};
  cout << s.canPlaceFlowers(flowerbed, 2) << endl;
  return 0;
}
