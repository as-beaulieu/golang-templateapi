package blockV5

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	if err := encode.Encode(tx); err != nil {
		log.Panic(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

//Coinbase - the transactions evolved around the genesis block, since there is no previous block or transactions
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data} //-1 because referencing no output
	txout := TxOutput{100, to}          //value here is the reward for the miner to complete the genesis block PoW

	//So when we create this, and an account named Jack mines this block, then they are rewarded for mining the block

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

func (tx *Transaction) IsCoinbase() bool {
	onlyOneInput := len(tx.Inputs) == 1
	noInputIDs := len(tx.Inputs[0].ID) == 0
	outputIndexNegOne := tx.Inputs[0].Out == -1
	return onlyOneInput && noInputIDs && outputIndexNegOne
}

type TxOutput struct {
	Value  int    //Value in tokens
	PubKey string //A value that is needed to unlock the tokens in the value field
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

type TxInput struct { //references to previous outputs
	ID  []byte //transaction that the output is inside of
	Out int
	Sig string
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("Error: insufficient funds for transaction")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs { //creating an input for each of the unspent transactions
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to}) //Where user is sending X tokens to an address
	if acc > amount {                               //Created if there are any leftover tokens in the sender's account
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
