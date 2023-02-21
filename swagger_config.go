package httpSwagger

import (
	"github.com/swaggo/swag"
	"html/template"
)

// URLsConfig stores multiple swagger json uris and names.
type URLsConfig struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// Config stores httpSwagger configuration variables.
type Config struct {
	// The url pointing to API definition (normally swagger.json or swagger.yaml). Default is `doc.json`.
	URL                      string
	URLs                     []URLsConfig
	DocExpansion             string
	DeepLinking              bool
	DomID                    string
	PersistAuthorization     bool
	DisplayOperationID       bool
	DefaultModelsExpandDepth int
	DefaultModelExpandDepth  int
	DefaultModelRendering    string
	DisplayRequestDuration   bool
	ShowExtensions           bool
	ShowCommonExtensions     bool
	SupportedSubmitMethods   []string
	TryItOutEnabled          bool

	InstanceName string
	BeforeScript template.JS
	AfterScript  template.JS
	Plugins      []template.JS
	UIConfig     map[template.JS]template.JS
}

// URL presents the url pointing to API definition (normally swagger.json or swagger.yaml).
func URL(url string) func(*Config) {
	return func(c *Config) {
		c.URL = url
	}
}

// URLs adds the url pointing to API definition (normally swagger.json or swagger.yaml) to the list of URLs.
func URLs(url string, name string) func(*Config) {
	return func(c *Config) {
		c.URLs = append(c.URLs, URLsConfig{URL: url, Name: name})
	}
}

// DeepLinking true, false.
func DeepLinking(deepLinking bool) func(*Config) {
	return func(c *Config) {
		c.DeepLinking = deepLinking
	}
}

// DocExpansion list, full, none.
func DocExpansion(docExpansion string) func(*Config) {
	return func(c *Config) {
		c.DocExpansion = docExpansion
	}
}

// DomID #swagger-ui.
func DomID(domID string) func(*Config) {
	return func(c *Config) {
		c.DomID = domID
	}
}

// InstanceName set the instance name that was used to generate the swagger documents
// Defaults to swag.Name ("swagger").
func InstanceName(name string) func(*Config) {
	return func(c *Config) {
		c.InstanceName = name
	}
}

// PersistAuthorization Persist authorization information over browser close/refresh.
// Defaults to false.
func PersistAuthorization(persistAuthorization bool) func(*Config) {
	return func(c *Config) {
		c.PersistAuthorization = persistAuthorization
	}
}

// DisplayOperationID Controls the display of operationId in operations list. The default is false.
func DisplayOperationID(displayOperationID bool) func(*Config) {
	return func(c *Config) {
		c.DisplayOperationID = displayOperationID
	}
}

// DefaultModelsExpandDepth Controls the default expansion setting for models (set to -1 completely hide the models).
// The default value is 1.
func DefaultModelsExpandDepth(defaultModelsExpandDepth int) func(*Config) {
	return func(c *Config) {
		c.DefaultModelsExpandDepth = defaultModelsExpandDepth
	}
}

// DefaultModelExpandDepth Controls the default expansion setting for the model on the model-example section.
// The default value is 1.
func DefaultModelExpandDepth(defaultModelExpandDepth int) func(*Config) {
	return func(c *Config) {
		c.DefaultModelExpandDepth = defaultModelExpandDepth
	}
}

// DefaultModelRendering Controls how the model is shown when the API is first rendered.
// The default is example. It can be set to example, schema or model.
func DefaultModelRendering(defaultModelRendering string) func(*Config) {
	return func(c *Config) {
		c.DefaultModelRendering = defaultModelRendering
	}
}

// DisplayRequestDuration Controls the display of the request duration (in milliseconds) for Try-It-Out requests.
// The default is false.
func DisplayRequestDuration(displayRequestDuration bool) func(*Config) {
	return func(c *Config) {
		c.DisplayRequestDuration = displayRequestDuration
	}
}

// ShowExtensions Controls the display of vendor extension (x-) fields and values for Operations, Parameters, and Schema.
// The default is false.
func ShowExtensions(showExtensions bool) func(*Config) {
	return func(c *Config) {
		c.ShowExtensions = showExtensions
	}
}

// ShowCommonExtensions Controls the display of common extension (x-) fields and values for Operations, Parameters, and Schema.
// The default is false.
func ShowCommonExtensions(showCommonExtensions bool) func(*Config) {
	return func(c *Config) {
		c.ShowCommonExtensions = showCommonExtensions
	}
}

// SupportedSubmitMethods Controls the display of the submit operation for Try-It-Out requests.
// The default is ['get', 'post', 'put', 'delete', 'options', 'head', 'patch', 'trace'].
func SupportedSubmitMethods(supportedSubmitMethods ...string) func(*Config) {
	return func(c *Config) {
		c.SupportedSubmitMethods = supportedSubmitMethods
	}
}

// Plugins specifies additional plugins to load into Swagger UI.
func Plugins(plugins []string) func(*Config) {
	return func(c *Config) {
		vs := make([]template.JS, len(plugins))
		for i, v := range plugins {
			vs[i] = template.JS(v)
		}
		c.Plugins = vs
	}
}

// UIConfig specifies additional SwaggerUIBundle config object properties.
func UIConfig(props map[string]string) func(*Config) {
	return func(c *Config) {
		vs := make(map[template.JS]template.JS, len(props))
		for k, v := range props {
			vs[template.JS(k)] = template.JS(v)
		}
		c.UIConfig = vs
	}
}

// BeforeScript holds JavaScript to be run right before the Swagger UI object is created.
func BeforeScript(js string) func(*Config) {
	return func(c *Config) {
		c.BeforeScript = template.JS(js)
	}
}

// AfterScript holds JavaScript to be run right after the Swagger UI object is created
// and set on the window.
func AfterScript(js string) func(*Config) {
	return func(c *Config) {
		c.AfterScript = template.JS(js)
	}
}

func newConfig(configFns ...func(*Config)) *Config {
	config := Config{
		URL:                      "doc.json",
		URLs:                     nil,
		DocExpansion:             "list",
		DomID:                    "swagger-ui",
		DeepLinking:              true,
		PersistAuthorization:     false,
		DisplayOperationID:       false,
		DefaultModelsExpandDepth: 1,
		DefaultModelExpandDepth:  1,
		DefaultModelRendering:    "example",
		DisplayRequestDuration:   false,
		ShowExtensions:           false,
		ShowCommonExtensions:     false,
		SupportedSubmitMethods:   []string{"get", "put", "post", "delete", "options", "head", "patch", "trace"},
		TryItOutEnabled:          false,
	}

	for _, fn := range configFns {
		fn(&config)
	}

	if config.InstanceName == "" {
		config.InstanceName = swag.Name
	}

	return &config
}
