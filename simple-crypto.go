package simple_crypto

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// from https://hackernoon.com/learn-blockchains-by-building-one-117428612f46
// playing with how a simple blockchain works...

type Transaction struct {
	Sender    string `json:"sender" yaml:"sender" bson:"sender"`
	Recipient string `json:"recipient" yaml:"recipient" bson:"recipient"`
	Amount    uint64 `json:"amount" yaml:"amount" bson:"amount"`
}

type Block struct {
	Index        uint64        `json:"index" yaml:"index" bson:"index"`
	Timestamp    int64         `json:"timestamp" yaml:"timestamp" bson:"timestamp"`
	Transactions []Transaction `json:"transactions" yaml:"transactions" bson:"transactions"`
	Proof        uint64        `json:"proof" yaml:"proof" bson:"proof"`
	PreviousHash string        `json:"previous_hash" yaml:"previous_hash" bson:"previous_hash"`
}

var current_transactions []Transaction
var chain []Block
var node_instance string = uuid.New().String()

func init() {
	// create the genesis block
	new_block(1, "100")
}

/**
 * returns the last item in the chain
 */
func Last_Block() Block {
	return chain[len(chain)-1]
}

/**
 * Adds a new transaction to the next mined block
 */
func New_Transaction(sender string, recipent string, amount uint64) uint64 {
	current_transactions = append(current_transactions, Transaction{
		Sender:    sender,
		Recipient: recipent,
		Amount:    amount,
	})
	return Last_Block().Index + 1
}

/**
 * Creates a new block in the block chain
 */
func new_block(proof uint64, previous_hash string) Block {
	blk := Block{
		Index:        uint64(len(chain) + 1),
		Timestamp:    time.Now().UTC().Unix(),
		Transactions: current_transactions,
		Proof:        proof,
		PreviousHash: previous_hash,
	}
	if previous_hash == "" {
		hash := asSha256(Last_Block())
		blk.PreviousHash = hash
	}

	// reset the array
	current_transactions = current_transactions[:0]

	chain = append(chain, blk)

	return blk
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
 * proof of work function
 * Find a number p' such that hash(pp') contains leading 4 zeroes, where p is the previous p'
 * - p is the previous proof, and p' is the new proof
 */
func Proof_of_work(last_proof uint64) uint64 {
	nextProof := uint64(0)
	for ok := false; !ok; ok = Valid_proof(last_proof, nextProof) {
		nextProof++
	}
	return nextProof
}

/**
 * Validates the proof
 * in this case, does it contain 4 leading zeros?
 */
func Valid_proof(last_proof uint64, proof uint64) bool {
	return strings.HasPrefix(asSha256((strconv.FormatUint(last_proof, 10) + strconv.FormatUint(proof, 10))), "0000")
}

func Mine() Block {
	last_proof := Last_Block().Proof
	proof := Proof_of_work(last_proof)
	New_Transaction(
		"0",
		node_instance,
		1,
	)

	return new_block(proof, "")
}
