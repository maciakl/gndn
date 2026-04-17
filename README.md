# gndn

GNDN stands for "Go Nowhere, Do Nothning". The purpose of this tool is precisely that: to do nothing while generating a visually interesting output.

This program will run continously, generating output in the form of progress bars, spinners, bar charts, debuging messages and various other visual elements while doing nothing. It is intended to be used a placeholder or visual representation of work in progress.

![gndn](https://github.com/user-attachments/assets/7b4a271a-7729-4b72-8a8b-c3afb049cdee)

The output is radomized, and unique for every run.

## Usage

```bash
Usage: gndn [options]
Options:
  -v, --version    Print version information and exit
  -h, --help       Print this message and exit
```

## Installation

### Multi-Platform:

You can install gndn using `go install`:

```bash
go install github.com/maciakl/gndn@latest
```

### macOS and Linux:

Use [grab](https://github.com/maciakl/grab):

```bash
grab maciakl/gndn
```

### Windows:

Use [scoop](https://scoop.sh):

```bash
scoop add maciak https://github.com/maciakl/bucket
scoop update
scoop install gndn
```
