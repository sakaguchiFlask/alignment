package tools

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func countNucleotide(seq string) int {
	var basepair int
	for i := 0; i < len(seq); i++ {
		if seq[i:i+1] == "A" || seq[i:i+1] == "T" || seq[i:i+1] == "G" || seq[i:i+1] == "C" {
			basepair++
		}
	}
	return basepair
}

func Display(seq1, seq2, matching, name1, name2 string, offset, lineWidth int) error {
	if seq1 == "" || seq2 == "" || matching == "" {
		var ErrorNoSeq = errors.New("seq(s) is blank")
		return ErrorNoSeq
	}
	if lineWidth == 0 {
		lineWidth = 60
	}
	seqName1 := (name1 + strings.Repeat(" ", 10))[:6] + strings.Repeat(" ", 1)
	seqName2 := (name2 + strings.Repeat(" ", 10))[:6] + strings.Repeat(" ", 1)

	complementSeq2 := strings.Repeat(" ", offset) + seq2 + strings.Repeat(" ", (len(seq1)-offset-len(seq2)))
	complementMatching := strings.Repeat(" ", offset) + matching + strings.Repeat(" ", (len(seq1)-offset-len(matching)))

	var sequence1, matchingSequence, sequence2, scale1, scale2 string
	count1 := 1
	count2 := 1
	for i := 0; i*lineWidth < len(seq1); i++ {
		if (i+1)*lineWidth > len(seq1) {
			matchingSequence = complementMatching[(i * lineWidth):len(seq1)]
			sequence2 = complementSeq2[(i * lineWidth):len(seq1)]
			if countNucleotide(sequence2) != 0 {
				//scale2 = strconv.Itoa(count2)
				scale2 = strconv.Itoa(count2 + countNucleotide(sequence2))
				count2++
			} else {
				scale2 = ""
			}
			sequence1 = seq1[(i * lineWidth):len(seq1)]
			if (len(sequence1) - len(strconv.Itoa(count1+countNucleotide(sequence1))) - len(strconv.Itoa(count1))) < 0 {
				scale1 = strconv.Itoa(count1) + " " + strconv.Itoa(count1+countNucleotide(sequence1)-1)
			} else {
				scale1 = strconv.Itoa(count1) + strings.Repeat(" ", len(sequence1)-len(strconv.Itoa(count1+countNucleotide(sequence1)))-len(strconv.Itoa(count1))) + strconv.Itoa(count1+countNucleotide(sequence1)-1)
			}
			count1 = count1 + countNucleotide(sequence1)
		} else {
			matchingSequence = complementMatching[(i * lineWidth):((i + 1) * lineWidth)]
			sequence2 = complementSeq2[(i * lineWidth):((i + 1) * lineWidth)]
			if countNucleotide(sequence2) != 0 {
				//scale2 = strconv.Itoa(count2)
				var indent int
				for indent = 0; sequence2[indent:indent+1] == " "; indent++ {
				}
				var lengthSeq int
				for lengthSeq = indent + 1; lengthSeq < lineWidth && sequence2[lengthSeq:lengthSeq+1] != " "; lengthSeq++ {
				}
				scale2 = strings.Repeat(" ", indent) + strconv.Itoa(count2) + strings.Repeat(" ", lengthSeq-indent-len(strconv.Itoa(count2))-len(strconv.Itoa(count2+countNucleotide(sequence2)-1))) + strconv.Itoa(count2+countNucleotide(sequence2)-1)
				count2 = count2 + countNucleotide(sequence2)
			} else {
				scale2 = ""
			}
			sequence1 = seq1[(i * lineWidth):((i + 1) * lineWidth)]
			scale1 = strconv.Itoa(count1) + strings.Repeat(" ", len(sequence1)-len(strconv.Itoa(count1+countNucleotide(sequence1)))-len(strconv.Itoa(count1))) + strconv.Itoa(count1+countNucleotide(sequence1)-1)
			count1 = count1 + countNucleotide(sequence1)
		}
		fmt.Println(strings.Repeat(" ", len(seqName1)+1) + scale1)
		fmt.Println(seqName1, sequence1)
		fmt.Println(strings.Repeat(" ", len(seqName1)), matchingSequence)
		fmt.Println(seqName2, sequence2)
		fmt.Println(strings.Repeat(" ", len(seqName1)+1) + scale2)
		fmt.Println()

	}
	return nil
}

/*
AGAGTTTGAT CCTGGCTCAG GATGAACGCT AGCGATAGGC TTAACACATG CAAGTCGAGG
*/
