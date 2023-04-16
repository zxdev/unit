package unit_test

import (
	"testing"

	"github.com/zxdev/unit"
)

func TestOne(t *testing.T) {

	var u unit.Unit
	u.Parse("./sample.unit", "sample")
	t.Log(u)

}

func TestTwo(t *testing.T) {

	var u unit.Unit
	u.Parse("./sample.unit", "sample2")
	t.Log(u)

}
