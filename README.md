<h1 align="center" style="display:flex;align-items:center;justify-content:center;">
<img alt="Logo" src="./.github/logo.png"/>
CNVRTR
</h1>

<p  align="center">
  <a  href="https://www.linkedin.com/in/luan-souza-6b07b1171/">
    <img  alt="Luan Souza"  src="https://img.shields.io/badge/-Luan Souza-282A36?style=for-the-badge&logo=Linkedin&logoColor=white"  />
  </a>

  <img  alt="Repository size"  src="https://img.shields.io/github/repo-size/LuanSilveiraSouza/cnvrtr?color=282A36&style=for-the-badge">

  <a  href="https://github.com/LuanSilveiraSouza/cnvrtr/commits/master">
    <img  alt="GitHub last commit"  src="https://img.shields.io/github/last-commit/LuanSilveiraSouza/cnvrtr?color=282A36&style=for-the-badge">
  </a>

  <img  alt="License"  src="https://img.shields.io/badge/license-MIT-282A36?&style=for-the-badge">
</p>

# :pushpin: Sumary

* [Introduction](#paperclip-introduction)
* [Setup and Instalation](#computer-setup-and-instalation)
* [How to Use](#rocket-how-to-use)
* [Bugs and Issues](#bug-bugs-and-issues)
* [Contributing](#man_mechanic-contributing)
* [License](#books-license)

# :paperclip: Introduction

CNVRTR is a Command Line Interface written in Golang to convert mainly plain text to another encoding formats and vice-versa.

# :computer: Setup and Instalation

```bash
# Clone Repository
$ git clone https://github.com/LuanSilveiraSouza/cnvrtr.git

# Make sure that you have in your machine the dependencies of the project
$ go mod download

# Execute the project with
$ go run main.go

# Compile with
$ go install
```

It is recommended that you use ```go install``` to access the cli in whatever place of your machine.

# :rocket: How to use

```bash
# To get all commands:
$ cnvrtr help

# Examples

$ cnvrtr ascii hey
# Output: 104 101 121

$ cnvrtr b64 aGVsbG8= -d
# Output: hello

$ cnvrtr b64 amazing -c
# Output: YW1hemluZw==

# -d/--decode: decode instead of encode
# -c/--clipboard: copy result to clipboard
```

# :bug: Bugs and Issues

Feel free to open new issues and colaborate with others issues in [CNVRTR Issues](https://github.com/LuanSilveiraSouza/cnvrtr/issues)

# :man_mechanic: Contributing

This project is completely open source, so if you want, consider forking it and making PRs with the changes you think its helpful to the project growns.

# :books: License

Released in 2020 under [MIT License](https://opensource.org/licenses/MIT)

Made with :heart: by Luan Souza.