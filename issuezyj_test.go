package mergo_test

import (
	"github.com/imdario/mergo"
	"testing"
)

type data struct {
	B *bool
	I *int
	S *string
}

func TestIssueZyj(t *testing.T) {
	vb := true
	vi := 12
	vs := "34"

	vi2 := 324

	d1 := data{
		B: &vb,
		I: &vi,
	}

	d2 := data{
		I: &vi2,
		S: &vs,
	}

	err := mergo.Merge(&d1, d2, mergo.WithOverride)
	if err != nil {
		t.Errorf("Error while merging %s", err)
	}

	if *d1.B != vb {
		t.Error("vb failed")
	}

	if *d1.I != vi2 {
		t.Error("vi failed")
	}

	if *d1.S != vs {
		t.Error("vs failed")
	}
}
