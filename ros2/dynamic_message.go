package ros2

// IMPORT REQUIRED PACKAGES.

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/justinscorringe/rclgo/libtypes"
	"github.com/pkg/errors"
)

// DEFINE PUBLIC STRUCTURES.

type DynamicMessageType struct {
	spec  *libtypes.MsgSpec         // Standard go type msg implementation
	cSpec *ROSIdlMessageTypeSupport // C Specification msg implementation
}

type DynamicMessage struct {
	MessageBase

	dynamicType *DynamicMessageType
	data        map[string]interface{}
}

// DEFINE PRIVATE STRUCTURES.

// DEFINE PUBLIC GLOBALS.

// DEFINE PRIVATE GLOBALS.

const Sep = "/"

var rosPkgPath string // Colon separated list of paths to search for message definitions on.

var context *libtypes.MsgContext // We'll try to preserve a single message context to avoid reloading each time.

// DEFINE PUBLIC STATIC FUNCTIONS.

// SetRuntimePackagePath sets the ROS package search path which will be used by DynamicMessage to look up ROS message definitions at runtime.
func SetRuntimePackagePath(path string) {
	// We're not going to check that the result is valid, we'll just accept it blindly.
	rosPkgPath = path
	// Reset the message context
	ResetContext()
	// All done.
	return
}

// GetRuntimePackagePath returns the ROS package search path which will be used by DynamicMessage to look up ROS message definitions at runtime.  By default, this will
// be equivalent to the ${ROS_PACKAGE_PATH} environment variable.
func GetRuntimePackagePath() string {
	// If a package path hasn't been set at the time of first use, by default we'll just use the ROS environment default.
	if rosPkgPath == "" {
		rosPkgPath = os.Getenv("ROS_PACKAGE_PATH")
	}
	// All done.
	return rosPkgPath
}

// ResetContext resets the package path context so that a new one will be generated
func ResetContext() {
	context = nil
}

// NewDynamicMessageType generates a DynamicMessageType corresponding to the specified typeName from the available ROS message definitions; typeName should be a fully-qualified
// ROS message type name.  The first time the function is run, a message 'context' is created by searching through the available ROS message definitions, then the ROS message to
// be used for the definition is looked up by name.  On subsequent calls, the ROS message type is looked up directly from the existing context.
func NewDynamicMessageType(typeName string) (*DynamicMessageType, error) {
	return newDynamicMessageTypeNested(typeName, "")
}

// newDynamicMessageTypeNested generates a DynamicMessageType from the available ROS message definitions.  The first time the function is run, a message 'context' is created by
// searching through the available ROS message definitions, then the ROS message type to use for the defintion is looked up by name.  On subsequent calls, the ROS message type
// is looked up directly from the existing context.  This 'nested' version of the function is able to be called recursively, where packageName should be the typeName of the
// parent ROS message; this is used internally for handling complex ROS messages.
func newDynamicMessageTypeNested(typeName string, packageName string) (*DynamicMessageType, error) {
	// Create an empty message type.
	m := new(DynamicMessageType)

	// If we haven't created a message context yet, better do that.
	if context == nil {
		// Create context for our ROS install.
		c, err := libtypes.NewMsgContext(strings.Split(GetRuntimePackagePath(), ":"))
		if err != nil {
			return nil, err
		}
		context = c
	}

	// We need to try to look up the full name, in case we've just been given a short name.
	fullname := typeName

	// The Header type has some special treatment.
	if typeName == "Header" {
		fullname = "std_msgs/Header"
	} else {
		_, ok := context.GetMsgs()[fullname]
		if !ok {
			// Seems like the package_name we were give wasn't the full name.

			// Messages in the same package are allowed to use relative names, so try prefixing the package.
			if packageName != "" {
				fullname = packageName + "/" + fullname
			}
		}
	}

	// Load context for the target message.
	spec, err := context.LoadMsg(fullname)
	if err != nil {
		return nil, err
	}

	// Now we know all about the message!
	m.spec = spec

	// TODO: Now we need to get the c spec of the message (rosidl)

	// We've successfully made a new message type matching the requested ROS type.
	return m, nil
}

// DEFINE PUBLIC RECEIVER FUNCTIONS.

//	DynamicMessageType

// Name returns the full ROS name of the message type; required for ros.MessageType.
func (t *DynamicMessageType) Name() string {
	return t.spec.FullName
}

// Text returns the full ROS message specification for this message type; required for ros.MessageType.
func (t *DynamicMessageType) Text() string {
	return t.spec.Text
}

