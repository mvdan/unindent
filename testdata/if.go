package foo

func BodyIf() {
	if cond {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
	}
}

func ScoreTooLow() {
	if cond {
		for i := 0; i < 10; i++ {
			println(num)
		}
	}
}

func IfEmpty() {
	if cond {
	}
}

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

func IfWithInit() {
	if a := "foo"; cond {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
		println(a)
	}
}
