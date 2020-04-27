package handlers

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func valueFromString(str string) (float64, error) {
	str = strings.ReplaceAll(str, ".", "")
	str = strings.Replace(str, ",", ".", 1)
	v, err := strconv.ParseFloat(str, 8)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse credit card bill value")
	}
	return v, nil
}

func getStringInBetween(str string, start string, end string) (result string) {
	t1 := strings.SplitAfterN(strings.TrimSpace(str), start, 2)
	if len(t1) < 2 {
		return ""
	}
	t1 = strings.SplitN(strings.TrimSpace(t1[1]), end, 2)
	if len(t1) < 1 {
		return ""
	}
	return strings.TrimSpace(t1[0])
}

type Transaction struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}
