package debugger

import (
	"log"
	"os"
	"time"
)

var DebugFunction = func(string) {}
var refTime = time.Now()

var logger = log.New(os.Stderr, "\033[0;34mDEBUG\033[0m: ", log.Ldate|log.Ltime)

func setReferenceTime() {
	refTime = time.Now()
}

func EnableDebugMode() {
	DebugFunction = debug
}

func Debug(desc string) {
	DebugFunction(desc)
}

// This method is set as the debug function if debugging is enabled
func debug(desc string) {
	logger.Printf("+%dms elapsed [%s]\n", time.Since(refTime).Milliseconds(), desc)
	setReferenceTime()
}
