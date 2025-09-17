# 🕵️‍♂️ Valkan
![Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)
![Cybersecurity](https://img.shields.io/badge/Focus-Cybersecurity-red?style=for-the-badge&logo=apache)
![Open Source](https://img.shields.io/badge/Open%20Source-Yes-brightgreen?style=for-the-badge&logo=github)
![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue?style=for-the-badge&logo=opensourceinitiative)

> ⚠️ **Disclaimer**
> **Valkan is a real tool**, developed for legitimate **network scanning** and **vulnerability exploitation** activities in controlled and authorized environments.
> Its use is exclusively for **offensive security testing and auditing** purposes with explicit permission.
> **Unauthorized or illegal use is prohibited** by applicable laws, including the [Marco Civil da Internet (Law No. 12,965/2014)](https://www.planalto.gov.br/ccivil_03/_ato2011-2014/2014/lei/l12965.htm).
> The author **is not responsible** for any misuse or damage caused by this tool.

---

<img width="1549" height="782" alt="image" src="https://github.com/user-attachments/assets/38f58661-15e7-471f-916a-e4f295c3361b" />

---

## 🎯 Objective

This project provides a **functional tool** for:

- **Network scanning** on ports **1–1024** or **full range**, detecting open services and potential vulnerabilities.
- Modular engine with components such as **Reporter**, **Detector**, **HTTP module**, and **CLI-based terminal interface**.
- Built in **Go (Golang)** for high performance, **concurrent scanning**, and low memory usage.
- Designed for **real-world offensive security assessments** in authorized and controlled environments.

## 🧰 Technologies

- **Language:** [Go (Golang)](https://golang.org) 

- **Core Libraries & Modules:**
  - [`cobra`](https://github.com/spf13/cobra) – CLI framework for structured command-line interfaces
  - [`net`](https://pkg.go.dev/net) – Network operations (TCP/UDP scanning, IP resolution)
  - [`http`](https://pkg.go.dev/net/http) – HTTP requests and header/banner grabbing
  - [`os`](https://pkg.go.dev/os) – OS-level access for system interaction and info gathering
  - [`os/exec`](https://pkg.go.dev/os/exec) – Execution of system commands when necessary
  - [`fmt`](https://pkg.go.dev/fmt) – Terminal output formatting
  - [`encoding/json`](https://pkg.go.dev/encoding/json) – JSON output for structured reporting
  - [`runtime`](https://pkg.go.dev/runtime) – System architecture and OS detection
  - [`time`](https://pkg.go.dev/time) – Timeout handling, scan delays, and timestamping

  [![My Skills](https://skillicons.dev/icons?i=golang,bash)](https://skillicons.dev)

---

## 🏗️ Architecture

The architecture is **modular**, **concurrent**, and **extensible**, designed for flexibility and performance in real-world security assessments:

- **🔎 Scanner Module**  
  Performs port scanning (range: `1–1024` or `1–65535`) using concurrent routines for fast network enumeration.

- **🧠 Detector Module**  
  Analyzes open ports and services, performs banner grabbing, and identifies potential vulnerabilities or weak configurations.

- **📝 Reporter Module**  
  Outputs results in structured formats (e.g., console, JSON), enabling easy parsing and documentation of scan results.

- **🌐 HTTP Module**  
  Sends requests to web services to extract HTTP headers, status codes, server info, and other metadata.

- **💻 CLI Interface**  
  Built with [`cobra`](https://github.com/spf13/cobra) to provide a clean and interactive command-line experience.

- **⚙️ System Info Layer**  
  Detects system architecture, OS type, kernel version, and other relevant environment data.

- **📦 Future Web UI (In Development)**  
  A user-friendly web dashboard is under development (started on `2025-09-06`) for managing scans and visualizing results.

## ⚙️ Download

### 1. Download in WebSIte

<img width="933" height="469" alt="image" src="https://github.com/user-attachments/assets/91e08e61-8339-4279-84df-28fe1b0b360d" />

### 2. File Valkan
A download will come to your files

<img width="893" height="547" alt="image" src="https://github.com/user-attachments/assets/ebd29f5c-0f9a-4413-ab67-97a3b8ab433e" />

## 💻 Run Valkan

### 1. List Itens
When you arrive at the terminal, you will direct the terminal to where the Valkan file is and use the command:
```bash
ls
```
This command is used to list items.

<img width="642" height="466" alt="image" src="https://github.com/user-attachments/assets/a2088ae6-e99f-457d-a886-7dbbadf6633e" />

### 2. Command
You will enter the command:
```bash
./valkan
```
and finish, congratulations 

<img width="1069" height="662" alt="image" src="https://github.com/user-attachments/assets/aeed9c80-c659-4274-8a2c-deb86df8f88a" />

---

## 🤝 Contribution Guidelines

Contributions are welcome, provided they align with the educational goals of the project.  
If you find bugs, have ideas for improvements, or want to add features, feel free to contribute via pull requests.

To contribute:

1. Fork the repository.  
2. Create a branch with your changes.  
3. Submit a pull request with a detailed description of your changes.

## 📷 General Screenshot

<img width="1329" height="863" alt="image" src="https://github.com/user-attachments/assets/2641ca31-d693-4d4d-95b4-6989bce280e4" />

<img width="1034" height="729" alt="image" src="https://github.com/user-attachments/assets/fae332f9-36e4-4b42-b682-b29b81f1a13a" />


## 📄 License
This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0).  
See the [LICENSE](./LICENSE) file for details.


## 📬 Contact
- Contact the author via [GitHub](https://github.com/Vyzer9)

<img width="351" height="383" alt="image" src="https://github.com/user-attachments/assets/1c883c12-9f16-4064-a752-40ed4edee172" />




>⚠️ Final Notice:
>This project is intended for research, testing, and development purposes in controlled environments only.
>Do not use this code in production systems, unauthorized networks, or for malicious activities.
>The author explicitly disclaims any responsibility for misuse and condemns any form of unethical or illegal usage.

