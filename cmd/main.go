package main

import (
    "fmt"
    "log"
    "password_cli/pkg/generator"
)

func main() {
    // Logo ASCII
    fmt.Println("===================================")
    fmt.Println("     ____                  _        ")
    fmt.Println("    |  _ \\ __ _ ___  ___  | | ___   ")
    fmt.Println("    | |_) / _` / __|/ _ \\ | |/ _ \\  ")
    fmt.Println("    |  __/ (_| \\__ \\  __/ | | (_) | ")
    fmt.Println("    |_|   \\__,_|___/\\___| |_|\\___/  ")
    fmt.Println("        Password CLI Generator       ")
    fmt.Println("===================================")
    fmt.Println()

    for {
        var length int
        fmt.Print("Quantos caracteres você quer na sua senha? (0 para sair)\n ")
        fmt.Print("De 0 a 128 caracteres.\n")
        fmt.Print("Digite: \n")
        _, err := fmt.Scanln(&length)
        if err != nil {
            fmt.Println("Entrada inválida! Digite um número.")
            continue
        }

        if length == 0 {
            fmt.Println("Saindo... Até logo!")
            break
        }

        if length < 4 || length > 128 {
            fmt.Println("Tamanho inválido! Use entre 4 e 128 caracteres.\n")
            continue
        }

        pass, err := generator.GeneratePasswordWithType(length, "all")
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Senha gerada:", pass)
        fmt.Println() // linha em branco para separar as execuções
    }
}