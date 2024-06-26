/*
Таймеры применяются, когда нужно сделать что-то один раз в будущем - tickers тогда,
когда нужно делать что-то повторно с постоянным интервалом.
Здесь пример счетчика тиков, который периодически тикает до тех пор, пока мы его не остановим.
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	/*
	   Счетчик тиков использует механизм, похожий на таймер:
	   канал, который отправляет значения.
	   Здесь мы используем range, встроенный в канал для перебора значений, поступающих каждые 500 мсек.
	*/
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	/*
	   Счетчики тиков могут быть остановлены подобно таймерам.
	   Как только счетчик будет остановлен, больше не будет получать значения в канале.
	   Остановим его после 1500 мсек.
	*/
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
