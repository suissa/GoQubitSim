# GoQubitSim

<div align="center">
  <img src="https://i.imgur.com/FWzye0m.png" alt="QubitSim" width="220"/>
  <h1>GoQubitSim</h1>
  <p><strong>Uma biblioteca Go para simulação de computação quântica e criptografia quântica</strong></p>
  
  <p>
    <a href="https://github.com/suissa/goqubitsim/stargazers"><img src="https://img.shields.io/github/stars/suissa/goqubitsim?style=flat-square" alt="Stars Badge"/></a>
    <a href="https://github.com/suissa/goqubitsim/network/members"><img src="https://img.shields.io/github/forks/suissa/goqubitsim?style=flat-square" alt="Forks Badge"/></a>
    <a href="https://github.com/suissa/goqubitsim/pulls"><img src="https://img.shields.io/github/issues-pr/suissa/goqubitsim?style=flat-square" alt="Pull Requests Badge"/></a>
    <a href="https://github.com/suissa/goqubitsim/issues"><img src="https://img.shields.io/github/issues/suissa/goqubitsim?style=flat-square" alt="Issues Badge"/></a>
    <a href="https://github.com/suissa/goqubitsim/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue?style=flat-square" alt="License Badge"/></a>
  </p>
  
  <p>
    <a href="#sobre-o-projeto">Sobre</a> •
    <a href="#funcionalidades">Funcionalidades</a> •
    <a href="#tecnologias-utilizadas">Tecnologias</a> •
    <a href="#instalação">Instalação</a> •
    <a href="#como-usar">Como Usar</a> •
    <a href="#estrutura-do-projeto">Estrutura</a> •
    <a href="#documentação">Documentação</a> •
    <a href="#contribuindo">Contribuir</a>
  </p>
</div>

---

## 📋 Sobre o Projeto

**QubitSim** é uma biblioteca JavaScript moderna que implementa conceitos fundamentais de computação quântica e criptografia quântica. Projetada para ser intuitiva e educacional, permite aos desenvolvedores e pesquisadores:

- Simular operações com qubits
- Implementar algoritmos quânticos clássicos
- Explorar protocolos de criptografia quântica
- Visualizar estados quânticos
- Monitorar e otimizar performance
- Desenvolver interfaces acessíveis e internacionalizadas

Ideal para estudantes, educadores e entusiastas que desejam explorar o fascinante mundo da computação quântica sem a necessidade de hardware quântico real.

---

## 🚀 Funcionalidades

### ⚛️ Operações Quânticas Básicas
- **Manipulação de qubits**: Criação, transformação e medição de qubits
- **Portas quânticas**: Implementação de portas fundamentais (H, X, Y, Z, CNOT)
- **Emaranhamento quântico**: Simulação de estados emaranhados
- **Medição quântica**: Colapso de superposições e obtenção de resultados clássicos
- **Visualização avançada**: Interface gráfica interativa para circuitos quânticos

### 🔐 Criptografia Quântica
- **Geração de chaves quânticas**: Criação segura de chaves usando princípios quânticos
- **Protocolo BB84**: Implementação do primeiro protocolo de distribuição quântica de chaves
- **Criptografia e descriptografia**: Proteção de mensagens usando chaves quânticas
- **Detecção de interferência**: Identificação de tentativas de espionagem
- **Autenticação quântica**: Verificação segura de identidades
- **Validação de segurança**: Verificações automáticas de vulnerabilidades

### 🧮 Algoritmos Quânticos
- **Deutsch-Jozsa**: Determinação de propriedades de funções booleanas
- **Grover**: Busca em bases de dados não estruturadas
- **Shor**: Fatoração de números inteiros
- **Bernstein-Vazirani**: Descoberta de strings ocultas
- **Visualização de algoritmos**: Interface interativa para acompanhamento de execução

### 📊 Monitoramento e Performance
- **Métricas em tempo real**: Monitoramento de uso de recursos
- **Profiling**: Análise detalhada de performance
- **Otimização automática**: Cache inteligente e lazy loading
- **Rastreamento**: Logs estruturados e análise de erros
- **Dashboards**: Visualização de métricas e alertas

