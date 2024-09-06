function getPos(squareNumber, n) {
  const startRow = Math.floor((squareNumber - 1) / n);
  const row = n - startRow - 1;
  let col = (squareNumber - 1) % n;
  if (startRow % 2 !== 0) {
    col = n - 1 - col;
  }
  return {
    row,
    col,
  };
}

function getNextSquares(currSquareNumber) {
  return {
    minSquare: currSquareNumber + 1,
    maxSquare: currSquareNumber + 6,
  };
}

/**
 * @param {number[][]} board
 * @return {number}
 */
var snakesAndLadders = function(board) {
  let q = [1];
  const n = board.length;
  const endSquare = n * n;
  let res = 0;
  let canReachToTheEnd = false;
  const visitedSquare = {};

  while (q.length > 0) {
    const newQ = [];

    for (const squareNumber of q) {
      if (visitedSquare[squareNumber]) {
        continue;
      }
      visitedSquare[squareNumber] = true;
      if (squareNumber === endSquare) {
        canReachToTheEnd = true;
        break;
      }
      const { minSquare, maxSquare } = getNextSquares(squareNumber);
      for (
        let nextSquare = minSquare;
        nextSquare <= maxSquare && nextSquare <= endSquare;
        nextSquare++
      ) {
        const { row, col } = getPos(nextSquare, n);
        if (board[row][col] === -1) {
          newQ.push(nextSquare);
        } else {
          newQ.push(board[row][col]);
        }
      }
    }
    if (canReachToTheEnd) {
      break;
    }
    res++;
    q = newQ;
  }
  if (canReachToTheEnd) {
    return res;
  }
  return -1;
};
