# 🕵️‍♂️ Valkan - Network Scanner & Exploration Vulnerability
[![License: AGPL v3](https://img.shields.io/badge/License-AGPLv3-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.20%2B-blue.svg)](https://go.dev/)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20windows-lightgrey)
![Status](https://img.shields.io/badge/type-cybersecurity--scanner-success)


> ⚠️ **Disclaimer**  
> Valkan is a powerful tool intended solely for security testing and vulnerability exploitation in controlled environments and with explicit authorization. Unauthorized use is prohibited and may violate laws (e.g., Marco Civil); the author is not responsible for any misuse or damage.

## 📌 Table of Contents

- [Objective](#-objective)
- [Website](#-website)
- [Features](#-features)
- [Technologies WebSite](#-technologies-website)
- [Technologies Tool](#-technologies-tool)
- [Architecture](#️-architecture)
- [Download](#️-download)
- [Run Valkan](#-run-valkan)
- [Contribution Guidelines](#-contribution-guidelines)
- [License](#-license)
- [Contact](#-contact)


## Objective

This project provides a **functional tool** for:

- **Network scanning** on ports **1–1024** or **full range**, detecting open services and potential vulnerabilities.
- Modular engine with components such as **Reporter**, **Detector**, **HTTP module**, and **CLI-based terminal interface**.
- Built in **Go (Golang)** for high performance, **concurrent scanning**, and low memory usage.
- Designed for **real-world offensive security assessments** in authorized and controlled environments.

## Website
Here: [WebSite](https://valkan.vercel.app/)
The website was designed with ease, focused on helping the community download and install Valkan easily. Just click Download and Valkan will be ready to go. We'll be creating full documentation in the future...

## Features
- [x] Concurrent port scanning (TCP/UDP)
- [x] Banner grabbing
- [x] Firewall detection (silent)
- [x] JSON and terminal reporting
- [x] CLI interface with Cobra
- [x] Web UI
- [ ] OS fingerprinting (coming soon)
- [ ] CVE matching (in development)
- [ ] Web technology detection
- [ ] Directory fuzzing
- [ ] Scriptable engine (Go or Lua-based)
- [ ] HTML and CSV export
- [ ] Host discovery (ICMP, ARP, TCP ping)
- [ ] Terminal UI (TUI)
- [ ] Vulnerable configuration detection (e.g. anonymous FTP)

## Technologies WebSite
The website is simple and purposefully designed as a tool, not a product meant to be displayed on display. If you'd like to contribute to the creation and updating of the website, send a message to Vyzer9.
The website was made with the languages: 

- **Language:** TypeScript, JavaScript, Node.js, React, Bun, Vite, HTML, CSS, Tailwind

## Technologies Tool

- **Language:** [Go (Golang)](https://golang.org) and others

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

---

## Architecture

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

## Download

### 1. Download in WebSIte

<img width="1908" height="921" alt="image" src="https://github.com/user-attachments/assets/76743f15-846c-4cdf-8df3-ab2043ea1378" />

## Run Valkan

### 1. List Itens
When you arrive at the terminal, you will direct the terminal to where the Valkan file is and use the command:
```bash
ls
chmod +x ./Valkan
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

## Contribution Guidelines

Contributions are welcome, provided they align with the educational goals of the project.  
If you find bugs, have ideas for improvements, or want to add features, feel free to contribute via pull requests.

To contribute:

1. Fork the repository.  
2. Create a branch with your changes.  
3. Submit a pull request with a detailed description of your changes.


## License
This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0).  
See the [LICENSE](./LICENSE) file for details.


## Contact
- Contact the author via [GitHub](https://github.com/Vyzer9)

<img width="339" height="365" alt="image" src="https://github.com/user-attachments/assets/68f5e6aa-4046-48ed-aecb-1bdf71567bc4" />




