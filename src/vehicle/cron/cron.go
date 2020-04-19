package cron

import "github.com/robfig/cron/v3"

/*
分钟					|0-59| 			* /，-
营业时间 				|0-23| 			* /，-
每月的一天			|1-31|			* /，-？
月					|1-12或JAN-DEC| * /，-
星期几				|0-6或SUN-SAT|  * /，-？
*/

const (
	//每隔一分钟
	PerMinuteSpec  = "*/10 * * * *"
	//从第49分钟开始，每隔1分钟
	FromMinute_DisMinuteSpec = "49/1 * * * *"
	//从第几小时，第54分钟开始，每隔1分钟
	FromHourMinute_DisMinuteSpec = "54/1 15 * * *"
)

func Setup() {
	c := cron.New()
	c.AddFunc(PerMinuteSpec, perMinuteFun)
	c.Start()
}

/**
 # ┌───────────── min (0 - 59)
 # │ ┌────────────── hour (0 - 23)
 # │ │ ┌─────────────── day of month (1 - 31)
 # │ │ │ ┌──────────────── month (1 - 12)
 # │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
 # │ │ │ │ │                  Saturday, or use names; 7 is also Sunday)
 # │ │ │ │ │
 # │ │ │ │ │
 # * * * * *  command to execute

　1）星号(*)
　　　　表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月

　　2）斜线(/)
　　　　表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15

　　3）逗号(,)
　　　　用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行

　　4）连字号(-)
　　　　表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）

　　5）问号(?)
　　　　只用于 日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替 *

　　6）L，W，#
　　　　Go中没有L，W，#的用法，下文作解释
 */
