# Use Mage instead. See [NOTES_MAGE.md](./NOTES_MAGE.md)
# Make notes
* Installed go in WSL, but ended up not using. Ran go in container instead (see below)
```bash
apt  install golang-go
go version go1.18.1 linux/amd64
```
* Had to update Makefile to get packr2 working. Go 1.16 and higher should be "go install" not "go get"
* Instead of running make commands in WSL / Ubuntu I ran in golang container using debian:bullseye, matching bundle template.Dockerfile
```bash
docker run -it -v /home/src/KurtSchenk/porter-yq:/go/src golang:bullseye bash
root@fc2ead9c8925:/go# ls      
bin  src
root@fc2ead9c8925:/go# cd src
root@fc2ead9c8925:/go/src# ls
LICENSE  Makefile  README.md  azure-pipelines.yml  build  cmd  go.mod  go.sum  pkg  tests

root@fc2ead9c8925:/go/src# make clean build test-unit xbuild-all test-integration
cd pkg/yq && packr2 clean
rm -fr bin/
GO111MODULE=on go mod tidy
[...]
```
* You may get this error running "make clean ..." If so do following
```bash
./pkg/mod/github.com/gobuffalo/packr/v2@v2.8.3/packr2/cmd/fix/imports.go:16:2: missing go.sum entry for module providing package golang.org/x/tools/go/ast/astutil (imported by github.com/gobuffalo/packr/v2/packr2/cmd/fix); to add:
        go get github.com/gobuffalo/packr/v2/packr2/cmd/fix@v2.8.3
make: *** [Makefile:52: packr2] Error 1

root@6ee80c2e8881:/go/src# go get github.com/gobuffalo/packr/v2/packr2/cmd/fix@v2.8.3
go: downloading golang.org/x/tools v0.13.0
```
* This created the binary for the mixin
* Next copy mixin binaries to .porter folder
```bash
root@fc2ead9c8925:/go/src# make install
mkdir -p /root/.porter/mixins/yq/runtimes/
install bin/mixins/yq/yq /root/.porter/mixins/yq/yq
install bin/mixins/yq/yq-runtime /root/.porter/mixins/yq/runtimes/yq-runtime
```
* Above copies mixin binaries in the container, so do the same outside of it.
```bash
(base) root@DESKTOP-3Q08DV2:~/.porter/mixins# mkdir yq
(base) root@DESKTOP-3Q08DV2:~/.porter/mixins# ls
cache.json  exec  helm  helm3  kubernetes  yq
(base) root@DESKTOP-3Q08DV2:~/.porter/mixins# cd /home/src/squillace/porter-yq
(base) root@DESKTOP-3Q08DV2:/home/src/squillace/porter-yq# install ./bin/mixins/yq/yq /root/.porter/mixins/yq/yq
(base) root@DESKTOP-3Q08DV2:/home/src/squillace/porter-yq# porter mixin list
----------------------------------------
  Name        Version  Author           
----------------------------------------
  exec        v1.0.1   Porter Authors   
  helm        v0.13.4  Porter Authors   
  helm3       v1.0.1   Mohamed Chorfa   
  kubernetes  v1.0.2   Porter Authors   
  yq          v0.1.0   ralph squillace  
```
* Now I can use yq mixin in porter.yaml
* And install and uninstall
```bash
(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq/tests/bundle# porter install
executing install action from porter-custom-mixin (installation: /porter-custom-mixin)
Install Hello World
Hello World
Test yq mixin
yq 3.2.3
yq data.xml
"My name"
Output name
My name
execution completed successfully!

(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq/tests/bundle# porter uninstall
uninstalling bundle
executing uninstall action from porter-custom-mixin (installation: /porter-custom-mixin)
Uninstall Hello World
Goodbye World
execution completed successfully!
```
  * To use the mixin in a bundle see [NOTES_BUNDLE.md](./NOTES_BUNDLE.md)

# TODO
Not that PKG is "PKG = github.com/squillace/porter-yq"
Leave that way but I updated tag to v0.1.1 so the "-X" will have the updated version

From
GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags '-w -X github.com/squillace/porter-yq/pkg.Version=v0.1.0-9-gc4be20d -X github.com/squillace/porter-yq/pkg.Commit=c4be20d' -o bin/mixins/yq/yq-runtime ./cmd/yq

To
GOARCH=amd64 GOOS=linux GO111MODULE=on go build -ldflags '-w -X github.com/squillace/porter-yq/pkg.Version=v0.1.1 -X github.com/squillace/porter-yq/pkg.Commit=c4be20d' -o bin/mixins/yq/yq-runtime ./cmd/yq

* mixin version now

(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq/bin/mixins/yq# ./yq-runtime version
yq v0.1.1-2-g9c972b1 (9c972b1) by ralph squillace
