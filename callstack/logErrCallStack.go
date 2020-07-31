package callstack

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// PrintErrStack is printing formatted call stack
func PrintErrStack(err error) {
	fmt.Printf("ERROR %s - %s\n", time.Now().Format("2006-01-02 15:04:05.000"), err)

	functionCalls := errStack(err)
	for _, fc := range functionCalls {
		fmt.Printf("\t%s %d -\t%s\n", fc.fileName, fc.line, fc.funcName)
	}

	fmt.Println()
}

type functionCall struct {
	filePath string
	fileName string
	funcName string
	line     int
}

func errStack(err error) (functionCalls []functionCall) {
	if err, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		for _, f := range err.StackTrace() {
			pc := uintptr(f) - 1
			fn := runtime.FuncForPC(pc)
			file, line := fn.FileLine(pc)
			fc := functionCall{
				filePath: file,
				fileName: fileName(file),
				funcName: fn.Name(),
				line:     line,
			}
			functionCalls = append(functionCalls, fc)
		}
	}

	return
}

func fileName(name string) string {
	i := strings.LastIndex(name, "/")
	return name[i+1:]
}
