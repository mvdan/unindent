package foo

func CannotEarly() {
	if cond {
		println(num)
		if cond2 {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
	}
	println(num)
}

var CanEarly = func() {
	if cond {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
	}
}
