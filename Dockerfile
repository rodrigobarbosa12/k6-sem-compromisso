FROM golang:1.23.3 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o myapp ./app

FROM ubuntu:22.04

# Instalar dependências necessárias (curl e gnupg2)
RUN apt-get update && apt-get install -y curl gnupg2 && rm -rf /var/lib/apt/lists/*

# Adicionar o repositório do k6 e instalar o k6 dentro do container
RUN curl -fsSL https://dl.k6.io/key.gpg | gpg --dearmor > /usr/share/keyrings/k6-archive-keyring.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | tee /etc/apt/sources.list.d/k6.list \
    && apt-get update \
    && apt-get install -y k6 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/myapp /usr/local/bin/myapp

COPY ./tests /app/tests

WORKDIR /app

EXPOSE 8080

CMD ["/usr/local/bin/myapp"]
