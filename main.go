package main

import (
	"fmt"
)

type jogador interface {
	stats()
}
type dados struct {
	nome  string
	pos   string
	idade int
	clube string
}
type statsPrinc struct {
	jogos   float64
	gols    float64
	assists float64
}
type zag struct {
	dados
	statsPrinc
	desarmes float64
	cabec    float64
}

func (z zag) stats() {
	fmt.Println(z.nome)
	fmt.Println(z.clube)
	fmt.Println(z.pos)
	fmt.Println("Total de Jogos:", z.jogos)
	fmt.Println("Gols por jogo:", z.gols/z.jogos)
	fmt.Println("Asists por jogo:", z.assists/z.jogos)
	fmt.Println("Desarmes por jogo:", z.desarmes/z.jogos)
	fmt.Println("Cabeceio por jogo:", z.cabec/z.jogos)
}

type meia struct {
	dados
	statsPrinc
	desarmes float64
	passes   float64
}

func (m meia) stats() {
	fmt.Println(m.nome)
	fmt.Println(m.clube)
	fmt.Println(m.pos)
	fmt.Println("Total de Jogos:", m.jogos)
	fmt.Println("Gols por jogo:", m.gols/m.jogos)
	fmt.Println("Asists por jogo:", m.assists/m.jogos)
	fmt.Println("Desarmes por jogo:", m.desarmes/m.jogos)
	fmt.Println("Passes por jogo:", m.passes/m.jogos)
}

func statsJogador(j jogador) {
	j.stats()
}

type atacante struct {
	dados
	statsPrinc
	dribles float64
	chutes  float64
}

func (a atacante) stats() {
	fmt.Println(a.nome)
	fmt.Println(a.clube)
	fmt.Println(a.pos)
	fmt.Println("Total de Jogos:", a.jogos)
	fmt.Println("Gols por jogo:", a.gols/a.jogos)
	fmt.Println("Asists por jogo:", a.assists/a.jogos)
	fmt.Println("Chute por jogo:", a.chutes/a.jogos)
	fmt.Println("Drible por jogo:", a.dribles/a.jogos)
}

func main() {
	jog1 := zag{
		dados: dados{
			nome:  "Pablo",
			pos:   "Zagueiro",
			idade: 23,
			clube: "Manchester United",
		}, statsPrinc: statsPrinc{
			jogos:   50,
			gols:    6,
			assists: 1,
		},
		desarmes: 65,
		cabec:    99,
	}
	jog2 := meia{
		dados: dados{
			nome:  "Marcelo",
			pos:   "Meio-Campo",
			idade: 26,
			clube: "Real Madrid",
		}, statsPrinc: statsPrinc{
			jogos:   54,
			gols:    3,
			assists: 8,
		},
		desarmes: 76,
		passes:   240,
	}
	jog3 := atacante{
		dados: dados{
			nome:  "Paulo",
			pos:   "Atacante",
			idade: 22,
			clube: "Juventus",
		}, statsPrinc: statsPrinc{
			jogos:   47,
			gols:    29,
			assists: 12,
		},
		dribles: 34,
		chutes:  211,
	}

	statsJogador(jog1)
	statsJogador(jog2)
	jogador(jog3).stats()
}
