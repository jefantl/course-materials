package logging

import "log"

var LOG_LEVEL = 1

func IfLevel(s string, level int) {
	if level <= LOG_LEVEL {
		log.Println(s)
	}
}
