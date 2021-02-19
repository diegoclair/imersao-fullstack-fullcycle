# Imersão Full Stack & FullCycle - CodePix

## Descrição

Repositório do Apache Kafka (Mensageria)

**Importante**: Roda-lo antes de qualquer apliação

## Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):
```
127.0.0.1 host.docker.internal
```
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

## Rodar a aplicação

Execute os comandos:

```
docker-compose down && docker-compose up
```
***Importante***: ```docker-compose down``` é para destruir os volumes anteriores, é necessário senão para não dar erro ao subir o container novamente 