/**
 * @param {string[]} details
 * @return {number}
 */
var countSeniors = function(details) {
  let res = 0;
  for (let i = 0; i < details.length; i++) {
    const age = details[i].slice(11, 13);
    if (Number(age) > 60) {
      res++;
    }
  }
  return res;
};
