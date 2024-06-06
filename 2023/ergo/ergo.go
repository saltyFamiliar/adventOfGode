package ergo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"unicode"
)

type ScannerCloser struct {
	fp      *os.File
	scanner *bufio.Scanner
}

func NewScannerCloser(filePath string) (*ScannerCloser, error) {
	fp, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(fp)
	newSC := &ScannerCloser{
		fp:      fp,
		scanner: scanner,
	}
	return newSC, nil
}

func (sc *ScannerCloser) Scan() (string, bool) {
	ok := sc.scanner.Scan()
	return sc.scanner.Text(), ok
}

// this ScannerCloser should no longer be available after Close() is called
// is there a way to delete it?
func (sc *ScannerCloser) Close() error {
	sc.fp.Close()
	return sc.scanner.Err()
}

func (sc *ScannerCloser) ScanLines() (lines []string, err error) {
	for {
		line, ok := sc.Scan()
		if !ok {
			err = sc.Close()
			if len(lines) > 0 && lines[len(lines)-1] == "" {
				lines = lines[:len(lines)-1]
			}
			return lines, err
		}
		lines = append(lines, line)
	}
}

func (sc *ScannerCloser) LineIter() func(func(string) bool) {
	return func(yield func(string) bool) {
		for {
			s, ok := sc.Scan()
			if !yield(s) || !ok {
				err := sc.Close()
				if err != nil {
					log.Println(err.Error())
				}
				return
			}
		}
	}
}

func Must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ":" + err.Error())
	}
}

func RuneToIntIfDigit(r rune) (int, error) {
	if !unicode.IsDigit(r) {
		return 0, fmt.Errorf("Not a digit")
	}

	return int(r - '0'), nil
}

func EzIntParse(digits string) int {
	num, err := strconv.Atoi(digits)
	Must("parse number", err)
	return num
}
