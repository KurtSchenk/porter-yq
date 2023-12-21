* In order to update the example to Mage instead of Makefile, created a skeleton mixin and copied mage.go and magefile.go here: ```porter mixin create [name]```
* Mage is working as an alternative to Make now. To run 
```bash
(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq# go run mage.go Clean Build TestUnit XBuildAll Test
?       github.com/squillace/porter-yq/pkg      [no test files]
ok      github.com/squillace/porter-yq/pkg/yq   (cached)
?       github.com/squillace/porter-yq/pkg      [no test files]
ok      github.com/squillace/porter-yq/pkg/yq   (cached)
yq  () by ralph squillace
```
* TODO NEed to get version information, which will be included in "yq  () by ralph squillace"
* This creates the following files
```bash
(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq# find ./bin -type f
./bin/mixins/yq/v0.1.1-10-g117e3f4/yq-windows-arm64.exe
./bin/mixins/yq/v0.1.1-10-g117e3f4/yq-linux-arm64
./bin/mixins/yq/v0.1.1-10-g117e3f4/yq-linux-amd64
./bin/mixins/yq/v0.1.1-10-g117e3f4/yq-windows-amd64.exe
./bin/mixins/yq/v0.1.1-10-g117e3f4/yq-darwin-amd64
./bin/mixins/yq/v0.1.1-10-g117e3f4/yq-darwin-arm64
./bin/mixins/yq/runtimes/yq-runtime
./bin/mixins/yq/yq
./bin/mixins/yq/dev/yq-windows-arm64.exe
./bin/mixins/yq/dev/yq-linux-arm64
./bin/mixins/yq/dev/yq-linux-amd64
./bin/mixins/yq/dev/yq-windows-amd64.exe
./bin/mixins/yq/dev/yq-darwin-amd64
./bin/mixins/yq/dev/yq-darwin-arm64
./bin/mixins/yq/canary-dev/yq-windows-arm64.exe
./bin/mixins/yq/canary-dev/yq-linux-arm64
./bin/mixins/yq/canary-dev/yq-linux-amd64
./bin/mixins/yq/canary-dev/yq-windows-amd64.exe
./bin/mixins/yq/canary-dev/yq-darwin-amd64
./bin/mixins/yq/canary-dev/yq-darwin-arm64
```
* Starting working on TestIntegration which is in Makefile (test-integration), but this is not working yet
* To install the mixin run the following
```bash
(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq# go run mage.go Install
Installing the yq mixin into /root/.porter
(base) root@DESKTOP-3Q08DV2:/home/src/KurtSchenk/porter-yq# porter mixin list
----------------------------------------
  Name        Version  Author           
----------------------------------------
  az          v1.0.2   Porter Authors   
  exec        v1.0.14  Porter Authors   
  helm        v0.13.4  Porter Authors   
  helm3       v1.0.1   Mohamed Chorfa   
  kubernetes  v1.0.2   Porter Authors   
  yq                   ralph squillace  
```
  * To use the mixin in a bundle see [NOTES_BUNDLE.md](./NOTES_BUNDLE.md)