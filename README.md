# Go Unsafe Scanner
Go Memory Pointer Scanner using the [Unsafe](https://pkg.go.dev/unsafe) Pointer Package

<img src="images/unsafe.jpg"/>

## Quick-start

1. Run the help for all avaliable commands

```bash
go run main.go help
```

2. To see a very basic pointer editing proof-of-concept example run:

```bash
go run main.go pointertest
```

2. Start a dummy application which we will use to test the unsafe pointer scanner

```bash
go run main.go dummyapp
```

3. In a seperate terminal window, try to scan the value

```bash
go run main.go pointerscan
```