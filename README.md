<div align="center">

![Hexify LOGO](https://i.imgur.com/8AWo5xR.png)

</div>
<h4 align="center">Simple, fast, and feature-rich hex dump tool with customizable output, highlighting.</h4>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/yourpwnguy/hexify">
<!-- <a href="https://github.com/yourpwnguy/hexify/releases"><img src="https://img.shields.io/github/downloads/yourpwnguy/hexify/total"> -->
<a href="https://github.com/yourpwnguy/hexify/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/yourpwnguy/hexify">
<!-- <a href="https://github.com/yourpwnguy/hexify/releases/"><img src="https://img.shields.io/github/release/yourpwnguy/hexify"> -->
<a href="https://github.com/yourpwnguy/hexify/issues"><img src="https://img.shields.io/github/issues-raw/yourpwnguy/hexify">
<a href="https://github.com/yourpwnguy/hexify/stars"><img src="https://img.shields.io/github/stars/yourpwnguy/hexify">
<!-- <a href="https://github.com/yourpwnguy/hexify/discussions"><img src="https://img.shields.io/github/discussions/yourpwnguy/hexify"> -->
</p>

---

`Hexify` is a command-line tool written in Go that generates a colored hex dump of a file. It supports custom width, line limits, and pattern highlighting, and it can output the result to a file.

## Features 💡

- Display the hexadecimal representation of a file.
- Highlight specific byte patterns in the hex dump (e.g., "MZ" header).
- Set the width of each line and limit the number of lines.
- Save the hex dump to a file, without ANSI color codes.

## Installation 🛠️
To install the hexify tool, you can simply use the following commands.

### Windows
```bash
go install -v "github.com/yourpwnguy/hexify/cmd/hexify@latest"
```

### Linux
```bash
go install -v "github.com/yourpwnguy/hexify/cmd/hexify@latest"

# Do the below step only if your "~/go/bin" is not in PATH
cp ~/go/bin/refine /usr/local/bin/
```

## Usage 📘
```yaml
Usage: hexify [options]

Options: [flag] [argument] [Description]

BASIC:
  -f string             Specify the input file to read data from.
  -o string             Specify the output file to write the hex dump to

PROCESSING:
  -l int                Specify the number of lines to print (default: all lines).
  -w int                Specify the number of hex bytes to print per line (default: 16).
  -ha string            Highlight a specific ASCII string pattern (default: 'MZ').

DEBUG:
  -h, --help            Display this help message.
  -v, --version         Check current version
```

## But Why Use Our Tool❓ 

I know there are many hex dump tools available (especially gui ones), but here's why I created Hexify: I wanted to understand how hex dumpers work under the hood, and I noticed a lack of flexible, easy-to-use hex dump CLI tools for Windows. With `Hexify`, you get a tool that’s simple, customizable, and designed to fit seamlessly into your workflow. Whether you're highlighting patterns, customizing output widths, or saving results to a file, `Hexify` is built to give you the control you need, without the bloat.

## Contributing 🤝

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, feel free to open an issue or submit a pull request.
