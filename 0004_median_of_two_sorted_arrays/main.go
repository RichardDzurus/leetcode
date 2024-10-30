package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	len1 := len(nums1)
	len2 := len(nums2)
	isOdd := (len1+len2)%2 == 1
	low := 0
	high := len(nums1)
	var mid1 int
	var mid2 int
	for low <= high {
		mid1 = low + (high-low)/2
		mid2 = (len1+len2+1)/2 - mid1

		maxLeft1 := -1000001
		if mid1 > 0 {
			maxLeft1 = nums1[mid1-1]
		}
		minRight1 := 1000001
		if mid1 < len1 {
			minRight1 = nums1[mid1]
		}
		maxLeft2 := -1000001
		if mid2 > 0 {
			maxLeft2 = nums2[mid2-1]
		}
		minRight2 := 1000001
		if mid2 < len2 {
			minRight2 = nums2[mid2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if isOdd {
				return float64(max(maxLeft1, maxLeft2))
			} else {
				return (float64(max(maxLeft1, maxLeft2)) + float64(min(minRight1, minRight2))) / 2
			}
		} else if maxLeft1 > minRight2 {
			high = mid1 - 1
		} else {
			low = mid1 + 1
		}
	}

	return 0.0
}

func main() {
	x := findMedianSortedArrays([]int{100000}, []int{100001})
	println(x)
}
