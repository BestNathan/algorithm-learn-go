package concurrency

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

// https://leetcode-cn.com/problems/building-h2o/
// 现在有两种线程，氢 oxygen 和氧 hydrogen，你的目标是组织这两种线程来产生水分子。

// 存在一个屏障（barrier）使得每个线程必须等候直到一个完整水分子能够被产生出来。

// 氢和氧线程会被分别给予 releaseHydrogen 和 releaseOxygen 方法来允许它们突破屏障。

// 这些线程应该三三成组突破屏障并能立即组合产生一个水分子。

// 你必须保证产生一个水分子所需线程的结合必须发生在下一个水分子产生之前。

// 换句话说:

// 如果一个氧线程到达屏障时没有氢线程到达，它必须等候直到两个氢线程到达。
// 如果一个氢线程到达屏障时没有其它线程到达，它必须等候直到一个氧线程和另一个氢线程到达。
// 书写满足这些限制条件的氢、氧线程同步代码。

//  

// 示例 1:

// 输入: "HOH"
// 输出: "HHO"
// 解释: "HOH" 和 "OHH" 依然都是有效解。
// 示例 2:

// 输入: "OOHHHH"
// 输出: "HHOHHO"
// 解释: "HOHHHO", "OHHHHO", "HHOHOH", "HOHHOH", "OHHHOH", "HHOOHH", "HOHOHH" 和 "OHHOHH" 依然都是有效解。
//  

// 限制条件:

// 输入字符串的总长将会是 3n, 1 ≤ n ≤ 50；
// 输入字符串中的 “H” 总数将会是 2n；
// 输入字符串中的 “O” 总数将会是 n。

func BuildH2O() {
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Printf("h2o print error: %v\n", err)
		}
	}()

	for _, s := range []string{"OOHHHH", "OHHOHH", "OOOOHHHHHHHH", "HOHOHOHHH", "OOOOOOOOOOOOOOOOOOOOOOOOOHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH"} {
		h2oGenerator(s)
	}

	fmt.Println()
	h2oGenerator("OH")
}

type H2OPrinter struct {
	c  *sync.Cond
	hc int
	oc int
}

func (hp *H2OPrinter) printH() {
	hp.c.L.Lock()
	for hp.hc == 2 {
		hp.c.Wait()
	}
	fmt.Print("H")
	hp.hc++
	if hp.hc == 2 {
		hp.oc = 0
	}
	hp.c.L.Unlock()
	hp.c.Broadcast()
}

func (hp *H2OPrinter) printO() {
	hp.c.L.Lock()
	for hp.oc == 1 {
		hp.c.Wait()
	}
	fmt.Print("O")
	hp.oc++
	hp.hc = 0
	hp.c.L.Unlock()
	hp.c.Broadcast()
}

func h2oGenerator(s string) {
	fmt.Printf("h2o generate with source string: %s\n", s)

	hcount := strings.Count(s, "H")
	ocount := strings.Count(s, "O")

	if hcount != 2*ocount {
		panic(fmt.Sprintf("invalid input: `H` count is %d, `O` count is %d, count of `H` should be doubled to count of `O`", hcount, ocount))
	}

	r := strings.NewReader(s)
	ch := make(chan byte, len(s))

	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		ch <- b
	}
	close(ch)
	h2oPrinter(ch)
}

func h2oPrinter(ch chan byte) {
	c := sync.NewCond(&sync.Mutex{})
	h2oP := H2OPrinter{
		c:  c,
		hc: 0,
		oc: 1,
	}

	var wg sync.WaitGroup

	fmt.Print("h2o generate result: ")
	for b := range ch {
		wg.Add(1)
		go func(s string) {
			switch s {
			case "H":
				h2oP.printH()
			case "O":
				h2oP.printO()
			}
			wg.Done()
		}(string(b))
	}

	wg.Wait()
	fmt.Println()
}
