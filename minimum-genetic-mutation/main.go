package minimumgeneticmutation

func canGeneChange(gene1, gene2 string) bool {
	changedTimes := 0
	for i := range gene1 {
		if gene1[i] != gene2[i] {
			changedTimes++
		}
		if changedTimes > 1 {
			return false
		}
	}
	return true
}

func minMutation(startGene string, endGene string, bank []string) int {
	visited := map[string]bool{}
	q := []string{startGene}
	minStep := -1
	canBeMutation := false

	for len(q) > 0 && !canBeMutation {
		newQ := []string{}
		minStep++
		for _, gene := range q {
			visited[gene] = true
			if gene == endGene {
				canBeMutation = true
				break
			}
			for _, bankGene := range bank {
				if visited[bankGene] {
					continue
				}
				if canGeneChange(gene, bankGene) {
					newQ = append(newQ, bankGene)
				}
			}
		}
		q = newQ
	}
	if canBeMutation {
		return minStep
	}
	return -1
}
