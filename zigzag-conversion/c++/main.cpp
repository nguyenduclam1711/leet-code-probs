#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  string convert(string s, int numRows) {
    if (numRows == 1) {
      return s;
    }
    vector<vector<char>> arr(numRows);
    for (int i = 0; i < arr.size(); i++) {
      arr[i] = vector<char>(s.size());
    }

    int row = 0, col = 0;
    bool isDiagnol = false;

    for (int i = 0; i < s.size(); i++) {
      arr[row][col] = s[i];

      if (!isDiagnol) {
        if (row < numRows - 1) {
          row++;
        } else {
          row--;
          col++;
          isDiagnol = true;
        }
      } else {
        if (row == 0) {
          row++;
          isDiagnol = false;
        } else {
          row--;
          col++;
        }
      }
    }
    string res = "";
    for (int i = 0; i < arr.size(); i++) {
      for (int j = 0; j < arr[i].size(); j++) {
        if (arr[i][j] != 0) {
          res += arr[i][j];
        }
      }
    }
    return res;
  }
};

int main() {
  Solution s;
  cout << s.convert("PAYPALISHIRING", 3) << endl;
  return 0;
}
