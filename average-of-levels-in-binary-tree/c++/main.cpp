#include <iostream>
#include <new>
#include <vector>

using namespace std;

struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

class Solution {
public:
  vector<double> averageOfLevels(TreeNode *root) {
    vector<double> res;
    vector<TreeNode *> q = {root};

    while (q.size() > 0) {
      vector<TreeNode *> newQ;
      double total = 0;

      for (int i = 0; i < q.size(); i++) {
        auto node = q[i];
        if (node->left != nullptr) {
          newQ.push_back(node->left);
        }
        if (node->right != nullptr) {
          newQ.push_back(node->right);
        }
        total += node->val;
      }
      res.push_back(total / q.size());
      q = newQ;
    }
    return res;
  }
};

int main() {
  Solution s;
  TreeNode *root = new TreeNode(
      3, new TreeNode(9, new TreeNode(15), new TreeNode(7)), new TreeNode(20));
  auto avgLevels = s.averageOfLevels(root);
  for (int i = 0; i < avgLevels.size(); i++) {
    cout << avgLevels[i] << " ";
  }
  cout << endl;
  return 0;
}
