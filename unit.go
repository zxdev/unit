package unit

import (
	"bufio"
	"os"
	"strings"
)

/*
reads and builds a map[string]string from a simple unit style resource file

./sample
[sample]
key1 = value1
key2 = value2

	var u unit.Unit
	u.Parse("./sample", "sample")
	// map[key1:value1,key2:value2]
	t.Log(u)

*/

type Unit map[string]string

// Parse a unit file; path,section,...
func (m *Unit) Parse(path string, section ...string) bool {

	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	var row string
	var extract = len(section) == 0
	var scanner = bufio.NewScanner(f)
	(*m) = make(map[string]string)
	for scanner.Scan() {
		row = scanner.Text()

		switch {
		case len(row) == 0: // empty line
		case strings.HasPrefix(row, "#"): // comment
		case strings.HasPrefix(row, "[") && strings.HasSuffix(row, "]"): // section
			if extract = len(section) == 0; !extract {
				for i := range section {
					if extract = section[i] == row[1:len(row)-1]; extract {
						break
					}
				}
			}

		default:
			if extract { // extract key:value
				var seg []string
				switch {
				case strings.Contains(row, "="):
					seg = strings.SplitN(row, "=", 2)
				case strings.Contains(row, ":"):
					seg = strings.SplitN(row, ":", 2)
				}
				if len(seg) > 0 {
					(*m)[strings.TrimSpace(seg[0])] = strings.TrimSpace(seg[1])
				}
			}
		}
	}

	return len(*m) > 0

}
