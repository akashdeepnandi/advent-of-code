# 🎄 Advent of Code

## 📑 Table of Contents

- [About](#about)
- [Languages](#languages)
- [Coding Principles](#coding-principles)
- [Usage](#usage)
- [References](#references)

---

## 📖 About

This repository tracks my solutions for [Advent of Code](https://adventofcode.com).

My primary language of choice this year is **Golang**, but I’m also experimenting with solutions in other languages to compare performance and styles.

---

## 💻 Languages

| Language   | Badge                                                                                |
| ---------- | ------------------------------------------------------------------------------------ |
| Go         | ![Go](https://img.shields.io/badge/Go-1.22-blue?logo=go)                             |
| Bash       | ![Bash](https://img.shields.io/badge/Bash-Shell-lightgrey?logo=gnu-bash)             |
| JavaScript | ![JavaScript](https://img.shields.io/badge/JavaScript-ES2023-yellow?logo=javascript) |
| Python     | ![Python](https://img.shields.io/badge/Python-3.12-blue?logo=python)                 |
| C++        | ![C++](https://img.shields.io/badge/C++-20-orange?logo=cplusplus)                    |
| Rust       | ![Rust](https://img.shields.io/badge/Rust-1.80-red?logo=rust)                        |

Each day’s solution is organized by **year/day/language**, for example:

```

y15/d01/golang/
y15/d01/python/

```

---

## 🧭 Coding Principles

When solving problems across different languages, I follow a few guiding principles:

- **Readability over cleverness**  
  Code should be easy to understand. I avoid overly dense one-liners or language tricks that hurt clarity.

- **Core logic written manually**  
  For the main algorithm, I don’t lean on built-in shortcuts (e.g., regex magic or library functions) — the point is to _learn by implementing_.  
  Utility functions like `min`/`max` are fair use.

- **Fair optimizations only**  
  I won’t over-optimize unless it’s simple and natural (e.g., using a `byte slice` instead of repeated string concatenation).  
  Extreme micro-optimizations are avoided.

- **Consistent naming**  
  Functions and variables are named for clarity and intent, not brevity.

- **Cross-language comparability**  
  I try to keep the structure of solutions similar across languages so I can study differences in style, idioms, and performance.

---

## ▶️ Usage

Run solutions in different languages with the following commands:

### JavaScript ![JavaScript](https://img.shields.io/badge/JavaScript-ES2023-yellow?logo=javascript)

```bash
cd yYY/dDD/javascript
node index.js input
```

### Go ![Go](https://img.shields.io/badge/Go-1.22-blue?logo=go)

```bash
cd yYY/dDD/golang
go run . -f input
```

### Python ![Python](https://img.shields.io/badge/Python-3.12-blue?logo=python)

```bash
cd yYY/dDD/python
python solve.py input
```

### Rust ![Rust](https://img.shields.io/badge/Rust-1.80-red?logo=rust)

```bash
cd yYY/dDD/rust
cargo run input
```

### C++ ![C++](https://img.shields.io/badge/C++-20-orange?logo=cplusplus)

```bash
cd yYY/dDD/cpp
g++ main.cpp helper.cpp -o solve && ./solve input
```

### Bash ![Bash](https://img.shields.io/badge/Bash-Shell-lightgrey?logo=gnu-bash)

```bash
cd yYY/dDD/bash
bash solve.sh input
```

---

## 📚 References

- [Golang Performance Benchmark](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
