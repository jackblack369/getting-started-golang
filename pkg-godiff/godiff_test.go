package pkg_godiff

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"testing"
)

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
)

func Test1(t *testing.T) {
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)

	fmt.Println(dmp.DiffPrettyText(diffs))
}

func Test2(t *testing.T) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, false)
	diffs = dmp.DiffCleanupSemantic(diffs)

	fmt.Println(dmp.DiffPrettyText(diffs))
}
