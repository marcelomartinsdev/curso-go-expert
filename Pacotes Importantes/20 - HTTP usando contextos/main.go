package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// Uma goroutine, em termos simples, é uma função que pode ser executada em paralelo com o restante do programa.
// Elas são frequentemente usadas para realizar tarefas concorrentes, como fazer solicitações HTTP, permitindo que o programa continue funcionando enquanto aguarda a conclusão de tarefas demoradas.
// A principal diferença entre uma goroutine e uma chamada HTTP comum é que uma goroutine não retorna um valor imediatamente, apenas executa uma função, enquanto uma chamada HTTP retorna uma resposta.

// O pacote "context" é usado para compartilhar informações entre várias partes do código e é especialmente útil para o cancelamento de tarefas e a definição de limites de tempo para a execução de tarefas.
// Em um nível mais profundo, o pacote "context" usa o pacote "sync/atomic" para facilitar a comunicação entre goroutines.
// O pacote "sync/atomic" é usado para realizar operações atômicas, garantindo que certas operações não sejam interrompidas ou corrompidas por várias goroutines ao mesmo tempo.
// Por exemplo, incrementar um contador pode ser uma operação atômica para evitar conflitos de concorrência.
// Um cenário real de uso do pacote "context" ocorre quando uma solicitação faz uma consulta a um banco de dados que pode levar muito tempo.
// Se o usuário decidir cancelar a solicitação, o pacote "context" pode ser usado para cancelar instantaneamente a consulta ao banco de dados.
// É uma boa prática permitir que as goroutines parem de executar com base no contexto para evitar o uso desnecessário de recursos quando não são mais necessárias.

func main() {
	ctx := context.Background() // Cria um contexto vazio.
	// Exemplo de criação de um contexto com um prazo (timeout) de 5 segundos.
	// Isso significa que qualquer tarefa que use esse contexto e leve mais de 5 segundos para ser concluída será cancelada automaticamente.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel() // Cancela o contexto quando a função "main" terminar de ser executada. Isso ajuda a evitar vazamentos de memória, liberando recursos não utilizados.

	// Cria uma nova solicitação HTTP usando o contexto. Se o contexto for cancelado, a solicitação também será cancelada.
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	// Realiza a solicitação HTTP usando o cliente padrão. O cliente padrão já está configurado com valores como cookies, cabeçalhos e tempo limite.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta e imprime no console.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}

// Explicando em detalhes o código acima:
// Primeiro, criamos um contexto vazio e, em seguida, criamos um contexto com um prazo de 5 segundos.
// Isso significa que qualquer tarefa executada com esse contexto será cancelada automaticamente se levar mais de 5 segundos.
// Em seguida, criamos uma solicitação HTTP usando esse contexto, o que significa que a solicitação também será cancelada se o contexto for cancelado.
// Depois, fazemos a solicitação HTTP usando o cliente padrão, que já vem com algumas configurações predefinidas, como cookies, cabeçalhos e tempo limite.
// Por fim, lemos o corpo da resposta da solicitação e o imprimimos no console.
