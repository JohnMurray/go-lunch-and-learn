# Installing Go

## Installation

#### Change Directory

    cd ${HOME}

### Install a binary distribution
 
#### OSX or Linux

```bash
pushd /tmp

# Download the appropriate file
wget https://storage.googleapis.com/golang/go1.8.3.darwin-amd64.tar.gz   # Mac
wget https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz    # Linux

# install to /usr/local
sudo tar -C /usr/local -xzf go1.8.3.darwin-amd64.tar.gz  # Mac
sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz   # Linux
```

## Setup the Workspace

### GOPATH

#### Edit

Edit the appropriate shell file. Likely one of:

```bash
${HOME}/.bash_profile   # Mac
${HOME}/.bashrc         # Linux - Bash
${HOME}/.zshrc          # Zsh (Mac & Linux)
```

And add the following:

```bash
export GOPATH="${HOME}/go"
export PATH="/usr/local/go/bin:${GOPATH}/bin:$PATH"
```

Then create the directory for our `$GOPATH`:

```bash
mkdir -p "${HOME}/go"
```

#### Activate:

Source the appropriate shell file (the one edited in the previous step) so we can make sure your
setup correctly and create our workspace.

```bash
source ${HOME}/.bashrc
```

### Workspace Directories

#### Create the directories on Unix

The `$GOPATH` is where all the magic happens with Go. It's both where we pull down dependencies too as well as
where we do our development. We'll go over this in the Lunch-and-Learn.

```bash
mkdir -p ${GOPATH}/src
mkdir -p ${GOPATH}/pkg
mkdir -p ${GOPATH}/bin
mkdir -p ${GOPATH}/src/github.com/${username}
```

### Check the Go Environment

The last thing is to make sure you can run the following command.

```bash
go env
```

On my machine, I see the following output:

```bash
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/jmurray/projects/go-workspace"
GORACE=""
GOROOT="/usr/local/go"
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build198393980=/tmp/go-build -gno-record-gcc-switches"
CXX="g++"
CGO_ENABLED="1"
PKG_CONFIG="pkg-config"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
```

My go-path is a little different, but this is more or less what you should see. The important parts are to make sure
`GOPATH` poitns to the directory we setup (`$HOME/go`) and `GOROOT` points to the correct install of go (`/usr/local/go`).
