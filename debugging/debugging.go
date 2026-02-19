package debugging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

var logger *log.Logger

func Log(s string) {
	if logger == nil {
		wd, err := os.Getwd()
		if err != nil {
			return
		}
		f, err := os.Create(filepath.Join(wd, "debug.log"))
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "", log.Default().Flags())
	}
	logger.Println(s)
}

func LogWithType(name string, component any) {
	Log(fmt.Sprintf("%v %s", reflect.TypeOf(component), name))
}
