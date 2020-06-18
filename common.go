package gota_common

import (
	"log"
	"os"
)

func GetProjectHomeDir() string {
	prjHome, err := os.Getwd()
	if err != nil {
		log.Fatal( err )
	}
	return prjHome
}
