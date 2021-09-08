# Arquitetura Hexagonal "Ports and Adapters"
Essa implementação tem como objetivo fixar conhecimento adquirido em Arquitetura Hexagonal.
Dessa forma, demonstrado a utilização de adaptadores para consumo de um serviço.
---

<div align="center">
  <img src="https://raw.githubusercontent.com/tiagonevestia/go-ports-and-adapters/main/.github/assets/golang-ports-and-adapters.png" width="400px" />
</div>

## Script de criação da tabela
```
CREATE TABLE planet (
  id uuid NOT NULL,
  name varchar(100) COLLATE pg_catalog."default" NOT NULL,
  climate varchar(30) NOT NULL,
  terrain varchar(30) NOT NULL,
  MovieAppearances integer,
  CONSTRAINT planet_pkey PRIMARY KEY (id)
);
```

## Exemplo utilizando Web Server

Inicalizar Web Server
```
go run main.go http
```

## Pontos importantes

> Crescimento sustentável.

> O software deve ser desenhado por você e não pelo seu framework.

> As peças precisam se encaixar e eventualmente substituídas.
