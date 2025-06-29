#include <iostream>
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

class BSTIterator {
private:
  int currentIndex = 0;
  vector<int> inorderArr;

public:
  BSTIterator(TreeNode *root) { addToInorderArr(root, inorderArr); }

  void addToInorderArr(TreeNode *curr, vector<int> &inorderArr) {
    if (!curr) {
      return;
    }
    addToInorderArr(curr->left, inorderArr);
    inorderArr.push_back(curr->val);
    addToInorderArr(curr->right, inorderArr);
  }

  int next() {
    int res = inorderArr[currentIndex];
    currentIndex++;
    return res;
  }

  bool hasNext() { return currentIndex < inorderArr.size(); }
};

int main() { return 0; }
