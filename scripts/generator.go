package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

const (
	folder   = "documents"
	fileName = "users.csv"
	rowsNum  = 100000
)

func main() {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	filePath := filepath.Join(folder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"id", "name", "age", "email"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing header to CSV:", err)
		return
	}

	for i := 1; i <= rowsNum; i++ {
		id := fmt.Sprintf("%d", i)
		age := fmt.Sprintf("%d", rand.Intn(45)+15)
		name := fmt.Sprintf("name%d", i)
		email := fmt.Sprintf("user%d@testmail.com", i)

		row := []string{id, name, age, email}

		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing row to CSV:", err)
			return
		}
	}

	fmt.Printf("CSV file generated successfully at %s\n", filePath)
}
