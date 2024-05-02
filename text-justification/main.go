package textjustification

func fullJustify(words []string, maxWidth int) []string {
	currStr := ""
	res := []string{}
	wordsStack := []string{}
	currOnlyWordsLen := 0

	for len(words) > 0 {
		for len(currStr) < maxWidth && len(words) > 0 {
			currWord := words[0]
			wordsStack = append(wordsStack, currWord)
			currOnlyWordsLen += len(currWord)
			if len(currStr) == 0 {
				currStr += currWord
			} else {
				currStr += " " + currWord
			}
			words = words[1:]
		}
		if len(currStr) > maxWidth {
			lastWord := wordsStack[len(wordsStack)-1]
			currOnlyWordsLen -= len(lastWord)
			wordsStack = wordsStack[:len(wordsStack)-1]
			words = append([]string{lastWord}, words...)
		}
		numberOfSpaces := len(wordsStack) - 1
		if len(words) > 0 && numberOfSpaces > 0 {
			// meaning this is not the last line and there are at least 2 words in words stack
			totalSpacesLength := maxWidth - currOnlyWordsLen
			evenSpaceLength := totalSpacesLength / numberOfSpaces
			leftSpaceLength := totalSpacesLength % numberOfSpaces

			// for storing spaces length for each word space to distribute all spaces evenly
			spacesLengths := []int{}

			for i := 0; i < numberOfSpaces; i++ {
				currSpaceLen := evenSpaceLength
				if leftSpaceLength > 0 {
					currSpaceLen += 1
					leftSpaceLength--
				}
				spacesLengths = append(spacesLengths, currSpaceLen)
			}

			// reset currStr for make the string to push to the result
			currStr = ""
			for i, j := 0, 0; i < len(wordsStack); i++ {
				currStr += wordsStack[i]
				if i < len(wordsStack)-1 {
					for k := 0; k < spacesLengths[j]; k++ {
						currStr += " "
					}
				}
				j++
			}
		} else {
			// this is the last line or there is only 1 word in the stack
			currStr = ""
			for i, word := range wordsStack {
				currStr += word
				if i < len(wordsStack)-1 {
					currStr += " "
				}
			}
			leftSpaces := maxWidth - len(currStr)
			for i := 0; i < leftSpaces; i++ {
				currStr += " "
			}
		}

		// append the result
		res = append(res, currStr)

		// reset everything after 1 loop
		currStr = ""
		currOnlyWordsLen = 0
		wordsStack = []string{}
	}
	return res
}
