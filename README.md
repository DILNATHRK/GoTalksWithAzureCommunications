# 📬 GoTalkWithACS

![Go Version](https://img.shields.io/badge/go-1.20+-blue)
![Azure](https://img.shields.io/badge/azure-communication_services-blueviolet)
![Contributions](https://img.shields.io/badge/contributions-welcome-brightgreen)
![License](https://img.shields.io/badge/license-MIT-yellow)

> A collection of use-case modules built with **Golang** to demonstrate how to work with [Azure Communication Services (ACS)](https://learn.microsoft.com/en-us/azure/communication-services/). 🚀

## 📁 Repository Structure

```
GoTalkWithACS/
│
├── SendEmailWithACS/                  # Email module with ACS
│   ├── handler/                       # Core email logic
│   ├── helper/                        # HMAC, hashing, formatting utilities
│   ├── model/                         # Payload and request models
│   └── email.go                       # Gin route to trigger email send
│   └── .env   
│
├── LICENSE
└── README.md
```

---

## ✨ Features

- ✅ Modular design for each ACS use-case
- ✅ Written in idiomatic Go
- ✅ Uses `gin-gonic` for RESTful APIs
- ✅ Follows `.env` pattern for clean secret management
- ✅ Supports contribution of new modules

---

## 🚀 Getting Started

### 1️⃣ Clone the repo

```bash
git clone https://github.com/your-username/GoTalkWithACS.git
cd GoTalkWithACS/SendEmailWithACS
```

### 2️⃣ Set up your environment

Create a `.env` file inside `SendEmailWithACS/`:

```env
Refer .env.example
```

### 3️⃣ Run the server

```bash
go run email.go
```

## 🛠️ Built With

- [Golang](https://golang.org/)
- [Azure Communication Services](https://azure.microsoft.com/en-us/products/communication-services/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)

---

## 🤝 Contribution Guide

Want to add your own ACS module (like SMS or WhatsApp)? Awesome!

### 📂 Folder Naming Convention

- Each module must be a top-level folder: `SendSmsWithACS/`, `SendChatWithACS/`, etc.

### 🔧 Development Guidelines

- Use `gin` for routes
- Follow `.env` pattern
- Isolate helpers and models into their respective folders

### 👇 To contribute:

1. Fork this repo
2. Create your feature branch: `git checkout -b feature/AddSmsModule`
3. Commit your changes: `git commit -m 'Add SMS send module'`
4. Push to the branch: `git push origin feature/AddSmsModule`
5. Open a pull request

---

## 👨‍💻 Authors

- https://github.com/DILNATHRK – Maintainer

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

## 🌐 Connect with Azure Communication Services

>  Want to explore more? Check out the official [Azure ACS Docs](https://learn.microsoft.com/en-us/azure/communication-services/).

```

---

