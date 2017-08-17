# unindent

	go get -u mvdan.cc/unindent

Reports code that is unnecessarily indented. Examples include:

```
for _, elem := range list {
	if cond {
		// here be many lines
	}
}
```

```
if cond1 {
	if cond2 {
		// here be many lines
	}
}
```

```
if cond1 {
} else {
	if cond2 {
		// here be many lines
	}
}
```
