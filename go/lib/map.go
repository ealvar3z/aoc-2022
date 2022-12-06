package lib

func SetMissing(val map[string]int64, index string, _default int64) {
	if _, ok := val[index]; !ok {
		val[index] = _default
	}
}
