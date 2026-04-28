# Desafio de Multithreading - Pós Go Expert

## Descrição

Este é um desafio da **Pós-Graduação Go Expert** que demonstra o uso avançado de **goroutines** e **canais** em Go. O objetivo é implementar um programa concorrente que busca informações de um CEP através de duas APIs diferentes, utilizando o padrão de corrida onde a primeira API a responder retorna o resultado.

## Objetivo do Desafio

- Praticar conceitos avançados de concorrência em Go
- Implementar comunicação segura entre goroutines usando canais
- Utilizar o statement `select` para coordenação de operações assíncronas
- Compreender e aplicar timeout em operações concorrentes
- Demonstrar boas práticas de tratamento de erros em código concorrente

## APIs Utilizadas

1. **ViaCEP** - https://viacep.com.br/ws/{CEP}/json/
2. **BrasilAPI** - https://brasilapi.com.br/api/cep/v1/{CEP}

## Como Rodar

### Pré-requisitos

- Go 1.16 ou superior instalado
- Conexão com a internet (para acessar as APIs)

### Executando o Projeto

1. Abra o terminal na pasta do projeto:

```bash
cd /multithreading-challange
```

2. Execute o programa com o comando abaixo:

```bash
# Opção 1: Rodar o arquivo principal
go run main.go
```

3. O resultado será exibido no console

## Exemplo de Saída

```
Endereço encontrado API: BrasilAPI: {Cep:89037500 Logradouro:Rua Benjamin Constant Bairro:Escola Agrícola Uf: Estado:SC Cidade:Blumenau}
```

ou

```
Endereço encontrado API: ViaCEP: {Cep:01001-000 Logradouro:Praça da Sé Bairro:Sé Uf: Estado:SP Cidade:São Paulo}
```

ou em caso de timeout:

```
Timeout: Nenhum endereço encontrado
```

## Como Funciona

1. O programa recebe um CEP como entrada (atualmente hardcoded como `"01000-000"`)
2. Cria dois canais para receber as respostas das APIs
3. Inicia duas goroutines, uma para cada API
4. Usa `select` para aguardar qual API responder primeiro (com timeout de 1 segundo)
5. Exibe o resultado da primeira API que responder

## Decisões de Implementação

### Estrutura do Código

Optou-se por implementar todo o código em um único arquivo (`main.go`) pois não havia necessidade de criar módulos separados apenas para as funções de requisição HTTP. Isso mantém o projeto simples e focado no objetivo principal do desafio.

### Tratamento de Erros

Não foi implementada uma tratativa específica de erros para cenários como CEP inválido ou falha nas APIs, pois o enunciado do desafio não pedia explicitamente nem especificava como lidar com essas situações. O foco estava nos conceitos de concorrência (goroutines, channels e select), não na validação de entrada ou tratamento avançado de erros.

**Isso fez sentido?** Sim, pois permite manter o código enxuto e alinhado com o escopo do desafio. Em um projeto real, seria importante adicionar validação de CEP e tratamento robusto de erros.

### Comportamento com CEP Inválido

Caso um CEP inválido seja informado:

- As APIs retornarão respostas vazias ou de erro
- O programa exibirá uma mensagem com endereço vazio (todos os campos em branco)
- Não será feito o parse dos dados, resultando em uma estrutura `Endereco` com valores padrão

## Estrutura do Código

- `main.go` - Todo o código: estruturas de dados, funções de busca nas APIs e lógica principal de coordenação

## Como Modificar o CEP

Edite o arquivo `main.go` e altere a linha:

```go
cep := "01000-000" //Coloque aqui seu CEP
```

E execute novamente com `go run main.go`
