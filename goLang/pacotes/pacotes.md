## pacotes

recomendações do pacote
    - deixar claro o nome do pacote
    - não ter muito subdiretorios
    
## pacote main

todas as funções que estão no pacote main tem a visibilidade global
para executar os pacotes é necessário - go run *.go

## pacotes diretorio

o nome do pacote é sempre o nome do diretorio em que o arquivo está exemplo operations/add o pacote deve ser chamado de add

## criar um pacote 

é necessário criar um pacote na raiz do projeto

go mod init github.com/fabriciosouza21/golang-pacotes

## utilização dos metodos

há uma padronização que os valores publicos e privados

os metodos que iniciar com letra minusculas são privadas
os metodos que iniciar com letra maiusculas são publicas

## Go pacotes importação

não são importados arquivos e sim pacotes, posso ter varias arquivos com codigo go, se eles estiverem no mesmo pacote

## modos de importação 

modo normal

``` go
import (
	"fmt"

	"github.com/fabriciosouza21/golang-pacotes/operations"
)

```

importação statica não é recomendada, pode gerar um conflito de nomes, cuidado !!!

``` go
import (
	"fmt"

	. "github.com/fabriciosouza21/golang-pacotes/operations"
)

``` 

utilizar alias, também pode ser util quando ha conflitos de pacote

``` go
import (
	"fmt"

	op "github.com/fabriciosouza21/golang-pacotes/operations"
)

``` 

da para chama somente a função de inicialização com _ {nome do pacote}

``` go
import (
	"fmt"

	"github.com/fabriciosouza21/golang-pacotes/operations"

	_ "github.com/fabriciosouza21/golang-pacotes/display"
)

``` 


não pode haver dependicias ciclicas entre pacotes, o compilador vai da error

subdiretorio em pacotes

``` go


import (
	"fmt"

	op "github.com/fabriciosouza21/golang-pacotes/operations"
	mfmt "github.com/fabriciosouza21/golang-pacotes/operations/fmt"

	_ "github.com/fabriciosouza21/golang-pacotes/display"
)

```

