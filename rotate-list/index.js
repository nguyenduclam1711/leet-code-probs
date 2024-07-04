var rotateRight = function (head, k) {
  if (!head) {
    return head;
  }
  const arr = [];
  let curr = head;

  while (curr) {
    arr.push(curr);
    curr = curr.next;
  }
  const numberOfRotate = k % arr.length;

  if (numberOfRotate === 0) {
    return head;
  }
  const lastPosInArr = arr.length - numberOfRotate - 1;
  arr[arr.length - 1].next = head;
  arr[lastPosInArr].next = null;
  head = arr[lastPosInArr + 1];
  return head;
};
