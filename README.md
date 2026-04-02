# Password CLI Generator 🔐

Um gerador de senhas interativo escrito em Go, focado em simplicidade e segurança.

## ✨ Funcionalidades

- **Segurança Criptográfica**: Utiliza `crypto/rand` para garantir a aleatoriedade das senhas.
- **Interface Interativa**: Solicita o tamanho da senha diretamente no terminal.
- **Validação**: Garante que o tamanho da senha esteja entre 4 e 128 caracteres.
- **Conjunto Completo de Caracteres**: Gera senhas incluindo letras (maiúsculas e minúsculas), números e caracteres especiais.

## 🚀 Como Executar

### Pré-requisitos

- [Go](https://go.dev/doc/install) 1.22 ou superior instalado.

### Passo a Passo

1. **Clone o repositório**:
   ```bash
   git clone (https://github.com/JeanVitorDSL/password-cli.git)
   cd password-cli
   ```

2. **Execute o projeto**:
   ```bash
   go run cmd/main.go
   ```

3. **(Opcional) Build para binário**:
   ```bash
   go build -o pass-gen cmd/main.go
   ./pass-gen
   ```

## 🛠️ Estrutura do Projeto

- `cmd/main.go`: Ponto de entrada da aplicação e lógica da interface CLI.
- `pkg/generator/`: Contém a lógica de geração de senhas.
- `pkg/storage/`: Espaço reservado para futuras implementações de persistência.

## 📝 Exemplo de Uso

```text
===================================
     ____                  _        
    |  _ \ __ _ ___  ___  | | ___   
    | |_) / _` / __|/ _ \ | |/ _ \  
    |  __/ (_| \__ \  __/ | | (_) | 
    |_|   \__,_|___/___| |_|\___/  
        Password CLI Generator       
===================================

Quantos caracteres você quer na sua senha? (0 para sair)
 De 0 a 128 caracteres.
Digite: 
16
Senha gerada: aB3!fG9#kL1@mN5&
```

---
Desenvolvido com ❤️ em Go.
