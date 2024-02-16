package ask

import "strconv"

func Int(format string, args ...interface{}) (int, error) {
	val, err := str(format, args)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(val)
}
