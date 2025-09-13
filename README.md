# Go Auth API

API básica em **Go** com autenticação via **JWT** e banco de dados **Postgres**.  
Atualmente inclui endpoints:

- `POST /signup` – criar usuário  
- `POST /login` – autenticar usuário  
- `GET /me` – pegar dados do usuário logado  
- `GET /ping` – endpoint de teste com data/hora

---

## Tecnologias

- Go (Golang)  
- Gin (framework HTTP)  
- GORM (ORM)  
- PostgreSQL  
- Docker & Docker Compose  

---

## Rodando o projeto

1. Clone o repositório:

```bash
git clone <URL_DO_REPOSITORIO>
cd go-auth-api
```

2. Configure variáveis de ambiente no `.env` (exemplo já incluso).  

3. Suba os containers:

```bash
docker compose up --build
```

4. Acesse os endpoints:

- `http://localhost:9000/ping`  
- `http://localhost:9000/signup`  
- `http://localhost:9000/login`  
- `http://localhost:9000/me`  

> A porta `9000` é a que a API escuta dentro do container.

---

## Observações

- O banco de dados é inicializado com Postgres 16 e volume persistente (`db_data`).  
- O JWT é usado para autenticação de endpoints privados.  
- Para desenvolvimento, a API tenta se conectar automaticamente ao banco e reinicia em caso de falha (configuração `restart: always` no Docker).

