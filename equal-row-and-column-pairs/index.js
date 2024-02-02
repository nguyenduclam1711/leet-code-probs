/**
 * @param {number[][]} grid
 * @return {number}
 */
var equalPairs = function(grid) {
  const mapRow = {};
  const colStrs = Array(grid.length);

  for (const row of grid) {
    let str = "";
    for (let i = 0; i < row.length; i++) {
      const addStr = `${row[i]},`;
      str += addStr;
      colStrs[i] = (colStrs[i] || "") + addStr;
    }
    mapRow[str] = (mapRow[str] || 0) + 1;
  }
  let res = 0;
  for (const str of colStrs) {
    if (mapRow[str] > 0) {
      res += mapRow[str];
    }
  }
  return res;
};
