module {{ .name }}

go 1.18

require github.com/spf13/cobra v1.5.0

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	k8s.io/client-go v0.19.3
)
