/**
 * @param {character[][]} maze
 * @param {number[]} entrance
 * @return {number}
 */
var nearestExit = function (maze, entrance) {
  const steps = Array(maze.length);
  for (let i = 0; i < steps.length; i++) {
    steps[i] = Array(maze[0].length).fill(0);
  }

  const queue = [entrance];
  let res = Infinity;

  const addToQueue = function (x, y, prevStep) {
    if (x < 0 || x === maze.length || y < 0 || y === maze[0].length) {
      return;
    }
    if (steps[x][y] > 0 || (x === entrance[0] && y === entrance[1])) {
      return;
    }
    if (maze[x][y] === "+") {
      return;
    }
    const currStep = prevStep + 1;
    steps[x][y] = currStep;
    if (
      x === 0 ||
      x === maze.length - 1 ||
      y === 0 ||
      y === maze[0].length - 1
    ) {
      if (currStep < res) {
        res = currStep;
      }
      return;
    }
    queue.push([x, y]);
  };

  while (queue.length > 0) {
    const currQueueLen = queue.length;
    for (const pos of queue) {
      const x = pos[0];
      const y = pos[1];
      // go up
      addToQueue(x - 1, y, steps[x][y]);
      // go right
      addToQueue(x, y + 1, steps[x][y]);
      // go down
      addToQueue(x + 1, y, steps[x][y]);
      // go left
      addToQueue(x, y - 1, steps[x][y]);
    }
    queue.splice(0, currQueueLen);
  }
  if (res === Infinity) {
    return -1;
  }
  return res;
};
