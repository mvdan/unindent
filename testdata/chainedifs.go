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
