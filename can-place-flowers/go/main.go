package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		bed := flowerbed[i]

		if bed == 1 {
			continue
		}

		isValidPos := true
		nextBedPos := i + 1
		prevBedPos := i - 1

		if i == 0 {
			if nextBedPos < len(flowerbed) && flowerbed[nextBedPos] == 1 {
				isValidPos = false
			}
		} else if i == len(flowerbed)-1 {
			if prevBedPos >= 0 && flowerbed[prevBedPos] == 1 {
				isValidPos = false
			}
		} else {
			if flowerbed[prevBedPos] == 1 || flowerbed[nextBedPos] == 1 {
				isValidPos = false
			}
		}
		if isValidPos {
			flowerbed[i] = 1
			n--
		}
	}
	return n == 0
}
