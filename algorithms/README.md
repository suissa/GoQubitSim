# Algoritmos

## 1. Algoritmo de Deutsch-Jozsa

### Fundamentação Teórica

Resolve o problema de determinar se uma função binária é:

- Constante (sempre retorna 0 ou sempre 1)
- Balanceada (retorna 0 para metade das entradas e 1 para a outra metade)

Implementação em Go:

```go
djOracle := func(input, output *core.Qubit) {
    gates.ApplyHadamard(output) // Função constante f(x) = 0
}
result := dj.Execute()
```

### Funcionamento

- Inicializa dois qubits em superposição
- Aplica o oráculo
- Mede o resultado


| Passo | Operação Quântica | Complexidade |
| -------- | ----- | ----------- |
| 1       | Hadamard em ambos qubits | O(1) |
| 2        |Aplicação do Oracle | O(1) |
| 3        | Hadamard e medição | O(1) |

### Casos de Uso

- Verificação de integridade de funções hash
- Teste de circuitos quânticos
- Demonstração de supremacia quântica para problemas específicos