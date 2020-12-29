package file

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

//将字节写到文件中
func WriteStringToFile() {
	file, err := os.Create("source/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := file.WriteString("Hello GO")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

//将字符串一行一行写入文件中
func WriteByteToFile() {
	f, err := os.Create("source/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(n2, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

//将字符串一行行写入文件
func WriteStringToFileByLine() {
	f, err := os.Create("source/lines.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	d := []string{"Welcom to the world of Go1.", "Go is a compiled language.", "It is easu to learn Go."}
	for _, v := range d {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
}

//追加到文件
func AddStringToFile() {
	f, err := os.OpenFile("source/lines.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handing is easy"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}

//并发写入文件
//1、创建一个channel用来读和写这个随机数
//2、创建100个生产者goroutine.每个goroutine将产生随机数并将随机数写入到channel里
//3、创建一个消费者goroutine用来从channel读取随机数并将它写入文件，这样的话我们就只有一个goroutine向文件中写数据，从而避免竞争条件
//4、一旦完成则关闭文件
func ConcurrentWriteToFile() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}

//产生随机数；并将data写入到channel中，之后通过调用waitGroup的Done方法来通知任务已经完成
func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

//将数据写入文件
func consume(data chan int, done chan bool) {
	f, err := os.Create("source/concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}
