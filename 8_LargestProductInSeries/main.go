package main

import (
	"fmt"
	"io"
	"os"
)

type Grid [][]byte

const GRID_COL int = 50
const GRID_ROW int = 20

func main() {
	//file, err := os.Open("input.txt")
	//
	//if err != nil {
	//	log.Fatalf("Could not open file: %s\n", err.Error())
	//}
	//
	//defer file.Close()
	//
	//grid, err := readGrid(file)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, line := range grid {
	//	for _, ch := range line {
	//		print(int(ch))
	//	}
	//	print("\n")
	//}
}

//https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func readGrid(file *os.File) (Grid, error) {
	grid := make([][]byte, GRID_COL)

	for i := 0; i < len(grid); i++ {

		grid[i] = make([]byte, GRID_ROW)

		if _, err := file.Read(grid[i]); err == io.EOF {
			return nil, fmt.Errorf("Could not read in full grid: %s\n", err.Error())
		}
	}

	return grid, nil
}
