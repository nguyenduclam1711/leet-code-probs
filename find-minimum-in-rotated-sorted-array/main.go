package findminimuminrotatedsortedarray

func FindMin(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mdl := (l + r) / 2

		if nums[mdl] > nums[r] {
			l = mdl + 1
		} else {
			r = mdl
		}
	}

	return nums[l]
}
