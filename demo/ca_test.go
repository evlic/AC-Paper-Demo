package demo

import (
	"AC-Paper-Demo/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	U128 "github.com/mengzhuo/uint128"
)

var (
	data = [][]int{
		{10, 15},
		{5, 3},
		{100, 50},
	}
	log = *common.GetLogger()
)

// 暴力测试
func TestPlain(t *testing.T) {
	log.Info(
		"\n|\t编号 \t|\t 输入  \t|\t 输出 \t|\t 「+」次数\t｜\t" +
			"\n|\t --- \t|\t ---  \t|\t ---- \t|\t ------\t\t|",
	)

	for idx, val := range data {
		quit := make(chan byte)
		go printTime(quit, idx)
		solution := Solution(val[0], val[1])

		fmt.Printf(
			"|\t%v\t|\t%v\t|\t%v\t|\t%v\t\t|\n", idx, val, solution, Cnt,
		)

		Cnt, _ = U128.NewFromString("0")
		close(quit)
	}
}

// TestByStdin 使用命令行输入
func TestByStdin(t *testing.T) {
	log.Infof("\n输入 (i,j) 对 （支持多组, 空格分隔） >> ")

	arg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	arg = strings.Trim(arg, "\n")

	args := strings.Split(arg, " ")
	fmt.Println(args)

	l := len(args)
	quit := make(chan byte)

	defer close(quit)

	if l >= 2 && l%2 == 0 {
		log.Info(
			"\n|\t编号 \t|\t 输入  \t|\t 输出 \t|\t 「+」次数\t｜\t" +
				"\n|\t --- \t|\t ---  \t|\t ---- \t|\t ------\t\t|",
		)

		for idx := 0; idx < l; idx += 2 {
			i, _ := strconv.Atoi(args[idx])
			j, _ := strconv.Atoi(args[idx+1])

			val := []int{i, j}

			// 异步输出运行时间
			go printTime(quit, idx/2)

			solution := Solution(val[0], val[1])
			fmt.Printf(
				"|\t%v\t|\t%v\t|\t%v\t|\t%v\t\t|\n", idx/2, val, solution, Cnt,
			)

			// 退出
			quit <- 1
			Cnt, _ = U128.NewFromString("0")
		}
	}
}

// TestSolutionMemo 备忘录
func TestSolutionMemo(t *testing.T) {
	log.Info(
		"\n|\t编号 \t|\t 输入  \t|\t 输出 \t|\t 「+」次数\t｜\t" +
			"\n|\t --- \t|\t ---  \t|\t ---- \t|\t ------\t\t|",
	)

	for idx, val := range data {
		quit := make(chan byte)
		go printTime(quit, idx)
		solution := SolutionMemo(val[0], val[1])

		fmt.Printf(
			"|\t%v\t|\t%v\t|\t%v\t|\t%v\t\t|\n", idx, val, solution, *Cnt,
		)

		close(quit)
	}
}

func printTime(quit chan byte, no int) {
	times := int64(0)
	tUnit := int64(time.Second)
	for {
		select {
		case <-quit:
			return
		default:
			break
		}
		times += tUnit
		time.Sleep(time.Duration(tUnit))
		fmt.Println("|", no, "| 已运行 >> |", times/tUnit, "(s)", " |计算「+」号：", Cnt, "次")
	}
}
