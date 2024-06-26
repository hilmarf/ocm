package compdesc

type conversionError struct {
	error
}

func ThrowConversionError(err error) {
	panic(conversionError{err})
}

func (e conversionError) Error() string {
	return "conversion error: " + e.error.Error()
}

func CatchConversionError(errp *error) {
	if r := recover(); r != nil {
		if je, ok := r.(conversionError); ok {
			*errp = je
		} else {
			panic(r)
		}
	}
}

func Validate(desc *ComponentDescriptor) error {
	data, err := Encode(desc)
	if err != nil {
		return err
	}
	_, err = Decode(data)
	return err
}
