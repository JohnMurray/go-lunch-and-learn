# Hello World

Open your favorite editor and create the following file:

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello, world!")
}
```

Save your changes as `hello.go` and now let's build and run our first go application:

```bash
go build -o hello hello.go
```

And to run:

```bash
./hello
```


## Cheating

We can shortcut the compile + run by using the `go run` command which essentially does this in
one step (although doesn't create a binary for you)

```bash
go run hello.go
```

For the rest of the lunch-and-learn, you should cheat. It's faster.
