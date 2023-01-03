package handlers

import (
	"net/http"

	"github.com/nikbat/bookings/cmd/pkg/config"
	"github.com/nikbat/bookings/cmd/pkg/models"
	"github.com/nikbat/bookings/cmd/pkg/render"
)

// Repository Pattern starts

var Repo *Repository

// repositiry type
type Repository struct {
	App *config.AppConfig
}

// creates a new respository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets a repository *instead of using name SetHandlers
func SetRepository(r *Repository) {
	Repo = r
}

//Repository pattern end

func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// n, err := fmt.Fprintf(w, "Jai Saraswati Mata!")
	// if err != nil {
	// 	log.Println("Error service request ", err)
	// }
	// log.Println("Number of bytes", n)

	re.App.SessionManager.Put(r.Context(), "remote_ip", r.RemoteAddr)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	//_, _ = fmt.Fprintf(w, "This is a about page")

	stringMap := make(map[string]string)
	stringMap["test"] = re.App.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = re.App.SessionManager.GetString(r.Context(), "remote_ip")

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

// func Devide(w http.ResponseWriter, r *http.Request) {
// 	x := 100.0
// 	y := 10.0

// 	n, err := devideNumber(x, y)

// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot devide by zero")
// 		return
// 	}
// 	fmt.Fprintf(w, fmt.Sprintf("%f device by %f is %f", x, y, n))
// }

// func devideNumber(x, y float64) (float64, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot devide by zero")
// 		return 0.0, err
// 	}
// 	return x / y, nil
// }
