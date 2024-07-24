package config

var (
	ApiKey string
	ApiSec string
)

// SetConfig sets the configuration variables
func SetConfig(key, secret string) {
	ApiKey = key
	ApiSec = secret
}
