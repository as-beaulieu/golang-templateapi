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

#Contents

## V1

Demonstrates hashing and setting up of a basic blockchain with no validation or proof of work

## V2 

Demonstrates V1 plus a basic proof of work and a validation before signing the hash to a new block

## V3

Demonstrates V2 plus adding a CLI and persistence to the blockchain

v2 still has new blocks hardcoded inside the application - CLI will allow users to try to input new ones

v2 loses the blockchain after the application closes - DB will add persistence

- Bitcoin and other cryptocurrencies use LevelDB - very low level key value store

    - Bitcoin core specification: Two main groups of data
    
        - Blocks
        
            - Stored with Metadata which describes all blocks on the chain
        
        - Chain State
        
            - Stores the state of a chain and all current unspent transaction outputs, with some metadata
            
    - Bitcoin specifications has each block be it's own file on the DB
    
        - For Performance: With each on it's own file, don't have to open up multiple
        blocks just to read one

- BadgerDB - key value store that is in native Go

    - Only accepts bytes or slices of bytes
    
## V4

adds transactions to V3

#Blockchain

before blockchain, there were other distributed options, but they all relied on **trust**

meaning that every piece of new data coming in had to have the correct data coming in

In blockchain, one of the new nodes coming in could have incorrect data, say 49% of nodes coming in
could have incorrect data, and the database can correct itself

Can use blockchain for a cryptocurrency, or for smart contracts

A real blockchain will compare hashes, and see if they have been changed

Secure element of a blockchain, is that for you to change a hash, you'll have to change the previous blocks as well for it to validate

## Wallets

## Proof of Work

### consensus algorithms 

#### proof of work

computational power

common in blockchain mining - powering the network

also goes hand in hand with validation

work must be hard to do, but proof of this work should be relatively easy

#### proof of stake

## Merkel Trees