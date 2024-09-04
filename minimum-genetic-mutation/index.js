function canGeneChange(gene1, gene2) {
  let changedTimes = 0;
  for (let i = 0; i < gene1.length; i++) {
    if (gene1[i] !== gene2[i]) {
      changedTimes++;
    }
    if (changedTimes > 1) {
      return false;
    }
  }
  return true;
}

/**
 * @param {string} startGene
 * @param {string} endGene
 * @param {string[]} bank
 * @return {number}
 */
var minMutation = function(startGene, endGene, bank) {
  const visited = {};
  let q = [startGene];
  let minStep = -1;
  let canBeMutation = false;

  while (q.length > 0 && !canBeMutation) {
    const newQ = [];
    minStep++;
    for (const gene of q) {
      visited[gene] = true;
      if (gene === endGene) {
        canBeMutation = true;
        break;
      }
      for (const bankGene of bank) {
        if (visited[bankGene]) {
          continue;
        }
        if (canGeneChange(gene, bankGene)) {
          newQ.push(bankGene);
        }
      }
    }
    q = newQ;
  }
  if (canBeMutation) {
    return minStep;
  }
  return -1;
};
