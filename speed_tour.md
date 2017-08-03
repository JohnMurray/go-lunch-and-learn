# Speed Tour

Originally borrowed from: [this workshop by Kelsey Hightower][workshop]

## Packages

```go
import(
   "fmt"
   "log"

   "github.com/kelseyhightower/targz"
   short "github.com/someone/with-a-really-long-package-name"
   _ "runtime/pprof"
)
```

- stdlib packages have no path
- import paths match repositories - central to dep-mgmnt in Go
- alias packages
- side-effecting imports


## Variables

```go
var (
  name     string         // package private
  Location = "Portland"   // package exported
)

func main() {
  name = "Kelsey Hightower"
  distro := "CoreOS"
  fmt.Printf("Name: %s\nLocation: %s\nDistro: %s\n", name, Location, distro)
}
```

- postfix type declarations
- exported vs non-exported names
- `=` vs `:=` declaration/assignment - tyep inference


## Arrays

```go
func main() {
  locations := [3]string{
    "Long Beach",
    "Atlanta",
    "Portland",
  } 
  fmt.Printf("Number of locations: %d", len(locations))
}
```

- fixed size, size part of type
- non-growable, like c-arrays

## Slices

```go
func main() {
  locations := []string{
    "Long Beach",
    "Atlanta",
    "Portland",
    "New York",
    "Denver",
    "Dallas",
  }
  for _, name := range locations {
    fmt.Printf("Name: %s\n", name)
  }

  middleTwo := locations[2:4]
  fmt.Printf("%#v", middleTwo)

    ints := make([]int, 0)
    for i := 0; i <= 100; i++ {
      ints = append(ints, i)
  }
    fmt.Printf("%#v", ints)
}
```

- array backed
- "view" of underlying
- length vs capacity
- still fixed capacity (by underlying)
- variable length


## Maps

```go
func main() {
	m := make(map[string]string)
	m["name"] = "Kelsey"

	name, ok := m["name"]
	if !ok {
		log.Fatal("Name does not exist")
	}
	fmt.Println(name)

  if name, ok = m["name"]; ok {
    fmt.Println(name)
  }

	for k, v := range m {
		fmt.Printf("Key: %s Value: %s", k, v)
	}
}
```

- no generics, except for map
- variable return
- flexible `if`
- range is a magical unicorn (!use w/ user-types)

## For, the only control-flow you could ever need

```go
// infinite
for {}

// while
blah := true
for blah {
  fmt.Println("runs as long as 'blah' is true")
  blah = false
}

// do-while
value := 0
for {
  value += 1
  if value > 5 {
    break
  }
}

// for over array
stuffs := [string] { "one", "two", "three" }
for i := range stuffs {
  fmt.Println(stuffs[i])
}

// for each
for _, v := range stuffs {
  fmt.Println(v)
}
```

## Functions

```go
func ping(url string) bool {
  resp, err := http.Get(string)
  if err != nil {
    return false
  }
  if resp.StatusCode != http.StatusOK {
    return false
  }
  return true
}
```

- post-fix return type
- no implicit return (not scala)
- avoid return from else-block

## Structs

```go
type Person struct {
  name     string
  Location string
}

func NewPerson(name string) *Person {
  p := Person{
    name: name,
  }
  return &p
}

func (p *Person) Name() string {
  return p.name
}

func (p *Person) SetName(name string) {
  p.name = name
}

func main() {
  p := NewPerson("Kelsey")
  fmt.Printf("Name: %s\n", p.Name())

  p.Location = "Portland"
  fmt.Printf("Location: %s\n", p.Location)
}
```

- `&` returns a pointer to an object
- `*` is a pointer type
- `NewPerson` dangling reference? NOPE! Escape analysis FTW!
- stack vs heap vs golang
- `new(Type)` is a thing - don't use it

## Channels and Goroutines

```go
func doubler(input, output chan int) {
  for {
    i := <-input
    output <- i * 2
  }
}

func printer(output chan int) {
  for {
    fmt.Printf("%d\n", <-output)
  }
}

func main() {
  input := make(chan int)
  output := make(chan int)

  go doubler(input, output)
  go printer(output)

  for i := 0; i <= 10; i++ {
    input <- i
  }
  time.Sleep(time.Duration(1) * time.Second)
}
```

- channels used between go-routines
- channels are typed
- `go` starts  new go-routine
- never sleep or busy-wait in `main` (read empty channel, or `runtime.Goexit()`)



  [workshop]: https://github.com/kelseyhightower/intro-to-go-workshop/blob/master/speed_tour.md
