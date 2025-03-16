# Guia Rápido

## Instalação
```bash
go get github.com/suissa/goqubitsim
```

Exemplo básico:

```go
package main

import (
    "fmt"
    "math"
    
    "github.com/suissa/goqubitsim/core"
    "github.com/suissa/goqubitsim/crypto"
)

func main() {
    // Criação de qubits
    q0 := core.NewQubit([]float64{1, 0})
    q1 := core.NewQubit([]float64{0, 1})

    // Aplicar rotação
    q0.RotateX(math.Pi / 2)
    
    // Emaranhamento
    q0.Entangle(q1)
    
    // Criptografia
    msg := []int{0, 1, 0}
    encrypted := crypto.EncryptMessage(msg)
    
    fmt.Printf("Estado final: %v\n", q0.Coefficients)
}
```