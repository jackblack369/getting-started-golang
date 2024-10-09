package file

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestJoin(t *testing.T) {
	res := filepath.Join("a", "b", "c")
	fmt.Println(res)
}
