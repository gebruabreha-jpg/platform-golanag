package common

func FormatRate(rate float64) (value float64, unit string) {
	switch {
	case rate >= (1024 * 1024 * 1024):
		value = rate / (1024 * 1024 * 1024)
		unit = "GB/s"
	case rate >= (1024 * 1024):
		value = rate / (1024 * 1024)
		unit = "MB/s"
	case rate >= 1024:
		value = rate / 1024
		unit = "KB/s"
	default:
		value = rate
		unit = "B/s"
	}

	return
}

