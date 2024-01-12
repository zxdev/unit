package unit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
reads and builds a map from a standard ini style file

./sample
[sample]
key1=value1
key2=value2

	var u unit.Unit
	u.Parse("./sample", "sample")
	t.Log(u)

*/

type Unit map[string]string

// Parse a unit file; path,section,...
func (m *Unit) Parse(path string, section ...string) bool {

	(*m) = Reader(path, section...)
	return len(*m) > 0

}

// Writer creates a unit file with a defined section using
// the map[string\string] to configure the key=value items
func Writer(path, section string, kv map[string]string) {

	f, _ := os.Create(path)
	defer f.Close()

	//fmt.Fprintf(f, "# %s created %s\n", filepath.Base(path), time.Now().UTC().Format(time.RFC3339)[:19])
	fmt.Fprintf(f, "[%s]\n", section)
	for k := range kv {
		fmt.Fprintf(f, "%s = %s\n", k, kv[k])
	}

}

// Reader will parse the unit file and build the key=value
// map[string]string from the stated section; multiple keys
// have their values joined as comma delimited
func Reader(path string, section ...string) map[string]string {

	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	var row string
	var extract = len(section) == 0
	var scanner = bufio.NewScanner(f)
	var m = make(map[string]string)
	for scanner.Scan() {

		// strip comments
		row = strings.TrimSpace(scanner.Text())
		if idx := strings.Index(row, "#"); idx > -1 {
			row = row[:idx]
		}

		// check length
		if len(row) == 0 {
			continue
		}

		// section
		if strings.HasPrefix(row, "[") && strings.HasSuffix(row, "]") {
			if extract = len(section) == 0; !extract {
				for i := range section {
					if extract = section[i] == row[1:len(row)-1]; extract {
						continue
					}
				}
			}
		}

		// extract key:value
		if extract {
			var seg []string
			switch {
			case strings.Contains(row, "="):
				seg = strings.SplitN(row, "=", 2)
			case strings.Contains(row, ":"):
				seg = strings.SplitN(row, ":", 2)
			}
			if len(seg) > 0 {
				seg[0] = strings.TrimSpace(seg[0])
				seg[1] = strings.TrimSpace(seg[1])
				if _, exist := m[seg[0]]; exist {
					m[seg[0]] += "," + seg[1]
				} else {
					m[seg[0]] = seg[1]
				}
				//m[strings.TrimSpace(seg[0])] = strings.TrimSpace(seg[1])
			}
		}

	}

	return m
}

// Append the key=value item under unit file section at the head
// along with a timestamp reference comment; does not validate
// or provide any assurances that the key is unique
func Append(path, section, key, value string) {

	src := path
	dst := path + ".tmp"
	section = "[" + section + "]"

	r, err := os.Open(src)
	if err == nil {

		w, err := os.Create(dst)
		if err == nil {

			var row string
			var scanner = bufio.NewScanner(r)
			for scanner.Scan() {

				row = strings.TrimSpace(scanner.Text())
				fmt.Fprintln(w, row)

				if row == section {
					// section, append key:value
					fmt.Fprintf(w, "%s=%s # %s\n",
						key, value, time.Now().UTC().Format("20060102"))
				}

			}
			w.Close()
			defer os.Rename(dst, src)
		}
		r.Close()

	}
}
