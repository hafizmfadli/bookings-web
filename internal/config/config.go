package config

// note config ini hanya boleh depend ke standard package
// jangan sampe ngiimport bagian package dari aplikasi (karena berpotensi menghasilkan error)
import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
