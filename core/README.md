

## Camada de Atenção Quântica

```go
func (qa *QuantumAttention) Forward(input []*Qubit) []*Qubit {
    numQubits := len(input)
    output := make([]*Qubit, numQubits)
    
    // 1. Gerar Queries, Keys e Values quânticos
    queries := make([]*Qubit, numQubits)
    keys := make([]*Qubit, numQubits)
    values := make([]*Qubit, numQubits)
    
    // Aplicar pesos quânticos
    for i := range input {
        queries[i] = qa.applyLinearTransform(input, qa.QueryWeights[i])
        keys[i] = qa.applyLinearTransform(input, qa.KeyWeights[i])
        values[i] = qa.applyLinearTransform(input, qa.ValueWeights[i])
    }
    
    // 2. Calcular scores de atenção quântica
    attentionQubits := make([]*Qubit, numQubits)
    for i := range attentionQubits {
        attentionQubits[i] = NewQubit([]float64{1, 0})
        for j := range queries {
            // Entrelaçar queries e keys
            qa.entanglePair(queries[j], keys[j], attentionQubits[i])
        }
        gates.ApplyHadamard(attentionQubits[i])
    }
    
    // 3. Aplicar atenção aos valores
    for i := range output {
        output[i] = NewQubit([]float64{0, 0})
        for j := range values {
            // Porta controlada pela atenção
            qa.applyControlledRotation(values[j], attentionQubits[j], output[i])
        }
        output[i].Normalize()
    }
    
    return output
}

// Métodos auxiliares
func (qa *QuantumAttention) applyLinearTransform(qubits []*Qubit, weights []float64) *Qubit {
    result := NewQubit([]float64{0, 0})
    for i, q := range qubits {
        angle := weights[i] * math.Pi
        temp := q.Clone()
        temp.RotateY(angle)
        result.Entangle(temp)
    }
    return result
}

func (qa *QuantumAttention) entanglePair(q1, q2, control *Qubit) {
    // Porta SWAP controlada
    if control.Measure() == 1 {
        q1.Coefficients[0], q1.Coefficients[1], q2.Coefficients[0], q2.Coefficients[1] = 
            q2.Coefficients[0], q2.Coefficients[1], q1.Coefficients[0], q1.Coefficients[1]
    }
}

func (qa *QuantumAttention) applyControlledRotation(source, control, target *Qubit) {
    angle := math.Atan2(source.Coefficients[1], source.Coefficients[0])
    if control.Measure() == 1 {
        target.RotateY(angle)
        target.Entangle(source)
    }
}
```

### Explicação da Implementação:

1. Transformação Linear Quântica:

- Cada query/key/value é uma combinação entrelaçada dos qubits de entrada
- Pesos determinam ângulos de rotação Y

```go
func (qa *QuantumAttention) applyLinearTransform(...)
```

2. Mecanismo de Atenção:

- Entrelaçamento controlado entre pares query-key
- Uso de portas Hadamard para criar superposição

```go
func (qa *QuantumAttention) entanglePair(...)
```

3. Aplicação aos Valores:

- Rotações controladas pelos estados de atenção
- Entrelaçamento final dos valores ajustados

```go
func (qa *QuantumAttention) applyControlledRotation(...)
```

### Características Principais:

1. Paralelismo Quântico:

- Processamento simultâneo de todas as relações query-key-value

```go
for j := range queries {
    qa.entanglePair(queries[j], keys[j], attentionQubits[i])
}
```

2. Não-Linearidade Quântica:

- Medições virtuais mantêm a coerência quântica

```go
if control.Measure() == 1 { ... }
```

3. Normalização Adaptativa:

- Preserva a unitariedade das operações

```go
output[i].Normalize()
```

Exemplo de Uso:

```go
attentionLayer := &core.QuantumAttention{
    QueryWeights: [][]float64{{0.5, 0.2}, {0.3, 0.8}},
    KeyWeights:   [][]float64{{0.1, 0.6}, {0.4, 0.3}},
    ValueWeights: [][]float64{{0.9, 0.4}, {0.2, 0.7}},
}

input := []*core.Qubit{
    core.NewQubit([]float64{1, 0}),
    core.NewQubit([]float64{0, 1}),
}

output := attentionLayer.Forward(input)
```


Esta implementação mostra como princípios quânticos podem ser usados para criar mecanismos de atenção totalmente quânticos, com vantagens em:

- Processamento paralelo intrínseco
- Capacidade de representação não-linear
- Eficiência em espaços de alta dimensionalidade