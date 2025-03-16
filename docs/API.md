
**3. docs/API.md**
```markdown
# Referência da API

## Pacote core

### type Qubit
```go
type Qubit struct {
    Coefficients []float64
}
```

Métodos:

- RotateX(angle float64)
- RotateY(angle float64)
- RotateZ(angle float64)
- Measure() int
- Entangle(other *Qubit)

### type QuantumRegister

```go
type QuantumRegister struct {
    Qubits []*Qubit
}
```

Métodos:

- AddQubit(qubit *Qubit)
- MeasureAll() []int
- CheckParity() int