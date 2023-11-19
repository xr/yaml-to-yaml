package controllers

import (
	"bytes"
	"os"
	"text/template"

	"github.com/Unity-Technologies/unity-gateway-y2y/models"
)

func RenderRateLimiterActions(config *models.Config) (string, error) {
	templateData, err := os.ReadFile("views/rate_limiter_actions.tpl")
	if err != nil {
		return "", err
	}

	var renderedTemplate bytes.Buffer

	tmpl := template.Must(template.New("rateLimiterActions").Parse(string(templateData)))
	err = tmpl.Execute(&renderedTemplate, config)
	if err != nil {
		return "", err
	}

	renderedActions := renderedTemplate.String()

	return renderedActions, nil
}
