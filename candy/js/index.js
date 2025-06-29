// using greedy to solve this
// find relative smallest position and then set the candy to that position to 1 and then from that position, distribute the rest of candies
/**
 * @param {number[]} ratings
 * @return {number}
 */
var candy = function(ratings) {
  if (ratings.length === 1) {
    return 1;
  }
  const distributions = Array(ratings.length).fill(0);
  const initPos = [];

  for (let i = 0; i < ratings.length; i++) {
    if (i === 0) {
      if (ratings[i + 1] >= ratings[i]) {
        distributions[i] = 1;
        initPos.push(i);
      }
    } else if (i === ratings.length - 1) {
      if (ratings[i - 1] >= ratings[i]) {
        distributions[i] = 1;
        initPos.push(i);
      }
    } else {
      if (ratings[i - 1] >= ratings[i] && ratings[i + 1] >= ratings[i]) {
        distributions[i] = 1;
        initPos.push(i);
      }
    }
  }

  for (const pos of initPos) {
    // moving backward
    for (let i = pos - 1; i >= 0; i--) {
      if (ratings[i] <= ratings[i + 1]) {
        break;
      }
      distributions[i] = Math.max(distributions[i], distributions[i + 1] + 1);
    }
    // moving forward
    for (let i = pos + 1; i < ratings.length; i++) {
      if (ratings[i] <= ratings[i - 1]) {
        break;
      }
      distributions[i] = Math.max(distributions[i], distributions[i - 1] + 1);
    }
  }
  let res = 0;
  for (const d of distributions) {
    res += d;
  }
  return res;
};
