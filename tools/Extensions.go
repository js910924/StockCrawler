package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GetFavoriteStockList() {
	defer func() {
		if err := recover(); err != nil {
			CreateFile("favoriteList.txt")
		}
	}()

	file, err := os.Open("favoriteList.txt")
	if err != nil {
		log.Panicln("Can't find favoriteList.txt file")
	}

	defer file.Close()

	b, err := ioutil.ReadAll(file)
	fmt.Println(b)
}

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Panicln("Create file failed")
	}

	defer file.Close()
	log.Printf("Create %s file successfully", fileName)
}