### 🌐 Interface e Acessibilidade
- **Interface responsiva**: Adaptação a diferentes dispositivos
- **Internacionalização**: Suporte a múltiplos idiomas (pt-BR, en, es)
- **Acessibilidade**: Conformidade com WCAG 2.1
- **Temas personalizáveis**: Suporte a temas claros e escuros
- **Componentes reutilizáveis**: Biblioteca de componentes React

### 🔄 DevOps e CI/CD
- **Pipeline automatizado**: Testes, build e deploy automáticos
- **Controle de qualidade**: Linting e análise estática
- **Monitoramento**: Integração com ferramentas de APM
- **Backup automático**: Rotinas de backup e recuperação
- **Segurança**: Verificações automáticas de vulnerabilidades

---

## 🛠️ Tecnologias Utilizadas

<div align="center">
  <img src="https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black" alt="JavaScript"/>
  <img src="https://img.shields.io/badge/Node.js-339933?style=for-the-badge&logo=nodedotjs&logoColor=white" alt="Node.js"/>
  <img src="https://img.shields.io/badge/React-61DAFB?style=for-the-badge&logo=react&logoColor=black" alt="React"/>
  <img src="https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white" alt="TypeScript"/>
  <img src="https://img.shields.io/badge/ESModules-007ACC?style=for-the-badge&logo=javascript&logoColor=white" alt="ES Modules"/>
  <img src="https://img.shields.io/badge/JSDoc-008CC1?style=for-the-badge&logo=javascript&logoColor=white" alt="JSDoc"/>
  <img src="https://img.shields.io/badge/Jest-C21325?style=for-the-badge&logo=jest&logoColor=white" alt="Jest"/>
  <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL"/>
  <img src="https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white" alt="Redis"/>
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker"/>
</div>

---

## 📦 Instalação

```bash
# Instalação básica
npm install qubitsim

# Instalação com suporte a internacionalização
npm install qubitsim i18next react-i18next

# Instalação com todas as dependências opcionais
npm install qubitsim i18next react-i18next @sentry/react @sentry/tracing redis
```

---

## 💻 Como Usar

### Exemplo Básico de Uso

```go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/suissa/goqubitsim/core"
	"github.com/suissa/goqubitsim/crypto"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Demonstração do Sistema Quântico GoQubitSim ===")

	// 1. Criação de Qubits Básicos
	fmt.Println("\n1. Criando Qubits Básicos:")
	q0 := core.NewQubit([]float64{1, 0}) // |0⟩
	q1 := core.NewQubit([]float64{0, 1}) // |1⟩
	fmt.Printf("Qubit 0 Estado Inicial: %v\n", q0.Coefficients)
	fmt.Printf("Qubit 1 Estado Inicial: %v\n", q1.Coefficients)

	// 2. Operações Quânticas
	fmt.Println("\n2. Aplicando Portas Quânticas:")
	q0.RotateX(math.Pi / 2)
	q1.RotateY(math.Pi / 4)
	fmt.Printf("Qubit 0 após Rotação X: %v\n", q0.Coefficients)
	fmt.Printf("Qubit 1 após Rotação Y: %v\n", q1.Coefficients)

	// 3. Entrelaçamento Quântico
	fmt.Println("\n3. Entrelaçando Qubits:")
	q0.Entangle(q1)
	fmt.Println("Qubits Entrelaçados:")
	fmt.Printf("Estado Combinado: %v\n", q0.Coefficients)

	// 4. Geração de Chave Quântica
	fmt.Println("\n4. Gerando Chave Quântica:")
	keyConfig := crypto.GenerateQuantumKey(128)
	fmt.Printf("Chave Gerada: %v...\n", keyConfig.SiftedKey[:16])
	fmt.Printf("Tamanho da Chave: %d bits\n", len(keyConfig.SiftedKey))

	// 5. Criptografia de Mensagem
	fmt.Println("\n5. Criptografando Mensagem:")
	message := []int{0, 1, 0, 0, 1, 1, 0, 1} // "01001101"
	encrypted := crypto.EncryptMessage(message)
	fmt.Printf("Paridade da Mensagem: %d\n", encrypted.Parity)
	fmt.Printf("Qubits Transmitidos: %d\n", len(encrypted.Qubits))

	// 6. Simulação de Transmissão com Interferência
	fmt.Println("\n6. Simulando Transmissão:")
	if rand.Float32() < 0.3 { // 30% chance de interferência
		fmt.Println(">>> Interferência Detectada! <<<")
		encrypted.Qubits[2].RotateX(math.Pi / 4) // Altera um qubit
	}

	// 7. Decriptação da Mensagem
	fmt.Println("\n7. Decifrando Mensagem:")
	decrypted, err := crypto.DecryptMessage(encrypted.Qubits, encrypted.Parity)

	if err != nil {
		fmt.Println("Erro na Decriptação:", err)
	} else {
		fmt.Printf("Mensagem Decifrada: %s\n", decrypted)
		fmt.Printf("Mensagem Original:  %v\n", message)
	}

	// 8. Verificação Quântica de Segurança
	fmt.Println("\n8. Verificação de Segurança:")
	errorRate := keyConfig.CheckEavesdropping(50)
	fmt.Printf("Taxa de Erro Detectada: %.2f%%\n", errorRate*100)
	if errorRate < 0.1 {
		fmt.Println("Status: Sistema Seguro!")
	} else {
		fmt.Println("Status: Possível Ataque Detectado!")
	}
}

```

