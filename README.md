# lab

Inspired by Tobi Lütke's [try](https://github.com/tobi/try)

Building a version of `try` in Go and calling it `lab`

Still Work in Progress. Will be adding more features!

## Description

All new project directories will be created in `~/lab/experiments` directory

New project creates a directory with dates as the prefix like `2025-12-15-python-experiment`

## Usage

```bash
# Create new experiment directory
lab

# List all lab experiments
lab list
```

"lab list" will show a TUI with commands to select, delete and go to directory.

## Quick Start

```bash
# Linux x86_64
curl -sL https://github.com/joonbak/lab/releases/latest/download/lab-linux-amd64 -o lab
chmod +x lab
sudo mv lab /usr/local/bin/

# Mac (Apple Silicon)
curl -sL https://github.com/joonbak/lab/releases/latest/download/lab-darwin-arm64 -o lab
chmod +x lab
sudo mv lab /usr/local/bin/

# Add to your shell (bash/zsh) '~/.zshrc'
lab() {
  local output

  if [ $# -eq 0 ]; then
    output="$(/usr/local/bin/lab 2>/dev/tty)"
  else
    output="$(/usr/local/bin/lab "$@" 3>&1 >/dev/tty 2>&1)"
  fi

  if printf "%s" "$output" | grep -q "^cd "; then
    eval "$output"
  elif [ -n "$output" ]; then
    printf "%s\n" "$output"
  fi
}
```
