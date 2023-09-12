# Go - Iniciando

## Introdução

- Aula 1: Uma visão geral da linguagem Go
    - Primeiras informações sobre a linguagem.

- Aula 2: Packages e Ciclo de vida
    - Site com os packages padrões utilizadas pelo Go <https://pkg.go.dev/>
    - Criando um primeiro arquivo em Go
    - Necessário sempre começar com a package main
    - Comando ```import``` para importar packages externos

- Aula 3: Varáveis e tipos de dados
    - Linguagem com tipagem forte, uma váriavel não muda ao decorrer do código.
    - Criando e usando a primeira váriavel
    - Váriaveis entre parenteses () e com letra minuscula são acessiveis somente dentro do mesmo package, caso coloque com letra maiuscula ela pode ser utilizada por outro package realizando um import.
    - Toda váriavel declarada em Go deve ser utilizada
    - ```:=``` declarando uma váriavel dessa forma já atribuimos um valor a uma váriavel.
    - Cada tipo de dado tem seu valor de inicialização, exemplo ```int=0``` e ```string=""```
    - Podemos criar váriaveis de mesmo tipo em uma unica linha utilizando ,
    - Podemos também criar diferentes váriaveis e atribuir valores em uma linha, usando a mesma sequencia e utilizando =.

- Aula 4: Funções
    - Funções são trechos de códigos que podem ser reutilizados em outros lugares do código, com funções a manutenção ao código fica mais fácil.
    - Criação de um novo arquivo .go com algumas funções.

- Aula 5: Fluxos condicionais
    - if e switchcase são fluxos condicionais
    - Criação de diretório (aula3) com alguns exemplos de condicionais

- Aula 6: As várias formas de fazer looping
    - Usando o for de várias formas

- Aula 7: Maps, Slices e Structs
    - Primeiro codigo falando sobre slice
    - diferença entre slice e array é que o array possui valor fixo, você não consegue incrementar um array a menos que você crie um outro array levando as informações do primeiro. O slice é dinâmico você consegue realizar um append ou iniciar um slice vazio e ir incrementando ao decorrer do código.
    - Segundo codigo falando sobre maps
    - maps não possui ordenação
    - maps precisam de validação de chave, se não ele retorna o valor padrão da várial, exemplo 0 para int.
    - Terceiro falando sobre structs
    - criando um type pessoas com as informações
    - No código pode ser passado valor com seu atributo ou na ordem, porém na ordem caso a struct seja alterada e inserido um novo valor o código quebrará.

- Aula 8: O que são métodos
    - Métodos são atrelados a structs
    - método "string" é utilizado com structs

- Aula 9: Herança
    - Heranca, necessário mencionar o nome de uma struct dentro da outra.
    - Ter atenção para não repetir valores nas struct pai e filhos do contrário ele não mostra o valor.

- Aula 10: O que são interfaces
    - Criando interfaces para Documento, pessoa fisica e pessoa juridica

- Aula 11: Asserções em interfaces
    - 

