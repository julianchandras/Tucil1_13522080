package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	fmt.Println("+==========================================================================+")
	fmt.Println("|                                                                          |")
	fmt.Println("|  ____                  __           ____        __                       |")
	fmt.Println("| /\\  _`\\               /\\ \\         /\\  _`\\   __/\\ \\                      |")
	fmt.Println("| \\ \\ \\L\\ \\__  __    ___\\ \\ \\/'\\     \\ \\,\\L\\_\\/\\_\\ \\ \\____     __   _ __   |")
	fmt.Println("|  \\ \\ ,__/\\ \\/\\ \\ /' _ `\\ \\ , <      \\/_\\__ \\\\/\\ \\ \\ '__`\\  /'__`\\/\\`'__\\ |")
	fmt.Println("|   \\ \\ \\/\\ \\ \\_\\ \\/\\ \\/\\ \\ \\ \\\\`\\      /\\ \\L\\ \\ \\ \\ \\ \\L\\ \\/\\  __/\\ \\ \\/  |")
	fmt.Println("|    \\ \\_\\ \\ \\____/\\ \\_\\ \\_\\ \\_\\ \\_\\    \\ `\\____\\ \\_\\ \\_,__/\\ \\____\\\\ \\_\\  |")
	fmt.Println("|     \\/_/  \\/___/  \\/_/\\/_/\\/_/\\/_/     \\/_____/\\/_/\\/___/  \\/____/ \\/_/  |")
	fmt.Println("|                                                                          |")
	fmt.Println("+==========================================================================+")
	fmt.Println()

	play := true

	for play {
		var paths listOfPath

		bufferSize, tokenMatrix, sequences, filePath := initiate()
		startTime := time.Now()
		findAllPaths(tokenMatrix, bufferSize, &paths)
		optimalPath, lastIdx, reward := findOptimalPath(bufferSize, tokenMatrix, paths, sequences)
		doSave, runtime := outputResult(tokenMatrix, reward, optimalPath, lastIdx, startTime)
		if doSave {
			saveResult(tokenMatrix, reward, optimalPath, lastIdx, runtime, filePath)
		}

		var input string
		for input != "y" && input != "n" {
			fmt.Print("\nApakah ingin bermain lagi? (y/n) ")
			fmt.Scanln(&input)
		}
		if input == "n" {
			play = false
			fmt.Println()
			fmt.Println("  ____                                                 ")
			fmt.Println(" /\\  _`\\                                          __    ")
			fmt.Println(" \\ \\,\\L\\_\\     __      ___ ___   _____      __   /\\_\\   ")
			fmt.Println("  \\/_\\__ \\   /'__`\\  /' __` __`\\/\\ '__`\\  /'__`\\ \\/\\ \\  ")
			fmt.Println("    /\\ \\L\\ \\/\\ \\L\\.\\_/\\ \\/\\ \\/\\ \\ \\ \\L\\ \\/\\ \\L\\.\\_\\ \\ \\ ")
			fmt.Println("    \\ `\\____\\ \\__/.\\_\\ \\_\\ \\_\\ \\_\\ \\ ,__/\\ \\__/.\\_\\\\ \\_\\")
			fmt.Println("     \\/_____/\\/__/\\/_/\\/_/\\/_/\\/_/\\ \\ \\/  \\/__/\\/_/ \\/_/")
			fmt.Println("                                   \\ \\_\\                ")
			fmt.Println("                                    \\/_/                ")
			fmt.Println("  _____                                        __       ")
			fmt.Println(" /\\___ \\                                      /\\ \\      ")
			fmt.Println(" \\/__/\\ \\  __  __    ___ ___   _____      __  \\ \\ \\     ")
			fmt.Println("    _\\ \\ \\/\\ \\/\\ \\ /' __` __`\\/\\ '__`\\  /'__`\\ \\ \\ \\    ")
			fmt.Println("   /\\ \\_\\ \\ \\ \\_\\ \\/\\ \\/\\ \\/\\ \\ \\ \\L\\ \\/\\ \\L\\.\\_\\ \\_\\   ")
			fmt.Println("   \\ \\____/\\ \\____/\\ \\_\\ \\_\\ \\_\\ \\ ,__/\\ \\__/.\\_\\\\/\\_\\  ")
			fmt.Println("    \\/___/  \\/___/  \\/_/\\/_/\\/_/\\ \\ \\/  \\/__/\\/_/ \\/_/  ")
			fmt.Println("                                 \\ \\_\\                  ")
			fmt.Println("                                  \\/_/                  ")
			fmt.Println()
		} else {
			fmt.Println()
		}
	}
}
