package carrot

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

// Carrot is a http.Handler
type Carrot struct {
	s        *http.Server
	handlers map[string]*handler

	cfg *Config
}

// New returns a new initialized Router.
func New(cfg *Config) *Carrot {
	return &Carrot{
		handlers: make(map[string]*handler),
		cfg:      cfg,
	}
}

// ServeHTTP implement http.Handle
func (c *Carrot) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
func (c *Carrot) Start(address string) error {
	s := &http.Server{
		Addr:         address,
		Handler:      c,
		ReadTimeout:  c.cfg.ReadTimeout,
		WriteTimeout: c.cfg.WriteTimeout,
	}

	return s.ListenAndServe()
}

// Handle regist a path route
func (c *Carrot) Handle(method, path string, h Handler) {
	c.handlers[path] = &handler{method, h}
}

// Get a shortcut for carrot.Handle(http.MethodGet, path, handle)
func (c *Carrot) Get(path string, handler Handler) {
	c.Handle(http.MethodGet, path, handler)
}

// Post a shortcut for carrot.Handle(http.MethodPost, path, handle)
func (c *Carrot) Post(path string, handler Handler) {
	c.Handle(http.MethodPost, path, handler)
}

// Put a shortcut for carrot.Handle(http.MethodPut, path, handle)
func (c *Carrot) Put(path string, handler Handler) {
	c.Handle(http.MethodPut, path, handler)
}

// Delete a shortcut for carrot.Handle(http.MethodDelete, path, handle)
func (c *Carrot) Delete(path string, handler Handler) {
	c.Handle(http.MethodDelete, path, handler)
}

// Get a shortcut for carrot.Handle(http.MethodGet, path, handle)
func (c *Carrot) GetFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodGet, path, handler)
}

// Post a shortcut for carrot.Handle(http.MethodPost, path, handle)
func (c *Carrot) PostFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodPost, path, handler)
}

// Put a shortcut for carrot.Handle(http.MethodPut, path, handle)
func (c *Carrot) PutFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodPut, path, handler)
}

// Delete a shortcut for carrot.Handle(http.MethodDelete, path, handle)
func (c *Carrot) DeleteFunc(path string, handler HandleFunc) {
	c.Handle(http.MethodDelete, path, handler)
}

// // Head a shortcut for carrot.Handle(http.MethodHead, path, handle)
// func (c *Carrot) Head(path string, handler Handler) {
// 	c.Handle(http.MethodHead, path, handler)
// }

// // Patch a shortcut for carrot.Handle(http.MethodPatch, path, handle)
// func (c *Carrot) Patch(path string, handler Handler) {
// 	c.Handle(http.MethodPatch, path, handler)
// }

// // Connect a shortcut for carrot.Handle(http.MethodConnect, path, handle)
// func (c *Carrot) Connect(path string, handler Handler) {
// 	c.Handle(http.MethodConnect, path, handler)
// }

// // Options a shortcut for carrot.Handle(http.MethodOptions, path, handle)
// func (c *Carrot) Options(path string, handler Handler) {
// 	c.Handle(http.MethodOptions, path, handler)
// }

// // Trace a shortcut for carrot.Handle(http.MethodTrace, path, handle)
// func (c *Carrot) Trace(path string, handler Handler) {
// 	c.Handle(http.MethodTrace, path, handler)
// }
