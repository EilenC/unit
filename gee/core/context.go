package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	W http.ResponseWriter
	R *http.Request

	// request info
	Path   string
	Method string
	Params map[string]string

	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r, Path: r.URL.Path, Method: r.Method}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

func (c *Context) PostForm(key string) string {
	return c.R.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.R.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, _ = io.WriteString(c.W, fmt.Sprintf(format, values...))
}

func (c *Context) JSON(code int, obj interface{}) {
	marshal, err := json.Marshal(obj)
	if err != nil {
		http.Error(c.W, err.Error(), http.StatusInternalServerError)
		return
	}
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	_, _ = c.W.Write(marshal)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, _ = c.W.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, _ = io.WriteString(c.W, html)
}
