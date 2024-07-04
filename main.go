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
	mediaAss := z.assists/z.jogos
	mediaGols := z.gols/z.jogos
	mediaDes := z.desarmes/z.jogos
	mediaCabec := z.cabec/z.jogos
	fmt.Println("Total de Jogos: ", z.jogos)
	fmt.Printf("Gols por jogo: %2.2f \n", mediaGols)
	fmt.Printf("Asists por jogo: %2.2f \n", mediaAss )
	fmt.Printf("Desarmes por jogo: %2.2f \n", mediaDes )
	fmt.Printf("Cabeceio por jogo: %2.2f \n", mediaCabec )
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
	mediaAss := m.assists/m.jogos
	mediaGols := m.gols/m.jogos
	mediaDes := m.desarmes/m.jogos
	mediaPass := m.passes/m.jogos
	fmt.Println("Total de Jogos: ", m.jogos)
	fmt.Printf("Gols por jogo: %2.2f \n", mediaGols)
	fmt.Printf("Asists por jogo: %2.2f \n", mediaAss )
	fmt.Printf("Desarmes por jogo: %2.2f \n", mediaDes )
	fmt.Printf("Passes por jogo: %2.2f \n", mediaPass )
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
	mediaAss := a.assists/a.jogos
	mediaGols := a.gols/a.jogos
	mediaChutes := a.chutes/a.jogos
	mediaDribles := a.dribles/a.jogos
	fmt.Println("Total de Jogos: ", a.jogos)
	fmt.Printf("Gols por jogo: %2.2f \n", mediaGols)
	fmt.Printf("Asists por jogo: %2.2f \n", mediaAss )
	fmt.Printf("Chute por jogo: %2.2f \n", mediaChutes )
	fmt.Printf("Drible por jogo: %2.2f \n", mediaDribles )
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
