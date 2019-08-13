module github.com/haproxytech/dataplaneapi

go 1.12

require (
	github.com/GehirnInc/crypt v0.0.0-20190301055215-6c0105aabd46
	github.com/docker/go-units v0.4.0
	github.com/go-openapi/errors v0.19.0
	github.com/go-openapi/loads v0.19.0
	github.com/go-openapi/runtime v0.19.0
	github.com/go-openapi/spec v0.19.0
	github.com/go-openapi/strfmt v0.19.0
	github.com/go-openapi/swag v0.19.0
	github.com/go-openapi/validate v0.19.0
	github.com/haproxytech/client-native v1.2.1-0.20190813110624-83c4d5051c1c
	github.com/haproxytech/config-parser v1.1.2
	github.com/haproxytech/models v1.2.1
	github.com/jessevdk/go-flags v1.4.0
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/rs/cors v1.6.0
	github.com/sirupsen/logrus v1.4.2
	golang.org/x/net v0.0.0-20190607181551-461777fb6f67
	golang.org/x/sys v0.0.0-20190528012530-adf421d2caf4 // indirect
)

replace github.com/haproxytech/models => ../models

replace github.com/haproxytech/config-parser => ../config-parser

replace github.com/haproxytech/client-native => ../client-native
