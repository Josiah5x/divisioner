package divisioner

import (
	"errors"
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
// How to call the function, varName, Error then Division
func Division(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		log.SetPrefix("LOG: ")
		log.Println("Can't divide by zero")
		return float64(0), errors.New(("can't divide by zero"))
	} else {

		log.SetPrefix("LOG: ")
		log.Println("Answer=", num1/num2)
		return float64(num1 / num2), nil
	}

}
