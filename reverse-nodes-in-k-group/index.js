function reverseBetween(node, prevGroupNode, left, right) {
  let i = left;
  let groupHead = null;
  let groupTail = null;
  let curr = node;

  while (i >= left && i <= right) {
    const nextCurr = curr.next;
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
    if (prevGroupNode) {
      prevGroupNode.next = groupHead;
    }
    curr = nextCurr;
    i++;
  }
  return { groupHead, groupTail };
}

/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var reverseKGroup = function(head, k) {
  if (k === 1) {
    return head;
  }
  let curr = head;
  let listLength = 0;
  while (curr) {
    listLength++;
    curr = curr.next;
  }

  curr = head;
  let res = null;
  let prev = null;
  for (let i = 1; i <= listLength && curr; i += k) {
    const leftLength = listLength - i + 1;
    if (k <= leftLength) {
      const { groupHead, groupTail } = reverseBetween(curr, prev, i, i + k - 1);
      curr = groupTail.next;
      prev = groupTail;

      if (i === 1) {
        res = groupHead;
      }
    }
  }
  return res;
};
