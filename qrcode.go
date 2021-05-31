package main

import (
	"bufio"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	println("Bem-vindo ao gerador de QRCode!")
	//cria um loop infinito com o gerador
	for k := 0; k == 0; {
		println("Digite o texto que deseja codificar: ")
		in := bufio.NewReader(os.Stdin)
		input, _ := in.ReadString('\n')
		imprimir(gerar(input[:len(input)-2]))

	}
}

//gera o qrcode como uma lista de listas de valores booleanos (preto=1 e branco=0)
func gerar(texto string) [][]bool {
	img, _ := qrcode.New(texto, qrcode.Low)
	//println(img)
	return img.Bitmap()
}

//imprime a lista de listas usando emojis quadrados para representar os 0s e 1s
func imprimir(png [][]bool) {
	for i := 0; i < len(png); i++ {
		for j := 0; j < len(png[i]); j++ {
			if png[i][j] {
				//print("⬛")
				//print(" ")
				print("██")
			} else {
				//print("⬜")
				//print("■")
				print("  ")
			}
		}
		print("\n")
	}
}
