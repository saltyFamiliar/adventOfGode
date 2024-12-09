package aocio

import (
	"adventOfGode/toolbelt"
	"bufio"
	"log"
	"os"
)

type ScannerCloser struct {
	fp      *os.File
	scanner *bufio.Scanner
}

func NewScannerCloser(filePath string) *ScannerCloser {
	fp, err := os.Open(filePath)
	toolbelt.Must("couldn't open file", err)
	scanner := bufio.NewScanner(fp)
	newSC := &ScannerCloser{
		fp:      fp,
		scanner: scanner,
	}
	return newSC
}

func (sc *ScannerCloser) Scan() (string, bool) {
	ok := sc.scanner.Scan()
	return sc.scanner.Text(), ok
}

func (sc *ScannerCloser) Close() error {
	sc.fp.Close()
	return sc.scanner.Err()
}

func (sc *ScannerCloser) ScanLines() (lines []string) {
	defer func() {
		log.Println(sc.Close().Error())
	}()
	for {
		line, ok := sc.Scan()
		if !ok {
			if len(lines) > 0 && lines[len(lines)-1] == "" {
				lines = lines[:len(lines)-1]
			}
			return lines
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
