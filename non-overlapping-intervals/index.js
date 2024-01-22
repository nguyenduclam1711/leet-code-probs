/**
 * @param {number[][]} intervals
 * @return {number}
 */
var eraseOverlapIntervals = function(intervals) {
  if (intervals.length < 2) {
    return 0;
  }
  let res = 0;
  intervals.sort((a, b) => a[0] - b[0]);
  let currEnd = intervals[0][1];

  for (let i = 1; i < intervals.length; i++) {
    if (intervals[i][0] >= currEnd) {
      currEnd = intervals[i][1];
    } else {
      // meaning overlap, must remove 1 interval
      res++;
      // take the interval that has earlier end => chance less likely to overlap the next one
      if (intervals[i][1] < currEnd) {
        currEnd = intervals[i][1];
      }
    }
  }
  return res;
};
