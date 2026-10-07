# 🔄 GoGameV3 — File Converter

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)


## 📋 Table of Contents

- [Features](#-features)
- [Requirements](#-requirements)
- [Quick Start](#-quick-start)
- [Run Without Docker](#-run-without-docker)
- [Usage](#-usage)
- [Project Structure](#-project-structure)

---

## ✨ Features

<!-- TODO: fill in real features -->

- Convert files from one format to another
- Supported formats: ...
- Runs in Docker, no Go installation needed

---

## 🧰 Requirements

| Method | What you need |
|---|---|
| Docker | [Docker](https://docs.docker.com/get-docker/) |
| Local | [Go](https://go.dev/dl/) 1.21+ |

---

## 🚀 Quick Start

### Run with Docker

**1. Build the image:**

```bash
docker build --tag gogamev3 .
```

**2. Run the container:**

```bash
docker run gogamev3
```

> 💡 The converter needs access to your files. If it reads and writes files on disk, mount a folder with `-v`:
>
> ```bash
> docker run -v "$(pwd)/data:/data" gogamev3
> ```
>
> Files from the local `data` folder will be visible inside the container as `/data`, and the conversion output will appear there too.

---

## 💻 Run Without Docker

```bash
git clone <repository-url>
cd GoGameV3

go run .
```

Build an executable:

```bash
go build -o gogamev3 .
./gogamev3
```

---

## 🛠 Usage

<!-- TODO: show a real example command -->

```bash
# example
./gogamev3 input.ext output.ext
```

| Option | Description |
|---|---|
| ... | ... |

---

## 🗂 Project Strucвture

```
GoGameV3/
├── main.go        # entry point
├── Dockerfile     # image build
├── go.mod         # dependencies
└── ...
```


---
