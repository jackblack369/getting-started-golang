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

func TestCountVowel(t *testing.T) {
	s := "hello world"
	countVowel(s)
}

func countVowel(s string) {
	vowels := "aeiou"
	count := 0
	for _, c := range s {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}
	fmt.Println(count)
}

func TestDANStrant(t *testing.T) {
	dna := "ATTGC"
	res := DNAStrand(dna)
	fmt.Println("res:", res)
}

func DNAStrand(dna string) string {
	// A -> T, T -> A, C -> G, G -> C
	// dna = "ATTGC" => "TAACG"
	// dna = "GTAT" => "CATA"
	fmt.Println("dna:", dna)
	res := strings.Map(func(r rune) rune {
		switch r {
		case 'A':
			return 'T'
		case 'T':
			return 'A'
		case 'C':
			return 'G'
		case 'G':
			return 'C'
		}
		return r
	}, dna)
	fmt.Println("map DNA:", res)
	return res
}
