/**
 * @Author: handsomejob
 * @WechatMp: 程序员小乔
 * @Version: 1.0.0
 * @IDE: GoLand
 * @Date: 2022/8/27 17:38
 */

package comm

import (
	"database/sql/driver"
	"strconv"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

type LocalTime struct {
	time.Time
}

// 获取当前时间 返回时间戳
func (t LocalTime) MarshalJSON() ([]byte, error) {
	second := t.Unix()
	return []byte(strconv.FormatInt(second, 10)), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 当前时间转换成日期格式
func (t LocalTime) String() string {
	return time.Now().Format(timeFormat)
}

// 任意日期格式转换成 时间戳
func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = LocalTime{Time: now}
	return
}

// 将指定格式的日期，转成 YYYY-MM-DD HH:ii:ss
func (t LocalTime) FormatDateString(MyDateString string) string {
	// 1. 设置时区
	// 2. time.LoadLocation() 返回time.Time 类型
	loc, _ := time.LoadLocation(timezone)
	// timeFormat = 2006-01-02 15:04:05  go的默认日期
	// 将日期字符串解析为Go的time对象第一个参数指定格式，第二个是日期字符串，转换成时间戳(秒)
	myDate, err := time.ParseInLocation(timeFormat, MyDateString, loc)
	if err != nil {
		return ""
	}
	// 获取时间戳
	// 1.把时间戳转使用time.Unix()转化为 time.Time 类型时间
	// 2.使用Format() 方法进行格式化
	DateString := time.Unix(myDate.Unix(), 0).Format(timeFormat)
	return DateString
}

// 时间戳格式化时间 YYYY-MM-DD HH:ii:ss
func (t LocalTime) FormatTimestamp(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(timeFormat)
}
