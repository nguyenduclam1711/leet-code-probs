#include <iostream>

struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
  ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
    ListNode *p1 = l1;
    ListNode *p2 = l2;
    ListNode *head = nullptr;
    ListNode *l = head;
    int needToAdd = 0;

    while (p1 != nullptr || p2 != nullptr) {
      int val1 = 0;
      int val2 = 0;

      if (p1) {
        val1 = p1->val;
        p1 = p1->next;
      } else {
        val1 = 0;
      }
      if (p2) {
        val2 = p2->val;
        p2 = p2->next;
      } else {
        val2 = 0;
      }
      int newVal = val1 + val2 + needToAdd;

      if (newVal < 10) {
        needToAdd = 0;
      } else {
        newVal = newVal % 10;
        needToAdd = 1;
      }
      if (!head) {
        head = new ListNode(newVal);
        l = head;
      } else {
        l->next = new ListNode(newVal);
        l = l->next;
      }
    }
    if (needToAdd == 1) {
      l->next = new ListNode(needToAdd);
    }
    return head;
  }
};

int main(int argc, char *argv[]) {
  ListNode *l1 = new ListNode(2);
  l1->next = new ListNode(4);
  l1->next->next = new ListNode(3);
  ListNode *l2 = new ListNode(5);
  l2->next = new ListNode(6);
  l2->next->next = new ListNode(4);

  Solution s;
  ListNode *result = s.addTwoNumbers(l1, l2);
  std::cout << result->val << "\n";
  std::cout << result->next->val << "\n";
  std::cout << result->next->next->val << "\n";

  return 0;
}
