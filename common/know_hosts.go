package common

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"regexp"
)

type KnowHost struct {
	SrcName string
	DstName string
}

// 读取文件
func (knowHost *KnowHost) ReadFile() (lines []string) {

	inputFile, ok := os.Open(knowHost.SrcName)
	if ok != nil {
		fmt.Print("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		os.Exit(1)
		return
	}

	defer  inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		line,readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		lines = append(lines, line)
	}

	return lines
}

// 处理文件
func (knowHost *KnowHost) FileHandler(lines []string) (done []string) {

	for _,line := range lines {
		ok, _ := regexp.MatchString("^\n$", line)
		if ok {
			continue
		}
		re, _ := regexp.Compile("\\[([\\w-.]+)]:(9922|9923|60022|8822|58522|22|9933|2222)")
		str := re.ReplaceAllString(line, "$1")
		done = append(done, str)
	}

	return done
}

// 备份文件
func (knowHost *KnowHost) BackupFile() (written int64, err error) {
	src, err := os.Open(knowHost.SrcName)
	if err != nil {
		fmt.Print("An error occurred on opening the file. Does the file exist?\n" +
			"Have you got access to it?\n")
		os.Exit(1)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(knowHost.DstName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Print("An error occurred on opening the file. Have you got access to it?\n")
		os.Exit(1)
		return
	}
	defer  dst.Close()

	return io.Copy(dst, src)

}

// 写文件
func (knowHost *KnowHost) WriteFile(lines []string) {

	outputFile, err := os.OpenFile(knowHost.SrcName, os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("An error occurred on opening the file. Have you got access to it?\n")
		os.Exit(1)
		return
	}
	defer  outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	for _, line := range lines {
		outputWriter.WriteString(line)
	}
	outputWriter.Flush()
}