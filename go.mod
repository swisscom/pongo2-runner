module github.com/swisscom/pongo2-runner

go 1.16

require (
	github.com/alexflint/go-arg v1.3.0
	github.com/flosch/pongo2/v4 v4.0.2
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.2.2
)

replace github.com/flosch/pongo2/v4 => github.com/swisscom/pongo2/v4 v4.0.3-0.20210330161743-b5eff3243df2
