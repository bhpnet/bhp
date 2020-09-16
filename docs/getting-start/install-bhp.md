---
order: 1
---

# Install BHP

## Install Go

Install `go` by following the [official docs](https://golang.org/doc/install).
Remember to set your `$PATH` environment variable, for example:

```bash
mkdir -p $HOME/go/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=$GOPATH/bin" >> ~/.bash_profile
echo "export PATH=$PATH:$GOBIN" >> ~/.bash_profile
source ~/.bash_profile
```
Under Windows, you may set environment variables(`HOME` or `GO111MODULE`) through the "Environment Variables" 
button on the "Advanced" tab of the "System" control panel. Some versions of Windows 
provide this control panel through the "Advanced System Settings" option inside the 
"System" control panel.

> _NOTE_: **Go 1.14+** is required for the BHP.

## Install `bhp`

After setting up `go` correctly, you should be able to compile and run `bhp`.

Make sure that your server can access to google.com because our project depends on some libraries provided by google. (If you are not able to access google.com, you can also try to add a proxy: `export GOPROXY=https://goproxy.io`)

```bash
git clone -b <latest-release-tag> https://github.com/bhpnet/bhp
export GO111MODULE=on
cd bhp && make install
```

If your environment variables have set up correctly, you should not get any errors by running the above commands. Now check your `bhp` version.

```shell script
bhpd version --long
bhpcli version --long
```
