package common

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// LineColumnsCsvFileFromLineNo return columns of line, line started from param lineNo.
// startLineNo start from 0.
// if endLineNo == -1 means last line.
func LineColumnsCsvFile(filePath string, startLineNo int, endLineNo int) (lineColumns [][]string, err error) {
	var lines []string
	lines, err = LinesCsvFile(filePath, startLineNo, endLineNo)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		cols := strings.Split(line, ",")
		trimCols := make([]string, 0)
		for _, col := range cols {
			trimCols = append(trimCols, strings.TrimSpace(col))
		}
		lineColumns = append(lineColumns, trimCols)
	}
	return
}

// LinesCsvFileFromLineNo return string array contains csv lines, line started from param lineNo.
// startLineNo start from 0.
// if endLineNo == -1 means last line.

func LinesCsvFile(filePath string, startLineNo int, endLinNo int) (lines []string, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bio := bufio.NewReader(file)

	for line := 0; ; line++ {
		var (
			l []byte
			e error
		)
		l, _, e = bio.ReadLine()
		if e != nil {
			if e == io.EOF {
				break
			}
			err = e
			return
		}

		if line < startLineNo {
			continue
		}
		if endLinNo != -1 && line > endLinNo {
			continue
		}

		if len(l) <= 0 {
			break
		}

		lines = append(lines, string(l))
	}

	return
}
