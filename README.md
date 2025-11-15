

# 🧹 Flutter Cleaner CLI  
A modern, fast, and interactive command‑line tool to clean Flutter projects, reclaim disk space, and optimize your development workflow.

Built with **Go**, powered by **Lipgloss**, **Survey**, **go-pretty**, and designed for speed, usability, and a premium developer experience.

---

## ✨ Features

- 🚀 **Clean Flutter build folders** (single or all projects)
- 🧙 **Interactive Wizard mode**
- 📦 **Suggestions engine** — ranks projects by space usage + age
- 📊 **Stats mode** — see total build space usage
- 🔍 **Scan mode** — fast or deep scanning of directories
- 💾 **Dry run mode** — preview what will be cleaned
- ⚡ **Parallel cleaning** (configurable)
- 🎨 **Beautiful UI** with tables, colors, and progress bars
- 🛠️ **Cross-platform** (macOS, Linux, Windows)

---

## 📥 Installation

### Homebrew (macOS / Linux) — *coming soon*
```
brew install vipinkashyap/tap/fclean
```

### Manual Install (local build)

```
git clone https://github.com/vipinkashyap/flutter-cleaner-cli
cd flutter-cleaner-cli
go build -o fclean .
sudo mv fclean /usr/local/bin/
```

Verify:

```
fclean --help
```

---

## 🧙 Wizard Mode (Recommended)

Start the interactive flow:

```
fclean wizard
```

You’ll get:

- Scan
- Suggest
- Clean all
- Stats
- Exit

---

## 🔍 Scan for Flutter Projects

```
fclean scan ~
```

Fast mode (skips caches/system dirs):

```
fclean scan ~ --fast
```

---

## 💡 Suggestions (Smart Ranking)

Analyze your machine for large Flutter build folders:

```
fclean suggest ~
```

It produces:

- age of build folder
- build size
- project path
- saved JSON for future clean

To clean a suggested project:

```
fclean clean --suggest 1
```

---

## 🧹 Clean Projects

### Clean one project
```
fclean clean /path/to/project
```

### Clean all projects under the current directory
```
fclean clean --all
```

### Dry run mode (preview)
```
fclean clean --all --dry-run
```

### Parallel cleaning (default: 4)
```
fclean clean --all --parallel 8
```

---

## 📊 Stats

Show total space of all Flutter build folders:

```
fclean stats ~
```

---

## 🧩 Architecture Overview

```
flutter-cleaner-cli/
│
├── cmd/         # All commands: scan, clean, stats, suggest, wizard
├── ui/          # Styled UI layer: colors, tables, prompts, progress bars
├── main.go      # Entry point
└── go.mod
```

- **cmd/**: Cobra-based modular CLI structure  
- **ui/**: Shared components using Lipgloss + go-pretty + Survey  
- **clean.go**: Parallel cleaning, timing, dry-run logic  
- **scan.go**: Progress bars + file traversal  
- **tables.go**: Beautiful pretty tables  
- **styles.go**: Global colors & styles  

---

## 🔧 Roadmap

- 🔄 Auto-update support
- 🍺 Homebrew tap publishing
- 🐧 .deb, .rpm packaging
- 🔧 Version command & semantic releases
- 🚀 GitHub Actions release pipeline
- 📦 GoReleaser integration for cross-platform binaries

---

## 🤝 Contributing

PRs are welcome — especially around:

- UX improvements
- Performance enhancements
- New scan heuristics
- Advanced cleaning rules

---

## 📄 License

MIT License.

---

## ⭐ Support

If this tool saves you time or cleans gigabytes from your machine,  
consider starring the repo ⭐ on GitHub — it helps visibility and adoption.
