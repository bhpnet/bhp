---
order: 1
---

# 安装

## 安装`go`

:::tip
编译安装 BHP 软件依赖 **Go 1.14 +**。
:::

按照 [官方文档](https://golang.org/doc/install)安装`go`。
别忘记设置您的$GOPATH，$GOBIN和$PATH环境变量，例如：

```bash
mkdir -p $HOME/go/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=$GOPATH/bin" >> ~/.bash_profile
echo "export PATH=$PATH:$GOBIN" >> ~/.bash_profile
source ~/.bash_profile
```
确认已成功安装`go`

```bash
go version
```
## 安装`bhp`

正确配置`go`之后，您应该可以编译并运行`bhp`了。

请确保您的服务器可以访问 google.com，因为我们的项目依赖于google提供的某些库（如果您无法访问google.com，也可以尝试添加代理：`export GOPROXY=https://goproxy.cn`）

```bash
git clone -b <latest-release-tag> https://github.com/bhpnet/bhp
cd bhp
make install
```
如果环境变量配置无误，则通过运行以上命令即可完成BHP的安装。现在检查您的bhp版本是否正确：

```bash
bhpd version --long
bhpcli version --long
```