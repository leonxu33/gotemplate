package semaphore

import "sync"

type Semaphore struct {
	c  chan int
	wg sync.WaitGroup
}

func New(n int) *Semaphore {
	s := &Semaphore{
		c:  make(chan int, n),
		wg: sync.WaitGroup{},
	}
	return s
}

func (s *Semaphore) Acquire() {
	s.c <- 0
	s.wg.Add(1)
}

func (s *Semaphore) Release() {
	<-s.c
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}

func (s *Semaphore) GetLen() int {
	return len(s.c)
}