---

## 📁 Estrutura do Projeto

```
goqubitsim/
├── src/
│   ├── core/           # Componentes fundamentais
│   │   ├── qubit.go
│   │   ├── quantumRegister.go
│   │   └── measurement.go
│   ├── crypto/         # Funcionalidades de criptografia
│   │   ├── generateKey.go
│   │   ├── encryptMessage.go
│   │   └── decryptMessage.go
│   ├── gates/          # Portas quânticas
│   │   ├── hadamard.js
│   │   ├── pauli.js
│   │   └── cnot.js
│   ├── algorithms/     # Algoritmos quânticos
│   │   ├── deutschJozsa.js
│   │   ├── grover.js
│   │   ├── shor.js
│   │   └── bernsteinVazirani.js
│   ├── ui/            # Interface do usuário
│   │   ├── components/
│   │   ├── hooks/
│   │   └── themes/
│   ├── monitoring/    # Sistema de monitoramento
│   │   ├── metrics/
│   │   ├── profiling/
│   │   └── alerts/
│   ├── i18n/         # Internacionalização
│   │   ├── locales/
│   │   └── config.js
│   └── index.js
├── tests/              # Testes unitários
├── docs/              # Documentação da API
├── examples/          # Exemplos de uso
└── package.json
```

---

## 📚 Documentação

A documentação completa da API está disponível em `docs/`. Para gerar a documentação localmente:

```bash
npm run docs
```

Após a geração, abra `docs/index.html` no seu navegador para explorar a documentação interativa.

### Guias Disponíveis
- [Visão Geral](docs/OVERVIEW.md)
- [Guia de Início Rápido](docs/quickstart.md)
- [Referência da API](docs/api.md)
- [Guia de Desenvolvimento](docs/development.md)
- [Guia de Segurança](docs/security.md)
- [Guia de Performance](docs/performance.md)
- [Guia de Acessibilidade](docs/accessibility.md)
- [Guia de Internacionalização](docs/i18n.md)
- [Guia de Monitoramento](docs/monitoring.md)
- [Guia de CI/CD](docs/ci-cd.md)
- [FAQ](docs/faq.md)

---

## 🧪 Testes

Para garantir a qualidade e confiabilidade, o QubitSim possui uma suíte abrangente de testes:

