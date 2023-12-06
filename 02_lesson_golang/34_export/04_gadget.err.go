package main

import "fmt"

type TapePlayer struct { //Тип TapePlayer содержит метод Play, моделирующий воспроизведение трека, и метод Stop для остановки виртуального воспроизведения
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct { //Тип TapeRecorder содержит методы Play и Stop, а также метод Record.
	Microphones int
}

func (t TapeRecorder) Play(song string) { //Содержит метод Play, как у TapePlayer.
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() { //Содержит метод Stop, как у TapePlayer.
	fmt.Println("Stopped!")
}

func playList(device TapePlayer, songs []string) {
	for _, song := range songs { //Перебираем все песни в цикле.
		device.Play(song) //Воспроизведение текущей песни.
	}
	device.Stop() //Плеер останавливается после завершения.
}

func main() {
	player := TapeRecorder{} //Создаем значение TapeRecorder вместо TapePlayer.
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape) //TapeRecorder передается функции playList.
}

//Ошибка:
/*
cannot use player (type gadget.TapeRecorder)
as type gadget.TapePlayer in argument to playListОшибка
*/


