package nomo

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-carrot/surf"
)

type SwaggerVersion string
type Scheme string

const (
	SwaggerVersion20 SwaggerVersion = "2.0"

	SchemeHTTP  Scheme = "http"
	SchemeHTTPS        = "https"
	SchemeWS           = "ws"
	SchemeWSS          = "wss"
)

// Swagger Object: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#swagger-object
type Swagger struct {
	Swagger     SwaggerVersion     `json:"swagger"`
	Info        Info               `json:"info"`
	Host        string             `json:"host,omitempty"`
	BasePath    string             `json:"basePath,omitempty"`
	Schemes     []Scheme           `json:"schemes,omitempty"`
	Consumes    []string           `json:"consumes,omitempty"`
	Produces    []string           `json:"produces,omitempty"`
	Paths       map[string]*Path   `json:"paths"`
	Definitions map[string]*Schema `json:"definitions,omitempty"`
	// TODO parameters
	// TODO responses
	// TODO securityDefinitions
	// TODO security
	// TODO tags
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
}

func (s *Swagger) ConsumeModels(models ...surf.Model) {
	for _, model := range models {
		s.consumeModel(model)
	}
}

func (s *Swagger) consumeModel(model surf.Model) {
	// Get model config
	schemaProperties := make(map[string]*Schema, 0)
	config := model.GetConfiguration()

	// Build Schema Properties
	value := reflect.TypeOf(model).Elem()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		json := field.Tag.Get("json")
		if json != "-" && json != "" {
			schemaProperties[json] = &Schema{
				Type:    getSwaggerType(field),
				Example: getSwaggerExample(field),
			}
		}
	}

	// Append schema
	if s.Definitions == nil {
		s.Definitions = make(map[string]*Schema)
	}
	s.Definitions[config.TableName] = &Schema{
		Type:       TypeObject,
		Properties: schemaProperties,
	}
}

func getSwaggerType(field reflect.StructField) Type {
	switch field.Type.Kind() {
	// Bools
	case reflect.Bool:
		return TypeBoolean

	// Integers
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		return TypeInteger

	// Numbers
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uintptr:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return TypeNumber

	// Arrays
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		return TypeArray

	// Strings
	case reflect.String:
		return TypeString

	// Structs
	case reflect.Struct:
		// Handle struct representations
		switch field.Type.Name() {
		case "Int":
			return TypeInteger
		case "String":
			return TypeString
		case "Time":
			return TypeString
		case "Bool":
			return TypeBoolean
		case "Float":
			return TypeNumber
		}
	}
	return TypeString
}

func getSwaggerExample(field reflect.StructField) interface{} {
	swaggerType := getSwaggerType(field)

	// Parse Example
	exampleTag := field.Tag.Get("example")
	fmt.Println(exampleTag)
	if exampleTag != "" {
		switch swaggerType {
		case TypeBoolean:
			example, _ := strconv.ParseBool(exampleTag)
			return example
		case TypeInteger:
			example, _ := strconv.ParseInt(exampleTag, 10, 64)
			return example
		case TypeNumber:
			example, _ := strconv.ParseFloat(exampleTag, 64)
			return example
		case TypeArray:
			return make([]string, 0)
		case TypeString:
			return exampleTag
		}
		return exampleTag
	}

	// Default for Types
	switch swaggerType {
	case TypeBoolean:
		return false
	case TypeInteger:
		return 1
	case TypeNumber:
		return 1.2
	case TypeArray:
		return make([]string, 0)
	case TypeString:
		return "Some string"
	}

	return field.Type.Name()
}
