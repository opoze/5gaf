# 5G Application Function (Português)

Este repositório é parte de uma [pesquisa](https://github.com/arieldll/trabalho-cmu) em redes móveis 5G

É uma prova de conceito (PoC) de uma "[Aplication Function (AF)](https://github.com/jdegre/5GC_APIs#af-application-function)" designada para se comunicar com uma [Network Exposure Function (NEF)](https://github.com/arieldll/trabalho-cmu) do core 5G.

Esta função foi desenvolvida com a **[linguagem Go Lang](https://go.dev/)** e depende do projeto [NEF](https://github.com/arieldll/trabalho-cmu).

**Possui duas funcionalidades:**
- Iniciar um servidor HTTP que escuta na requests POST na URI /notify
- Encaminhar request POST HTTP de inscreção para a NEF


### Com as seguintes dependências:
- **[Docker](https://www.docker.com/)**
- **[Docker Composer](https://docs.docker.com/engine/reference/commandline/compose/)**
- **[GoLang 1.6](https://go.dev/)**
- **[NEF](https://github.com/arieldll/trabalho-cmu)**

### Configurações:
Existem dois arquivos "const" no projeto `/src/http-client/consts` e `/src/http-server/consts`
com as seguintes configurações:

- Configuração da URL da NEF em /src/http-client/consts
- Configuração da Port HTTP da AF em /src/http-server/consts
- Configuração NEFid em /src/http-client/consts

### How to run
Utilizando Docker e Docker Compose:
`docker-compose up`

Por conveniência o projeto pode rodar sobre Docker mas, pode também rodar utilizando apenas Go Lang:
`go run main.go`

# 5G Application Function (English)

This repository is part of a [research](https://github.com/arieldll/trabalho-cmu) on 5G Mobile Network.

It is an [Aplication Function (AF)](https://github.com/jdegre/5GC_APIs#af-application-function) prove of concept (PoC) designed to comunicate with 5GC [Network Exposure Function (NEF)](https://github.com/arieldll/trabalho-cmu).
The function is developed with **[Go Lang Language](https://go.dev/)** and depends on [NEF](https://github.com/arieldll/trabalho-cmu).

**It has two functionalities:**
- Start an HTTP server listen to POST requests on /notify URI
- Make a HTTP POST to NEF Subscription URL


### With the following dependencies:
- **[Docker](https://www.docker.com/)**
- **[Docker Composer](https://docs.docker.com/engine/reference/commandline/compose/)**.
- **[GoLang 1.6](https://go.dev/)**
- **[NEF](https://github.com/arieldll/trabalho-cmu)**


### Configurations:
There are two const files int the projetct `/src/http-client/consts` and `/src/http-server/consts` with following configurations:

- Configure NEF URL in /src/http-client/consts
- Configure AF HTTP Port in /src/http-server/consts
- Configure NEFid in /src/http-client/consts

### How to run
With Docker and Docker Compose:
`docker-compose up`

For convenience the project runs over docker but can also be run with Go only with:
`go run main.go`








