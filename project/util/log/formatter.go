package log

import (
	"fmt"
	"strings"
)

type Formatter struct {
	formatStrings []string
	params        []interface{}
}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Add(formatStr string, params ...interface{}) {
	f.formatStrings = append(f.formatStrings, formatStr)
	f.params = append(f.params, params...)
}

func (f *Formatter) ToString(separator string) string {
	formatStr := strings.Join(f.formatStrings, separator)

	return fmt.Sprintf(formatStr, f.params...)
}
