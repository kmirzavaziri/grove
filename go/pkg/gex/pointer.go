package gex

func P[T any](v T) *T {
	return &v
}

func UintP(v uint64) *uint64 {
	return P(v)
}

func FloatP(v float64) *float64 {
	return P(v)
}
