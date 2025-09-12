package logger

import (
	"os"
)

func init() {
	var logfile, err = os.Open("~/.config/gohttp/log")
	if err != nil {
		logfile , err = os.Create("~/.config/gohttp/log")
	}	
}

func log(fiele os.File){

	
}
