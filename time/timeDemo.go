package main

import (
	"fmt"
	"time"
)

func timeTimerTicker() {
	timer := time.NewTimer(5 * time.Second)   // 指定时间执行一次
	ticker := time.NewTicker(1 * time.Second) // 指定时间段循环执行

	fmt.Println("now: ", time.Now())
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("timer Done at: ", time.Now())
			return
		case t := <-ticker.C:
			fmt.Println("ticker at: ", t)
		}
	}
}

func timeDemo() {
	// 当前时间
	now := time.Now()
	fmt.Println(now.Weekday().String()) // 星期几 Tuesday

	// 当前时间的UTC格式
	fmt.Println(now.UTC())

	// Time转时间戳
	fmt.Println(now.Unix())     // 以秒为单位
	fmt.Println(now.UnixNano()) // 以纳秒为单位

	// 时间戳转Time
	fmt.Println(time.Unix(now.Unix(), 0))

	// Time格式化 格式化的模板为Go的诞生时间2006年1月2号15点04分 Mon Jan
	fmt.Printf(
		"%d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
	)
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))    // 24小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan")) // 12小时制
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	// Time字符串解析	（默认待解析的字符串为UTC时间）
	timeString := "2021-07-27 15:11:44.006829"
	t, _ := time.Parse("2006-01-02 15:04:05.000000", timeString)
	fmt.Println(t, t.UTC())

	// 加载时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t, _ = time.ParseInLocation("2006-01-02 15:04:05.000000", timeString, loc)
	fmt.Println(t, t.UTC())
	fmt.Println("now in Asia/Shanghai:", now.In(loc)) // 转换时区

	// UTC时区
	loc, _ = time.LoadLocation("UTC")
	t, _ = time.ParseInLocation("2006-01-02 15:04:05.000000", timeString, loc)
	fmt.Println(t, t.UTC())
	fmt.Println("now in UTC:", now.In(loc)) // 转换时区

	// 转换时区

	// Time操作 Add Sub Equal Before After
	fmt.Println(now, now.Add(time.Hour))
	fmt.Println(now.Add(-time.Hour))            // 加减时间
	fmt.Println(now.Add(time.Hour).Sub(now))    // 1h0m0s 时间差
	fmt.Println(now.Equal(now.Add(time.Hour)))  // 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较
	fmt.Println(now.After(now.Add(time.Hour)))  // false
	fmt.Println(now.Before(now.Add(time.Hour))) // true

}

func main() {
	//timeTimerTicker()
	timeDemo()
}
