package tools

import (
	"errors"
)

func Align(seq1, seq2 string, Penalty, weightScore int) (alignedSeq1, alignedSeq2, matchSymbol string, score int, err error) {
	if seq1 == "" || seq2 == "" {
		var ErrorNoSeq = errors.New("seq(s) is blank")
		return "", "", "", 0, ErrorNoSeq
	}
	d := Penalty // evaluation score
	matchscore := weightScore
	unmatchscore := -Penalty

	imax := len(seq1)
	jmax := len(seq2)
	var matrix = make([][]int, jmax)
	for j := 0; j < imax; j++ {
		matrix[j] = make([]int, imax)
	}
	for i := 0; i < imax; i++ {
		for j := 0; j < jmax; j++ {
			if seq1[i:i+1] == seq2[j:j+1] {
				matrix[i][j] = matchscore
			} else {
				matrix[i][j] = unmatchscore
			}
		}
	}

	imax++
	jmax++
	var Dmatrix = make([][]int, jmax)
	for j := 0; j < imax; j++ {
		Dmatrix[j] = make([]int, imax)
	}
	for i := 0; i < imax; i++ {
		Dmatrix[i][0] = (0 - d) * i
	}
	for j := 0; j < jmax; j++ {
		Dmatrix[0][j] = (0 - d) * j
	}
	for i := 1; i < imax; i++ {
		for j := 1; j < jmax; j++ {
			Dmatrix[i][j] = max((Dmatrix[i][j-1] - d), (Dmatrix[i-1][j] - d), (Dmatrix[i-1][j-1] + matrix[i-1][j-1]))
		}
	}

	var i, j = imax - 1, jmax - 1
	align1, align2, matching := "", "", ""

	for i > 0 && j > 0 {
		if max(Dmatrix[i-1][j-1], Dmatrix[i][j-1], Dmatrix[i-1][j]) == Dmatrix[i-1][j-1] {
			align1 = align1 + seq1[i-1:i]
			align2 = align2 + seq2[j-1:j]
			if seq1[i-1:i] == seq2[j-1:j] {
				matching = matching + "*"
			} else {
				matching = matching + " "
			}
			score = score + Dmatrix[i-1][j-1]
			i--
			j--
		} else {
			if max(Dmatrix[i-1][j-1], Dmatrix[i][j-1], Dmatrix[i-1][j]) == Dmatrix[i][j-1] {
				align1 = align1 + "-"
				align2 = align2 + seq2[j-1:j]
				matching = matching + " "
				score = score + Dmatrix[i][j-1]
				j--
			} else {
				align1 = align1 + seq1[i-1:i]
				align2 = align2 + "-"
				matching = matching + " "
				score = score + Dmatrix[i-1][j]
				i--
			}
		}
	}

	alignedSeq1, _ = Reverse(align1)
	alignedSeq2, _ = Reverse(align2)
	matchSymbol, _ = Reverse(matching)
	return
}
