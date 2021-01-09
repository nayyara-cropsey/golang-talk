package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Log string

type LogShipper interface {
	Start(ctx context.Context) (chan<- Log, error)
	Done() <-chan struct{}
	Err() error
}

func NewLogShipper() LogShipper {
	return &stdoutLogShipper{
		done: make(chan struct{}),
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	l := NewLogShipper()
	logCh, err := l.Start(ctx)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			logCh <- "hello"
		}
		fmt.Println("I'm done")
	}()

	time.Sleep(time.Second)
	cancel()

	<-l.Done()
	if l.Err() != nil {
		fmt.Println("error occurred:", l.Err())
	}
}

type stdoutLogShipper struct {
	done chan struct{}
	err  error
	file *os.File
}

func (s *stdoutLogShipper) Start(ctx context.Context) (chan<- Log, error) {
	ch := make(chan Log)
	file, err := os.Create("./test.txt")
	if err != nil {
		return nil, err
	}

	s.file = file
	go func() {
		for {
			select {
			case <-ctx.Done():
				goto done
			case data, ok := <-ch:
				if !ok {
					goto done
				} else {
					_, err := s.file.Write([]byte(fmt.Sprintf("%v\n", data)))
					if err != nil {
						fmt.Println("error", err)
					}
				}
			}
		}

	done:
		s.err = s.file.Close()
		close(s.done)
	}()

	return ch, nil
}

func (s *stdoutLogShipper) Done() <-chan struct{} {
	return s.done
}

func (s *stdoutLogShipper) Err() error {
	return s.err
}
