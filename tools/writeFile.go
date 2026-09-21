package tools

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func WriteFile(fileName, seq1, seq2, matching, name1, name2 string, offset, lineWidth int) error {
	if seq1 == "" || seq2 == "" {
		var ErrorNoSeq = errors.New("seq(s) is blank")
		return ErrorNoSeq
	}
	fp, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	writer := bufio.NewWriter(fp)
	defer writer.Flush()

	// _, err := writer.WriteString(line)

	if lineWidth == 0 {
		lineWidth = 60
	}
	seqName1 := (name1 + strings.Repeat(" ", 10))[:6] + strings.Repeat(" ", 1)
	seqName2 := (name2 + strings.Repeat(" ", 10))[:6] + strings.Repeat(" ", 1)

	complementSeq2 := strings.Repeat(" ", offset) + seq2 + strings.Repeat(" ", (len(seq1)-offset-len(seq2)))
	complementMatching := strings.Repeat(" ", offset) + matching + strings.Repeat(" ", (len(seq1)-offset-len(matching)))

	outputLine := make([]string, 6)
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
			sequence1 = seq1[(i * lineWidth):]
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
		outputLine[0] = strings.Repeat(" ", len(seqName1)) + scale1 + "\n"
		outputLine[1] = seqName1 + sequence1 + "\n"
		outputLine[2] = strings.Repeat(" ", len(seqName1)) + matchingSequence + "\n"
		outputLine[3] = (seqName2 + sequence2) + "\n"
		outputLine[4] = (strings.Repeat(" ", len(seqName1)) + scale2) + "\n"
		outputLine[5] = "\n"
		for i := 0; i < 6; i++ {
			_, err := writer.WriteString(outputLine[i])
			if err != nil {
				panic(err)
			}
		}
		/* 		_, err := writer.WriteString(strings.Repeat(" ", len(seqName1)+1) + scale1)
		   		_, err := writer.WriteString(seqName1, sequence1)
		   		_, err := writer.WriteString(strings.Repeat(" ", len(seqName1)), matchingSequence)
		   		_, err := writer.WriteString(seqName2, sequence2)
		   		_, err := writer.WriteString(strings.Repeat(" ", len(seqName1)+1) + scale2)
		   		_, err := writer.WriteString()
		*/
	}
	return nil

}
