package main

import (
	"fmt"
	"os"

	"main/tetris"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] != "" {
		result, err := Start(os.Args[1])
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Error: command is not correct")
		fmt.Println("Example: go run . tetris.txt")
	}
}

func Start(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()
	myArray, err := tetris.ReadInputFile(file)
	if err != nil {
		return "File Error", err
	} else {
		tetris.Solve(myArray)
		return tetris.PrintSolution(), nil
	}
}
