# Instalação

1- Rode o arquivo docker-compose.yaml no terminal

    docker-compose up -d

2- Acesse o terminal do container do mysql

    docker-compose exec mysql bash

3- Faça o login do mysql

    mysql -uroot -p goexpert

4- Crie a seguinte tabela products

    create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));