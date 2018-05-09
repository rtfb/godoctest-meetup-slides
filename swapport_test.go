package foo

import "testing"

func TestSwapPort(t *testing.T) {
	tests := []struct {
		h string // HLstruct
		p string // HLstruct
		e string // HLstruct
	}{
		{h: "", p: "666", e: "localhost:666"}, // HLinteresting
		{h: "foo", p: "8080", e: "foo:8080"},  // HLinteresting
		{h: "bar:80", p: "81", e: "bar:81"},   // HLinteresting
	}
	for _, test := range tests {
		actual := swapPort(test.h, test.p)
		if actual != test.e {
			t.Errorf("Expected %q, but got %q", test.e, actual)
		}
	}
}
