package twosum

func twoSum(nums []int, target int) []int {
    hm := make(map[int]int)
    for i, v := range nums {
        stored, ve := hm[v]
        if ve {
            return []int {stored, i}
        }
        
        sub := target - v
        _, e := hm[sub]
        if !e {
            hm[sub] = i
        }
    }
    // wont ever hit since guaranteed an answer.
    return []int{}
}
