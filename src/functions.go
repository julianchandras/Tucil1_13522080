package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func initiate() (int, matrix, listOfSequence, string) {
	var input string

	fmt.Println("Pilih cara memberi masukkan\n1. Melalui file \".txt\"\n2. Melalui Command Line (Matriks dan sekuens akan dihasilkan secara acak)")
	fmt.Print("\nPilihan (1/2): ")
	fmt.Scanln(&input)
	fmt.Println()

	for input != "1" && input != "2" {
		fmt.Println("Pilihan tidak tersedia!")
		fmt.Println("Pilih cara memberi masukkan\n1. Melalui file \".txt\"\n2. Melalui Command Line (Matriks dan sekuens akan dihasilkan secara acak)")
		fmt.Print("\nPilihan (1/2): ")
		fmt.Scanln(&input)
		fmt.Println()
	}

	if input == "1" {
		var fileFound bool = false
		var relativePath string
		for !fileFound {
			fmt.Print("Masukkan nama file (.txt): ")
			fmt.Scanln(&input)

			exePath, err := os.Executable()
			if err != nil {
				fmt.Println("Error:", err)
				return 0, matrix{}, listOfSequence{}, ""
			}

			exeDir := filepath.Dir(exePath)
			projectRoot := filepath.Dir(exeDir)
			relativePath = filepath.Join(projectRoot, "test", input)

			_, err = os.Stat(relativePath)

			if err == nil {
				fileFound = true
			} else if os.IsNotExist(err) {
				fmt.Println("\nFile tidak tersedia!")
			}
		}
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

	fmt.Print("Masukkan jumlah token: ")
	fmt.Scanln(&numOfTokens)
	for numOfTokens <= 0 {
		fmt.Print("Jumlah token harus positif! Masukkan jumlah token: ")
		fmt.Scanln(&numOfTokens)
	}

	fmt.Print("Masukkan token: ")
	inputScanner := bufio.NewScanner(os.Stdin)
	if inputScanner.Scan() {
		inputString = inputScanner.Text()
	}
	tokens = strings.Split(inputString, " ")

	for len(tokens) != numOfTokens || !checkToken(tokens) {
		fmt.Printf("Token tidak sesuai! Masukkan token (Sebanyak %d dan alfanumerik dua karakter): ", numOfTokens)
		inputScanner := bufio.NewScanner(os.Stdin)
		if inputScanner.Scan() {
			inputString = inputScanner.Text()
		}
		tokens = strings.Split(inputString, " ")
	}

	fmt.Print("Masukkan ukuran buffer: ")
	fmt.Scanln(&bufferSize)

	for bufferSize <= 0 {
		fmt.Print("Ukuran buffer harus positif! Masukkan ukuran buffer: ")
		fmt.Scanln(&bufferSize)
	}

	fmt.Print("Masukkan lebar dan tinggi matriks: ")
	fmt.Scanf("%d %d\n", &width, &height)

	for width <= 0 || height <= 0 {
		fmt.Print("Lebar dan tinggi matriks harus positif! Masukkan lebar dan tinggi matriks: ")
		fmt.Scanf("%d %d\n", &width, &height)
	}

	fmt.Print("Masukkan jumlah sekuens: ")
	fmt.Scanln(&numOfSequences)

	for numOfSequences <= 0 {
		fmt.Print("Jumlah sekuens harus positif! Masukkan jumlah sequence: ")
		fmt.Scanln(&numOfSequences)
	}

	minSequenceLength := 2
	totalCombination := int(math.Pow(float64(numOfTokens), float64(minSequenceLength)))
	for totalCombination < numOfSequences {
		minSequenceLength++
		totalCombination += int(math.Pow(float64(numOfTokens), float64(minSequenceLength)))
	}

	fmt.Printf("Masukkan panjang sekuens maksimal (minimal %d): ", minSequenceLength)
	fmt.Scanln(&maxSequenceLength)

	for maxSequenceLength < minSequenceLength {
		fmt.Printf("Panjang maksimal kurang! Masukkan panjang sekuens maksimal (minimal %d): ", minSequenceLength)
		fmt.Scanln(&maxSequenceLength)
	}

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

	for sequences.Neff < numOfSequences {
		var seq sequence
		currentSeqLength := random.Intn(maxSequenceLength-1) + 2
		for j := 0; j < currentSeqLength; j++ {
			seq.Tokens = append(seq.Tokens, tokens[random.Intn(numOfTokens)])
		}
		seq.Reward = (random.Intn(10) + 1) * 5 // kelipatan 10, 0..50
		if !isInSequences(sequences, seq) {
			sequences.Buffer = append(sequences.Buffer, seq)
			sequences.Neff++
		}
	}

	return bufferSize, tokenMatrix, sequences
}

func isInSequences(sequences listOfSequence, seq sequence) bool {
	found := false
	i := 0
	for i < sequences.Neff && !found {
		currentSequence := sequences.Buffer[i]
		if len(currentSequence.Tokens) == len(seq.Tokens) {
			same := true
			j := 0
			for j < len(currentSequence.Tokens) && same {
				if currentSequence.Tokens[j] != seq.Tokens[j] {
					same = false
				} else {
					j++
				}
			}
			if same {
				found = true
			} else {
				i++
			}
		} else {
			i++
		}
	}
	return found
}

func checkToken(tokens []string) bool {
	pattern := "^[a-zA-Z0-9][a-zA-Z0-9]$"
	regex := regexp.MustCompile(pattern)

	for i := 0; i < len(tokens); i++ {
		if !regex.MatchString(tokens[i]) {
			return false
		}
	}
	return true
}

func findAllPaths(tokenMatrix matrix, bufferSize int, paths *listOfPath) {
	if bufferSize > 0 {
		width := tokenMatrix.Width
		for i := 0; i < width; i++ {
			initialPath := path{[]point{{i, 0}}, 1}
			findPath(tokenMatrix, bufferSize, 0, &initialPath, paths)
		}
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
					tempPath.Neff++
					findPath(tokenMatrix, bufferSize, iterator+1, tempPath, paths)
					tempPath.Coordinates = tempPath.Coordinates[:tempPath.Neff-1]
					tempPath.Neff--
				}
			}
		} else { // horizontal
			for j := 1; j < width; j++ {
				tempPoint := point{(currentPoint.Column + j) % width, currentPoint.Row}
				if !isInPath(tempPoint, *tempPath) {
					tempPath.Coordinates = append(tempPath.Coordinates, tempPoint)
					tempPath.Neff++
					findPath(tokenMatrix, bufferSize, iterator+1, tempPath, paths)
					tempPath.Coordinates = tempPath.Coordinates[:tempPath.Neff-1]
					tempPath.Neff--
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
			i++
		}
	}
	return found
}

