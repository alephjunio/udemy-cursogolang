// Concorrência vs Paralelismo
//
// Concorrência: Execução de várias tarefas ao mesmo tempo, mas não necessariamente ao mesmo tempo.
// Paralelismo: Execução de várias tarefas ao mesmo tempo, ao mesmo tempo.
//
// Go é uma linguagem concorrente, mas não paralela.
// Isso significa que Go pode usar múltiplos núcleos de CPU,
// mas não executa várias tarefas ao mesmo tempo.
// Em vez disso, Go usa uma abordagem de "goroutines" para
// executar várias tarefas concorrentemente.
//
// Goroutines são como threads, mas mais leves.
// Elas são gerenciadas pelo runtime do Go,
// o que significa que você não precisa se preocupar com a criação,
// o gerenciamento e a sincronização de goroutines.
//
// Goroutines são usadas para executar tarefas concorrentemente,
// como fazer solicitações de rede, acessar bancos de dados,
// ou processar grandes volumes de dados.
//
// Goroutines são criadas usando a palavra-chave 'go'.
// Por exemplo:
// go minhaFuncao()
//
// Isso criará uma nova goroutine que executará a função 'minhaFuncao'.
//
// Goroutines são leves, o que significa que você pode criar milhares
// de goroutines sem problemas.
//
// Goroutines são sincronizadas usando canais (channels).
// Canais são usados para comunicação entre goroutines.
//
// Por exemplo:
// c := make(chan int)
// go minhaFuncao(c)
//
// Isso criará uma nova goroutine que executará a função 'minhaFuncao'
// e passará o canal 'c' como argumento.
//
// A função 'minhaFuncao' pode enviar valores para o canal usando a expressão 'c <- valor'.
//
// A função 'minhaFuncao' pode receber valores do canal usando a expressão 'valor := <-c'.
//
// Por exemplo:
// func minhaFuncao(c chan int) {
//     valor := <-c
//     fmt.Println("Valor recebido:", valor)
// }
//
// Isso criará uma função 'minhaFuncao' que receberá um valor do canal 'c'
// e imprimirá o valor recebido.
//
// Canais são seguros para uso em goroutines, o que significa que você pode
// enviar e receber valores do canal de várias goroutines ao mesmo tempo.
//
// Canais também podem ser usados para sincronizar a execução de goroutines.
// Por exemplo, você pode usar um canal para sinalizar quando uma goroutine
// terminou de executar.
//
// Canais também podem ser usados para implementar padrões de comunicação,
// como o padrão "produtor-consumidor".
//
// Canais são uma parte fundamental da concorrência em Go.
// Elas permitem que você comunique entre goroutines de forma segura e eficiente.
//
// Canais são usados para sincronizar a execução de goroutines,
// implementar padrões de comunicação,
// e comunicar entre diferentes partes de um programa.
//
// Canais são uma parte fundamental da concorrência em Go.
// Elas permitem que você comunique entre goroutines de forma segura e eficiente.
//
// Canais são usados para sincronizar a execução de goroutines,
// implementar padrões de comunicação,
// e comunicar entre diferentes partes de um programa.
//
// Canais são uma parte fundamental da concorrência em Go.
// Elas permitem que você comunique entre goroutines de forma segura e eficiente.
//
// Canais são usados para sincronizar a execução de goroutines,
// implementar padrões de comunicação,
// e comunicar entre diferentes partes de um programa.
//
