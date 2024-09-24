package main

import (
	"encoding/json"
	"fmt"
)

type Faturamento struct {
	FaturamentoDiario []float64 `json:"faturamento_diario"`
}

// 1) Calcular valor da variável SOMA.
func calcularSoma() {
	INDICE := 13
	SOMA := 0
	K := 0

	for K < INDICE {
		K++
		SOMA += K
	}

	fmt.Println("Soma:", SOMA)
}

// 2) Verificar se um número pertence à sequência de Fibonacci.
func pertenceFibonacci(num int) {
	a, b := 0, 1

	for b < num {
		a, b = b, a+b
	}

	if b == num || num == 0 {
		fmt.Println(num, "pertence à sequência de Fibonacci.")
	} else {
		fmt.Println(num, "não pertence à sequência de Fibonacci.")
	}
}

// 3) Calcular menor, maior e dias acima da média de faturamento.
func calcularFaturamento(dadosJson string) {
	var faturamento Faturamento

	err := json.Unmarshal([]byte(dadosJson), &faturamento)
	if err != nil {
		fmt.Println("Erro ao ler dados de faturamento:", err)
		return
	}

	var menor, maior, soma float64
	var diasValidos int
	for _, valor := range faturamento.FaturamentoDiario {
		if valor > 0 {
			if diasValidos == 0 || valor < menor {
				menor = valor
			}
			if diasValidos == 0 || valor > maior {
				maior = valor
			}
			soma += valor
			diasValidos++
		}
	}

	media := soma / float64(diasValidos)
	diasAcimaMedia := 0
	for _, valor := range faturamento.FaturamentoDiario {
		if valor > media {
			diasAcimaMedia++
		}
	}

	fmt.Println("Menor faturamento:", menor)
	fmt.Println("Maior faturamento:", maior)
	fmt.Println("Dias acima da média:", diasAcimaMedia)
}

// 4) Calcular percentual de faturamento por estado.
func calcularPercentual() {
	faturamentoEstados := map[string]float64{
		"SP":     67836.43,
		"RJ":     36678.66,
		"MG":     29229.88,
		"ES":     27165.48,
		"Outros": 19849.53,
	}

	total := 0.0

	for _, valor := range faturamentoEstados {
		total += valor
	}

	for estado, valor := range faturamentoEstados {
		percentual := (valor / total) * 100
		fmt.Printf("%s: %.2f%%\n", estado, percentual)
	}
}

// 5) Inverter string sem usar funções prontas.
func inverterString(s string) string {
	resultado := ""

	for i := len(s) - 1; i >= 0; i-- {
		resultado += string(s[i])
	}

	return resultado
}

func main() {
	// 1) Exemplo de execução da função calcularSoma.
	calcularSoma()

	// 2) Exemplo de execução da função pertenceFibonacci.
	num := 21 // Você pode alterar esse valor para testar diferentes números.
	pertenceFibonacci(num)
	pertenceFibonacci(num + 1) // Testa o número seguinte.

	// 3) Exemplo de execução da função calcularFaturamento.
	dadosJson := `{"faturamento_diario": [1500.50, 0, 2000.75, 2200.60, 0, 1700.80, 1600.30]}`
	calcularFaturamento(dadosJson)

	// 4) Exemplo de execução da função calcularPercentual.
	calcularPercentual()

	// 5) Exemplo de execução da função inverterString.
	str := "Sou uma string invertida"
	fmt.Println("String invertida:", inverterString(str))
}
