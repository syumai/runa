# Runa

* Runa is a utility to show character from given range of code points of Unicode.

## Installation

```console
go get -u github.com/syumai/runa
```

## Usage

* Runa accepts
  - Basic number (like: 33)
  - Hex number (like: 0xaa)

```console
runa 33 126    # show characters in 33 to 126
runa 0x41 0x5a # show characters in 0x41 to 0x5a
```

