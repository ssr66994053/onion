package carrot

import (
	"net/http"
	"time"
)

// Handle handle request registe to route
type Handle func(w http.ResponseWriter, req *http.Request)

type handler struct {
	method string
	handle Handle
}

// Carrot is a http.Handler
type Carrot struct {
	s        *http.Server
	handlers map[string]*handler

	// config
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// New returns a new initialized Router.
func New() *Carrot {
	return &Carrot{
		handlers: make(map[string]*handler),

		readTimeout:  10 * time.Second,
		writeTimeout: 10 * time.Second,
	}
}

// ServeHTTP implement http.Handle
func (c *Carrot) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if h, ok := c.handlers[path]; ok {
		if req.Method == h.method {
			h.handle(w, req)

			return
		}
	}

	w.Write([]byte("hello world! " + path))
}

// SetTimeout OPTIONAL customize some detail Config
func (c *Carrot) SetTimeout(readTimeout, writeTimeout time.Duration) {
	c.readTimeout = readTimeout
	c.writeTimeout = writeTimeout
}

// Start run http server
func (c *Carrot) Start(address string) error {
	s := &http.Server{
		Addr:         address,
		Handler:      c,
		ReadTimeout:  c.readTimeout,
		WriteTimeout: c.writeTimeout,
	}

	return s.ListenAndServe()
}

// Handle regist a path route
func (c *Carrot) Handle(method, path string, handle Handle) {
	c.handlers[path] = &handler{method, handle}
}

// Get a shortcut for carrot.Handle(http.MethodGet, path, handle)
func (c *Carrot) Get(path string, handle Handle) {
	c.Handle(http.MethodGet, path, handle)
}

// Head a shortcut for carrot.Handle(http.MethodHead, path, handle)
func (c *Carrot) Head(path string, handle Handle) {
	c.Handle(http.MethodHead, path, handle)
}

// Post a shortcut for carrot.Handle(http.MethodPost, path, handle)
func (c *Carrot) Post(path string, handle Handle) {
	c.Handle(http.MethodPost, path, handle)
}

// Put a shortcut for carrot.Handle(http.MethodPut, path, handle)
func (c *Carrot) Put(path string, handle Handle) {
	c.Handle(http.MethodPut, path, handle)
}

// Patch a shortcut for carrot.Handle(http.MethodPatch, path, handle)
func (c *Carrot) Patch(path string, handle Handle) {
	c.Handle(http.MethodPatch, path, handle)
}

// Delete a shortcut for carrot.Handle(http.MethodDelete, path, handle)
func (c *Carrot) Delete(path string, handle Handle) {
	c.Handle(http.MethodDelete, path, handle)
}

// Connect a shortcut for carrot.Handle(http.MethodConnect, path, handle)
func (c *Carrot) Connect(path string, handle Handle) {
	c.Handle(http.MethodConnect, path, handle)
}

// Options a shortcut for carrot.Handle(http.MethodOptions, path, handle)
func (c *Carrot) Options(path string, handle Handle) {
	c.Handle(http.MethodOptions, path, handle)
}

// Trace a shortcut for carrot.Handle(http.MethodTrace, path, handle)
func (c *Carrot) Trace(path string, handle Handle) {
	c.Handle(http.MethodTrace, path, handle)
}
