#!/bin/bash

# Iniciar os contêineres do Docker Compose em segundo plano
docker-compose up -d

# Esperar alguns segundos para que os contêineres tenham tempo para iniciar
sleep 4

# Executar o projeto Go após os contêineres estarem em execução
go run cmd/main.go

# Encerrar os contêineres quando o projeto Go for finalizado (opcional)
docker-compose down
