package taguk

import (
	"reflect"
	"strings"
)


// Server represents a taguk server
type Server struct {
	Resources ResourceMap
}

type ResourceMap map[string]Resource
type ActionMap map[string]Action

type CustomServer interface {
	// TODO: add options here
	Configure(r ResourceMap)
	//Serve serves with sensible defaults
	Serve()
	//Serve(listener string)
}

type Resource struct {
	Name string
	Version string
	r interface{}
	Actions ActionMap
}

type Item struct {
	ID int64
	i interface{}
}
type Action struct {
	Name string

	// Individual is true when the action is run on an individual item and false when it is run on a resource
	Individual bool

	// Base represents the base type that the action needs to receive on. The zero value of this type will be initialized
	// and the action will be called on it
	Base reflect.Type

	reflect.Method
}

// NewServer initializes a new server
func NewServer() Server {
	return Server{
		Resources: make(ResourceMap),
	}
}

// AddResources adds several resources
func (s Server) AddResources(res ...interface{}) {
	for _, r := range res {
		s.AddResource(r)
	}
}

// hasIdArgument returns true if the function provided has a second (after receiver) argument that is an int64
func hasIdArgument(f reflect.Type) bool {
	// The function needs at least 2 arguments, 1 for the struct receiver, another for the id
	if f.NumIn() > 1 {
		return f.In(1).Kind() == reflect.Int64
	}
	return false
}
func getBaseValue(f reflect.Type) reflect.Type {
	// The function needs at least 2 arguments, 1 for the struct receiver, another for the id
	if f.NumIn() > 0 {
		return f.In(0)
	}
	return nil
}

// AddResource adds the resource to the server, registering all actions for it as well
func (s *Server) AddResource(r interface{}) {
	// The following gets the type of the value underlying the interface{}
	v := reflect.ValueOf(r)
	if reflect.TypeOf(r).Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	
	rName := t.Name()

	res := Resource{Name:rName, Version:"0.1.0", r: r}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		action := Action{
			Name:   m.Name,
			Method: m,
			Base: getBaseValue(m.Func.Type()),
			Individual: hasIdArgument(m.Func.Type()),
		}
		if res.Actions == nil {
			res.Actions = make(ActionMap)
		}
		res.Actions[strings.ToLower(m.Name)] = action
	}
	//TODO: don't hack this, actually use the json struct field name to json key decoding code
    //https://golang.org/src/encoding/json/encode.go?s=6471:6514#L148
	s.Resources[strings.ToLower(rName)] = res
}

//type optionsFunc func()

//func (s Server) ServeHTTP() {
//	p := os.Getenv("PORT")
//	if p == "" {
//		p = "8080"
//	}
//
//	hostname := "127.0.0.1"
//
//	listener := hostname + ":" + p
//
//	srv := httpjson.NewServer()
//	srv.Configure(s.resources)
//	srv.Serve(listener)
//}

