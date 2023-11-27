var getSwapPosition = function (x, y, matrixLen) {
  const maxPos = matrixLen - 1;
  // top border
  if (x === 0) {
    return {
      x: maxPos - y,
      y: 0,
    };
  }
  // right border
  if (y === maxPos) {
    return {
      x: 0,
      y: x,
    };
  }
  // bottom border
  if (x === maxPos) {
    return {
      x: maxPos - y,
      y: maxPos,
    };
  }
  // left border
  if (y === 0) {
    return {
      x: maxPos,
      y: x,
    };
  }
  return -1, -1;
};

var getMapKey = function (x, y) {
  return `${x},${y}`;
};

var rotate = function (matrix) {
  let matrixLen = matrix.length;
  let loop = 0;
  const m = {};

  while (matrixLen > 0) {
    const totalNumberOfOuterLinePos = matrixLen * 4 - 4;
    if (totalNumberOfOuterLinePos > 0) {
      let x = 0;
      let y = 0;
      for (let i = 0; i < totalNumberOfOuterLinePos; i++) {
        const { x: swapX, y: swapY } = getSwapPosition(x, y, matrixLen);
        const posX = x + loop;
        const posY = y + loop;
        const posSwapX = swapX + loop;
        const posSwapY = swapY + loop;

        m[getMapKey(posX, posY)] = matrix[posX][posY];
        if (getMapKey(posSwapX, posSwapY) in m) {
          matrix[posX][posY] = m[getMapKey(posSwapX, posSwapY)];
        } else {
          matrix[posX][posY] = matrix[posSwapX][posSwapY];
        }
        if (x <= y) {
          if (x == y && y == matrixLen - 1) {
            y--;
            continue;
          }
          if (y < matrixLen - 1) {
            y++;
          } else {
            x++;
          }
        } else {
          if (y > 0) {
            y--;
          } else {
            x--;
          }
        }
      }
    }
    // the inside outer line always smaller than the larger on by 2
    matrixLen -= 2;
    loop++;
  }
};
