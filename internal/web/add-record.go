package web

import (
	// "log"
	"net/http"
	"strconv"
	"time"

	"github.com/aceberg/red/internal/db"
	"github.com/aceberg/red/internal/models"
)

func addRecordHandler(w http.ResponseWriter, r *http.Request) {
	var rec models.Record

	date := r.FormValue("date")
	rec.Name = r.FormValue("name")
	minus := r.FormValue("minus")
	plus := r.FormValue("plus")

	rec.Minus, _ = strconv.Atoi(minus)
	rec.Plus, _ = strconv.Atoi(plus)
	rec.Total = rec.Plus - rec.Minus

	if date == "" {
		currentTime := time.Now()
		rec.Date = currentTime.Format("2006-01-02")
	} else {
		rec.Date = date
	}

	// log.Println("INFO: new record", rec)

	db.Insert(AppConfig.DB, rec)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
