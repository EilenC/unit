package src

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Block chain block item
type Block struct {
	Index     int64
	TimeStamp int64
	PrevHash  string
	Hash      string

	Data string
}

// calcHash
func (b *Block) calcHash() {
	var builder strings.Builder
	builder.WriteString(fmt.Sprint(b.Index))
	builder.WriteString(strconv.FormatInt(b.TimeStamp, 10))
	builder.WriteString(b.PrevHash)
	builder.WriteString(b.Data)
	hashBytes := sha256.Sum256([]byte(builder.String()))
	b.Hash = hex.EncodeToString(hashBytes[:])
}

// GenerateBlock add chain block
func (b *Block) GenerateBlock(data string) Block {
	nb := Block{
		Index:     b.Index + 1,
		TimeStamp: time.Now().Unix(),
		PrevHash:  b.Hash,
		Data:      data,
	}
	nb.calcHash() //calc block hash
	return nb
}

// InitBlock init block
func InitBlock() Block {
	block := Block{Index: -1}
	return block.GenerateBlock("Init Block")
}
