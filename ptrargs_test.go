import (
	"errors"
	"testing"
)

func Test_ptrargs_gdt5(t *testing.T) {
	ptrData := []struct {
		f0 string // HL
		f2 int    // HL
	}{
		{f2: 13},
		{f0: ""},
	}
	tests := []struct {
		f0 *string // HL
		f1 string
		f2 *int // HL
		f3 float32
		e0 error
	}{
		{nil, "", &ptrData[0].f2, 42.0, errors.New("pstr is nil")},
		{&ptrData[1].f0, "x", nil, 42.0, nil},
	}
	// ...
}
