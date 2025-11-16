# Kong Gateway Exporter

![Logo](logo.png)

**Développé par : Najeh TOUMI**  
**Date : Novembre 2025**  
**Technologies :** Go, Docker, Prometheus, Grafana, systemd

---

## 1️⃣ Objectif

Kong Gateway Exporter permet de :

- Scrapper l’API Admin de Kong et vérifier l’état du service.
- Exposer les métriques Prometheus via `/metrics`.
- Être déployé sur VM ou Docker.
- S’intégrer facilement avec Prometheus et Grafana.

---

## 2️⃣ Architecture

![Architecture](images/architecture.png)

**Composants :**

1. Kong Gateway (API Admin sur port 8001)  
2. Exporter (Go) → `/metrics`  
3. Prometheus scrape l’exporter  
4. Grafana visualise les métriques  

---

## 3️⃣ Installation sur VM Debian

```bash
sudo apt update
sudo apt install -y golang
git clone https://github.com/NajehToumi/kong-gateway-exporter.git
cd kong-gateway-exporter
go build -o kong-gateway-exporter main.go
sudo mv kong-gateway-exporter /usr/local/bin/
sudo cp kong-gateway-exporter.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable kong-gateway-exporter
sudo systemctl start kong-gateway-exporter
sudo systemctl status kong-gateway-exporter
curl http://localhost:9542/metrics
