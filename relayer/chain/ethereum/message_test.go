package ethereum_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	gethTrie "github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/snowfork/snowbridge/relayer/chain/ethereum"
	"github.com/stretchr/testify/assert"

	"github.com/snowfork/snowbridge/relayer/chain/parachain"
)

type TestProof parachain.ProofData

// For interface gethTrie.KeyValueReader
func (tp *TestProof) Get(key []byte) ([]byte, error) {
	for i, k := range tp.Keys {
		if bytes.Equal(k, key) {
			return tp.Values[i], nil
		}
	}
	return nil, fmt.Errorf("Value for key %s does not exist", key)
}

// For interface gethTrie.KeyValueReader
func (tp *TestProof) Has(key []byte) (bool, error) {
	_, err := tp.Get(key)
	return err == nil, nil
}

func TestMessage_Proof(t *testing.T) {
	block := block11408438()
	receipts := receipts11408438()
	// We'll prove inclusion for this event by proving inclusion for
	// the encapsulating receipt

	event5_5 := receipts[5].Logs[5]

	receipt5Encoded, err := rlp.EncodeToBytes(receipts[5])
	if err != nil {
		panic(err)
	}

	// Construct Merkle Patricia Trie for receipts
	receiptTrie := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))

	types.DeriveSha(receipts, receiptTrie)

	fmt.Println("Hash", receiptTrie.Hash())

	if receiptTrie.Hash() != block.ReceiptHash() {
		panic("Receipt trie does not match block receipt hash")
	}

	msg, err := ethereum.MakeMessageFromEvent(event5_5, receiptTrie)
	assert.Nil(t, err)
	assert.NotNil(t, msg)

	key, err := rlp.EncodeToBytes(uint(5))
	if err != nil {
		panic(err)
	}
	proofNodes := TestProof(*msg.Proof.ReceiptProof)
	provenReceipt, err := gethTrie.VerifyProof(block.ReceiptHash(), key, &proofNodes)
	assert.Nil(t, err)
	assert.Equal(t, provenReceipt, receipt5Encoded)
}
