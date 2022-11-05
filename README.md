<p align="center">
  <img src="./logo.png" alt="Wordl - Web like Terminal Wordle">
</p>

> Terminal Wordle.

**WORDL** aims to provide web like `Wordle` experience in the terminal, mainly
- `Slow Reveal` letter colors after guessing.
- Keyboard Hints

<img src="./wordl.gif" alt="Wordl - Web like Terminal Wordle">


### Installation

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

Right now, Go needs to be installed on the user's machine to get the binary executable. There are plans to provide binaries for different platforms, so that people without `Go` installed can try out **wordl.** Please refer this [issue](https://github.com/palerdot/wordl/issues/1) for more details.

#### Wordle Words list

Data for Wordle words is present in [`guess/data`](./guess/data) directory. The data was initially seeded from [here](https://gist.github.com/cfreshman/a7b776506c73284511034e63af1017ee) and [here](https://gist.github.com/cfreshman/d5fb56316158a1575898bba1eed3b5da). Right now, wordle words list is not synced. If you want any words to be added or removed, please submit a PR.

