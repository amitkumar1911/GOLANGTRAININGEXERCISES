package stringutil

func CommonSubstring(s1 string, s2 string) string {

	var ans string = ""
	if len(s1) > len(s2) {
		for i := 0; i < len(s1); i++ {

			var temp string = ""

			for j := i; j < len(s1); j++ {

				temp1 := string(s1[j])

				temp = temp + temp1

				// fmt.Println(temp)

				if temp == s2 {

					ans = temp

				}

			}

		}

	} else {

		var temp string = ""
		for i := 0; i < len(s2); i++ {

			for j := i; j < len(s2); j++ {

				temp1 := string(s2[j])

				temp = temp + temp1

				if temp == s1 {

					ans = temp

				}

			}

		}
	}
	return ans
}

func main() {

	CommonSubstring("amitkumar", "kumar")

}
