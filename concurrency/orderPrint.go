package concurrency

import (
	"fmt"
	"sync"
)

// https://leetcode-cn.com/problems/print-in-order/

// public class Foo {
//   public void one() { print("one"); }
//   public void two() { print("two"); }
//   public void three() { print("three"); }
// }


// random order print
type randOrderPrinter struct {
	
}

func (p *randOrderPrinter) first() {
	fmt.Print("first;")
}

func (p *randOrderPrinter) second() {
	fmt.Print("second;")
}

func (p *randOrderPrinter) third() {
	fmt.Print("third;")
}

func RandOrderPrint()  {
	for index := 0; index < 10; index++ {
		randOrderRunner()
		fmt.Println()
	}
}

func randOrderRunner()  {
	var p randOrderPrinter
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		p.first()
		wg.Done()
	}()

	go func() {
		p.second()
		wg.Done()
	}()

	go func() {
		p.third()
		wg.Done()
	}()

	wg.Wait()
}


// order print
type orderPrinter struct {
	cl *sync.Cond
	i int
}

func (p *orderPrinter) first() {
	p.cl.L.Lock()
	for p.i != 0 {
		p.cl.Wait()
	}
	fmt.Print("first;")
	p.i++
	p.cl.L.Unlock()
	p.cl.Broadcast()
}

func (p *orderPrinter) second() {
	p.cl.L.Lock()
	for p.i != 1 {
		p.cl.Wait()
	}
	fmt.Print("second;")
	p.i++
	p.cl.L.Unlock()
	p.cl.Broadcast()
}

func (p *orderPrinter) third() {
	p.cl.L.Lock()
	for p.i != 2 {
		p.cl.Wait()
	}
	fmt.Print("third;")
	p.cl.L.Unlock()
	p.cl.Broadcast()
}

func OrderPrint()  {
	for index := 0; index < 10; index++ {
		orderRunner()
		fmt.Println()
	}
}

func orderRunner()  {
	cl := sync.NewCond(&sync.Mutex{})

	op := orderPrinter{
		cl: cl,
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		op.first()
		wg.Done()
	}()

	go func() {
		op.second()
		wg.Done()
	}()

	go func() {
		op.third()
		wg.Done()
	}()

	wg.Wait()
}