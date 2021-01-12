package main

import (
	"context"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("Arch\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Max procs\t", runtime.GOMAXPROCS(runtime.NumCPU()))

	//monitorCh := make(chan bool)
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go log("foo", ch, ctx, wg)
	go log("bar", ch, ctx, wg)

	go func() {
		for {
			ch <- randomdata.SillyName()
			waitRand()
		}
	}()

	/*go func() {
		for {
			fmt.Println("Goroutines\t", runtime.NumGoroutine())
			g := runtime.NumGoroutine()
			if g == 2 {
				monitorCh <- true
			}
			waitRand()
		}
	}()*/

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan bool)
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	for {
		select {
		case <-quit:
			cancel()
		case <-doneCh:
			fmt.Println("all loggers done")
			return
		}
	}

	//<-monitorCh

	fmt.Println("exiting")
	time.Sleep(time.Second)
}

func log(name string, ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	file, err := os.Create(fmt.Sprintf("./%v.log", name))
	if err != nil {
		fmt.Println("file creation error: ", name, err)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("exiting", name)
			err := file.Close()
			if err != nil {
				fmt.Println("file closing error: ", name, err)
			}
			waitRandSec()
			wg.Done()
			fmt.Println("closed file", name)
			return
		case msg := <-ch:
			_, err := file.Write([]byte(fmt.Sprintf("%v\n", msg)))
			if err != nil {
				fmt.Println("message sending error: ", name, err)
			}
		}
	}
}

func fanOut(in chan string, ch1, ch2 chan string) {
	for {
		select {
		case msg := <-in:
			ch1 <- msg
			ch2 <- msg
		}
	}
}

func fanIn(out chan string, ch1, ch2 chan string) {
	for {
		select {
		case msg := <-ch1:
			out <- msg
		case msg := <-ch2:
			out <- msg
		}
	}
}

func waitRand() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func waitRandSec() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}
