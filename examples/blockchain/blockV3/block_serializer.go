package blockV3

import (
	"bytes"
	"encoding/gob"
	"log"
)

func (b *Block) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	if err := encoder.Encode(b); err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()

}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	if err := decoder.Decode(&block); err != nil {
		log.Panic(err)
	}

	return &block
}
