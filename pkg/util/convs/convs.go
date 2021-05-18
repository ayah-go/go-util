package convs

import "strconv"

func String2Int64(source string) int64 {
	parseInt, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		return 0
	}
	return parseInt
}
