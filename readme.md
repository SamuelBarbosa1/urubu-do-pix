## Urubu do Pix

Descrição
O Urubu do Pix é um sistema fictício de gestão financeira em Go. O sistema simula um banco com funcionalidades de criação de contas, depósito e transferência de dinheiro, participação em uma loteria, e outros recursos para entretenimento e brincadeiras.

<div>
  Funcionalidades
Criação de Conta: Permite aos usuários criar uma conta com um nome e senha, gerando uma chave Pix exclusiva.

````
## Login: Usuários podem fazer login utilizando sua chave Pix e senha.
````

## Depósito de Dinheiro: Usuários podem depositar dinheiro em suas contas.
````
## Transferência de Dinheiro: Permite transferências de dinheiro entre u
uários.
````

## Participação na Loteria: Usuários podem participar de uma loteria, apostando um valor e recebendo um resultado.
````

## Ranking de Usuários: Exibe um ranking dos usuários com base no saldo.

````
## Conquistas: Sistema de conquistas para desbloquear diferentes objetivos e marcos.
````

## Transferência Surpresa: Realiza transferências aleatórias entre usuários como uma surpresa.
````
## Empréstimo com Juros Altos: Simula um sistema de empréstimos com juros elevados.
````

## Doações Caridosas: Permite aos usuários fazer doações para uma causa fictícia.
````
## Estrutura do Projeto
````

## main.go: Ponto de entrada principal do aplicativo.
````

## user/user.go: Define a estrutura de dados do usuário e funções relacionadas.
````

## user/user_repository.go: Implementa o repositório de usuários, incluindo criação, atualização e recuperação de usuários.
````

## transaction/transaction.go: Gerencia as transações financeiras entre usuários.
````

## lottery/lottery_service.go: Gerencia a participação dos usuários na loteria.
````

## ranking/ranking_service.go: Calcula e exibe o ranking dos usuários com base no saldo.
````

## troll/troll_service.go: Implementa funcionalidades de transferência surpresa e outras funções divertidas.

````
## achievements/achievements_service.go: Gerencia as conquistas dos usuários.
````
Instalação e Execução
Clone o Repositório

bash
Copiar código
git clone https://github.com/SamuelBarbosa1/urubu-do-pix
Navegue para o Diretório do Projeto

bash
Copiar código
cd urubu-do-pix
Compile o Projeto

bash
Copiar código
go build -o urubu-do-pix
Execute o Projeto

bash
Copiar código
./urubu-do-pix
```````
