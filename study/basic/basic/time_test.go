package basic

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTime1(t *testing.T) {
	now := time.Now() // 获取当前时间
	fmt.Printf("concurrent time:%v\n", now)
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
}

func TestTime2(t *testing.T) {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// 如果当前系统有时区数据库，则可以加载一个位置得到对应的时区
	// 例如，加载纽约所在的时区
	newYork, err := time.LoadLocation("America/New_York") // UTC-05:00
	if err != nil {
		fmt.Println("load America/New_York location failed", err)
		return
	}
	fmt.Println()
	// 加载上海所在的时区
	//shanghai, err := time.LoadLocation("Asia/Shanghai") // UTC+08:00
	// 加载东京所在的时区
	//tokyo, err := time.LoadLocation("Asia/Tokyo") // UTC+09:00
	// 创建时间对象需要指定位置。常用的位置是 time.Local（当地时间） 和 time.UTC（UTC时间）。
	// timeInLocal := time.Date(2009, 1, 1, 20, 0, 0, 0, time.Local) // 系统本地时间
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)
	sameTimeInNewYork := time.Date(2009, 1, 1, 7, 0, 0, 0, newYork)

	// 北京时间（东八区）比UTC早8小时，所以上面两个时间看似差了8小时，但表示的是同一个时间
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)
	fmt.Println(timesAreEqual)
	// 纽约（西五区）比UTC晚5小时，所以上面两个时间看似差了5小时，但表示的是同一个时间
	timesAreEqual = timeInUTC.Equal(sameTimeInNewYork)
	fmt.Println(timesAreEqual)
}

func TestTime3(t *testing.T) {
	//沟槽的 pattern
	str := "2006/01/02 15:04:05.000"
	fmt.Println(time.Now().Format(str))
}

func TestTime4(t *testing.T) {
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeFormat := "2006/01/02 15:04:05"
	timeObj, err := time.Parse(timeFormat, "2022/10/05 11:25:20")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2022-10-05 11:25:20 +0000 UTC
	fmt.Println(timeObj)
	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	timeObj, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 2022-10-05 11:25:20 +0800 CST
	fmt.Println(timeObj)
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(2 * time.Second)
	time.Sleep(3 * time.Second)
	<-timer.C
	fmt.Println("timer fired")
	timer.Reset(3 * time.Second)
	<-timer.C
	fmt.Println("timer fired")
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		fmt.Println("hello")
	}
}

func consumer1(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			handle(event)
		case <-time.After(time.Hour):
			log.Println("warning: no messages received")
		}
	}
}

func consumer2(ch <-chan Event) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
		select {
		case event := <-ch:
			cancel()
			handle(event)
		case <-ctx.Done():
			log.Println("warning: no messages received")
		}
	}
}

func consumer3(ch <-chan Event) {
	timerDuration := 1 * time.Hour
	timer := time.NewTimer(timerDuration)

	for {
		timer.Reset(timerDuration)
		select {
		case event := <-ch:
			handle(event)
		case <-timer.C:
			log.Println("warning: no messages received")
		}
	}
}

type Event struct{}

func handle(Event) {
}
