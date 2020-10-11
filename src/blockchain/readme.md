to create this as a standalone project with its own go mod

`go mod init github.com/_project_subject_/_project_name_`

creates a go.mod file, declares a `module` with module name

to import a package with go.mod

```go
import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
```

how the import looks in the go.mod

```
module github.com/_project_subject_/_project_name_

require (
    rsc.io/quote v1.5.2
)
```

also creates a go.sum file - similar to a yarn.lock file

#Blockchain origin

before blockchain, there were other distributed options, but they all relied on **trust**

meaning that every piece of new data coming in had to have the correct data coming in

In blockchain, one of the new nodes coming in could have incorrect data, say 49% of nodes coming in
could have incorrect data, and the database can correct itself

Can use blockchain for a cryptocurrency, or for smart contracts