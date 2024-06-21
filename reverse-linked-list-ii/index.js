function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

/**
 * @param {ListNode} head
 * @param {number} left
 * @param {number} right
 * @return {ListNode}
 */
var reverseBetween = function (head, left, right) {
  let i = 1;
  let curr = head;
  let groupHead = null;
  let groupTail = null;
  let prevGroupHead = null;

  while (curr) {
    const nextCurr = curr.next;
    if (i === left - 1) {
      prevGroupHead = curr;
    } else if (i >= left && i <= right) {
      if (i === left) {
        groupHead = curr;
      } else if (i === left + 1) {
        groupTail = curr;
        groupHead.next = groupTail.next;
        groupTail.next = groupHead;
        const temp = groupTail;
        groupTail = groupHead;
        groupHead = temp;
      } else {
        groupTail.next = curr.next;
        curr.next = groupHead;
        groupHead = curr;
      }
      if (prevGroupHead) {
        prevGroupHead.next = groupHead;
      }
    }
    curr = nextCurr;
    i++;
  }
  if (left === 1) {
    return groupHead;
  }
  return head;
};
