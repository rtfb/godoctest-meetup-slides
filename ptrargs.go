import "errors"

func ptrargs(pstr *string, str string, i *int, f float32) error {
	/*
		[]test{
			{nil, "", 13, 42.0, errors.New("pstr is nil")},
			{"", "x", nil, 42.0,  nil},
		}
	*/
	if pstr == nil {
		return errors.New("pstr is nil")
	}
	*pstr = "result"
	return nil
}
