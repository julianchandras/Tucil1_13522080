package main

import (
	// "fmt"
	"time"
)

func main() {
	var paths listOfPath

	bufferSize, tokenMatrix, sequences, filePath := loadData()
	startTime := time.Now()
	findAllPaths(tokenMatrix, bufferSize, &paths)
	optimalPath, lastIdx, reward := findOptimalPath(bufferSize, tokenMatrix, paths, sequences)
	doSave, runtime := outputResult(tokenMatrix, reward, optimalPath, lastIdx, startTime)
	if doSave {
		saveResult(tokenMatrix, reward, optimalPath, lastIdx, runtime, filePath)
	}
	// OUTPUT
	
	// OUTPUT
	
	// Test
	// testPath := path{
	// 	Neff: 7,
	// 	Coordinates: []point{
	// 		{1, 0},
	// 		{1, 1},
	// 		{2, 1},
	// 		{2, 2},
	// 		{3, 2},
	// 		{3, 3},
	// 		{4, 3},
	// 	},
	// }
	// test1, test2 := findReward(tokenMatrix, testPath, sequences)
	// fmt.Println(test1, test2)
}
