package foo

func IfErrNotNil() {
	n, err := fnErr()
	if err != nil {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
		println(err.Error())
		return
	}
	println(n)
}

func IfErrIsNil() {
	n, err := fnErr()
	if err == nil {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
		println(n)
		return
	}
	println(err.Error())
}

func IfObjNotNil() {
	n, err := fnErr()
	if n != nil {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
		println(n)
		return
	}
	println(err.Error())
}

func IfErrOther() {
	n, err := fnErr()
	if err != errVar {
		for i := 0; i < 10; i++ {
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
			num++; println(num); num++; println(num)
		}
		println(err.Error())
		return
	}
	println(n)
}
