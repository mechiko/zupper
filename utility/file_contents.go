package utility

import (
	"fmt"
	// "zupper/domain"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const dirTemplate = "templates"
const moduleName = `zupper/`

func FilePackageContents(root, filename string) string {
	filename = filepath.Join(root, retrieveCallPackageName(), dirTemplate, filename)
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("%s %s", err.Error(), filename))
	}
	return string(b)
}

// пример имени пакета от корня "zupper/webapp/pages/setup"
func retrieveCallPackageName() string {
	pc, _, _, _ := runtime.Caller(2)
	// _, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}
	return strings.TrimPrefix(packageName, moduleName)
}
