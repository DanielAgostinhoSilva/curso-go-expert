### Comando para geração de codigo com PROTOC

### geração dos servico e entidades com o proto
```shell
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

### executando um cliente para realizar chamadas para o GPRC
````shell
evans -r repl
````