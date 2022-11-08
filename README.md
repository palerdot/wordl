<p align="center">
  <img src="./logo.png" alt="Wordl - Web like Terminal Wordle">
</p>

> Terminal Wordle.

**WORDL** aims to provide web like `Wordle` experience in the terminal, mainly
- `Slow Reveal` letter colors after guessing.
- Keyboard Hints

<img src="./wordl.gif" alt="Wordl - Web like Terminal Wordle">


### Installation

Binaries are available for different platforms from latest release page - https://github.com/palerdot/wordl/releases/latest

If you have `Go` installed, you can choose to install via one of the methods below.

```
go install github.com/palerdot/wordl@latest
```

`Go` version `1.16`+ is required. 

You can also build from source if you have `Go` installed.

```
git clone https://github.com/palerdot/wordl
go build .
./wordl
```

#### Wordle Words list

Data for Wordle words is present in [`guess/data`](./guess/data) directory. The data was initially seeded from [here](https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee) and [here](https://gist.github.com/cfreshman/d5fb56316158a1575898bba1eed3b5da). Right now, wordle words list is not synced. If you want any words to be added or removed, please submit a PR.

