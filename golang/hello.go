package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var data *bufio.Reader
	data = bufio.NewReader(os.Stdin)
	var lineCount int = 1
	var blocksCount int = 0
	var goodsCount int = 0
	var commissionPersent int = 0
	var blockProfit int = 0
	var blockLostProfit float64 = 0.00

	for i := 1; true; i++ {
		line, err := data.ReadString('\n')
		line = strings.TrimSpace(line)
		contain := strings.Contains(line, " ")

		if contain {
			stringSplit := strings.Split(line, " ")
			numberCount, _ := strconv.Atoi(stringSplit[0])
			commissionCount, _ := strconv.Atoi(stringSplit[1])

			goodsCount = numberCount
			commissionPersent = commissionCount

			lineCount = lineCount + 1 + goodsCount
			blocksCount--
		}

		if i == 1 {
			number, _ := strconv.Atoi(line)
			blocksCount = number
		}

		if i != 1 && !contain {
			number, _ := strconv.Atoi(line)
			numberProfit := number * commissionPersent / 100
			numberLostProfit := float64(number) * float64(commissionPersent) / 100

			blockProfit += numberProfit
			blockLostProfit += numberLostProfit
		}

		if i > 1 && goodsCount == 0 {
			lostMoney := math.Floor((blockLostProfit-float64(blockProfit))*100) / 100

			fmt.Printf("%.2f\n", lostMoney)

			blockProfit = 0
			blockLostProfit = 0
		}

		if goodsCount != 0 {
			goodsCount--
		}

		if i == lineCount && blocksCount == 0 && goodsCount == 0 {
			break
		}

		if err == io.EOF {
			break
		}
	}
}
