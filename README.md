
## Executar o programa Go
```bash
go run main.go
```

### Inicializar o módulo Go
```bash
go mod init github.com/KayoRonald/exemplo-go
```

### Compilar e gerar o executável para execução no sistema atual
```bash
./exemplo-go
```

## Escolher OS
| Comando | Descrição |
|---------|-----------|
| GOOS  | Nome do sistema operacional alvo |
| GOARCH | amd64 para 64 bits, 386 para 32 bits |
| -o | Especifica o nome do executável gerado |

### Gerar executável para macOS
```bash
GOOS=darwin GOARCH=amd64 go build -o nome_do_executavel
```

### Gerar executável para Linux
```bash
GOOS=linux GOARCH=amd64 go build -o nome_do_executavel
```

### Gerar executável para Windows
```bash
GOOS=windows GOARCH=amd64 go build -o nome_do_executavel.exe
```