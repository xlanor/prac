package roman

func romanToInt(s string) int {
    lookup := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    res := lookup[string(s[len(s)-1])]
    for i := len(s)-1; i > 0; i-- {
        if lookup[string(s[i-1])] < lookup[string(s[i])] {
            res = res - lookup[string(s[i-1])]
        }else {
            res = res + lookup[string(s[i-1])]
        }
    }
    return res
    
}
