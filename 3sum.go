package threesum

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) >= 3 {
		sort.Ints(nums)
        //fmt.Printf("Sorted: %v\n", nums)
		for i := 0; i <= len(nums)-2; i++ {
            //fmt.Printf("Current i: %d\n", i)
			// check for duplicates with the one before.
            if i != 0 && (nums[i] == nums[i-1]) {
                //fmt.Printf ("Index %d and %d have the same value (%d:%d), continuing\n", i, i-1, nums[i], nums[i-1]) 
				// dont check previous if you are the first.
				continue
			}

			// a + b = -c
			c := 0 - nums[i]
			aPtr := i + 1
			bPtr := len(nums) - 1
			//fmt.Printf("aPtr %d bPtr %d\n", aPtr, bPtr)
            //fmt.Println("======")
			// I wish I did this in c++ instead.
			for {
               // fmt.Printf("Comparing aPtr: %d, bPtr: %d\n", aPtr, bPtr)
                
				if aPtr >= bPtr {
					break
				}
                //fmt.Printf("Comparing %d == %d\n", nums[aPtr]+nums[bPtr], c)
				// 2 sum.
				if nums[aPtr]+nums[bPtr] == c {
                    //fmt.Printf("Index %d and %d summation matched\n", aPtr, bPtr)
					res = append(res, []int{nums[i], nums[aPtr], nums[bPtr]})
					for {
						if aPtr >= bPtr {
							break
						}
						if nums[aPtr] == nums[aPtr+1] {
							aPtr = aPtr + 1
						} else {
							break
						}
					}
					for {
						if aPtr >= bPtr {
							break
						}
						if nums[bPtr] == nums[bPtr-1] {
							bPtr = bPtr - 1
						} else {
							break
						}
					}
					aPtr = aPtr + 1
					bPtr = bPtr - 1
				} else if nums[aPtr]+nums[bPtr] < c {
					aPtr = aPtr + 1
				} else {
					bPtr = bPtr - 1
				}
			}
		}

	}
	return res
}
