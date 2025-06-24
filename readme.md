# 🌐 Microservices with NGINX Reverse Proxy

This project demonstrates a microservices architecture with **Golang** and **Python (Flask)** backend services. Both services run on the **same internal port (8080)** and are reverse-proxied by **NGINX** through Docker containers.

---

## 📦 Tech Stack

- ⚙️ Golang HTTP service (`service_1`)
- 🐍 Python Flask service (`service_2`)
- 🌐 NGINX (reverse proxy and logger)
- 🐳 Docker + Docker Compose
- 🔁 Bridge networking (no host network)

---


## 🔧 How It Works

- Both `service_1` and `service_2` run on the **same internal port (8080)** inside their containers.
- NGINX reverse proxy also runs in its own container and listens on **port 80**, which maps to `localhost:8080` on the host.
- NGINX uses **path-based routing** to forward requests:
  - `/service1/*` → routed to `http://service1:8080`
  - `/service2/*` → routed to `http://service2:8080`
- All inter-container communication is done using **container names as hostnames** (e.g., `service1`, `service2`), enabled by Docker Compose’s default bridge network.
- Health checks ensure both services are up before NGINX starts routing traffic.


---


## 🚀 How to Run

### 🧱 Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### 🔥 Start Everything

<pre lang="markdown">docker-compose up --build</pre>


---

## 📊 Endpoints

| Service     | Path                    | Description      |
|-------------|-------------------------|------------------|
| service_1   | `/service1/ping`        | Health check     |
|             | `/service1/hello`       | Greeting (Go)    |
| service_2   | `/service2/ping`        | Health check     |
|             | `/service2/hello`       | Greeting (Flask) |


---


## 📜 Logs & Health Checks

- NGINX logs incoming requests with time and path.
- You can check logs:

<pre lang="markdown">bash docker logs nginx-proxy</pre>
