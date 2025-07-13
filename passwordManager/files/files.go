package files

import (
	"fmt"
	"os"
)

func ReadFile() {}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully")
	file.Close()
}
