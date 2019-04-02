package main

import (
	"fmt"
	"os"
)

func main() {
	data := []byte(`
{
  "alphabet": [
    {
      "number": 2
    }
  ]
}

`)
	file, err := os.Create("api.json")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)

	fmt.Println("Done.")
}
