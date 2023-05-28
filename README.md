# book-service
![gopher](gopher.png)
## Prerequisites
If not already installed, install Golang from [https://golang.org/](https://golang.org/)


Running/building this app requires installing [GCC](https://gcc.gnu.org/).

* Linux

GCC is included on most distros, you can validate it with:
  ```sh
    $ gcc --version
  ```
If not installed, use for example apt:
```sh
    $ sudo apt update
    $ sudo apt install build-essential
  ```

* Mac

GCC is included in every recent version, if using older one easiest way to install is with [Homebrew](https://brew.sh/)
  ```sh
    $ brew install gcc
  ```

* Windows

Install GCC, such as [TDM GCC](https://jmeubank.github.io/tdm-gcc/download/)

## Building & Running

  ```sh
  git clone https://github.com/teepli/book-service.git
  # Run the app
  make run
  # Or build binaries
  make build
  # Or build and run with docker
  make docker-build && make docker-run
  ```
