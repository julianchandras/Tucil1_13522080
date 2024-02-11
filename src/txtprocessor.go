package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func parseFile(filename string) (int, matrix, listOfSequence) {
	var err error
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}

	str := string(data)

	var bufferSize int
	var gameMatrix matrix
	var listOfSeq listOfSequence

	lines := splitLines(str)

	bufferSize, err = strconv.Atoi(lines[0])
	if err != nil {
		fmt.Println("Error:", err)
	}

	temp := strings.Split(lines[1], " ")
	gameMatrix.Width, err = strconv.Atoi(temp[0])
	if err != nil {
		fmt.Println("Error:", err)
	}

	gameMatrix.Height, err = strconv.Atoi(temp[1])
	if err != nil {
		fmt.Println("Error:", err)
	}

	currentLine := 2
	for i := 0; i < gameMatrix.Height; i++ {
		gameMatrix.Buffer = append(gameMatrix.Buffer, strings.Split(lines[currentLine], " "))
		currentLine += 1
	}

	numOfSequences, err := strconv.Atoi(lines[currentLine])
	if err != nil {
		fmt.Println("Error:", err)
	}
	currentLine += 1

	listOfSeq.Neff = 0
	for i := 0; i < numOfSequences; i++ {
		tempToken := strings.Split(lines[currentLine], " ")
		currentLine += 1
		tempReward, err := strconv.Atoi(lines[currentLine])
		if err != nil {
			fmt.Println("Error:", err)
		}

		newSeq := sequence{
			Tokens: tempToken,
			Reward: tempReward,
		}
		listOfSeq.Buffer = append(listOfSeq.Buffer, newSeq)
		listOfSeq.Neff += 1
		currentLine += 1
	}
	return bufferSize, gameMatrix, listOfSeq
}

func splitLines(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n' || r == '\r'
	})
}

func saveResult(tokenMatrix matrix, reward int, optimalPath path, lastIdx int, runtime time.Duration, filePath string) {
	var newPath string
	var err error
	if filePath != "" {
		extension := filepath.Ext(filePath)
		newPath = filePath[:len(filePath)-len(extension)]
		newPath += "-result" + ".txt"
	} else {
		var filename string
		fmt.Print("Masukkan nama file untuk menyimpan hasil: ")
		fmt.Scanln(&filename)
		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		exeDir := filepath.Dir(exePath)
		projectRoot := filepath.Dir(exeDir)
		newPath = filepath.Join(projectRoot, "test", filename+".txt")
	}
	file, err := os.Create(newPath)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()
	if reward != 0 {
		_, err = file.WriteString(fmt.Sprintf("%d\n", reward))
		if err != nil {
			fmt.Println("Error:", err)
		}

		var tempString string
		for i := 0; i <= lastIdx; i++ {
			tempString += tokenMatrix.Buffer[optimalPath.Coordinates[i].Row][optimalPath.Coordinates[i].Column]
			if i != lastIdx {
				tempString += " "
			}
		}
		tempString += "\n"

		_, err = file.WriteString(tempString)
		if err != nil {
			fmt.Println("Error:", err)
		}

		for i := 0; i <= lastIdx; i++ {
			_, err = file.WriteString(fmt.Sprintf("%d, %d\n", optimalPath.Coordinates[i].Column+1, optimalPath.Coordinates[i].Row+1))
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		tempString = "\n" + runtime.String()

		_, err = file.WriteString(tempString)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
