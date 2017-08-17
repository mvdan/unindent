# unindent

	go get -u mvdan.cc/unindent

Reports code that is unnecessarily indented. For example:

```
for _, elem := range list {
	if cond {
		// here
		// be
		// many
		// lines
	}
}
```

Can be rewritten as:

```
for _, elem := range list {
	if !cond {
		continue
	}
	// here
	// be
	// many
	// lines
}
```
