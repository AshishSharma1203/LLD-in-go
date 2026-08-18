package main

type Payment interface{
	Pay(amount float32)
}
