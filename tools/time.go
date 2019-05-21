package tools

import (
	"fmt"
	"time"
)

//TF 时间格式化
type TF struct {
	//年月日之间的分隔符
	YMD string
	//日期和时间之间的分隔符
	DT string
	//时分秒之间的分隔符
	HMS string

	dateFormatter     string
	timeFormatter     string
	datetimeFormatter string
}

//NewTimeFormatter 创建时间格式化对象
//默认使用 yyyy-mm-dd hh:mm:ss 形式
func NewTimeFormatter() *TF {
	return &TF{
		YMD: "-", DT: " ", HMS: ":",
	}
}

func (tf *TF) initFormatter() {
	if tf.dateFormatter == "" {
		tf.dateFormatter = fmt.Sprintf("2006%s01%s02", tf.YMD, tf.YMD)
	}
	if tf.timeFormatter == "" {
		tf.timeFormatter = fmt.Sprintf("15%s04%s05", tf.HMS, tf.HMS)
	}
	if tf.datetimeFormatter == "" {
		tf.datetimeFormatter = fmt.Sprintf("2006%s01%s02%s15%s04%s05", tf.YMD, tf.YMD, tf.DT, tf.HMS, tf.HMS)
	}
}

//FormatDate 格式化日期
func (tf *TF) FormatDate(time time.Time) string {
	tf.initFormatter()
	return time.Format(tf.dateFormatter)
}

//FormatDateTime 格式化时间
func (tf *TF) FormatDateTime(time time.Time) string {
	tf.initFormatter()
	return time.Format(tf.datetimeFormatter)
}

//FormatTime 格式化时间
func (tf *TF) FormatTime(time time.Time) string {
	tf.initFormatter()
	return time.Format(tf.timeFormatter)
}
