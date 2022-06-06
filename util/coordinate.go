package util

import "math"

func GetDistance(x1 int64, y1 int64, x2 int64, y2 int64) float64 {
	return math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)
}
