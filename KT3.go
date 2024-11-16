package main

import (
    "fmt"
    "sync"
    "time"

    "github.com/gen2brain/beeep"
)

type Monkey struct {
    HighSpeed   int
    ClimbTree   bool
    IntelligenceLevel int
}

type Skunk struct {
    SprayRange int
    Size       int
    Nocturnal  bool
}

type Elephant struct {
    Size         int
    MemoryPower  int
    RecognizeDiseases bool
}

func processMonkey(wg *sync.WaitGroup) {
    defer wg.Done()
    monkey := Monkey{HighSpeed: 55, ClimbTree: true, IntelligenceLevel: 8}
    fmt.Printf("Обезьяна - Скорость: %d км/ч, Залезает на деревья: %v, Уровень интеллекта: %d\n", monkey.HighSpeed, monkey.ClimbTree, monkey.IntelligenceLevel)

    go func() {
        beeep.Notify("Обезьяна", "Обработана информация об обезьяне", "")
    }()
    time.Sleep(1 * time.Second) 
}

func processSkunk(wg *sync.WaitGroup) {
    defer wg.Done()
    skunk := Skunk{SprayRange: 3, Size: 5, Nocturnal: true}
    fmt.Printf("Скунс - Дальность распыления: %d м, Размер: %d кг, Ночной: %v\n", skunk.SprayRange, skunk.Size, skunk.Nocturnal)

    go func() {
        beeep.Notify("Скунс", "Обработана информация о скунсе", "")
    }()
    time.Sleep(1 * time.Second) 
}

func processElephant(wg *sync.WaitGroup) {
    defer wg.Done()
    elephant := Elephant{Size: 5000, MemoryPower: 10, RecognizeDiseases: true}
    fmt.Printf("Слон - Размер: %d кг, Память: %d/10, Распознаёт болезни: %v\n", elephant.Size, elephant.MemoryPower, elephant.RecognizeDiseases)

    go func() {
        beeep.Notify("Слон", "Обработана информация о слоне", "")
    }()
    time.Sleep(1 * time.Second) 
}

func main() {
    var wg sync.WaitGroup

    wg.Add(3)
    go processMonkey(&wg)
    go processSkunk(&wg)
    go processElephant(&wg)

    wg.Wait()
    fmt.Println("Все животные обработаны.")
}
