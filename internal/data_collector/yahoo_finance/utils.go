package yahoo_finance

import (
	"fmt"
	"math"
	"time"
)

const Interval1M = "1m"
const Interval5M = "5m"
const Interval15M = "15m"
const Interval30M = "30m"
const Interval1H = "1h"
const Interval3H = "3h"
const Interval6H = "6h"
const Interval12H = "12h"
const Interval1D = "1d"
const Interval1WK = "1wk"
const Interval1MO = "1mo"

const Range1D = "1d"
const Range5D = "5d"
const Range1MO = "1mo"
const Range3MO = "3mo"
const Range6MO = "6mo"
const Range1Y = "1y"
const Range5Y = "5y"
const RangeMAX = "max"

type DaysRange struct {
	days        float64
	rangeString string
}

var rangeItems = []DaysRange{
	{1, Range1D},
	{5, Range5D},
	{30, Range1MO},
	{30 * 3, Range3MO},
	{30 * 6, Range6MO},
	{365, Range1Y},
	{365 * 5, Range5Y},
}

var intervalMap = map[uint64]string{
	60:                Interval1M,
	60 * 5:            Interval5M,
	60 * 15:           Interval15M,
	60 * 30:           Interval30M,
	60 * 60:           Interval1H,
	60 * 60 * 3:       Interval3H,
	60 * 60 * 6:       Interval6H,
	60 * 60 * 12:      Interval12H,
	60 * 60 * 24:      Interval1D,
	60 * 60 * 24 * 7:  Interval1WK,
	60 * 60 * 24 * 30: Interval1MO,
}

func GetRange(tsStart int64, tsEnd int64) string {
	tsStart_ := time.Unix(tsStart, 0).UTC()
	tsEnd_ := time.Unix(tsEnd, 0).UTC()
	differenceInDays := math.Ceil(tsEnd_.Sub(tsStart_).Hours() / 24)

	for _, rI := range rangeItems {
		if differenceInDays < rI.days {
			return rI.rangeString
		}
	}
	return RangeMAX
}

func GetInterval(seconds uint64) (string, error) {
	if interval, ok := intervalMap[seconds]; ok {
		return interval, nil
	}
	return "", fmt.Errorf("can not find interval for %d seconds", seconds)
}
