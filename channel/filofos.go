package main

import (
	"fmt"
	"sync"
	"time"
)

const n = 5

var (
	waitGroup sync.WaitGroup
	garfos    [n]chan bool
	filosofos [n]Filosofo
)

type Filosofo struct {
	Posicao  int
	Nome     string
	GarfoEsq chan bool
	GarfoDir chan bool
}

func main() {

	garfos[0] = make(chan bool)
	garfos[1] = make(chan bool)
	garfos[2] = make(chan bool)
	garfos[3] = make(chan bool)
	garfos[4] = make(chan bool)

	pitagoras := Filosofo{Posicao: 1, Nome: "Pitagoras"}
	aristoteles := Filosofo{Posicao: 2, Nome: "Aristoteles"}
	platao := Filosofo{Posicao: 3, Nome: "Platao"}
	socrates := Filosofo{Posicao: 4, Nome: "Socrates"}
	epicuro := Filosofo{Posicao: 5, Nome: "Epicuro"}

	filosofos[0] = pitagoras
	filosofos[1] = aristoteles
	filosofos[2] = platao
	filosofos[3] = socrates
	filosofos[4] = epicuro

	// aloca os canais garfos aos respectivos filosofos
	for i := 0; i < n; i++ {
		var garfoEsq = &garfos[i]
		var garfoDir = &garfos[(i+1)%n]

		filosofos[i].GarfoEsq = *garfoEsq
		filosofos[i].GarfoDir = *garfoDir
	}

	var (
		start time.Time
		total time.Duration
	)

	start = time.Now()
	waitGroup.Add(5)

	go channelGarfos(garfos[0])
	go channelGarfos(garfos[1])
	go channelGarfos(garfos[2])
	go channelGarfos(garfos[3])
	go channelGarfos(garfos[4])

	go jantar(&filosofos[0])
	go jantar(&filosofos[1])
	go jantar(&filosofos[2])
	go jantar(&filosofos[3])
	go jantar(&filosofos[4])

	waitGroup.Wait()
	total = time.Since(start)
	fmt.Printf("Tempo total: %f", total.Seconds())
}

func jantar(filosofo *Filosofo) {
	for i := 0; i < 10; i++ {
		filosofo.pensar()
		filosofo.pegarGarfo()
		filosofo.comer()
		filosofo.largarGarfo()
	}
	waitGroup.Done()
}

func (f *Filosofo) pegarGarfo() {
	// para evitar o deadlock foi adicionado uma condicao para que o ultimo filosofo pege o garfo direito primeiro
	// assim ele ficara disputando com o primeiro filosofo e o quarto filosofo poderá pegar o garfo a direita
	// quebrando o ci
	if f.Posicao == 5 {
		<-f.GarfoDir
		<-f.GarfoEsq
	} else {
		<-f.GarfoEsq
		<-f.GarfoDir
	}
	fmt.Println(f.Nome + " pegou o garfo ")
}

func (f *Filosofo) largarGarfo() {
	// destrava os processos lendo do canal
	f.GarfoEsq <- true
	f.GarfoDir <- true
	fmt.Println(f.Nome + " largou o garfo ")
}

func (f *Filosofo) pensar() {
	fmt.Println("Filosofo " + f.Nome + " pensando")
	time.Sleep(5 * time.Second)
}

func (f *Filosofo) comer() {
	fmt.Println("Filosofo " + f.Nome + " comendo")
	time.Sleep(5 * time.Second)
}

func channelGarfos(c chan bool) {
	for true {
		c <- true
		<-c
	}
}
