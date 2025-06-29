#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  int candy(vector<int> &ratings) {
    if (ratings.size() == 1) {
      return 1;
    }
    vector<int> distributions(ratings.size());
    vector<int> initDistributionPos;

    for (int i = 0; i < ratings.size(); i++) {
      int r = ratings[i];
      if (i == 0) {
        if (ratings[i + 1] >= r) {
          distributions[i] = 1;
          initDistributionPos.push_back(i);
        }
      } else if (i == ratings.size() - 1) {
        if (ratings[i - 1] >= r) {
          distributions[i] = 1;
          initDistributionPos.push_back(i);
        }
      } else {
        if (ratings[i - 1] >= r && ratings[i + 1] >= r) {
          distributions[i] = 1;
          initDistributionPos.push_back(i);
        }
      }
    }
    for (int pos : initDistributionPos) {
      for (int i = pos - 1; i >= 0; i--) {
        if (ratings[i] <= ratings[i + 1]) {
          break;
        }
        distributions[i] = max(distributions[i], distributions[i + 1] + 1);
      }
      for (int i = pos + 1; i < ratings.size(); i++) {
        if (ratings[i] <= ratings[i - 1]) {
          break;
        }
        distributions[i] = max(distributions[i], distributions[i - 1] + 1);
      }
    }
    int res = 0;
    for (int d : distributions) {
      res += d;
    }
    return res;
  }
};

int main() {
  Solution s;
  vector<int> ratings = {3, 2, 1, 1, 4, 3, 3};
  cout << s.candy(ratings) << endl;
  return 0;
}
