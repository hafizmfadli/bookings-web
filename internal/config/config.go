package config

// note config ini hanya boleh depend ke standard package
// jangan sampe ngiimport bagian package dari aplikasi (karena berpotensi menghasilkan error)
import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/hafizmfadli/bookings-web/internal/models"
)

// AppConfig holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
