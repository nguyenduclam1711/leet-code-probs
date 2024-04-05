/**
 * @param {number[][]} points
 * @return {number}
 */
var findMinArrowShots = function (points) {
  points.sort((a, b) => {
    return a[1] - b[1];
  });
  let res = 1;
  let curMaxPoint = points[0][1];

  for (let i = 1; i < points.length; i++) {
    if (points[i][0] <= curMaxPoint) {
      continue;
    }
    res++;
    curMaxPoint = points[i][1];
  }
  return res;
};
