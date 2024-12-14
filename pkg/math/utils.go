package math

import (
	"fmt"
	"strconv"
	"strings"
)

func Float64ArrayToString(a []float64, braces ...string) string {
	retArray := make([]string, len(a))
	for i, v := range a {
		retArray[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	if len(braces) > 1 {
		return fmt.Sprintf("%s%v%s", braces[0], strings.Join(retArray, ", "), braces[1])
	}
	return fmt.Sprintf("[%v]", retArray)
}
