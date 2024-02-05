/**
 * @param {string} senate
 * @return {string}
 */
var predictPartyVictory = function (senate) {
  const stack = senate.split("");
  let i = 0;
  let allSameParty = false;

  while (!allSameParty) {
    allSameParty = true;
    if (i >= stack.length) {
      i = 0;
    }
    let j = i + 1;
    for (let loop = 0; loop < senate.length; loop++) {
      if (j >= stack.length) {
        j = 0;
      }
      if (stack[i] !== stack[j]) {
        allSameParty = false;
        stack.splice(j, 1);
        break;
      }
      j++;
    }
    i++;
  }
  if (stack[0] === "D") {
    return "Dire";
  }
  return "Radiant";
};
