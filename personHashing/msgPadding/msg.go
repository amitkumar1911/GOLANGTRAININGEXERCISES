package msgPadding

import "fmt"

func MakeMsg(c1 chan []string, c2 chan string) {

	for i := range c1 {
		msg := fmt.Sprintf("%s%s%s%s", i[0], i[1], i[2], i[3])
		msgSignature := fmt.Sprintf("%-*s", 100, msg)
		c2 <- msgSignature
	}
	close(c2)
}
