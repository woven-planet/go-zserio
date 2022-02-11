module github.com/woven-planet/go-zserio

go 1.18

require (
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20211101200231-0802afb9c160
	github.com/cucumber/godog v0.12.4
	github.com/google/go-cmp v0.5.6
	github.com/iancoleman/strcase v0.2.0
	github.com/icza/bitio v1.1.0
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/x448/float16 v0.8.4
)

replace (
	github.com/miekg/dns => github.com/miekg/dns v1.1.45 // avoid CVE-2017-15133
	golang.org/x/mod => golang.org/x/mod v0.5.1
	golang.org/x/net => golang.org/x/net v0.0.0-20220121210141-e204ce36a2ba
	golang.org/x/sys => golang.org/x/sys v0.0.0-20220114195835-da31bd327af9
	golang.org/x/tools => golang.org/x/tools v0.1.8
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/cucumber/gherkin-go/v19 v19.0.3 // indirect
	github.com/cucumber/messages-go/v16 v16.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.0 // indirect
	github.com/hashicorp/go-memdb v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
