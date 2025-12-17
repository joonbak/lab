# lab

Inspired by Tobi Lütke's [try](https://github.com/tobi/try)

Building a version of `try` in Go and calling it `lab`

Still Work in Progress. Will be adding more features!

# Description

All new project directories will be created in `~/lab/experiments` directory

New project creates a directory with dates as the prefix like `2025-12-15-python-experiment`

# Quick Start

```bash
# Clone the repo
git clone https://github.com/joonbak/lab.git

# Build the App
go build .

# Move the executable
sudo mv lab /usr/local/bin/

# Add to your shell (bash/zsh) '~/.zshrc'
lab() {
  eval "$(/usr/local/bin/lab "$@" 2>/dev/tty)"
}
```
