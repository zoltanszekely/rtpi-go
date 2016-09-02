package types

type RealTimeBusInformationResult struct {
	Route       *string
	Origin      *string
	Destination *string
	DueTime     *string
}

type RealTimeBusInformation struct {
	Result
	StopID  *string
	Results *[]RealTimeBusInformationResult
}

func (obj *RealTimeBusInformation) parseResult(data *interface{}) (RealTimeBusInformationResult, error) {
	m := (*data).(map[string]interface{})
	r := RealTimeBusInformationResult{}

	var v interface{}
	var found bool

	v, found = m["route"]
	if found {
		s, err := toString(v)
		if err != nil {
			return r, err
		}
		r.Route = &s
	}

	v, found = m["origin"]
	if found {
		s, err := toString(v)
		if err != nil {
			return r, err
		}
		r.Origin = &s
	}

	v, found = m["destination"]
	if found {
		s, err := toString(v)
		if err != nil {
			return r, err
		}
		r.Destination = &s
	}

	v, found = m["duetime"]
	if found {
		s, err := toString(v)
		if err != nil {
			return r, err
		}
		r.DueTime = &s
	}

	return r, nil
}

func (obj *RealTimeBusInformation) parse(data *interface{}) error {
	obj.Result.parse(data)

	m := (*data).(map[string]interface{})

	var v interface{}
	var found bool

	v, found = m["stopid"]
	if found {
		s, err := toString(v)
		if err != nil {
			return err
		}
		obj.StopID = &s
	}

	v, found = m["results"]
	if found {
		results, err := toSlice(v)
		if err != nil {
			return err
		}
		slice := make([]RealTimeBusInformationResult, len(results))
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

func NewRealTimeBusInformation(data *interface{}) (*RealTimeBusInformation, error) {
	obj := RealTimeBusInformation{}
	err := obj.parse(data)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
