## O QUE É GO LANG?
<img src="https://blog.jetbrains.com/wp-content/uploads/2021/02/Go_8001611039611515.gif" width="500px" />
É uma linguagem de programação criada pela Google para simplificar processos e ações que interagem diretamente com o Hardware. A linguagem é bem simples e fácil de aprender, já que boa parte dela segue a mesma estrutura das linuagens modernas.

Inclusive, é possível notar uma **grande** familiaridade com a C (pelo menos nesses passos iniciais).

## ESTRUTURA BASE DE UM PROJETO GO
Primeiramente, deve-se baixar a Go Lang e instalar na sua máquina [clicando aqui](https://golang.org/dl/).
Depois de instalar, poderá utilizar o VS Code para programar em Go. O VS Code tem uma integração incrível com Go. 
Por padrão, devemos criar uma pasta nomeada ``go``. Dentro dessa pasta, é necessário criar mais três ``src``, ``bin`` e ``pkg``.

## ONDE CRIAR MEUS PROJETOS?
Os projetos ficam dentro da pasta *SRC*, onde podemos criar todos os nossos arquivos.

## CRIANDO UM HELLO WORLD EM GO
Dentro de *SRC*, cria-se uma pasta chamada ``Hello``, e dentro podemos criar um arquivo chamado **hello.go**.
Dentro dele, escreve-se a seguinte estrutura:
```
package main

import ("fmt")

func main(){
  fmt.Println("Olá mundo.")
}
```

Para executar o arquivo, abre-se o terminal e digita-se (dentro do diretório onde está o hello.go):
```
go run hello.go
```
