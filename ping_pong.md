# Ping Pong

A big reason people love Go is the concurrency primitives. Let's try them out
with a coding challenge.

### Goal_
Launch 2 go-routines that send messages to each other over a channel. One
go-routine will send "Ping" and the other will send "Pong." Each go-routine
can only respond after they've received a message from the other go-routine.
The exception is that the go-routine sending "Ping" can send an initial message
un-prompted.

### Cheat Sheet

__create a channel__
```go
unbuffered_channel = make(chan string)

buffered_channel = make(chan string, 5)
```

__write to channel__
```go
muffin_chan = make(chan string, 10)
msg := "muffins"

muffin_chan <- msg
for {
  // will block go-routine if channel buffer is full
  muffin_chan <- "infinite muffins"
}

```

__launch a go-routine w/ a chan__
```go
func AsyncPrinter(msg_stream chan string) {
  var msg string
  for {
    msg <- msg_stream
    fmt.Println(msg)
  }
}

func main() {
  muffin_stream := make(chan string, 100)
  go AsyncPrinter(muffin_stream)

  for {
    muffin_stream <- "MUFFINS!!!"
  }
}
```

__sleep a go-routine for X millis__
```go
import (
  "fmt"
  "time"
)

func main() {
  for {
    time.Sleep(100) // 100ms
    fmt.Println("witness my delay...")
  }
}
```

__"hand off" main thread to go-routines__
```go
import (
  "fmt"
  "runtime"
  "time"
)

func IRunForever() {
  for {
    time.Sleep(100)
    fmt.Println("running...")
  }
}

func main() {
  go IRunForever()

  runtime.Goexit()
}
