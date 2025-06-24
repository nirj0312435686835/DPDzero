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

## 🏗️ Project Structure

.
├── docker-compose.yml # Manages all services
├── nginx/
│ ├── nginx.conf # Reverse proxy config
│ └── Dockerfile # Builds NGINX container
├── service_1/
│ ├── main.go # Go service with two endpoints
│ ├── Dockerfile
│ └── README.md # Local docs for service_1
├── service_2/
│ ├── app.py # Flask service with two endpoints
│ ├── Dockerfile
│ └── README.md # Local docs for service_2
└── README.md # You are here

yaml
Copy
Edit

---

## 🔧 Services Overview

### 🟢 `service_1` (Go HTTP Server)
- Runs on port `8080` inside its container
- Endpoints:
  - `GET /ping` → health check
  - `GET /hello` → greeting message

### 🟠 `service_2` (Flask App)
- Also runs on port `8080` inside its container
- Endpoints:
  - `GET /ping` → health check
  - `GET /hello` → greeting message

> Note: Both services run on port 8080 internally. NGINX handles the routing based on path.

---

## 🌐 NGINX Reverse Proxy

The NGINX container listens on **port 80** internally and maps to **host port 8080**. It routes requests as follows:

| URL                          | Routed To           |
|-----------------------------|---------------------|
| `/service1/*`               | `service_1:8080/*`  |
| `/service2/*`               | `service_2:8080/*`  |

It also logs incoming requests with timestamps using a custom format.

---

## 🔍 Health Checks

Docker Compose uses built-in `healthcheck` on both services to ensure they are up before NGINX connects.

---

## 🚀 How to Run

### 🧱 Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### 📦 One-Command Build & Run

```bash
docker-compose up --build
Docker will:

Build both services and NGINX container

Start everything in a bridge network

Make the app available at http://localhost:8080

🔗 Available Endpoints
Once the services are running, you can test them via:

Service 1:

http://localhost:8080/service1/ping

http://localhost:8080/service1/hello

Service 2:

http://localhost:8080/service2/ping

http://localhost:8080/service2/hello

📜 Example Logs
You can view request logs in the terminal or inspect the logs with:

bash
Copy
Edit
docker logs nginx-proxy
Example:

bash
Copy
Edit
01/Jul/2025:10:00:00 +0000 - 172.18.0.1 - GET /service1/hello HTTP/1.1
💡 Notes
Both services use port 8080 inside their container

NGINX handles path-based routing (/service1, /service2)

You can scale services independently and extend this for more microservices

📚 Related Files
Each service also contains its own README.md explaining how to run it independently.

🙌 Credits
Developed by [Your Name].
Feel free to fork, improve, or use in your own microservice architectures.


