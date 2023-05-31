package helpers

func ToFloat32(a *float32) float32 {
	if a != nil {
		return *a
	}
	return 0
}

func ToFloat64(a *float64) float64 {
	if a != nil {
		return *a
	}
	return 0
}
