package callstack

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// LogErrStack is a call stack logging function
func LogErrStack(err error) {
	fmt.Printf("ERROR %s - %s\n", time.Now().Format("2006-01-02 15:04:05.000"), err)
	if err, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		for _, f := range err.StackTrace() {
			pc := uintptr(f) - 1
			fn := runtime.FuncForPC(pc)
			file, line := fn.FileLine(pc)
			fmt.Printf("\t%s %d -\t%s\n", fileName(file), line, fn.Name())
		}
	}
	fmt.Println()
}

func fileName(name string) string {
	i := strings.LastIndex(name, "/")
	return name[i+1:]
}
