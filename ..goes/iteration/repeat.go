package iteration

// Repeat value passed 5 times
func Repeat(v string, repeatNtimes int) string {
	var repeated string
	for i := 1; i <= repeatNtimes; i++ {
		repeated += v
	}
	return repeated
}
