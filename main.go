package main

import (
	"fmt"
	"os"

	"github.com/KalimaraPeleteiro/CLI-FileEncryption/criptography"
	"golang.org/x/term"
)

func main() {

	if len(os.Args) < 2 {
		helpMessage()
		os.Exit(0)
	}
	action := os.Args[1]

	switch action {
	case "help":
		helpMessage()
	case "encrypt":
		encryptHandler()
	case "decrypt":
		decryptHandler()
	default:
		fmt.Println("Use encrypt para criptografar um arquivo, ou decrypt para descriptografá-lo.")
	}

}

func helpMessage() {
	fmt.Println("Programa de Encriptação")
	fmt.Println("Use para criptografar arquivos cujo conteúdo deseja manter secreto, e acessar somente uma senha prévia.")
	fmt.Println("")
	fmt.Println("Como Usar?")
	fmt.Println("go run . encrypt caminho/para/o/arquivo")
	fmt.Println("")
	fmt.Println("Comandos Disponíveis")
	fmt.Println("encrypt \t Para criptografar arquivos (é necessário passar uma senha).")
	fmt.Println("decrypt \t Para descriptografar arquivos com base na senha entregue.")
	fmt.Println("help    \t Para obter informações sobre o programa (é esse texto).")
	fmt.Println("")
}

func encryptHandler() {
	if len(os.Args) < 3 {
		fmt.Println("Você não passou o caminho para o arquivo. Para entender melhor o funcionamento do programa, tente [go run . help]")
		os.Exit(0)
	}

	file := os.Args[2]
	if !isFileValid(file) {
		panic("Não pude encontrar o arquivo.")
	}

	password := getPassword()
	fmt.Println("\nCriptografando seu arquivo...")
	criptography.Encrypt(file, password)
	fmt.Println("\nArquivo seguro!")
}

func decryptHandler() {
	if len(os.Args) < 3 {
		fmt.Println("Você não passou o caminho para o arquivo. Para entender melhor o funcionamento do programa, tente [go run . help]")
		os.Exit(0)
	}

	file := os.Args[2]
	if !isFileValid(file) {
		panic("Não pude encontrar o arquivo.")
	}

	fmt.Print("Passe a senha do arquivo: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nTentando descriptografar o arquivo...")
	criptography.Decrypt(file, password)
	fmt.Println("\nArquivo Descriptografado!")
}

func getPassword() []byte {
	fmt.Print("Insira a senha: ")
	password, _ := term.ReadPassword(0)

	fmt.Print("\nConfirme sua senha: ")
	passwordConfirmation, _ := term.ReadPassword(0)

	if !isPasswordValid(password, passwordConfirmation) {
		fmt.Print("\nVocê confirmou com uma senha diferente. Talvez escreveu errado?\n")
	}

	return password
}
