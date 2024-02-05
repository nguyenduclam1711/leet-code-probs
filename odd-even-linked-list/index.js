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
var oddEvenList = function(head) {
  if (!head) {
    return head;
  }

  let oddPointer = head;
  let evenPointer = head.next;
  const firstEven = evenPointer;

  while (oddPointer && evenPointer) {
    if (!evenPointer.next) {
      break;
    }
    oddPointer.next = evenPointer.next;
    oddPointer = oddPointer.next;
    evenPointer.next = oddPointer.next;
    evenPointer = evenPointer.next;
  }
  oddPointer.next = firstEven;
  return head;
};
