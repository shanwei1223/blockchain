package main

import "BlockChain/core"

func main() {
	bc := core.CreateNewBlockChain()
	bc.SetData("这是第一个块.")
	bc.SetData("这是第二个块.")
	bc.Print()
}