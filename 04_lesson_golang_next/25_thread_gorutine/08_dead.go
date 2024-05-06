package main

//1. Deadlock - ситуация, когда кто-то пишет в канал НО ИЗ НЕГО НИКОГДА НИКТО НИЧЕГО НЕ ПРОЧИТАЕТ, или когда кто-то читает из канала
// НО В НЕГО НИКТО НИКОГДА НЕ ЗАПИШЕТ!

// По сути ситуация означает, что для отправляющей стороны отсутствует получатель (с другой стороны никто не ждет данных). И наоборот.

//fatal error: all goroutines are asleep - deadlock!

func main() {
	ch := make(chan int)
	ch <- 10
	//<-ch
}
