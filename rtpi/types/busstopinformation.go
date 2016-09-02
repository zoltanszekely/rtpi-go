package types

type BusStopInformationResult struct {
	StopID    *string
	ShortName *string
	FullName  *string
}

type BusStopInformation struct {
	Result
	Results *[]BusStopInformationResult
}

func (obj *BusStopInformation) parseResult(data *interface{}) (BusStopInformationResult, error) {
	m := (*data).(map[string]interface{})
	r := BusStopInformationResult{}

	var v interface{}
	var found bool

	v, found = m["stopid"]
	if found {
		s, err := toString(v)
		if err != nil {
			return r, err
		}
		r.StopID = &s
	}

	v, found = m["shortname"]
	if found {
		s, err := toString(v)
		if err != nil {
			return r, err
		}
		r.ShortName = &s
	}

	v, found = m["fullname"]
	if found {
		s, err := toString(v)
		if err != nil {
			return r, err
		}
		r.FullName = &s
	}

	return r, nil
}

func (obj *BusStopInformation) parse(data *interface{}) error {
	obj.Result.parse(data)

	m := (*data).(map[string]interface{})

	var v interface{}
	var found bool

	v, found = m["results"]
	if found {
		results, err := toSlice(v)
		if err != nil {
			return err
		}
		slice := make([]BusStopInformationResult, len(results))
		obj.Results = &slice
		for i, result := range results {
			r, err := obj.parseResult(&result)
			if err != nil {
				return err
			}
			(*obj.Results)[i] = r
		}
	}
	return nil
}

func NewBusStopInformation(data *interface{}) (*BusStopInformation, error) {
	obj := BusStopInformation{}
	err := obj.parse(data)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
