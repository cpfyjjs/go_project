package fib

import (
	"bufio"
	"fmt"
	"learngo/errorhanding/fib"
	"os"
)

func tryDefer(){
	for i:=0;i<100;i++{
		defer fmt.Println(i)
		if i >= 30{
			panic("too many")
		}
	}
}


func writeFile(filename string){
	file,err := os.Create(filename)
	//os.OpenFile(filename,os.O_CREATE|os.O_EXCL,0666)
	if err != nil{
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	f  := fib.Fibonacci()
	for i:=0;i<20;i++{
		fmt.Fprintln(writer,f())
	}
	writer.Flush()

}

func main() {
	writeFile("fib.txt")
	tryDefer()
}
