module github.com/squillace/porter-yq

go 1.18

require (
	get.porter.sh/porter v0.29.1
	github.com/ghodss/yaml v1.0.0
	github.com/gobuffalo/logger v1.0.7 // indirect
	github.com/gobuffalo/packd v1.0.2 // indirect
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/hashicorp/go-plugin => github.com/carolynvs/go-plugin v1.0.1-acceptstdin
