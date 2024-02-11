package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func loadData() (int, matrix, listOfSequence, string) {
	var input string
	fmt.Println("Pilih cara memberi masukkan\n1. Melalui file \".txt\"\n2. Melalui Command Line (Matriks dan sekuens akan dihasilkan secara acak)")
	fmt.Print("\nPilihan (1/2): ")
	fmt.Scanln(&input)
	fmt.Println()
	if input == "1" {
		fmt.Print("Masukkan nama file (.txt): ")
		fmt.Scanln(&input)

		exePath, err := os.Executable()
		if err != nil {
			fmt.Println("Error:", err)
			return 0, matrix{}, listOfSequence{}, ""
		}

		exeDir := filepath.Dir(exePath)
		projectRoot := filepath.Dir(exeDir)
		relativePath := filepath.Join(projectRoot, "test", input)

		bufferSize, tokenMatrix, sequences := parseFile(relativePath)
		return bufferSize, tokenMatrix, sequences, relativePath
	} else {
		bufferSize, tokenMatrix, sequences := initiateRandomly()

		fmt.Println("\nMatriks:")
		for i := 0; i < tokenMatrix.Height; i++ {
			for j := 0; j < tokenMatrix.Width; j++ {
				fmt.Print(tokenMatrix.Buffer[i][j])
				if j != tokenMatrix.Width-1 {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}

		fmt.Println("\nSekuens:")
		for i := 0; i < sequences.Neff; i++ {
			for j := 0; j < len(sequences.Buffer[i].Tokens); j++ {
				fmt.Printf("%s ", sequences.Buffer[i].Tokens[j])
			}
			fmt.Printf("Reward: %d\n", sequences.Buffer[i].Reward)
		}
		return bufferSize, tokenMatrix, sequences, ""
	}
}

func initiateRandomly() (int, matrix, listOfSequence) {
	var numOfTokens int
	var inputString string
	var bufferSize int
	var width, height int
	var tokenMatrix matrix
	var numOfSequences int
	var maxSequenceLength int
	var sequences listOfSequence
	var tokens []string

	fmt.Scanln(&numOfTokens)
	inputScanner := bufio.NewScanner(os.Stdin)
	if inputScanner.Scan() {
		inputString = inputScanner.Text()
	}
	fmt.Scanln(&bufferSize)
	fmt.Scanf("%d %d\n", &width, &height)
	fmt.Scanln(&numOfSequences)
	fmt.Scanln(&maxSequenceLength)

	tokens = strings.Split(inputString, " ")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	tokenMatrix.Buffer = make([][]string, height)
	for i := 0; i < height; i++ {
		tokenMatrix.Buffer[i] = make([]string, width)
		for j := 0; j < width; j++ {
			tokenMatrix.Buffer[i][j] = tokens[random.Intn(numOfTokens)]
		}
	}
	tokenMatrix.Width, tokenMatrix.Height = width, height

	for i := 0; i < numOfSequences; i++ {
		var seq sequence
		currentSeqLength := random.Intn(maxSequenceLength-1) + 1
		for j := 0; j < currentSeqLength; j++ {
			seq.Tokens = append(seq.Tokens, tokens[random.Intn(numOfTokens)])
		}
		seq.Reward = (random.Intn(10) + 1) * 5 // kelipatan 10, 0..50
		sequences.Buffer = append(sequences.Buffer, seq)
		sequences.Neff += 1
	}

	return bufferSize, tokenMatrix, sequences
}

func findAllPaths(tokenMatrix matrix, bufferSize int, paths *listOfPath) {
	width := tokenMatrix.Width
	for i := 0; i < width; i++ {
		initialPath := path{[]point{{i, 0}}, 1}
		findPath(tokenMatrix, bufferSize, 0, &initialPath, paths)
	}
}

func findPath(tokenMatrix matrix, bufferSize int, iterator int, tempPath *path, paths *listOfPath) {
	if iterator == bufferSize-1 {
		newPath := path{
			Coordinates: make([]point, len(tempPath.Coordinates)),
			Neff:        tempPath.Neff,
		}
		copy(newPath.Coordinates, tempPath.Coordinates)
		paths.Buffer = append(paths.Buffer, newPath)
		paths.Neff++
		return
	} else {
		width := tokenMatrix.Width
		height := tokenMatrix.Height

		var currentPoint point = tempPath.Coordinates[tempPath.Neff-1]
		// gerak vertikal
		if iterator%2 == 0 {
			for j := 1; j < height; j++ {
				tempPoint := point{currentPoint.Column, (currentPoint.Row + j) % height}
				if !isInPath(tempPoint, *tempPath) {
					tempPath.Coordinates = append(tempPath.Coordinates, tempPoint)
					tempPath.Neff += 1
					findPath(tokenMatrix, bufferSize, iterator+1, tempPath, paths)
					tempPath.Coordinates = tempPath.Coordinates[:tempPath.Neff-1]
					tempPath.Neff -= 1
				}
			}
		} else { // horizontal
			for j := 1; j < width; j++ {
				tempPoint := point{(currentPoint.Column + j) % width, currentPoint.Row}
				if !isInPath(tempPoint, *tempPath) {
					tempPath.Coordinates = append(tempPath.Coordinates, tempPoint)
					tempPath.Neff += 1
					findPath(tokenMatrix, bufferSize, iterator+1, tempPath, paths)
					tempPath.Coordinates = tempPath.Coordinates[:tempPath.Neff-1]
					tempPath.Neff -= 1
				}
			}
		}
	}
}

func isInPath(coordinate point, p path) bool {
	i := 0
	found := false
	for i < p.Neff && !found {
		if p.Coordinates[i] == coordinate {
			found = true
		} else {
			i += 1
		}
	}
	return found
}

func findOptimalPath(bufferSize int, tokenMatrix matrix, paths listOfPath, sequences listOfSequence) (path, int, int) {
	maxReward := 0
	minIdx := 0
	idxPath := 0
	for idx, path := range paths.Buffer {
		reward, lastIdx := findReward(tokenMatrix, path, sequences)
		if reward > maxReward {
			maxReward = reward
			minIdx = lastIdx
			idxPath = idx
		} else if reward == maxReward {
			if lastIdx < minIdx {
				minIdx = lastIdx
				idxPath = idx
			}
		}
	}
	return paths.Buffer[idxPath], minIdx, maxReward
}

func findReward(tokenMatrix matrix, path path, sequences listOfSequence) (int, int) {
	lastIdx := 0
	reward := 0
outerLoop:
	for _, val := range sequences.Buffer {
		// fmt.Println("----------------------------------------")
		// mencoba semua sequence
		maxIterationIdx := path.Neff - len(val.Tokens)
		for i := 0; i <= maxIterationIdx; i++ {
			// i menunjukkan indeks path, titik awal menyocokkan sequence
			j := 0
			stillTrue := true
			for j < len(val.Tokens) && stillTrue {
				if val.Tokens[j] != tokenMatrix.Buffer[path.Coordinates[i+j].Row][path.Coordinates[i+j].Column] {
					// fmt.Println(i, j, val.Tokens[j], tokenMatrix.Buffer[path.Coordinates[i+j].Row][path.Coordinates[i+j].Column], "tetot")
					stillTrue = false
				} else {
					// fmt.Println(i, j, val.Tokens[j], tokenMatrix.Buffer[path.Coordinates[i+j].Row][path.Coordinates[i+j].Column])
					j++
				}
			}
			if stillTrue {
				tempLastIdx := j + i - 1
				// fmt.Println("FOUND", tempLastIdx)
				if tempLastIdx > lastIdx {
					lastIdx = tempLastIdx
				}
				reward += val.Reward
				continue outerLoop
			}
		}
	}
	return reward, lastIdx
}

func outputResult(tokenMatrix matrix, reward int, optimalPath path, lastIdx int, startTime time.Time) (bool, time.Duration) {
	fmt.Printf("\n%d\n", reward)
	for i := 0; i <= lastIdx; i++ {
		fmt.Printf("%s", tokenMatrix.Buffer[optimalPath.Coordinates[i].Row][optimalPath.Coordinates[i].Column])
		if i != lastIdx {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
	for i := 0; i <= lastIdx; i++ {
		fmt.Printf("%d, %d\n", optimalPath.Coordinates[i].Column+1, optimalPath.Coordinates[i].Row+1)
	}
	endTime := time.Now()
	runtime := endTime.Sub(startTime)
	fmt.Println()
	fmt.Println(runtime)

	fmt.Print("\nApakah ingin menyimpan hasil? (y/n) ")
	var save string
	fmt.Scanln(&save)

	doSave := (save == "y")
	return doSave, runtime
}
