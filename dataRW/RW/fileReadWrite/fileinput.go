package fileReadWrite

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileInput() {
	inputFile, err := os.Open("file.txt")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
		"Dose the file exist\n" + "Have you got acces to it\n")
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, err := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if err == io.EOF {
			return
		}
	}
}
