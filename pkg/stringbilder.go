package pkg

import (
	"fmt"
	"strings"
)

func StringBuild(v ...any) string {
	sbuilder := strings.Builder{}
	for _, item := range v {
		sbuilder.WriteString(fmt.Sprintf("%v ", item))
	}
	return sbuilder.String()
}
