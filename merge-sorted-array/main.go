package mergesortedarray

import "sort"

// Time: O(nlog(n))
// Space: O(1)
func Merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	sort.Ints(nums1)
}

// Time: O(m + n)
// Space: O(m + n)
func Merge2(nums1 []int, m int, nums2 []int, n int) {
	s := []int{}
	i, j := 0, 0

	for i < m || j < n {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				s = append(s, nums1[i])
				i++
			} else {
				s = append(s, nums2[j])
				j++
			}
		} else if i < m && j >= n {
			s = append(s, nums1[i])
			i++
		} else if j < n && i >= m {
			s = append(s, nums2[j])
			j++
		}
	}
	copy(nums1, s)
}

// Time: O(m + n)
// Space: O(1)
func Merge3(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	curIndex := m + n - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[curIndex] = nums1[p1]
			p1--
		} else {
			nums1[curIndex] = nums2[p2]
			p2--
		}
		curIndex--
	}
	for p1 >= 0 {
		nums1[curIndex] = nums1[p1]
		p1--
		curIndex--
	}
	for p2 >= 0 {
		nums1[curIndex] = nums2[p2]
		p2--
		curIndex--
	}
}
