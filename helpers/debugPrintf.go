package helpers

import (
	"fmt"
	"runtime"
	"strings"
)

//Log this function show a customize printf
func Log(msg string, args ...interface{}) {
	programCounter, _, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	last := strings.Split(fn.Name(), "/")
	pathFile := last[len(last)-1]

	prefix := fmt.Sprintf("[%s %d] %s", pathFile, line, msg)
	fmt.Printf(prefix, args...)
	fmt.Println()
}
