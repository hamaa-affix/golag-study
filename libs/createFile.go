package libs

import (
	"log"
	"os"
)

func Create() {
	file, err := os.Create("text.txt")
	if err != nil {
		log.Fatal("fatal to create file")
	}

	file.Write([]byte("create したぜ\n"))
	file.Close()
}
