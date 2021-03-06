package blockV2

import (
	"TemplateApi/examples/blockchain/common"
	"bytes"
)

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			common.ToHex(int64(nonce)),
			common.ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}
