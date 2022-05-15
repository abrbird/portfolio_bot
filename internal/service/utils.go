package service

import "math"

func AbsInt64(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetIntervals(startTimeStamp int64, endTimeStamp int64, interval int64) []int64 {
	intervalsCount := int64(0)
	difference := endTimeStamp - startTimeStamp
	if difference > 0 {
		intervalsCount = int64(math.Ceil(float64(difference) / float64(interval)))
	}

	intervals := make([]int64, intervalsCount)
	for i := int64(0); i < intervalsCount; i++ {
		intervals[i] = startTimeStamp + i*interval
	}

	return intervals
}
