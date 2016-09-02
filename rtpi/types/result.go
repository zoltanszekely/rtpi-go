package types

type Result struct {
	ErrorCode    *int
	ErrorMessage *string
}

func (result *Result) parse(data *interface{}) error {
	m := (*data).(map[string]interface{})

	var v interface{}
	var found bool

	v, found = m["errorcode"]
	if found {
		i, err := toInt(v)
		if err != nil {
			return nil
		}
		result.ErrorCode = &i
	}

	v, found = m["errormessage"]
	if found {
		s, err := toString(v)
		if err != nil {
			return nil
		}
		result.ErrorMessage = &s
	}

	return nil
}
