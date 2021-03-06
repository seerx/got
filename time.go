package got

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

	dateFormatter                    string
	timeFormatter                    string
	timeWithMillisecondFormatter     string
	datetimeFormatter                string
	datetimeWithMillisecondFormatter string
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
	if tf.timeWithMillisecondFormatter == "" {
		tf.timeWithMillisecondFormatter = tf.timeFormatter + ".000"
	}
	if tf.datetimeFormatter == "" {
		tf.datetimeFormatter = fmt.Sprintf("2006%s01%s02%s%s15%s04%s05%s", tf.Y, tf.M, tf.D, tf.DT, tf.H, tf.Mn, tf.S)
	}
	if tf.datetimeWithMillisecondFormatter == "" {
		tf.datetimeWithMillisecondFormatter = tf.datetimeFormatter + ".000"
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

//FormatDateTimeM 格式化时间，末尾带毫秒
func (tf *TF) FormatDateTimeM(time time.Time) string {
	tf.initFormatter()
	return time.Format(tf.datetimeWithMillisecondFormatter)
}

//FormatTimeM 格式化时间，末尾带毫秒
func (tf *TF) FormatTimeM(time time.Time) string {
	tf.initFormatter()
	return time.Format(tf.timeWithMillisecondFormatter)
}

// ParseDate 解析日期
func (tf *TF) ParseDate(date string) (time.Time, error) {
	tf.initFormatter()
	return time.Parse(tf.dateFormatter, date)
}

// ParseDatetime 解析日期及时间
func (tf *TF) ParseDatetime(date string) (time.Time, error) {
	tf.initFormatter()
	return time.Parse(tf.datetimeFormatter, date)
}

// ParseTime 解析时间
func (tf *TF) ParseTime(date string) (time.Time, error) {
	tf.initFormatter()
	return time.Parse(tf.timeFormatter, date)
}

// ParseDatetimeM 解析日期及时间（含毫秒）
func (tf *TF) ParseDatetimeM(date string) (time.Time, error) {
	tf.initFormatter()
	return time.Parse(tf.datetimeWithMillisecondFormatter, date)
}

// GetDateTemplate 获取日期格式模板
func (tf *TF) GetDateTemplate() string {
	tf.initFormatter()
	return tf.dateFormatter
}

// GetDatetimeTemplate 获取日期及时间格式模板
func (tf *TF) GetDatetimeTemplate() string {
	tf.initFormatter()
	return tf.datetimeFormatter
}

// GetTimeTemplate 获取时间格式模板
func (tf *TF) GetTimeTemplate() string {
	tf.initFormatter()
	return tf.timeFormatter
}

// GetDatetimeMTemplate 获取日期及时间(含毫秒)格式模板
func (tf *TF) GetDatetimeMTemplate() string {
	tf.initFormatter()
	return tf.datetimeWithMillisecondFormatter
}
