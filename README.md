# Go Sidecar for Kubernetes ConfigMap Sync

Este projeto implementa um **container sidecar** em Go que monitora arquivos `.properties` em um diretório compartilhado e atualiza dinamicamente um **ConfigMap** no Kubernetes com o conteúdo desses arquivos.

## Funcionalidades

- Monitora arquivos `.properties` em um diretório (ex: `/config`)
- Atualiza ou cria automaticamente um `ConfigMap` no Kubernetes
- Projetado para rodar como sidecar ao lado do container principal
- Docker-ready e pronto para uso em ambientes Kubernetes

---

## Como Funciona

1. O sidecar verifica periodicamente o diretório `/config` em busca de arquivos `.properties`.
2. O conteúdo de cada arquivo é lido e incluído como `key-value` no ConfigMap.
3. Se o ConfigMap já existir, ele será atualizado; caso contrário, será criado.

---

## Estrutura do Projeto
    .
    ├── main.go # Código-fonte principal
    ├── Dockerfile # Imagem Docker do sidecar
    ├── go.mod # Dependências do Go
    └── README.md # Este arquivo

---

## Executando Localmente

### Pré-requisitos

- Go 1.24
- Docker
- Acesso ao cluster Kubernetes (via `~/.kube/config` ou via serviço in-cluster)

### Build

```bash
    go build -o sidecar main.go

### Docker

```bash
    docker build -t meu-sidecar:latest .
