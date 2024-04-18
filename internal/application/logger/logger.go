package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var Info *log.Logger
var Error *log.Logger

func init() {

	absPath, err := filepath.Abs("./log")
	if err != nil {
		fmt.Print("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"General-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	Info = log.New(generalLog, "INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(generalLog, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
