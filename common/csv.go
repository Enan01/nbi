package common

import (
	"bufio"
	"os"
	"strings"
)

// LineColumnsCsvFileFromLineNo return columns of line, line started from param lineNo.
// First lineNo is 0.
func LineColumnsCsvFileFromLineNo(filePath string, lineNo int) (lineColumns [][]string, err error) {
	var lines []string
	lines, err = LinesCsvFileFromLineNo(filePath, lineNo)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		cols := strings.Split(line, ",")
		lineColumns = append(lineColumns, cols)
	}
	return
}

// LinesCsvFileFromLineNo return string array contains csv lines, line started from param lineNo.
// First lineNo is 0.
func LinesCsvFileFromLineNo(filePath string, lineNo int) (lines []string, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bio := bufio.NewReader(file)

	for line := 0; ; line++ {
		var l []byte
		l, _, err = bio.ReadLine()
		if len(l) <= 0 {
			break
		}
		if line < lineNo {
			continue
		}

		lines = append(lines, string(l))
	}

	return
}
