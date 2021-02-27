package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stdout, "Using: %s file\n", os.Args[0])
		os.Exit(1)
	}

	fileInfo, err = os.Stat(os.Args[1])
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}

	fmt.Println("Dosya Adı       :", fileInfo.Name())
	fmt.Println("Dizin           :", fileInfo.IsDir())	
	fmt.Println("Bayt Boyutu     :", fileInfo.Size())
	fmt.Println("İzinler         :", fileInfo.Mode())
	fmt.Println("Son Değiştirilme:", fileInfo.ModTime())
	fmt.Printf("Sistem Arayüzü   : %T\n", fileInfo.Sys())
	fmt.Printf("Sistem Bilgisi   : %+v\n\n", fileInfo.Sys())
}
