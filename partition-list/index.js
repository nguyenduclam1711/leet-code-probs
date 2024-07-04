/**
 * @param {ListNode} head
 * @param {number} x
 * @return {ListNode}
 */
var partition = function (head, x) {
  if (!head) {
    return head;
  }
  let leftHead = null;
  let leftTail = null;
  let rightHead = null;
  let rightTail = null;
  let curr = head;

  while (curr) {
    const nextCurr = curr.next;
    if (curr.val < x) {
      if (!leftHead) {
        leftHead = curr;
      }
      if (leftTail) {
        leftTail.next = curr;
      }
      leftTail = curr;
    } else {
      if (!rightHead) {
        rightHead = curr;
      }
      if (rightTail) {
        rightTail.next = curr;
      }
      rightTail = curr;
    }
    curr = nextCurr;
  }
  if (leftTail) {
    leftTail.next = rightHead;

    if (rightTail) {
      rightTail.next = null;
    }
    return leftHead;
  }
  if (rightTail) {
    rightTail.next = null;
  }
  return rightHead;
};
