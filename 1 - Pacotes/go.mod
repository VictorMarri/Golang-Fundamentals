//arquivo que vai centralizar todas as dependencias do meu projeto
//Bibliotecas externas, versoes etc tudo fica registrado aqui..\
module awesomeProject

go 1.24

//Dependencias externas ficam aqui. Nós instalamos ela via cmd `go get {link do pacote}`
//Tudo q for gerado nesse arquivo é gerado automaticamente pelo GO
//Se a gente usar o comando 'go mod tidy' ele vai remover desse arquivo todas as dependencias que nao estão sendo usadas na aplicacão
require github.com/badoux/checkmail v1.2.4 // indirect
