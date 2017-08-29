package foo

func IfIf() {
	if cond {
		if cond2 == cond3 {
			println(num)
		}
	}
}

func IfIfWouldNeedParen() {
	if cond {
		if cond2 || cond3 {
			println(num)
		}
	}
}

func IfIfNoParen() {
	if (cond || num == 0) {
		if cond2 && cond3 {
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

func IfIfInList() {
	if num == 0 {
		println("foo")
	}
	if cond {
		if cond2 == cond3 {
			println(num)
		}
	}
}

func IfAndAction() {
	if cond {
		if action() {
			println(num)
		}
	}
}

func IfAndBuiltin() {
	if cond {
		if len(slice) > 0 {
			println(num)
		}
	}
}

func IfAndConversion() {
	if cond {
		if uint(num) == 4 {
			println(num)
		}
	}
}
