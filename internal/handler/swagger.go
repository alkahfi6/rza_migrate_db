package handler

import (
	_ "embed"
	"fmt"
	"html"
	"net/http"
)

var (
	//go:embed swagger.yaml
	swaggerSpecYAML []byte
)

func (h *MigrationHandler) handleSwaggerSpec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeMethodNotAllowed(w)
		return
	}
	w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(swaggerSpecYAML)
}

func (h *MigrationHandler) handleSwaggerUI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeMethodNotAllowed(w)
		return
	}

	page := fmt.Sprintf(`<!doctype html>
<html>
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>%s - Swagger UI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.ui = SwaggerUIBundle({
      url: '/swagger.yaml',
      dom_id: '#swagger-ui',
      deepLinking: true,
      defaultModelsExpandDepth: 1,
      defaultModelExpandDepth: 1
    });
  </script>
</body>
</html>`, html.EscapeString(h.serviceName))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(page))
}
