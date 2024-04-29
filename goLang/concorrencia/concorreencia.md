## Concorrência vs paralelismo

## paralismo

- paralismo é ter processamento simultaneo, com multiprocessamento
- paralismo é mais custoso do que a cocorrência
- overheard de controler(cuidado) nem sempre o fato de adicionar mais processadores nem sempre melhorar melhores

## cocorrência

gerenciamento de uma tarefa, que concorrêm por um mesmo recurso, escalonamento

-------------------------------------

Go é uma linguagem concorrente

- paralelismo - executar código simultaneamente em processadores fisicos diferentes
- concorrência - intercalar (administra) vários processos ao mesmo tempo e isso o ocorre em um único processador físico


Concorrência viabiliza paralelismo

é possível que a concorrência seja melhor que o paralelismo !

paralelismo é exige muito mais do SO e do hardware.


Concorrência é a forma de estrutura o seu programa, modelar o sistema para que tenha os beneficios da concorrência.

é a composição de processos (tipicamente funcões) que executam de forma independente


