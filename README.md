# MNEE SDK API Wrapper (Go)

A high-performance Golang REST API wrapper for the [MNEE SDK](https://github.com/mnee-xyz/go-mnee-1sat-sdk). This server exposes endpoints for MNEE token operations including Balance, History, UTXOs, and Transfers.

## 🚀 Quick Start

You can run this server instantly using Docker. No Go installation required.

### Prerequisites
- [Docker](https://www.docker.com/get-started) installed on your machine.

### 1. Run with Docker CLI

Replace `your_api_key_here` with your actual MNEE API key.

```bash
docker run -d \
  -p 8080:8080 \
  -e MNEE_API_KEY="your_api_key_here" \
  -e MNEE_ENV="sandbox" \
  --name mnee-api \
  princerockwallet/mnee-go-api:latest