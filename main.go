package main

import "fmt"

// import "path/filepath"

func main() {
	fmt.Println("Choose Feature: ")
	fmt.Println("1. OBJ Viewer")
	fmt.Println("2. Voxelization")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		ObjViewer()
	case 2:
		fmt.Println("Source Obj Filepath:")
		var filepath string
		fmt.Scanln(&filepath)
		voxelization(filepath)
	}
}
