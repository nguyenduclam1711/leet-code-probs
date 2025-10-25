class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}

function swapPairs(head: ListNode | null): ListNode | null {
  let prev: ListNode | null = null;
  let currNode = head;
  let newHead = head;
  while (currNode && currNode.next) {
    const nextNode = currNode.next;
    currNode.next = nextNode.next;
    nextNode.next = currNode;
    if (prev) {
      prev.next = nextNode;
    }
    currNode = currNode.next;
    prev = nextNode.next;
    if (newHead === head) {
      newHead = nextNode;
    }
  }
  return newHead
};
