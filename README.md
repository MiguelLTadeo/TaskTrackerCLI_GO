# TaskTrackerCLI_GO

CLI simples em Go para gerenciar tarefas pelo terminal.

## Requisitos

- Go 1.27.1 ou superior

## Como rodar

Todos os comandos abaixo devem ser executados na raiz do projeto.

### Rodar sem compilar

```bash
go run . add Comprar pão
go run . list
```

### Compilar e executar

```bash
go build -o task-cli .
./task-cli add Estudar Go
```

### Instalar globalmente (opcional)

```bash
go build -o ~/go/bin/task-cli .
task-cli list
```

> `~/go/bin` precisa estar no seu `PATH`.

## Comandos

| Comando | Descrição |
| --- | --- |
| `task-cli add <descrição>` | Cria uma nova tarefa |
| `task-cli update <id> <nova descrição>` | Atualiza a descrição de uma tarefa |
| `task-cli delete <id>` | Remove uma tarefa |
| `task-cli mark-in-progress <id>` | Marca a tarefa como em andamento |
| `task-cli mark-done <id>` | Marca a tarefa como concluída |
| `task-cli list` | Lista todas as tarefas |
| `task-cli list <status>` | Lista por status: `todo`, `in-progress` ou `done` |

## Armazenamento

As tarefas são salvas em `~/tasks.json`, no diretório home do usuário.
