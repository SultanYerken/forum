package app

import (
	"fmt"
	"log"
	"os"
)

func CreateDB() {
	_, err := os.Stat("./internal/usecase/repo/database.db")
	if err == nil {
		return
	} else {
		file, err := os.Create("./internal/usecase/repo/database.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Println("DB created")
		file.Close()
	}
}
