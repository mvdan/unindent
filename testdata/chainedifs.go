package foo

func IfIf() {
	if cond {
		if cond2 {
			println(num)
		}
	}
}

func ElseIf() {
	if cond {
	} else {
		if cond2 {
			println(num)
		}
	}
}

func ElseNoIf() {
	if cond {
		if cond2 {
			println(num)
		}
	} else {
	}
}

func ElseIfSymmetry() {
	if cond {
		if !cond2 {
			println(num + 1)
		}
	} else {
		if cond2 {
			println(num)
		}
	}
}
