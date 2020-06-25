package utilities

import(
	"strings"
)

// CreateKeyString - string concat
func CreateKeyString(keys ...string) string {
	var b strings.Builder
	for i := 0; i < len(keys); i++ {
	  b.WriteString(keys[i])
	}
	return b.String()
  }