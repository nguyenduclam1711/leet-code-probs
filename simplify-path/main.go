package simplifypath

func simplifyPath(path string) string {
	stack := []rune{}
	// init path have the last / for preventing edge case have dots in in the path last position
	path += "/"

	for _, r := range path {
		if r != '/' {
			stack = append(stack, r)
			continue
		}

		// init root folder
		if len(stack) == 0 {
			stack = append(stack, r)
			continue
		}

		// case multiple slashes together
		if stack[len(stack)-1] == '/' {
			continue
		}

		folderNameLen, numberOfDots := 0, 0
		for i := len(stack) - 1; i >= 0; i-- {
			currStackRune := stack[i]
			if currStackRune == '/' {
				break
			}
			if currStackRune == '.' {
				numberOfDots++
			}
			folderNameLen++
		}

		if folderNameLen == numberOfDots {
			// case ./
			if numberOfDots == 1 {
				stack = stack[:len(stack)-1]
				continue
			}

			// case ../
			if numberOfDots == 2 {
				// remove folder
				numberOfSlash := 0

				for len(stack) > 0 {
					if stack[len(stack)-1] == '/' {
						numberOfSlash++
					}
					if numberOfSlash == 2 {
						break
					}
					stack = stack[:len(stack)-1]
				}

				// stack must have at least root folder
				if len(stack) == 0 {
					stack = append(stack, '/')
				}
				continue
			}

			// valid folder name here
			stack = append(stack, r)
			continue
		}
		// valid folder name here
		stack = append(stack, r)
	}
	// remove last slash
	if stack[len(stack)-1] == '/' && len(stack) > 1 {
		stack = stack[:len(stack)-1]
	}
	return string(stack)
}
