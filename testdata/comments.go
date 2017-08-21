package foo

func IfIfInside() {
	if cond {
		if cond2 {
			// comment inside
			println(num)
		}
	}
}

func IfIfBefore() {
	if cond {
		// comment before
		if cond2 {
			println(num)
		}
	}
}

func IfIfAfter() {
	if cond {
		if cond2 {
			println(num)
		}
		// comment after
	}
}
