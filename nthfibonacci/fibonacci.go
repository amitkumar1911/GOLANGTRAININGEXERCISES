package nthfibonacci

func Fibonacci() func(int) []int {

	// temp := n

	ans := []int{0, 1}

	return func(temp int) []int {

		if temp > 2 {

			if len(ans) > temp {
				return ans[:temp]
			}

			length := temp - len(ans)

			for i := 1; i <= length; i++ {
				ans = append(ans, ans[len(ans)-1]+ans[len(ans)-2])
			}

			return ans

		} else if temp <= 2 && temp >= 0 {

			return ans[:temp]
		} else {
			return []int{}
		}

	}

}
