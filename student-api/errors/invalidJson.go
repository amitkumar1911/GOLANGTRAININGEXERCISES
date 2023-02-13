package errors

import "fmt"

type InvalidJson struct {
	Param string
}

func (i InvalidJson) Error() string {
	return fmt.Sprintf("%v is invalid", i.Param)
}
