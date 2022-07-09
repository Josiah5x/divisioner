package divisioner

import (
	"log"
	"os"
)

// Logging take one parameter which can be true or false///
func Logging(opt bool) {
	if opt {
		logfile, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		}
		log.SetOutput(logfile)

	}
}

// this is a division which takes two parameter num1 & num2 float64 number
func Division(num1, num2 float64) float64 {
	log.SetPrefix("LOG: ")
	log.Println("Answer=", num1/num2)
	return num1 / num2

}
