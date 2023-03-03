package utils

func GetXPRequired(level int) int64 {
	levelFloat := float64(level)
	return int64((5.0 / 6.0) * levelFloat * (2*levelFloat*levelFloat + 27*levelFloat + 91))
}
