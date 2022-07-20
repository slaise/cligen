package config


// Config ...
type Config struct {
	AllNamespaces bool      json: `allNamespaces`
    Namespace     string    json: `namespace`
}