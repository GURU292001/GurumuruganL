# GurumuruganL

Programming Language : # go language

## 🛠️ Install the Latest Go Version

Follow these steps to download and install the updated Go programming language.

---

### 📅 1. Download Go

Visit the official Go downloads page:

[https://go.dev/dl/](https://go.dev/dl/)

Download the installer or archive appropriate for your system:

* **Windows:** `.msi` installer
* **Linux:** `.tar.gz` archive

> **Tip:** Always choose the latest stable release (e.g., Go 1.xx.x)

---

### 🪟 Windows Installation

1. Download the `.msi` installer from the downloads page.
2. Run the installer.
3. Follow the setup wizard instructions to complete installation.

---

### 🐧 Linux Installation

1. **Remove any previous Go installation:**

   ```bash
   sudo rm -rf /usr/local/go
   ```

2. **Extract the downloaded archive (replace `go1.xx.x.linux-amd64.tar.gz` with the file you downloaded):**

   ```bash
   sudo tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
   ```

3. **Add Go to your `PATH`:**

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```

   To make this permanent, add the above line to your `~/.bashrc`, `~/.profile`, or `~/.zshrc`.

---

### ✅ Verify Installation

After installation, open a terminal or command prompt and run:

```bash
go version
```

You should see output similar to:

```
go version go1.xx.x linux/amd64
```

or

```
go version go1.xx.x windows/amd64
```

---

### 📚 References

* [Official Go Installation Guide](https://go.dev/doc/install)


## ▶️ Running the Go Program

1. **Open the folder** where your Go file is located.
2. **Open a terminal** in that folder.
3. **Run the Go file using:**
   ```bash
   go run <filename>.go
   ```
   ### Example
   ```bash
   go run Problem-1.go
   ```

---
