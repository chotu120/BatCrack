<div align="center">

### 🦇 BatCrack  
A fast & simple hash brute-forcer written in Go

</div>

---

### ⚡ Features
- 🔑 Supports multiple algorithms:
  - `MD5` `SHA-1` `SHA-256` `SHA-384` `SHA-512`
- ⚙️ Cracks hashes with brute force method
- 🐹 Blazing-fast performance with Go
- 🧩 Clean & minimal CLI interface
- 🔄 Built-in charset:
  - lowercase letters (a-z)  
  - uppercase letters (A-Z)  
  - digits (0-9)  
  - common special characters (!@#$%^&*...)  
---

### ⚠️ Limitations
- 🔒 Charset **cannot be customized**  
- ⏳ Length range **cannot be limited or fixed** (always 1 → ∞)  
- 📂 No wordlist support (bruteforce only)

---

### 📦 Installation
```bash
git clone https://github.com/MomboteQ/BatCrack.git
cd BatCrack
go build -o batcrack
```

---

### 🚀 Usage
```bash
# Example: crack an MD5 hash using 8 threads
./batcrack -H 098f6bcd4621d373cade4e832627b4f6 -a md5 -t 8
```

---