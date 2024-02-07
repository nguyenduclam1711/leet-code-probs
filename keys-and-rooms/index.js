/**
 * @param {number[][]} rooms
 * @return {boolean}
 */
var canVisitAllRooms = function(rooms) {
  const visited = Array(rooms.length).fill(false);
  let totalVisited = 0;
  const dfs = function(roomNum) {
    if (visited[roomNum]) {
      return false;
    }
    visited[roomNum] = true;
    totalVisited++;
    if (totalVisited == rooms.length) {
      return true;
    }
    for (const key of rooms[roomNum]) {
      if (dfs(key)) {
        return true;
      }
    }
    return false;
  };
  return dfs(0);
};
