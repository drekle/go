package main

/*
https://www.hackerrank.com/challenges/coin-change/problem
*/
import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var debug = `166 23
5 37 8 39 33 17 22 32 13 7 10 35 40 2 43 49 46 19 41 1 12 11 28
`

func printArr(arr []int, delimiter string) {
	var buf bytes.Buffer
	for index, j := range arr {
		buf.WriteString(strconv.FormatInt(int64(j), 10))
		if index < len(arr)-1 {
			buf.WriteString(" " + delimiter + " ")
		}
	}
	println(buf.String())
}

type solver struct {
	wg            sync.WaitGroup
	lock          sync.Mutex
	possibilities int64
}

func (s *solver) incPossibilities() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.possibilities++
}

func solveForRemainder(coins *[]int, offset int, remainder int64, s *solver) {
	for ; offset < len(*coins); offset++ {
		//exhaust all possibilities for this coin
		if int64((*coins)[offset]) <= remainder {
			newRemainder := remainder - int64((*coins)[offset])
			for {
				if newRemainder < 0 {
					break
				} else if newRemainder == 0 {
					s.incPossibilities()
					break
				}
				s.wg.Add(1)
				solveForRemainder(coins, offset+1, newRemainder, s)
				newRemainder = newRemainder - int64((*coins)[offset])
			}
		}
	}
	s.wg.Done()
}

func main() {
	r := bytes.NewReader([]byte(debug))
	_ = r
	reader := bufio.NewReader(r)
	// reader := bufio.NewReader(os.Stdin)
	lineBytes, _, _ := reader.ReadLine()
	lineOne := strings.Split(string(lineBytes), " ")
	lineBytes, _, _ = reader.ReadLine()
	lineTwo := strings.Split(string(lineBytes), " ")

	changeDue, _ := strconv.ParseInt(lineOne[0], 10, 64)
	arrLen, _ := strconv.ParseInt(lineOne[1], 10, 64)
	//Fail Fast scenarios
	if changeDue == 0 || arrLen == 0 {
		fmt.Println(0)
		return
	}

	coins := make([]int, 0)
	for i := arrLen - 1; i >= 0; i-- {
		coin64, _ := strconv.ParseInt(lineTwo[i], 10, 64)
		coin := int(coin64)
		coins = append(coins, coin)
	}

	//Create an index for every integer less than the change due
	//  IFF a remainder exists to reach the changeDue
	s := solver{
		lock:          sync.Mutex{},
		possibilities: 0,
		wg:            sync.WaitGroup{},
	}
	s.wg.Add(1)
	solveForRemainder(&coins, 0, changeDue, &s)
	s.wg.Wait()
	fmt.Println(strconv.FormatInt(s.possibilities, 10))
}
