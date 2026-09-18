package tools

import "errors"

func Complementary(seq string) (string, error) {
	if seq == "" {
		var ErrorNoSeq = errors.New("seq is blank")
		return "", ErrorNoSeq
	}
	complement := ""
	for i := 0; i < len(seq); i++ {
		switch seq[i : i+1] {
		case "A", "a":
			complement = complement + "T"
		case "T", "t":
			complement = complement + "A"
		case "G", "g":
			complement = complement + "C"
		case "C", "c":
			complement = complement + "G"
		}
	}
	return complement, nil
}
