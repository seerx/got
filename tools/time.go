package tools

import (
	"fmt"
	"time"
)

//TF 时间格式化
type TF struct {
	// //年月日之间的分隔符
	// YMD string
	// //日期和时间之间的分隔符
	// DT string
	// //时分秒之间的分隔符
	// HMS string

	//年月日之间的分隔符
	Y string // 年
	M string // 月
	D string // 日
	// YMD string
	//日期和时间之间的分隔符
	DT string
	//时分秒之间的分隔符
	H  string
	Mn string
	S  string

	dateFormatter     string
	timeFormatter     string
	datetimeFormatter string
}

//NewTimeFormatter 创建时间格式化对象
//默认使用 yyyy-mm-dd hh:mm:ss 形式
func NewTimeFormatter() *TF {
	return &TF{
		Y: "-", M: "-", D: "",
		DT: " ",
		H:  ":", Mn: ":", S: "",
	}
}

//NewChineseTimeFormatter 创建中文格式化日期
func NewChineseTimeFormatter() *TF {
	return &TF{
		Y: "年", M: "月", D: "日",
		DT: "",
		H:  "时", Mn: "分", S: "秒",
	}
}

func (tf *TF) initFormatter() {
	if tf.dateFormatter == "" {
		tf.dateFormatter = fmt.Sprintf("2006%s01%s02%s", tf.Y, tf.M, tf.D)
	}
	if tf.timeFormatter == "" {
		tf.timeFormatter = fmt.Sprintf("15%s04%s05%s", tf.H, tf.Mn, tf.S)
	}
	if tf.datetimeFormatter == "" {
		tf.datetimeFormatter = fmt.Sprintf("2006%s01%s02%s%s15%s04%s05%s", tf.Y, tf.M, tf.D, tf.DT, tf.H, tf.M, tf.S)
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
