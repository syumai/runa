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

```sh
$ runa 33 90     # show characters in 33 to 90
# !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ

$ runa 0x21 0x5a # show characters in 0x21 to 0x5a
# !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ

$ runa -c abcde  # show code point of `abcde`
# 97 98 99 100 101 102 103
```

