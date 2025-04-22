package num

import (
	"strconv"
	"testing"
)

func InitNum(name string, num int64) (string, int64) {
	return name, num
}

func TestInitNum(t *testing.T) {
	name, num := InitNum("dongwei", 0)
	t.Log("name is ", name)
	t.Log("num is ", num)
}

func TestConvert(t *testing.T) {
	capacity := 1048576 // Example capacity value

	// Convert capacity to string
	capacityStr := strconv.Itoa(capacity)
	t.Log("Capacity value is: ", capacityStr)
	// init int64 num
	var num int64
	num = 100209397489274927
	capacityStr2 := strconv.FormatInt(num, 10)
	t.Log("Capacity value is: ", capacityStr2)

	// Convert capacity string to int
	numText := "1048576" // Example capacity string
	res, err := strconv.ParseInt(numText, 10, 64)
	if err != nil {
		t.Log("Error converting capacity to int: ", err)
	}
	t.Log("res value is: ", res)
}

func Test1(t *testing.T) {
	t.Log("Test1")
	const (
		// checker
		CHECK_TOPOLOGY int = iota
		CHECK_SSH_CONNECT
	)
	t.Log(CHECK_TOPOLOGY)
	t.Log(CHECK_SSH_CONNECT)
}
