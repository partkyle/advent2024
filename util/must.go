package util

func Must[E any](e E, err error) E {
	if err != nil {
		panic(err)
	}
	return e
}
