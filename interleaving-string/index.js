/**
 * @param {string} s1
 * @param {string} s2
 * @param {string} s3
 * @return {boolean}
 */
var isInterleave = function(s1, s2, s3) {
  if (s1.length + s2.length !== s3.length) {
    return false;
  }
  if (s3.length === 0) {
    return true;
  }
  let q = [[s1, s2]];
  let compareStr = "";
  for (let i = 0; i < s3.length; i++) {
    if (q.length === 0) {
      break;
    }
    const c = s3[i];
    let newQ = [];
    const existedS1 = {};
    const existedS2 = {};

    for (let j = 0; j < q.length; j++) {
      const currS1 = q[j][0];
      const currS2 = q[j][1];

      if (currS1.length === 0 && currS2.length === 0) {
        continue;
      }
      if (currS1.length > 0 && currS1[0] === c) {
        const nextS1 = currS1.slice(1);
        if (!existedS2[nextS1] || !existedS1[currS2]) {
          newQ.push([nextS1, currS2]);
          existedS2[nextS1] = true;
          existedS1[currS2] = true;
        }
      }
      if (currS2.length > 0 && currS2[0] === c) {
        const nextS2 = currS2.slice(1);
        if (!existedS2[nextS2] || !existedS1[currS1]) {
          newQ.push([currS1, nextS2]);
          existedS2[nextS2] = true;
          existedS1[currS1] = true;
        }
      }
    }
    if (newQ.length > 0) {
      compareStr += c;
    }
    q = newQ;
  }
  return compareStr === s3;
};
