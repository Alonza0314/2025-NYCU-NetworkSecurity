# 2025-NetworkSecurity

## Prerequisites

- Golang

    - AMD

        ```bash
        wget https://dl.google.com/go/go1.24.5.linux-amd64.tar.gz
        sudo tar -C /usr/local -zxvf go1.24.5.linux-amd64.tar.gz
        mkdir -p ~/go/{bin,pkg,src}
        echo 'export GOPATH=$HOME/go' >> ~/.zprofile
        echo 'export GOROOT=/usr/local/go' >> ~/.zprofile
        echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.zprofile
        echo 'export GO111MODULE=auto' >> ~/.zprofile
        source ~/.zprofile
        ```

    - ARM / Mac

        ```bash
        wget https://dl.google.com/go/go1.24.5.linux-arm64.tar.gz
        sudo tar -C /usr/local -zxvf go1.24.5.linux-arm64.tar.gz
        mkdir -p ~/go/{bin,pkg,src}
        echo 'export GOPATH=$HOME/go' >> ~/.zprofile
        echo 'export GOROOT=/usr/local/go' >> ~/.zprofile
        echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.zprofile
        echo 'export GO111MODULE=auto' >> ~/.zprofile
        source ~/.zprofile
        ```

## Clone the repo

```bash
git clone https://github.com/Alonza0314/2025-NYCU-NetworkSecurity.git
```

## Homeworks

- [Homework2](homework2/README.md)
