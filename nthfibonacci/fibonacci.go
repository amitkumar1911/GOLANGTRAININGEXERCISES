package nthfibonacci

type Slc struct {
	ans []int
}

func (s *Slc) Fibonacci(temp int) []int {

	s.ans = []int{0, 1}

	if temp > 2 {

		if len(s.ans) > temp {
			return s.ans[:temp]
		}

		length := temp - len(s.ans)

		for i := 1; i <= length; i++ {
			s.ans = append(s.ans, s.ans[len(s.ans)-1]+s.ans[len(s.ans)-2])
		}

		return s.ans

	} else if temp <= 2 && temp >= 0 {

		return s.ans[:temp]
	} else {
		return []int{}
	}
}

// func Fibonacci() func(int) []int {

// 	// temp := n

// 	ans := []int{0, 1}

// 	return func(temp int) []int {

// 		if temp > 2 {

// 			if len(ans) > temp {
// 				return ans[:temp]
// 			}

// 			length := temp - len(ans)

// 			for i := 1; i <= length; i++ {
// 				ans = append(ans, ans[len(ans)-1]+ans[len(ans)-2])
// 			}

// 			return ans

// 		} else if temp <= 2 && temp >= 0 {

// 			return ans[:temp]
// 		} else {
// 			return []int{}
// 		}

// 	}

// }
