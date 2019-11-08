package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)

	if length%2 == 0 {
		return getStat(length/2-1, true, nums1, nums2)
	}

	return getStat(length/2, false, nums1, nums2)
}

func getStat(index int, odd bool, nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	var el1, i, j, k int
	for {
		if (i < len1) && (j >= len2 || nums1[i] < nums2[j]) {
			if odd {
				if k == index {
					el1 = nums1[i]
				} else if k == index+1 {
					return float64(el1+nums1[i]) / 2
				}
			} else if k == index {
				return float64(nums1[i])
			}
			i++
		} else {
			if odd {
				if k == index {
					el1 = nums2[j]
				} else if k == index+1 {
					return float64(el1+nums2[j]) / 2
				}
			} else if k == index {
				return float64(nums2[j])
			}
			j++
		}
		k++
	}
}

func main() {}
