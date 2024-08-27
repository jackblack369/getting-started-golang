package string

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestOne(t *testing.T) {
	url := "sqlite:///Users/dongwei/.curveadm/data/curveadm.db"
	dataSourceName := strings.TrimPrefix(url, "sqlite://")
	println("dataSourceName is ", dataSourceName)
}

// FirstCharToCamelCase converts the first character of a string to uppercase
func TestFirstCharToCamelCase(t *testing.T) {
	s := "hello-world"

	s = strings.ReplaceAll(s, "_", "-")
	parts := strings.Split(s, "-")

	res_words := []string{}
	for i, part := range parts {
		if i == 0 {
			if strings.ToUpper(part[0:1]) == part[0:1] {
				res_words = append(res_words, strings.ToUpper(part[0:1])+strings.ToLower(part[1:]))
			} else {
				res_words = append(res_words, strings.ToLower(part))
			}
		} else {
			res_words = append(res_words, strings.ToUpper(part[0:1])+strings.ToLower(part[1:]))
		}
	}
	fmt.Println(res_words)
	res := strings.Join(res_words, "")
	fmt.Println(res)
}

func TestToCamelCase(t *testing.T) {
	s := "Hello-world-Brook"
	words := regexp.MustCompile("-|_").Split(s, -1)

	for i, w := range words[1:] {
		words[i+1] = strings.Title(w)
	}

	res := strings.Join(words, "")
	fmt.Println(res)
}

func TestTitle(t *testing.T) {
	s := "hello-world"
	res := strings.Title(s)
	fmt.Println(res)
}
