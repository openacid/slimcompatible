package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	seg := strings.Repeat("a", 8*1024)
	ss := []string{
		seg + "x" + seg + "a",
		seg + "x" + seg + "b",
		seg + "x" + seg + "bc",

		seg + "y" + seg + "a",
		seg + "y" + seg + "ab",
		seg + "y" + seg + "d",

		seg + "z" + seg + "u",
		seg + "z" + seg + "v",
		seg + "z" + seg + "w",
		seg + "z" + seg + "x",
	}

	// ss := RandVarStrings(5, 65536, nil)

	tmpl := `package testkeys
var (
	testKeys5vl64k = %#v
 )`
	fmt.Printf(tmpl, ss)
}

var alphas = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var runes = []rune("~!@#$%^&*()_+`-=[]{};:<>?,./abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandVarStrings(n, maxLen int, from []rune) []string {

	if from == nil {
		from = alphas
	}

	rlen := len(from)

	mp := make(map[string]bool, 0)

	for i := 0; i < n; i++ {
		l := rand.Intn(maxLen + 1)
		b := make([]rune, l)
		for j := 0; j < l; j++ {
			k := rand.Intn(rlen)

			b[j] = from[k]
		}
		s := string(b)
		if _, ok := mp[s]; ok {
			i--
		} else {
			mp[s] = true
		}
	}

	rst := make([]string, 0, n)
	for k := range mp {
		rst = append(rst, k)
	}

	sort.Strings(rst)
	return rst
}
