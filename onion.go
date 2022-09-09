package onion

import (
	"net/http"
)

// Handler a interface to handle request
type Handler interface {
	handle(w http.ResponseWriter, req *http.Request, params map[string]string)
}

// HandleFunc a adapter implements Handler
type HandleFunc func(w http.ResponseWriter, req *http.Request, params map[string]string)

func (f HandleFunc) handle(w http.ResponseWriter, req *http.Request, params map[string]string) {
	f(w, req, params)
}

type handler struct {
	method  string
	handler Handler
}

// Onion is a http.Handler
type Onion struct {
	s        *http.Server
	handlers map[string]*handler

	cfg *Config
}

// New returns a new initialized Router.
func New(cfg *Config) *Onion {
	return &Onion{
		handlers: make(map[string]*handler),
		cfg:      cfg,
	}
}

// ServeHTTP implement http.Handle
func (c *Onion) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if h, ok := c.handlers[path]; ok {
		if req.Method == h.method {
			h.handler.handle(w, req, nil)

			return
		}
	}

	w.Write([]byte("hello world! " + path))
}

// Start run http server
func (c *Onion) Start(address string) error {
	s := &http.Server{
		Addr:         address,
		Handler:      c,
		ReadTimeout:  c.cfg.ReadTimeout,
		WriteTimeout: c.cfg.WriteTimeout,
	}

	return s.ListenAndServe()
}

// Handle regist a path route
func (c *Onion) Handle(method, path string, h Handler) {
	c.handlers[path] = &handler{method, h}
}

// Get a shortcut for Onion.Handle(http.MethodGet, path, handle)
func (c *Onion) Get(path string, handler Handler) {
	c.Handle(http.MethodGet, path, handler)
}

// Post a shortcut for Onion.Handle(http.MethodPost, path, handle)
func (c *Onion) Post(path string, handler Handler) {
	c.Handle(http.MethodPost, path, handler)
}

// Put a shortcut for Onion.Handle(http.MethodPut, path, handle)
func (c *Onion) Put(path string, handler Handler) {
	c.Handle(http.MethodPut, path, handler)
}

// Delete a shortcut for Onion.Handle(http.MethodDelete, path, handle)
func (c *Onion) Delete(path string, handler Handler) {
	c.Handle(http.MethodDelete, path, handler)
}

// Get a shortcut for Onion.Handle(http.MethodGet, path, handle)
func (c *Onion) GetFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodGet, path, handler)
}

// Post a shortcut for Onion.Handle(http.MethodPost, path, handle)
func (c *Onion) PostFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodPost, path, handler)
}

// Put a shortcut for Onion.Handle(http.MethodPut, path, handle)
func (c *Onion) PutFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodPut, path, handler)
}

// Delete a shortcut for Onion.Handle(http.MethodDelete, path, handle)
func (c *Onion) DeleteFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodDelete, path, handler)
}

// // Head a shortcut for Onion.Handle(http.MethodHead, path, handle)
// func (c *Onion) Head(path string, handler Handler) {
// 	c.Handle(http.MethodHead, path, handler)
// }

// // Patch a shortcut for Onion.Handle(http.MethodPatch, path, handle)
// func (c *Onion) Patch(path string, handler Handler) {
// 	c.Handle(http.MethodPatch, path, handler)
// }

// // Connect a shortcut for Onion.Handle(http.MethodConnect, path, handle)
// func (c *Onion) Connect(path string, handler Handler) {
// 	c.Handle(http.MethodConnect, path, handler)
// }

// // Options a shortcut for Onion.Handle(http.MethodOptions, path, handle)
// func (c *Onion) Options(path string, handler Handler) {
// 	c.Handle(http.MethodOptions, path, handler)
// }

// // Trace a shortcut for Onion.Handle(http.MethodTrace, path, handle)
// func (c *Onion) Trace(path string, handler Handler) {
// 	c.Handle(http.MethodTrace, path, handler)
// }
