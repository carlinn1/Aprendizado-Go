package main

import "fmt"

func main() {

	// Tipos de dados
	var inteiro int
	var flutuante float64
	var booleano bool
	var texto string

	// Atribuindo valores
	fmt.Println("Valor do inteiro:", inteiro)
	fmt.Println("Valor do flutuante:", flutuante)
	fmt.Println("Valor do booleano:", booleano)
	fmt.Println("Valor do texto:", texto)

	// Atribuindo valores específicos
	var nome string = "Carlos"
	var idade int = 30
	var altura float64 = 1.75
	var estudante bool = true

	// Imprimindo os valores
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Estudante:", estudante)

	// Funções
	soma := func(a, b int) int {
		return a + b
	}

	// Chamando a função de soma
	resultado := soma(5, 3)
	fmt.Println("Resultado da soma:", resultado)

	//Arrays e Slices
	var numeros [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array de números:", numeros)

	// Slices são mais flexíveis que arrays e não têm tamanho fixo
	slice := []string{"Go", "Python", "Java"}
	fmt.Println("Slice de linguagens:", slice)

	// Structs são tipos de dados compostos que agrupam campos relacionados
	type Pessoa struct {
		Nome  string
		Idade int
	}

	pessoa1 := Pessoa{Nome: "Ana", Idade: 25}
	fmt.Println("Pessoa:", pessoa1)

	// If e Else
	// Verificar se a pessoa é maior de idade
	if pessoa1.Idade >= 18 {
		fmt.Println(pessoa1.Nome, "é maior de idade.")
	} else {
		fmt.Println(pessoa1.Nome, "é menor de idade.")
	}

	// Loop For
	fmt.Println("Números de 1 a 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	//Loop While (usando for como while)
	fmt.Println("Contagem regressiva de 5 a 1:")
	contador := 5
	for contador > 0 {
		fmt.Println(contador)
		contador--
	}

	// Switch
	dia := "Segunda-feira"
	switch dia {
	case "Segunda-feira":
		fmt.Println("Hoje é segunda-feira.")
	case "Terça-feira":
		fmt.Println("Hoje é terça-feira.")
	default:
		fmt.Println("Hoje é outro dia da semana.")
	}

	//Maps são coleções de pares chave-valor
	estudantes := make(map[string]int)
	estudantes["Alice"] = 85
	estudantes["Bob"] = 90
	estudantes["Charlie"] = 78

	fmt.Println("Notas dos estudantes:", estudantes)

	//Switch Case
	nota := 85
	switch {
	case nota >= 90:
		fmt.Println("Excelente!")
	case nota >= 80:
		fmt.Println("Muito bom!")
	case nota >= 70:
		fmt.Println("Bom!")
	default:
		fmt.Println("Precisa melhorar.")
	}
}
