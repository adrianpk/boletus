package kabestan

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config struct.
type Config struct {
	namespace string // i.e.: MWZ, MYCVS, APP, etc...
	values    map[string]string
}

// Load configuration.
func LoadConfig(namespace string) *Config {
	cfg := &Config{}
	cfg.SetNamespace(namespace)
	cfg.loadNamespaceEnvVars()
	return cfg
}

// SetNamespace for configuration.
func (c *Config) SetNamespace(namespace string) {
	c.namespace = strings.ToUpper(namespace)
}

func (c *Config) namespacePrefix() string {
	return fmt.Sprintf("%s_", c.namespace)
}

func (c *Config) SetValues(values map[string]string) {
	c.values = values
}

// Get reads all visible environment variables
// that belongs to the namespace.
// An optional reload parameter lets re-read
// the values from environment.
func (c *Config) Get(reload ...bool) map[string]string {
	if len(reload) > 1 && reload[0] {
		return c.get(true)
	}
	return c.get(false)
}

func (c *Config) get(reload bool) map[string]string {
	if reload || len(c.values) == 0 {
		return c.readNamespaceEnvVars()
	}
	return c.values
}

// Val read a specific namespaced  environment variable
// And return its value.
// An optional reload parameter lets re-read
// the values from environment.
func (c *Config) Val(key string, reload ...bool) (value string, ok bool) {
	vals := c.get(false)
	if len(reload) > 1 && reload[0] {
		vals = c.get(true)
	}
	val, ok := vals[key]
	return val, ok
}

// ValOrDef read a specific namespaced environment variables
// And return its value.
// A default value is returned if key value is not found.
// An optional reload parameter lets re-read
// the value from environment.
func (c *Config) ValOrDef(key string, defVal string, reload ...bool) (value string) {
	vals := c.get(false)
	if len(reload) > 1 && reload[0] {
		vals = c.get(true)
	}
	val, ok := vals[key]
	if !ok {
		val = defVal
	}
	return val
}

// ValAsInt read a specific namespaced  environment variables
// and return its value as an int if it can be parsed as such type.
// A default value is returned if key value is not found.
// An optional reload parameter lets re-read
// the values from environment before.
func (c *Config) ValAsInt(key string, defVal int64, reload ...bool) (value int64) {
	vals := c.get(false)
	if len(reload) > 1 && reload[0] {
		vals = c.get(true)
	}
	val, ok := vals[key]
	if !ok {
		return defVal
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defVal
	}
	return i
}

// ValAsFloat read a specific namespaced  environment variables
// and return its value as a float if it can be parsed as such type.
// A default value is returned if key value is not found.
// An optional reload parameter lets re-read
// the values from environment before.
func (c *Config) ValAsFloat(key string, defVal float64, reload ...bool) (value float64) {
	vals := c.get(false)
	if len(reload) > 1 && reload[0] {
		vals = c.get(true)
	}
	val, ok := vals[key]
	if !ok {
		return defVal
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return defVal
	}
	return f
}

// ValAsBool read a specific namespaced  environment variables
// and return its value as a bool if it can be parsed as such type.
// A default value is returned if key value is not found.
// An optional reload parameter lets re-read
// the values from environment before.
func (c *Config) ValAsBool(key string, defVal bool, reload ...bool) (value bool) {
	vals := c.get(false)
	if len(reload) > 1 && reload[0] {
		vals = c.get(true)
	}
	val, ok := vals[key]
	if !ok {
		return defVal
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return defVal
	}
	return b
}

// loadNamespaceEnvVars load all visible environment variables
// that belongs to the namespace.
func (c *Config) loadNamespaceEnvVars() {
	c.values = c.readNamespaceEnvVars()
}

// readNamespaceEnvVars reads all visible environment variables
// that belongs to the namespace.
func (c *Config) readNamespaceEnvVars() map[string]string {
	nevs := make(map[string]string)
	np := c.namespacePrefix()

	for _, ev := range os.Environ() {
		if strings.HasPrefix(ev, np) {
			varval := strings.SplitN(ev, "=", 2)

			if len(varval) < 2 {
				continue
			}

			key := c.keyify(varval[0])
			nevs[key] = varval[1]
		}
	}

	return nevs
}

// keyify enviroment variable names
// i.e.: NAMESPACE_CONFIG_VALUE becomes config.value
func (c *Config) keyify(name string) string {
	split := strings.Split(name, "_")
	if len(split) < 1 {
		return ""
	}
	// Without namespace prefix
	wnsp := strings.Join(split[1:], ".")
	// Dot separated lowercased
	dots := strings.ToLower(strings.Replace(wnsp, "_", ".", -1))
	return dots
}

// getEnvOrDef returns the value of environment variable or a default value.
// if this value is empty or an empty string.
// An empty string is returned if environment variable is inexistent
// and a default was not provided.
func getEnvOrDef(envar string, def ...string) string {
	val := os.Getenv(envar)
	if val != "" {
		return val
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
