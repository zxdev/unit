package unit_test

import (
	"cloudlab/pkg/unit"
	"testing"
)

func TestCreate(t *testing.T) {

	unit.Writer("./sample.unit", "sample",
		map[string]string{
			"key1": "value1",
			"key2": "value2",
		})

}

func TestRead(t *testing.T) {

	var u unit.Unit
	u.Parse("./sample.unit", "sample")
	t.Log(u)

}

func TestAppend(t *testing.T) {

	unit.Append("./sample.unit", "sample", "key", "append")

}

func TestMulti(t *testing.T) {
	var u unit.Unit
	u.Parse("./sample.unit")
	t.Log(u)
}
