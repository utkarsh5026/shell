# 🐚 Go Shell Implementation

[![Go Version](https://img.shields.io/badge/Go-1.22-blue.svg)](https://golang.org/doc/go1.22)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![CodeCrafters Progress](https://backend.codecrafters.io/progress/shell/35120eb7-e255-4ee5-a200-744469f3cfc4)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

A POSIX-compliant shell implementation in Go, featuring built-in commands and external program execution capabilities. This project is part of the CodeCrafters "Build Your Own Shell" challenge.

## 📚 Table of Contents

- [🐚 Go Shell Implementation](#-go-shell-implementation)
  - [📚 Table of Contents](#-table-of-contents)
  - [✨ Features](#-features)
  - [🚀 Installation](#-installation)
    - [Prerequisites](#prerequisites)
  - [📋 Built-in Commands](#-built-in-commands)
  - [🔮 Future Prospects](#-future-prospects)
  - [🙏 Big Thanks To](#-big-thanks-to)

## ✨ Features

- Interactive command-line interface
- Built-in command support
- External program execution
- Path resolution
- Home directory expansion
- Error handling and status reporting

## 🚀 Installation

### Prerequisites

- Go 1.22 or higher
- Git

## 📋 Built-in Commands

| Command            | Description                     | Example            |
| ------------------ | ------------------------------- | ------------------ |
| `exit [code]`    | Exit the shell with status code | `exit 0`         |
| `echo [text]`    | Display text to stdout          | `echo Hello`     |
| `type [command]` | Show command type               | `type echo`      |
| `pwd`            | Print working directory         | `pwd`            |
| `cd [directory]` | Change directory                | `cd ~/Documents` |

## 🔮 Future Prospects

- [ ] Command history and search
- [ ] Tab completion with smart suggestions
- [ ] Pipeline support (`|`, `>`, `>>`)
- [ ] Background process handling (`&`)
- [ ] Signal handling (`Ctrl+C`, `Ctrl+Z`)


## 🙏 Big Thanks To

- The awesome folks at [Codecrafters](https://codecrafters.io/)
