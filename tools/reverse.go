package tools

import (
	"errors"
	"fmt"
)

func Reverse(seq string) (string, error) {
	if seq == "" {
		var ErrorNoSeq = errors.New("seq is blank")
		return "", ErrorNoSeq
	}
	reverse := ""
	for i := 0; i < len(seq); i++ {
		reverse = reverse + seq[len(seq)-1-i:len(seq)-i]
	}
	return reverse, nil
}

func main() {
	var seq = "ATGC"
	reverse, _ := Reverse(seq)
	fmt.Println("original sequence : ", seq)
	fmt.Println("reverse sequence  : ", reverse)
}
