/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
  if (!head) {
    return head;
  }
  let newHead = null;
  let prevNewHead = null;
  let curr = null;
  let currNumNode = null;
  let count = 0;
  currNumNode = head;
  curr = head;

  while (curr) {
    if (curr.val === currNumNode.val) {
      count++;
    } else {
      if (count === 1) {
        if (!newHead) {
          newHead = currNumNode;
        }
        if (prevNewHead) {
          prevNewHead.next = currNumNode;
        }
        prevNewHead = currNumNode;
      }
      count = 1;
      currNumNode = curr;
    }
    curr = curr.next;
  }
  if (count === 1) {
    if (!newHead) {
      newHead = currNumNode;
    }
    if (prevNewHead) {
      prevNewHead.next = currNumNode;
    }
    prevNewHead = currNumNode;
  }
  if (prevNewHead) {
    prevNewHead.next = null;
  }
  return newHead;
};
