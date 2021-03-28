package main 

import (
	"fmt"
	"os"
	"log"
)

var (
	info	os.FileInfo
	err	error
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stdout, "Using: %s file\n", os.Args[1])
		os.Exit(1)
	}

	info, err = os.Stat(os.Args[2])
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File Not Entered. Enter File Name.")
		}
	}

	fmt.Println("Project Name : FileForensicer")
	fmt.Println("Created : https://github.com/EyupErgin")
	fmt.Println("Github  : https://github.com/IntelSights/FileForensicer/")
	fmt.Println("   ")
	fmt.Println("File Name        :", info.Name())
	fmt.Println("Directory        :", info.IsDir())
	fmt.Println("Byte Size        :", info.Size())
	fmt.Println("Permissions      :", info.Mode())
	fmt.Println("Last Modified    :", info.ModTime())
	fmt.Printf("System Interface : %T\n", info.Sys())
	fmt.Printf("System Info      : %+v\n\n", info.Sys())
}
