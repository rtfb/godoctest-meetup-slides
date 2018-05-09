import "strings"

func varargs(strs ...string) string {
	/*
		[]test{
			{[]string{}, ""},
			{[]string{"a"}, "a"},
			{[]string{"a", "b"}, "a b"},
		}
	*/
	return strings.Join(strs, " ")
}
