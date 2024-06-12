/**
 * @param {number[][]} intervals
 * @param {number[]} newInterval
 * @return {number[][]}
 */
var insert = function(intervals, newInterval) {
  if (intervals.length < 1) {
    return [newInterval];
  }
  const res = [];
  let lo = newInterval[0];
  let hi = newInterval[1];

  for (let i = 0; i < intervals.length; i++) {
    const interval = intervals[i];

    if (interval[0] <= hi && interval[1] >= lo) {
      if (interval[0] < lo) {
        lo = interval[0];
      }
      if (interval[1] > hi) {
        hi = interval[1];
      }
    } else {
      if (interval[0] > hi) {
        res.push([lo, hi]);
        lo = interval[0];
        hi = interval[1];
      } else if (interval[1] < lo) {
        res.push(interval);
      }
    }
  }
  res.push([lo, hi]);
  return res;
};
