/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let area = 0;
  let l = 0;
  let r = height.length - 1;

  while (l < r) {
    const width = r - l;
    if (height[l] < height[r]) {
      area = Math.max(area, width * height[l]);
      l++;
    } else {
      area = Math.max(area, width * height[r]);
      r--;
    }
  }
  return area;
};

console.log(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]));
