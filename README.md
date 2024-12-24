# URL Shortener

A simple URL shortener application written in Go. This application allows you to configure short paths mapped to full URLs, validate the URLs, and serve them via an HTTP server with redirection functionality.

---

## **Features**
- Redirects short URLs to their target URLs.
- Validates URLs with supported schemes (`http`, `https`, `ftp`, `ws`, `wss`, `mailto`, `tel`, `data`, `sftp`).
- Logs total URLs, valid URLs, and invalid URLs from the configuration.
- Provides a default `/hello` endpoint for testing.
- Graceful logging of redirections and errors.

---

## **Getting Started**

### **Prerequisites**
- Go installed on your system (version 1.18 or later).

### **Installation**
1. Clone this repository:
   ```bash
   git clone https://github.com/ah-naf/urlshort.git
   cd urlshort
   ```

2. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

3. Create a configuration file (e.g., `redirect.yaml`) with the following format:
   ```yaml
   - path: "/google"
     url: "https://www.google.com"
   - path: "/github"
     url: "https://github.com"
   - path: "/stackoverflow"
     url: "https://stackoverflow.com"
   ```

---

## **Usage**

### **Run the Application**
1. Start the server:
   ```bash
   go run main.go --config yaml --path redirect.yaml
   ```

2. Open your browser or use `curl` to test the endpoints:
   - Visit `http://localhost:8080/google` to be redirected to Google.
   - Visit `http://localhost:8080/github` to be redirected to GitHub.

---

## **Configuration**

### **Supported Schemes**
The application supports the following URL schemes:
- `http`
- `https`
- `ftp`
- `ws`
- `wss`
- `mailto`
- `tel`
- `data`
- `sftp`

### **Command-Line Flags**
- `--config`: Specify the configuration format (`yaml` or `json`).
- `--path`: Path to the configuration file.

Example:
```bash
go run main.go --config yaml --path redirect.yaml
```

---

## **Code Structure**
- **`main.go`**: The main entry point of the application. Sets up routes, handles redirection, and starts the HTTP server.
- **`utils`**: A package for parsing configuration files (`yaml` and `json`).

---

