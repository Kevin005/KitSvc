package wsutil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

// Engine stores the Gin engine and the handler functions.
type Engine struct {
	Gin *gin.Engine
}

// New creates the new WebSocket handler engine.
func New(e *gin.Engine) *Engine {
	return &Engine{Gin: e}
}

// Handle handles the incoming websocket request with the specified path.
func (e *Engine) Handle(relativePath string, handler gin.HandlerFunc) {
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	e.Gin.GET(relativePath, func(c *gin.Context) {
		// Put the melody in context.
		c.Set("websocket", m)
		handler(c)
	})
}
