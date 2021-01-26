package longest

func lengthOfLongestSubstring(s string) int {
    windowStart := 0
    windowEnd := 0
    max := 0
    
    seenMap := make(map[string]int, 0)
    
    for {
        if windowEnd < len(s){
            currentChar := string(s[windowEnd])
            _, exists := seenMap[currentChar]
            if !exists{
                seenMap[currentChar] = windowEnd
                if max < len(seenMap){
                    max = len(seenMap)
                }
                windowEnd = windowEnd+1
            }else{
                delete(seenMap, string(s[windowStart]))
                windowStart = windowStart+1
            }
        }else {
		break
	}
    }
    return max
}
