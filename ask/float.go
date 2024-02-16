package ask

import "strconv"

func Float(format string, args ...interface{}) (float64, error) {
	val, err := str(format, args)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(val, 64)
}
