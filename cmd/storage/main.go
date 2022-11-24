package main

import (
	"fmt"
	"log"
)
import "github.com/robertgarayshin/ozonLearning/storage"

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("text.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("It's restored", restoredFile)
}