```bash
# Executar todos os testes
npm test

# Executar testes com cobertura
npm run test:coverage

# Executar testes específicos
npm test -- --testPathPattern=qubit

# Executar testes de performance
npm run test:performance

# Executar testes de acessibilidade
npm run test:a11y

# Executar testes de integração
npm run test:integration
```

---

## 🤝 Contribuindo

Contribuições são bem-vindas e muito apreciadas! Siga estes passos:

1. 🍴 Faça um fork do projeto
2. 🌿 Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. 💾 Commit suas mudanças (`git commit -m 'feat(scope): add some AmazingFeature'`)
4. 📤 Push para a branch (`git push origin feature/AmazingFeature`)
5. 🔍 Abra um Pull Request

Consulte nossos guias detalhados:
- [Guia de Contribuição](CONTRIBUTING.md)
- [Guia de Estilo](docs/style-guide.md)
- [Guia de Desenvolvimento](docs/development.md)
- [Guia de Testes](docs/testing.md)

---

## 🔒 Segurança

O QubitSim leva a segurança a sério. Implementamos:

- Validação rigorosa de inputs
- Criptografia de dados sensíveis
- Proteção contra ataques comuns (XSS, CSRF)
- Auditoria automática de dependências
- Monitoramento de vulnerabilidades

Para reportar vulnerabilidades de segurança, por favor envie um email para security@qubitsim.com.

Consulte nosso [Guia de Segurança](docs/security.md) para mais detalhes.

---

## 📈 Performance

O QubitSim é otimizado para:

- Execução eficiente de algoritmos quânticos
- Renderização rápida de interfaces
- Gerenciamento inteligente de memória
- Cache automático de resultados
- Carregamento lazy de componentes

Recursos de monitoramento incluem:
- Métricas em tempo real
- Profiling detalhado
- Alertas automáticos
- Dashboards personalizáveis
- Logs estruturados

Consulte nosso [Guia de Performance](docs/performance.md) para mais detalhes.

---

## 🌐 Internacionalização

O QubitSim suporta múltiplos idiomas:

- 🇧🇷 Português (Brasil)
- 🇺🇸 English
- 🇪🇸 Español

Para adicionar um novo idioma ou melhorar traduções existentes, consulte nosso [Guia de Internacionalização](docs/i18n.md).

---

## ♿ Acessibilidade

O QubitSim segue as diretrizes WCAG 2.1:

- Navegação por teclado
- Suporte a leitores de tela
- Alto contraste
- Textos redimensionáveis
- Legendas e descrições

Consulte nosso [Guia de Acessibilidade](docs/accessibility.md) para mais detalhes.

---

## 📄 Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE) - veja o arquivo LICENSE para detalhes.

<div align="center">
  <img src="https://img.shields.io/badge/License-MIT-yellow.svg" alt="MIT License"/>
</div>

A licença MIT é uma licença de software permissiva que coloca poucas restrições de uso, modificação e distribuição. Ela permite:

- ✅ Uso comercial
- ✅ Modificação
- ✅ Distribuição
- ✅ Uso privado

A única exigência é manter o aviso de copyright e a licença em qualquer cópia do software/código fonte.

---

## 👥 Autores

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/suissa">
        <img src="https://github.com/suissa.png" width="100px;" alt="Suissa"/>
        <br />
        <sub><b>Suissa</b></sub>
      </a>
      <br />
      <sub>Desenvolvimento inicial</sub>
    </td>
  </tr>
</table>

---

## 📝 Notas

Este é um projeto educacional para estudo e simulação de conceitos de computação quântica. Não deve ser usado para criptografia em produção sem uma revisão de segurança adequada.

### Status do Projeto
- **Fase Atual**: 3 - Funcionalidades Core
- **Progresso**: 80%
- **Próximos Passos**:
  1. Implementação do algoritmo de Shor
  2. Melhorias na visualização de circuitos
  3. Otimizações de performance
  4. Expansão da documentação

---

<div align="center">
  <sub>Construído com ❤️ por entusiastas da computação quântica.</sub>
</div> 
