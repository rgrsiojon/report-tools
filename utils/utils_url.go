package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UtilUrl struct{}

func (UtilUrl) HandelResQuery(c *gin.Context) ([]string, int) {
	var count int
	q := c.Request.URL.Query()
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
