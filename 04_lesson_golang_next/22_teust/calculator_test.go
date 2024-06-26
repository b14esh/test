package main

import (
	"log"
	"testing"
)

//1. Файл с модульными тестами приянто называть:
// * <script_name>_test.go
// * <package_name>_test.go

// 2. Для того, чтобы тестировать функции (методы, стркутуры, интерфейсы и т.д.)
// На каждый юнит создаем по своей тестирующей функции (Test)
// Приянто каждую такую функию начинать со слова Test....
func TestAdd(t *testing.T) {
	//1. 1-ый test-case
	if res := Add(10, 20); res != 30 {
		log.Fatalf("Add(10, 30) fail. expected %d, got %d\n", 30, res) // log.Fatal провоцирует завершение работы кода
	}

	if res := Add(30, 30); res != 60 {
		log.Fatalf("Add(30, 30) fail. expected %d, got %d\n", 60, res)
	}
}

func TestSub(t *testing.T) {
	if res := Sub(30, 15); res != 15 {
		log.Fatalf("Sub(30, 15) fail. expected %d, got %d\n", 60, res)
	}
}

//func TestSub(t *testing.T) {}

func TestMult(t *testing.T) {}

//3. Для запуска тестов используем команду go test
// Детально : go test -v

//4. Coverage (покрытие) - показывает сколько % кода покрыто модульными тестами
// go test -cover -v
// 75~80% coverage - этого бывает более чем достаточно!

//5. Напоследок :
// Все что начинается с слова Test - будет запущено командой go test
// В Go приянто, что создается 1 модуль с тестами на весь пакет (вне зависимости от количества модулей в нем)
// Не тестируйте Getattr/Setattr в общем пайплайне (только специфика)
// Обязательно посмотрите, как происходит связка с CI для Go (TravisCI/CircleCI)
