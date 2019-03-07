package utils

import (
	"net/url"
	"strconv"
	"strings"
)

type UtilUrl struct{}

func (UtilUrl) HandelResQuery(q url.Values) ([]string, int) {
	var count int
	if q["time"] == nil {
		return nil, 7
	}
	if q["list"] == nil {
		return nil, 7
	}
	listNames := strings.Split(q["list"][0], ",")
	count, _ = strconv.Atoi(q["time"][0])
	return listNames, count
}
