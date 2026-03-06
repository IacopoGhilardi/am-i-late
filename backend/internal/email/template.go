package email

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"sync"
)

type TemplateName string

const (
	TemplateResetPasswordEmail TemplateName = "reset_password"
)

type TemplateEngine struct {
	templatesDir string
	cache        map[TemplateName]*template.Template
	mu           sync.RWMutex
}

func NewTemplateEngine(templatesDir string) *TemplateEngine {
	return &TemplateEngine{
		templatesDir: templatesDir,
		cache:        make(map[TemplateName]*template.Template),
	}
}

func (e *TemplateEngine) Render(name TemplateName, data any) (string, error) {
	tmpl, err := e.load(name)
	if err != nil {
		return "", fmt.Errorf("errore caricamento template %s: %w", name, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("errore rendering template %s: %w", name, err)
	}

	return buf.String(), nil
}

func (e *TemplateEngine) load(name TemplateName) (*template.Template, error) {
	e.mu.RLock()
	if tmpl, ok := e.cache[name]; ok {
		e.mu.RUnlock()
		return tmpl, nil
	}
	e.mu.RUnlock()

	path := filepath.Join(e.templatesDir, string(name)+".html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("template non trovato: %s", path)
	}
	
	e.mu.Lock()
	e.cache[name] = tmpl
	e.mu.Unlock()

	return tmpl, nil
}
