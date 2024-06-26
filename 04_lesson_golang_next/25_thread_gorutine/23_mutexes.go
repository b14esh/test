package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	//Пусть состояние (state) будет картой.
	var state = make(map[int]int)

	//Этот mutex будет синхронизировать доступ к state.
	var mutex = &sync.Mutex{}

	//Для сравнения подхода на основе мьютексов с другим, который увидим позже,
	//ops будет считать количество операций над состоянием.
	var ops int64 = 0

	//Запускаем 100 горутин для выполнения повторяющихся чтений состояния.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				/*
				   Для каждой операции чтения выбираем ключ для доступа,
				   закрываем (Lock()) этот mutex для получения эксклюзивного доступа к state,
				   читаем значение для выбранного ключа, Unlock() мьютекс и увеличиваем счётчик ops.
				*/
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				/*
				   Для гарантии того, что горутина не остановит планировщик, явно выполним после каждой операции runtime.Gosched().
				   Такое выполнение производится автоматически, например,
				   каждой операцией канала или при блокирующих вызовах подобных time.Sleep,
				   но в нашем случае нужно выполнить его вручную.
				*/
				runtime.Gosched()
			}
		}()
	}
	/*
	   Запускаем 10 горутин для имитации записи.
	   Используем тот же подход, что и для имитации чтения.
	*/
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	//Дадим время в 1 сек. для работы 10 горутин с state и mutex.
	time.Sleep(time.Second)

	//Отчёт об окончательном числе операций.
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	//Фиксируем окончательно состояние state и показываем его.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
