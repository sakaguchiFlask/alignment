package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sakaguchiFlask/alignment/tools"
)

func main() {
	var (
		inputFile    string
		displaywidth int
		gapPenalty   int
		scoreWeight  int
		textMode     bool
		outputFile   string
	)
	flag.StringVar(&inputFile, "i", "test.faa", "input multiple fasta format sequence file")
	flag.IntVar(&displaywidth, "l", 60, "Display line width (sequence length)")
	flag.IntVar(&gapPenalty, "g", 1, "Gap Penalty")
	flag.IntVar(&scoreWeight, "w", 1, "match score")
	flag.BoolVar(&textMode, "q", false, "Display the result on console")
	flag.StringVar(&outputFile, "o", "", "Output result text file")
	flag.Parse()

	fmt.Println("inputFile       : ", inputFile)
	fmt.Println("displaywidth    : ", displaywidth)
	fmt.Println("gapPenalty      : ", gapPenalty)
	fmt.Println("scoreWeight     : ", scoreWeight)
	fmt.Println("textMode        : ", textMode)
	fmt.Println("outputFile      : ", outputFile)
	seqs := tools.FileRead(inputFile)

	Penalty := gapPenalty
	matchscore := scoreWeight

	seq1 := strings.ReplaceAll(seqs[2], " ", "")
	seq2 := strings.ReplaceAll(seqs[4], " ", "")
	if len(seq1) < len(seq2) {
		tmp := seq1
		seq1 = seq2
		seq2 = tmp
	}

	var scores = make([]int, (len(seq1) - len(seq2)))

	for i := 0; i < (len(seq1) - len(seq2)); i++ {
		fragmentSeq1 := seq1[i : i+len(seq2)]

		_, _, _, score, _ := tools.Align(fragmentSeq1, seq2, Penalty, matchscore)
		scores[i] = score
	}

	maxScore, maxi := 0, 0

	for i, score := range scores {
		if score > maxScore {
			maxScore = score
			maxi = i
		}
	}

	fragmentSeq1 := seq1[maxi : maxi+len(seq2)]
	resultSeq1, resultSeq2, resultMatching, score, _ := tools.Align(fragmentSeq1, seq2, Penalty, matchscore)
	resultSeq := seq1[:maxi] + resultSeq1 + seq1[maxi+len(seq2):]
	if !textMode {
		err := tools.Display(resultSeq, resultSeq2, resultMatching, "seq1", "seq2", maxi, displaywidth)
		if err != nil {
			panic(err)
		}
	}
	if outputFile != "" {
		err := tools.WriteFile(outputFile, resultSeq, resultSeq2, resultMatching, "seq1", "seq2", maxi, displaywidth)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(maxi, score)
}
