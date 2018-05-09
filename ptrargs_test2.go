import "testing"

func Test_ptrargs_gdt5(t *testing.T) {
	// ...
	for _, test := range tests {
		r0 := ptrargs(test.f0, test.f1, test.f2, test.f3)
		if !ObjectsAreEqual(test.e0, r0) {
			t.Errorf("Expected %q, but got %q", test.e0, r0)
		}
	}
}
