package common

import (
	"os"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

/** ========================= 数学工具 ========================= */

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxInSlice >> 返回不含重复元素切片中最大值下标
func MaxInSlice(n []int) (idx int) {
	max := ^0xffffffff

	for i, v := range n {
		if v > max {
			max = v
			idx = i
		}
	}
	return
}

func MinIntSliceVal(a []int) int {
	min := 0x0ffffff
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func init() {
	logger.SetOutput(os.Stdout)

	logger.SetFormatter(
		&logrus.TextFormatter{
			ForceColors:               true,
			FullTimestamp:             true,
			ForceQuote:                true,
			EnvironmentOverrideColors: true,
			TimestampFormat:           " 2006-01-02 | 15:04:05 ",
			// 定义方法调用输出格式
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				function = frame.Function + "() >> "
				file = " : " + strconv.Itoa(frame.Line) + "|>"
				return
			},
		},
	)

	logger.SetReportCaller(true)
}

func GetLogger() *logrus.Logger {
	return logger
}
