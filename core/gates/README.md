# Portas Quânticas

## 1. Porta Hadamard (H)

### O que é?

A porta quântica mais fundamental para criar superposição. Transforma estados básicos em estados de superposição igualmente ponderados.

### Para que serve?

- Criar superposições quânticas
- Inicializar algoritmos quânticos
- Implementar transformadas de Fourier quântica

Matriz de Hadamard:

```
[ 1/√2   1/√2 ]
[ 1/√2  -1/√2 ]
```

Exemplo em Go:

```go
q := core.NewQubit([]float64{1, 0}) // |0⟩
gates.ApplyHadamard(q)
// Estado resultante: [0.7071, 0.7071] → (|0⟩ + |1⟩)/√2
```

## 2. Portas Pauli (X, Y, Z)

### O que são?

Operações básicas de rotação no espaço de Bloch:

| Porta | Efeito | Matriz |
| -------- | ----- | ----------- |
| X        | Rotação de 180° no eixo X     | [[0, 1], [1, 0]]     |
| Y        |Rotação de 180° no eixo Y     |     [[0, -i], [i, 0]]        |
| Z        | Rotação de 180° no eixo Z    |      [[1, 0], [0, -1]]       |


### Para que servem?

- X: Inversão de estados (equivalente ao NOT clássico)
- Y: Combinação de inversão e mudança de fase
- Z: Inversão de fase (marca estados |1⟩)

Exemplo em Go:

```go
pauli := gates.NewPauliGates()

// Pauli-X em |0⟩
q := core.NewQubit([]float64{1, 0})
pauli.ApplyX(q) // Estado: |1⟩

// Pauli-Z em superposição
q = core.NewQubit([]float64{0.7071, 0.7071})
pauli.ApplyZ(q) // Estado: (|0⟩ - |1⟩)/√2
```

## 3. Porta CNOT (Controlled-NOT)

### O que é?

Operação de dois qubits onde um qubit controle determina a operação no qubit alvo.

### Para que serve?

- Criar emaranhamento quântico
- Implementar operações condicionais
- Construir portas universais

Tabela Verdade:

| Controle | Alvo | Saída Controle | Saída Alvo |
| -------- | ----- | ----------- |
| 0       | 0 | 0 | 0 |
| 0        | 1 | 0 | 1 |
| 1        | 0 | 1 | 1 |
| 1        | 1 | 1 | 0 |

Exemplo em Go:

```go
controle := core.NewQubit([]float64{1, 0}) // |0⟩
alvo := core.NewQubit([]float64{0, 1})     // |1⟩

gates.ApplyCNOT(controle, alvo)
// Saída: controle=|0⟩, alvo=|1⟩ (não muda)

controle.RotateX(math.Pi) // |0⟩ → |1⟩
gates.ApplyCNOT(controle, alvo)
// Saída: controle=|1⟩, alvo=|0⟩ (alvo invertido)
```

## Casos de Uso Reais

### 1. Superposição para Algoritmos

```go
// Inicializa registro quântico
qr := core.NewQuantumRegister(3)

// Aplica Hadamard em todos os qubits
for _, q := range qr.Qubits {
    gates.ApplyHadamard(q)
}
// Cria estado: (|000⟩ + |001⟩ + ... + |111⟩)/√8
```


### 2. Teleporte Quântico

```go
// Qubit a ser teleportado
psi := core.NewQubit([]float64{0.6, 0.8}) // Estado arbitrário

// Par emaranhado
alice := core.NewQubit([]float64{1, 0})
bob := core.NewQubit([]float64{0, 1})
gates.ApplyCNOT(alice, bob)

// Protocolo de teleporte
gates.ApplyCNOT(psi, alice)
gates.ApplyHadamard(psi)
```

### 3. Correção de Erros

```go
// Codificação de 3 qubits
qData := core.NewQubit([]float64{0.8, 0.6})
q1 := core.NewQubit([]float64{1, 0})
q2 := core.NewQubit([]float64{1, 0})

gates.ApplyCNOT(qData, q1)
gates.ApplyCNOT(qData, q2)
// Estado codificado: α|000⟩ + β|111⟩
```

Estas portas formam a base para:

- Algoritmos quânticos (Shor, Grover)
- Protocolos de criptografia quântica (BB84)
- Correção de erros quânticos
- Computação quântica universal

A implementação em Go permite simular esses comportamentos fundamentalmente quânticos usando operações matriciais clássicas, servindo como ferramenta educacional e base para prototipagem de algoritmos.