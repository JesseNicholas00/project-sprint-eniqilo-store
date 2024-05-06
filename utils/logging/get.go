package logging

import (
	"fmt"
	"log"
	"strings"
)

func GetLogger(context string, subcontext ...string) *log.Logger {
	return log.New(
		log.Writer(),
		fmt.Sprintf("[%s] | %s: ", context, strings.Join(subcontext, ".")),
		log.Ldate|log.Ltime|log.Lmsgprefix|log.Lshortfile,
	)
}
