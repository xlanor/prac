package multiply

func multiply(num1 string, num2 string) string {
   x := make([]int ,len(num1)+len(num2))
	
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
		
			val := int(num1[i] - '0') * int(num2[j] - '0')
			val = val+x[i+j+1]
			x[i + j + 1] = val %10
			x[i+j] +=  val / 10
		}
	}
	
	idx := 0
	for _, v := range x {
		if v == 0 {
			idx ++
		}else{
			break
		}
	}
    if idx == len(num1)+len(num2){
        return "0"
    }
	// build.
	x = x[idx:]
    r := ""
    for _, v := range x {
        r += strconv.Itoa(v)
    }
    return r
}
