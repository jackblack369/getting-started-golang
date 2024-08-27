package stage1

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return // If request is an empty string, the function sets err to a new error and then uses a naked return to return the current values of response (which is an empty string) and err.
	}
	response = fmt.Sprintf("echo: %s", request)
	return // same usage as above (naked return)
}

func main() {
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}
