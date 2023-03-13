package src

import (
	"fmt"
	"log"
	"time"
)

// Chain blocks
type Chain struct {
	Blocks []*Block
}

// InitChain init new chain
func InitChain() *Chain {
	nb := InitBlock()
	c := new(Chain)
	c.Append(&nb)
	return c
}

// Display show blocks
func (c *Chain) Display() {
	for _, b := range c.Blocks {
		fmt.Printf("Index: %d\n", b.Index)
		fmt.Printf("PrevHash: %s\n", b.PrevHash)
		fmt.Printf("Hash: %s\n", b.Hash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Time: %s\n\n", time.Unix(b.TimeStamp, 0).Format("2006.01.02 15:04:05"))
	}
}

// SendData send data
func (c *Chain) SendData(data string) {
	nb := c.Blocks[len(c.Blocks)-1].GenerateBlock(data)
	c.Append(&nb)
}

// Append item to blocks
func (c *Chain) Append(nb *Block) {
	if len(c.Blocks) == 0 || validate(*nb, *c.Blocks[len(c.Blocks)-1]) {
		c.Blocks = append(c.Blocks, nb)
		return
	}
	log.Fatalln("invalid block")
}

// validate block
func validate(nb, ob Block) bool {
	if nb.Index-1 != ob.Index {
		return false
	}
	if nb.PrevHash != ob.Hash {
		return false
	}
	if nb.Hash == "" {
		return false
	}
	return true
}