func findOptimalPath(bufferSize int, tokenMatrix matrix, paths listOfPath, sequences listOfSequence) (path, int, int) {
	if paths.Neff != 0 {
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
	} else {
		return path{}, 0, 0
	}
}

func findReward(tokenMatrix matrix, path path, sequences listOfSequence) (int, int) {
	lastIdx := 0
	reward := 0
outerLoop:
	for _, val := range sequences.Buffer {
		// mencoba semua sequence
		maxIterationIdx := path.Neff - len(val.Tokens)
		for i := 0; i <= maxIterationIdx; i++ {
			// i menunjukkan indeks path, titik awal menyocokkan sequence
			j := 0
			stillTrue := true
			for j < len(val.Tokens) && stillTrue {
				if val.Tokens[j] != tokenMatrix.Buffer[path.Coordinates[i+j].Row][path.Coordinates[i+j].Column] {
					stillTrue = false
				} else {
					j++
				}
			}
			if stillTrue {
				tempLastIdx := j + i - 1
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
	var runtime time.Duration

	fmt.Printf("\n%d\n", reward)
	if optimalPath.Neff != 0 {
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
		runtime = endTime.Sub(startTime)
		fmt.Println()
		fmt.Println(runtime)
		fmt.Println()
	}
	var save string

	for save != "y" && save != "n" {
		fmt.Print("Apakah ingin menyimpan hasil? (y/n) ")
		fmt.Scanln(&save)
	}

	doSave := (save == "y")
	return doSave, runtime
}
