# 🚀 K6 sem compromisso

Este projeto implementa uma API simples em Go que permite a execução de testes de carga utilizando o K6. A API está configurada para executar um teste K6 localizado em `tests/index.mjs` sempre que o endpoint `/run-test` for chamado. Esse projeto é ideal para validar o desempenho e a resiliência de endpoints, garantindo que o sistema funcione adequadamente sob carga.

## 📋 Funcionalidades

- **Executar Testes de Carga com K6**: A API aceita requisições HTTP e utiliza o K6 para executar um script de teste de carga. 
- **Script Customizável**: O script do K6, que define os cenários de teste, está localizado em `tests/index.mjs` e pode ser ajustado para atender a diferentes necessidades de teste.

## 📂 Estrutura do Projeto


```plaintext
Clean Architecture Go Lang (basic project) 
myproject/
├── app
│   └── main.go                # Código principal da aplicação
├── interfaces
│   ├── authApi.go             # Endpoint que executa o script de teste K6
│   └── helloWorld.go          # Exemplo de endpoint Hello World (opcional)
├── tests
│   └── index.mjs              # Script de teste principal do K6
├── Dockerfile                 # Dockerfile para construir e rodar a aplicação
├── go.mod                     # Arquivo de dependências do Go
└── README.md                  # Documentação do projeto
```
## 🛠️ Como Executar
**Pré-requisitos**
- Docker: Certifique-se de ter o Docker instalado na sua máquina para uma execução fácil e eficiente.
- K6: Já configurado e instal

**Passo a Passo para Executar a API**
1. Clonar o Repositório:
```bash
git clone https://github.com/rodrigobarbosa12/k6-sem-compromisso.git
cd k6-sem-compromisso
```
2. Construir a Imagem Docker:
```bash
docker build -t k6-sem-compromisso .
```
3. Iniciar o Container:
```bash
docker run -p 8080:8080 k6-sem-compromisso
```

4. A API estará disponível em http://localhost:8080.

## 🚦 Executando o Teste de Carga
Uma vez que o container esteja em execução, você pode testar o endpoint chamando /run-test para disparar o teste de carga com o script index.mjs:

**Endpoint**
- URL: http://localhost:8080/run-test
- Método: GET

**Exemplo de Requisição**
Use o curl ou qualquer ferramenta de sua escolha:

```bash
curl -X GET http://localhost:8080/run-test
```

**Resultado**
A API retorna o resultado do teste em formato JSON, contendo detalhes sobre o desempenho do sistema com o teste executado.

## 🔧 Personalização do Script de Teste
O script do K6 pode ser editado no arquivo tests/index.mjs para criar cenários específicos de teste, ajustando parâmetros como número de usuários virtuais (VUs), duração do teste, endpoints, e mais.
