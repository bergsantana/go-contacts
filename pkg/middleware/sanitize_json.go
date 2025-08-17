package middleware

import (
	"encoding/json"

	"github.com/bergsantana/go-contacts/pkg/sanitize"
	"github.com/gofiber/fiber/v2"
)

func recursivelySanitize(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		return sanitize.StrictHTML(v)
	case map[string]interface{}:
		for key, val := range v {
			v[key] = recursivelySanitize(val)
		}
		return v
	case []interface{}:
		for i, val := range v {
			v[i] = recursivelySanitize(val)
		}
		return v
	default:
		return v
	}
}

func SanitizeJSONBody() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Sanitizar body de requests com json
		if c.Get("Content-Type") == "application/json" {
			body := c.Body()
			if len(body) > 0 {
				var data interface{}
				if err := json.Unmarshal(body, &data); err == nil {
					data = recursivelySanitize(data)

					newBody, _ := json.Marshal(data)
					c.Request().SetBodyRaw(newBody)
				}
			}
		}
		return c.Next()
	}
}
