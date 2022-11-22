# 5G Application Function

This repository is part of a research on 5G Mobile Network.

It is an Aplication Function AF) prove of concept(PoC) designed to comunicate with 5GC Network Exposure Function(NEF).
The function is developed with **Go Language**

**It has two functions:**
- Start an HTTP server listen to POST requests on /notify URI
- Make a HTTP POST to NEF Subscription URL


### Dependencies
Project uses **docker** and **docker-composer**.

### Configurations
There are two const files int the projetct `/src/http-client/consts` and `/src/http-server/consts`
with following configurations:

- Configure NEF URL in /src/http-client/consts
- Configure AF HTTP Port in /src/http-server/consts
- Configure NEFid in /src/http-client/consts

### How to run
`docker-compose up`








