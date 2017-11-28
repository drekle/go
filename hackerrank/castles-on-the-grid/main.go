package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

/*
   3
   _ X _
   _ X _
   _ _ _
   (0,0) (0,2)
   3
*/

const (
	nonTraversableChar = "X"
)

var boardSize int64
var buildTreeChan chan *vertex

//Tree vertecies
type vertex struct {
	X           int64
	Y           int64
	Traversable bool
	Depth       int64
	Parent      *vertex
}

func (vert *vertex) getLeft(board map[int64]map[int64]*vertex) []*vertex {
	var ret []*vertex
	for i := int64(0); ; i++ {
		if col, ok := board[vert.X-i]; ok {
			if val, ok := col[vert.Y]; ok && val.Traversable {
				ret = append(ret, val)
			} else {
				break
			}
		} else {
			break
		}
	}
	return ret
}

func (vert *vertex) getRight(board map[int64]map[int64]*vertex) []*vertex {
	var ret []*vertex
	for i := int64(0); ; i++ {
		if col, ok := board[vert.X+i]; ok {
			if val, ok := col[vert.Y]; ok && val.Traversable {
				ret = append(ret, val)
			} else {
				break
			}
		} else {
			break
		}
	}
	return ret
}

func (vert *vertex) getUp(board map[int64]map[int64]*vertex) []*vertex {
	var ret []*vertex
	for i := int64(0); ; i++ {
		if col, ok := board[vert.X]; ok {
			if val, ok := col[vert.Y-i]; ok && val.Traversable {
				ret = append(ret, val)
			} else {
				break
			}
		} else {
			break
		}
	}
	return ret
}

func (vert *vertex) getDown(board map[int64]map[int64]*vertex) []*vertex {
	var ret []*vertex
	for i := int64(0); ; i++ {
		if col, ok := board[vert.X]; ok {
			if val, ok := col[vert.Y+i]; ok && val.Traversable {
				ret = append(ret, val)
			} else {
				break
			}
		} else {
			break
		}
	}
	return ret
}

func parseCoord(x string, y string, board map[int64]map[int64]*vertex) *vertex {
	xVal, _ := strconv.ParseInt(x, 10, 64)
	yVal, _ := strconv.ParseInt(y, 10, 64)
	return board[xVal][yVal]
}

func main() {

	start := time.Now()

	r := bytes.NewReader([]byte(debug2))
	reader := bufio.NewReader(r)
	// reader := bufio.NewReader(os.Stdin)
	//Parse the boardsize
	boardSizeLine, _, _ := reader.ReadLine()
	boardSize, _ = strconv.ParseInt(strings.TrimSpace(string(boardSizeLine)), 10, 64)
	buildTreeChan = make(chan *vertex, 65535)
	board := make(map[int64]map[int64]*vertex, boardSize)
	for i := int64(0); i < boardSize; i++ {
		board[i] = make(map[int64]*vertex, boardSize)
	}
	//Parse the traversable pieces
	for i := int64(0); i < boardSize; i++ {
		line, _, _ := reader.ReadLine()
		runes := strings.Split(string(line), "")
		for index, char := range runes {
			if char == nonTraversableChar {
				board[int64(index)][i] = &vertex{
					Traversable: false,
					X:           int64(index),
					Y:           i,
					Parent:      nil,
				}
			} else {
				board[int64(index)][i] = &vertex{
					X:           int64(index),
					Y:           i,
					Depth:       int64(-1),
					Parent:      nil,
					Traversable: true,
				}
			}
		}
	}
	// Parse the coordinates
	coordLine, _, _ := reader.ReadLine()
	coordStrings := strings.Split(string(coordLine), " ")

	startVertex := parseCoord(coordStrings[1], coordStrings[0], board)
	startVertex.Depth = 0
	endVertex := parseCoord(coordStrings[3], coordStrings[2], board)

	buildTreeChan <- startVertex
	func() {
		for {
			workingVertex, ok := <-buildTreeChan
			if !ok {
				break
			}
			if workingVertex == endVertex {
				break
			}
			children := append(workingVertex.getUp(board), append(workingVertex.getDown(board), append(workingVertex.getLeft(board), workingVertex.getRight(board)...)...)...)
			for _, child := range children {
				if nil != child && (child.Depth < 0 || child.Depth > workingVertex.Depth+1) {
					child.Parent = workingVertex
					child.Depth = workingVertex.Depth + 1
					buildTreeChan <- child
				}
			}
		}
	}()
	fmt.Println(endVertex.Depth)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
