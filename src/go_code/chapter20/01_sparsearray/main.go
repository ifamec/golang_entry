package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// save to disk - persistence
	filePath := "./src/go_code/chapter20/01_sparsearray/data"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// sparse array
	var cm [11][11]int
	cm[1][2] = 1
	cm[2][3] = 2

	fmt.Println("Ori 2D Array")
	for _, v := range cm {
		for _, val := range v {
			fmt.Print(val, "  ")
		}
		fmt.Println()
	}

	// Turn into Sparse Array
	writer := bufio.NewWriter(file)
	var sa []ValNode
	vNode := ValNode{
		row: len(cm),
		col: len(cm[0]),
	}
	sa = append(sa, vNode)
	writer.WriteString(fmt.Sprintf("%d %d %d\n", vNode.row, vNode.col, vNode.val))
	for i, v := range cm {
		for j, val := range v {
			if val != 0 {
				vNode = ValNode{
					row: i,
					col: j,
					val: val,
				}
				sa = append(sa, vNode)
				writer.WriteString(fmt.Sprintf("%d %d %d\n", vNode.row, vNode.col, vNode.val))
			}
		}
	}
	writer.Flush()

	// Print Sparse Array in Code
	fmt.Println("\nSparse Array")
	fmt.Println(sa)

	fmt.Println("\nnRow:", sa[0].row, "nCol:", sa[0].col, "Sparse Array")
	var r [11][11]int
	for i, v := range sa {
		if i == 0 {
			continue
		}
		r[v.row][v.col] = v.val
	}
	fmt.Println("Recover From Sparse Array")
	for _, v := range cm {
		for _, val := range v {
			fmt.Print(val, "  ")
		}
		fmt.Println()
	}

	fmt.Println("\nRecover From File:")
	srcFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open Error:", err)
	}
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)
	var r2 [11][11]int
	for i:=0; ; i++ {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if i == 0 {
			continue
		}
		str = strings.TrimSuffix(str, "\n")
		slc := strings.Split(str, " ")
		dRow, _ := strconv.Atoi(slc[0])
		dCol, _ := strconv.Atoi(slc[1])
		dVal, _ := strconv.Atoi(slc[2])
		r2[dRow][dCol] = dVal
	}
	for _, v := range r2 {
		for _, val := range v {
			fmt.Print(val, "  ")
		}
		fmt.Println()
	}
}
