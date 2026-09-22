package tools

import (
	"bufio"
	"os"
	"strings"
)

func chop(s string) string {
	s = strings.TrimRight(s, "\n")
	if strings.HasSuffix(s, "\r") {
		s = strings.TrimRight(s, "\r")
	}
	return s
}

func FileRead(fileName string) []string {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	var seqs = make([]string, 6)
	scanner := bufio.NewScanner(fp)
	var readLine, sequence string
	counter := 0
	for scanner.Scan() {
		if counter > 4 {
			break
		}
		readLine = scanner.Text()
		if readLine[0:1] == ">" {
			seqs[counter] = sequence
			sequence = ""
			counter++
			seqs[counter] = readLine
			counter++
		} else {
			sequence = sequence + chop(readLine)
		}
	}
	return seqs[0:5]
}
