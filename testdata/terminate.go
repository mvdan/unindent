package foo

func NonTerminatingEmpty() {
	if cond {
	}
	num--; println(num); num--; println(num)
}

func NonTerminating() {
	if cond {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
	}
	num--; println(num); num--; println(num)
}

func Terminating() {
	if cond {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
		return
	}
	num--; println(num)
}

