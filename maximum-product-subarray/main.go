package main

import "fmt"

type object struct {
	min int
	max int
}

func maxProduct(nums []int) int {
	max := 0
	prefixArr := make([]object, len(nums))

	for i, num := range nums {
		if num == 0 {
			if num > max {
				max = num
			}
			prefixArr[i] = object{
				min: 1,
				max: 1,
			}
			continue
		}
		if i == 0 {
			max = num
			prefixArr[i] = object{
				min: num,
				max: num,
			}
			continue
		}
		x := prefixArr[i-1].max * num
		y := prefixArr[i-1].min * num

		if x > y {
			if num > x {
				prefixArr[i] = object{
					min: y,
					max: num,
				}
			} else if y < num {
				prefixArr[i] = object{
					min: y,
					max: x,
				}
			} else {
				prefixArr[i] = object{
					min: num,
					max: x,
				}
			}
		} else {
			if num > y {
				prefixArr[i] = object{
					min: x,
					max: num,
				}
			} else if x < num {
				prefixArr[i] = object{
					min: x,
					max: y,
				}
			} else {
				prefixArr[i] = object{
					min: num,
					max: y,
				}
			}
		}
		if prefixArr[i].max > max {
			max = prefixArr[i].max
		}
	}
	return max
}

func main() {
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
}
