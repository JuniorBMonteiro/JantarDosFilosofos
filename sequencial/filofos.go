package main

import (
	"fmt"
	"time"
)

const n = 5

var (
	garfos    [n]Garfo
	filosofos [n]Filosofo
)

type Filosofo struct {
	Posicao  int
	Nome     string
	GarfoEsq *Garfo
	GarfoDir *Garfo
}

type Garfo struct {
	Posicao int
}

func main() {
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

	garfos[0] = Garfo{Posicao: 1}
	garfos[1] = Garfo{Posicao: 2}
	garfos[2] = Garfo{Posicao: 3}
	garfos[3] = Garfo{Posicao: 4}
	garfos[4] = Garfo{Posicao: 5}

	var (
		start time.Time
		total time.Duration
	)

	start = time.Now()

	for i := 0; i < 10; i++ {
		jantar(&filosofos[0])
		jantar(&filosofos[1])
		jantar(&filosofos[2])
		jantar(&filosofos[3])
		jantar(&filosofos[4])
	}
	total = time.Since(start)
	fmt.Printf("Tempo total: %f", total.Seconds())
}

func jantar(filosofo *Filosofo) {
	filosofo.pensar()
	filosofo.pegarGarfo()
	filosofo.comer()
	filosofo.largarGarfo()
}

func (f *Filosofo) pegarGarfo() {
	var garfoEsq = &garfos[f.Posicao-1]
	var garfoDir = &garfos[(f.Posicao)%n]
	f.GarfoEsq = garfoEsq
	f.GarfoDir = garfoDir
	fmt.Println(f.Nome + " pegou o garfo direito")
}

func (f *Filosofo) largarGarfo() {
	f.GarfoEsq = nil
	f.GarfoDir = nil
	fmt.Println(f.Nome + " largou o garfo")
}

func (f *Filosofo) pensar() {
	fmt.Println("Filoso " + f.Nome + " pensando")
	time.Sleep(5 * time.Second)
}

func (f *Filosofo) comer() {
	fmt.Println("Filoso " + f.Nome + " comendo")
	time.Sleep(5 * time.Second)
}