func (t *DynamicMessageType) NewMessage()

// GenerateJSONSchema generates a (primitive) JSON schema for the associated DynamicMessageType; however note that since
// we are mostly interested in making schema's for particular _topics_, the function takes a string prefix, and string topic name, which are
// used to id the resulting schema.
func (t *DynamicMessageType) GenerateJSONSchema(prefix string, topic string) ([]byte, error) {
	// The JSON schema for a message consist of the (recursive) properties names/types:
	schemaItems, err := t.generateJSONSchemaProperties(prefix + topic)
	if err != nil {
		return nil, err
	}

	// Plus some extra keywords:
	schemaItems["$schema"] = "https://json-schema.org/draft-07/schema#"
	schemaItems["$id"] = prefix + topic

	// The schema itself is created from the map of properties.
	schemaString, err := json.Marshal(schemaItems)
	if err != nil {
		return nil, err
	}

	// All done.
	return schemaString, nil
}

func (t *DynamicMessageType) generateJSONSchemaProperties(topic string) (map[string]interface{}, error) {
	// Each message's schema indicates that it is an 'object' with some nested properties: those properties are the fields and their types.
	properties := make(map[string]interface{})
	schemaItems := make(map[string]interface{})
	schemaItems["type"] = "object"
	schemaItems["title"] = topic
	schemaItems["properties"] = properties

	// Iterate over each of the fields in the message.
	for _, field := range t.spec.Fields {
		if field.IsArray {
			// It's an array.
			propertyContent := make(map[string]interface{})
			properties[field.Name] = propertyContent

			if field.GoType == "uint8" {
				propertyContent["title"] = topic + Sep + field.Name
				propertyContent["type"] = "string"
			} else {
				// Arrays all have a type of 'array', regardless of that the hold, then the 'item' keyword determines what type goes in the array.
				propertyContent["type"] = "array"
				propertyContent["title"] = topic + Sep + field.Name
				arrayItems := make(map[string]interface{})
				propertyContent["items"] = arrayItems

				// Need to handle each type appropriately.
				if field.IsBuiltin {
					if field.Type == "string" {
						arrayItems["type"] = "string"
					} else if field.Type == "time" {
						timeItems := make(map[string]interface{})
						timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
						timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
						arrayItems["type"] = "object"
						arrayItems["properties"] = timeItems
					} else if field.Type == "duration" {
						timeItems := make(map[string]interface{})
						timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
						timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
						arrayItems["type"] = "object"
						arrayItems["properties"] = timeItems
					} else {
						// It's a primitive.
						var jsonType string
						if field.GoType == "int8" || field.GoType == "int16" || field.GoType == "uint16" ||
							field.GoType == "int32" || field.GoType == "uint32" || field.GoType == "int64" || field.GoType == "uint64" {
							jsonType = "integer"
						} else if field.GoType == "float32" || field.GoType == "float64" {
							jsonType = "number"
						} else if field.GoType == "bool" {
							jsonType = "bool"
						} else {
							// Something went wrong.
							return nil, errors.New("we haven't implemented this primitive yet")
						}
						arrayItems["type"] = jsonType
					}

				} else {
					// It's another nested message.

					// Generate the nested type.
					msgType, err := newDynamicMessageTypeNested(field.Type, field.Package)
					if err != nil {
						return nil, errors.Wrap(err, "Schema Field: "+field.Name)
					}

					// Recursively generate schema information for the nested type.
					schemaElement, err := msgType.generateJSONSchemaProperties(topic + Sep + field.Name)
					if err != nil {
						return nil, errors.Wrap(err, "Schema Field:"+field.Name)
					}
					arrayItems["type"] = schemaElement
				}
			}
		} else {
			// It's a scalar.
			if field.IsBuiltin {
				propertyContent := make(map[string]interface{})
				properties[field.Name] = propertyContent
				propertyContent["title"] = topic + Sep + field.Name

				if field.Type == "string" {
					propertyContent["type"] = "string"
				} else if field.Type == "time" {
					timeItems := make(map[string]interface{})
					timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
					timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
					propertyContent["type"] = "object"
					propertyContent["properties"] = timeItems
				} else if field.Type == "duration" {
					timeItems := make(map[string]interface{})
					timeItems["sec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "sec"}
					timeItems["nsec"] = map[string]string{"type": "integer", "title": topic + Sep + field.Name + Sep + "nsec"}
					propertyContent["type"] = "object"
					propertyContent["properties"] = timeItems
				} else {
					// It's a primitive.
					var jsonType string
					if field.GoType == "int8" || field.GoType == "uint8" || field.GoType == "int16" || field.GoType == "uint16" ||
						field.GoType == "int32" || field.GoType == "uint32" || field.GoType == "int64" || field.GoType == "uint64" {
						jsonType = "integer"
						jsonType = "integer"
						jsonType = "integer"
					} else if field.GoType == "float32" || field.GoType == "float64" {
						jsonType = "number"
					} else if field.GoType == "bool" {
						jsonType = "bool"
					} else {
						// Something went wrong.
						return nil, errors.New("we haven't implemented this primitive yet")
					}
					propertyContent["type"] = jsonType
				}
			} else {
				// It's another nested message.

				// Generate the nested type.
				msgType, err := newDynamicMessageTypeNested(field.Type, field.Package)
				if err != nil {
					return nil, errors.Wrap(err, "Schema Field: "+field.Name)
				}

				// Recursively generate schema information for the nested type.
				schemaElement, err := msgType.generateJSONSchemaProperties(topic + Sep + field.Name)
				if err != nil {
					return nil, errors.Wrap(err, "Schema Field:"+field.Name)
				}
				properties[field.Name] = schemaElement
			}
		}
	}

	// All done.
	return schemaItems, nil
}

//	DynamicMessage

// Data returns the data map field of the DynamicMessage
func (m *DynamicMessage) Data() map[string]interface{} {
	return m.data
}

// MarshalJSON provides a custom implementation of JSON marshalling, so that when the DynamicMessage is recursively
// marshalled using the standard Go json.marshal() mechanism, the resulting JSON representation is a compact representation
// of just the important message payload (and not the message definition).  It's important that this representation matches
// the schema generated by GenerateJSONSchema().
func (m *DynamicMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

//UnmarshalJSON provides a custom implementation of JSON unmarshalling. Using the dynamicMessage provided, Msgspec is used to
//determine the individual parsing of each JSON encoded payload item into the correct Go type. It is important each type is
//correct so that the message serializes correctly and is understood by the ROS system
func (m *DynamicMessage) UnmarshalJSON(buf []byte) error {

	//Delcaring temp variables to be used across the unmarshaller
	var err error
	var goField libtypes.Field
	var keyName []byte
	var oldMsgType string
	var msg *DynamicMessage
	var msgType *DynamicMessageType
	var data interface{}
	var fieldExists bool

	//Declaring jsonparser unmarshalling functions
	var arrayHandler func([]byte, jsonparser.ValueType, int, error)
	var objectHandler func([]byte, []byte, jsonparser.ValueType, int) error

	//JSON key is an array
	arrayHandler = func(key []byte, dataType jsonparser.ValueType, offset int, err error) {
		switch dataType.String() {
		//We have a string array
		case "string":
			if goField.GoType == "float32" || goField.GoType == "float64" {
				data, err = strconv.ParseFloat(string(key), 64)
				if err != nil {
					errors.Wrap(err, "Field: "+goField.Name)
				}
				if goField.GoType == "float32" {
					m.data[goField.Name] = append(m.data[goField.Name].([]JsonFloat32), JsonFloat32{F: float32((data.(float64)))})
				} else {
					m.data[goField.Name] = append(m.data[goField.Name].([]JsonFloat64), JsonFloat64{F: data.(float64)})
				}
			} else {
				m.data[goField.Name] = append(m.data[goField.Name].([]string), string(key))
			}
		//We have a number or int array.
		case "number":
			//We have a float to parse
			if goField.GoType == "float64" || goField.GoType == "float32" {
				data, err = strconv.ParseFloat(string(key), 64)
				if err != nil {
					errors.Wrap(err, "Field: "+goField.Name)
				}
			} else {
				data, err = strconv.ParseInt(string(key), 0, 64)
				if err != nil {
					errors.Wrap(err, "Field: "+goField.Name)
				}
			}
			//Append field to data array
			switch goField.GoType {
			case "int8":
				m.data[goField.Name] = append(m.data[goField.Name].([]int8), int8((data.(int64))))
			case "int16":
				m.data[goField.Name] = append(m.data[goField.Name].([]int16), int16((data.(int64))))
			case "int32":
				m.data[goField.Name] = append(m.data[goField.Name].([]int32), int32((data.(int64))))
			case "int64":
				m.data[goField.Name] = append(m.data[goField.Name].([]int64), int64((data.(int64))))
			case "uint8":
				m.data[goField.Name] = append(m.data[goField.Name].([]uint8), uint8((data.(int64))))
			case "uint16":
				m.data[goField.Name] = append(m.data[goField.Name].([]uint16), uint16((data.(int64))))
			case "uint32":
				m.data[goField.Name] = append(m.data[goField.Name].([]uint32), uint32((data.(int64))))
			case "uint64":
				m.data[goField.Name] = append(m.data[goField.Name].([]uint64), uint64((data.(int64))))
			case "float32":
				m.data[goField.Name] = append(m.data[goField.Name].([]JsonFloat32), JsonFloat32{F: float32((data.(float64)))})
			case "float64":
				m.data[goField.Name] = append(m.data[goField.Name].([]JsonFloat64), JsonFloat64{F: data.(float64)})
			}
		//We have a bool array
		case "boolean":
			data, err := jsonparser.GetBoolean(buf, string(key))
			_ = err
			m.data[goField.Name] = append(m.data[goField.Name].([]bool), data)
		//We have an object array
		case "object":
			switch goField.GoType {
			//We have a time object
			case "ros.Time":
				tmpTime := Time{}
				sec, err := jsonparser.GetInt(key, "Sec")
				nsec, err := jsonparser.GetInt(key, "NSec")
				if err == nil {
					tmpTime.Sec = uint32(sec)
					tmpTime.NSec = uint32(nsec)
				}
				m.data[goField.Name] = append(m.data[goField.Name].([]Time), tmpTime)
			//We have a duration object
			case "ros.Duration":
				tmpDuration := Duration{}
				sec, err := jsonparser.GetInt(key, "Sec")
				nsec, err := jsonparser.GetInt(key, "NSec")
				if err == nil {
					tmpDuration.Sec = uint32(sec)
					tmpDuration.NSec = uint32(nsec)
				}
				m.data[goField.Name] = append(m.data[goField.Name].([]Duration), tmpDuration)
			//We have a nested message
			default:
				newMsgType := goField.GoType
				//Check if the message type is the same as last iteration
				//This avoids generating a new type for each array item
				if oldMsgType != "" && oldMsgType == newMsgType {
					//We've already generated this type
				} else {
					msgType, err = newDynamicMessageTypeNested(goField.Type, goField.Package)
					_ = err
				}
				msg = msgType.NewMessage().(*DynamicMessage)
				err = msg.UnmarshalJSON(key)
				m.data[goField.Name] = append(m.data[goField.Name].([]Message), msg)
				//Store msg type
				oldMsgType = newMsgType
				//No error handling in array, see next comment
				_ = err

			}
		}

		//Null error as it is not returned in ArrayEach, requires package modification
		_ = err
		//Null keyName to prevent repeat scenarios of same key usage
		_ = keyName

	}

	//JSON key handler
	objectHandler = func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		//Store keyName for usage in ArrayEach function
		keyName = key
		fieldExists = false
		//Find message spec field that matches JSON key
		for _, field := range m.dynamicType.spec.Fields {
			if string(key) == field.Name {
				goField = field
				fieldExists = true
			}
		}
		if fieldExists == true {
			//Scalars First
			switch dataType.String() {
			//We have a JSON string
			case "string":
				//Special case where we have a byte array encoded as JSON string
				if goField.GoType == "uint8" {
					data, err := base64.StdEncoding.DecodeString(string(value))
					if err != nil {
						return errors.Wrap(err, "Byte Array Field: "+goField.Name)
					}
					m.data[goField.Name] = data
					//Case where we have marshalled a special float as a string
				} else if goField.GoType == "float32" || goField.GoType == "float64" {
					data, err = strconv.ParseFloat(string(value), 64)
					if err != nil {
						errors.Wrap(err, "Field: "+goField.Name)
					}
					if goField.GoType == "float32" {
						m.data[goField.Name] = JsonFloat32{F: float32(data.(float64))}
					} else {
						m.data[goField.Name] = JsonFloat64{F: data.(float64)}
					}
				} else {
					m.data[goField.Name] = string(value)
				}
			//We have a JSON number or int
			case "number":
				//We have a float to parse
				if goField.GoType == "float64" || goField.GoType == "float32" {
					data, err = jsonparser.GetFloat(buf, string(key))
					if err != nil {
						return errors.Wrap(err, "Field: "+goField.Name)
					}
					//We have an int to parse
				} else {
					data, err = jsonparser.GetInt(buf, string(key))
					if err != nil {
						return errors.Wrap(err, "Field: "+goField.Name)
					}
				}
				//Copy number value to message field
				switch goField.GoType {
				case "int8":
					m.data[goField.Name] = int8(data.(int64))
				case "int16":
					m.data[goField.Name] = int16(data.(int64))
				case "int32":
					m.data[goField.Name] = int32(data.(int64))
				case "int64":
					m.data[goField.Name] = int64(data.(int64))
				case "uint8":
					m.data[goField.Name] = uint8(data.(int64))
				case "uint16":
					m.data[goField.Name] = uint16(data.(int64))
				case "uint32":
					m.data[goField.Name] = uint32(data.(int64))
				case "uint64":
					m.data[goField.Name] = uint64(data.(int64))
				case "float32":
					m.data[goField.Name] = JsonFloat32{F: float32(data.(float64))}
				case "float64":
					m.data[goField.Name] = JsonFloat64{F: data.(float64)}
				}
			//We have a JSON bool
			case "boolean":
				data, err := jsonparser.GetBoolean(buf, string(key))
				if err != nil {
					return errors.Wrap(err, "Field: "+goField.Name)
				}
				m.data[goField.Name] = data
			//We have a JSON object
			case "object":
				switch goField.GoType {
				//We have a time object
				case "ros.Time":
					tmpTime := Time{}
					sec, err := jsonparser.GetInt(value, "Sec")
					nsec, err := jsonparser.GetInt(value, "NSec")
					if err == nil {
						tmpTime.Sec = uint32(sec)
						tmpTime.NSec = uint32(nsec)
					}
					m.data[goField.Name] = tmpTime
				//We have a duration object
				case "ros.Duration":
					tmpDuration := Duration{}
					sec, err := jsonparser.GetInt(value, "Sec")
					nsec, err := jsonparser.GetInt(value, "NSec")
					if err == nil {
						tmpDuration.Sec = uint32(sec)
						tmpDuration.NSec = uint32(nsec)
					}
					m.data[goField.Name] = tmpDuration
				default:
					//We have a nested message
					msgType, err := newDynamicMessageTypeNested(goField.Type, goField.Package)
					if err != nil {
						return errors.Wrap(err, "Field: "+goField.Name)
					}
					msg := msgType.NewMessage().(*DynamicMessage)
					if err = msg.UnmarshalJSON(value); err != nil {
						return errors.Wrap(err, "Field: "+goField.Name)
					}
					m.data[goField.Name] = msg
				}
			//We have a JSON array
			case "array":
				//Redeclare message array fields incase they do not exist
				switch goField.GoType {
				case "bool":
					m.data[goField.Name] = make([]bool, 0)
				case "int8":
					m.data[goField.Name] = make([]int8, 0)
				case "int16":
					m.data[goField.Name] = make([]int16, 0)
				case "int32":
					m.data[goField.Name] = make([]int32, 0)
				case "int64":
					m.data[goField.Name] = make([]int64, 0)
				case "uint8":
					m.data[goField.Name] = make([]uint8, 0)
				case "uint16":
					m.data[goField.Name] = make([]uint16, 0)
				case "uint32":
					m.data[goField.Name] = make([]uint32, 0)
				case "uint64":
					m.data[goField.Name] = make([]uint64, 0)
				case "float32":
					m.data[goField.Name] = make([]JsonFloat32, 0)
				case "float64":
					m.data[goField.Name] = make([]JsonFloat64, 0)
				case "string":
					m.data[goField.Name] = make([]string, 0)
				case "ros.Time":
					m.data[goField.Name] = make([]Time, 0)
				case "ros.Duration":
					m.data[goField.Name] = make([]Duration, 0)
				default:
					//goType is a nested Message array
					m.data[goField.Name] = make([]Message, 0)
				}
				//Parse JSON array
				jsonparser.ArrayEach(value, arrayHandler)
			default:
				//We do nothing here as blank fields may return value type NotExist or Null
				err = errors.Wrap(err, "Null field: "+string(key))
			}
		} else {
			return errors.New("Field Unknown: " + string(key))
		}
		return err
	}
	//Perform JSON object handler function
	err = jsonparser.ObjectEach(buf, objectHandler)
	return err
}

// DEFINE PRIVATE STATIC FUNCTIONS.

// DEFINE PRIVATE RECEIVER FUNCTIONS.

// ALL DONE.