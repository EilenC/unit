package main

import "chain/src"

func main() {
	c := src.InitChain()
	c.SendData("Send 1 BTC")
	c.SendData("Send 2 BTC")
	c.SendData("Send 3 BTC")
	c.Display()
}
