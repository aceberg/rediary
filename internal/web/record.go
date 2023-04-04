package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aceberg/rediary/internal/check"
	"github.com/aceberg/rediary/internal/db"
	"github.com/aceberg/rediary/internal/models"
)

func addRecordHandler(w http.ResponseWriter, r *http.Request) {
	var rec models.Record

	date := r.FormValue("date")
	name := r.FormValue("name")
	minus := r.FormValue("minus")
	plus := r.FormValue("plus")

	if name != "" {
		n := strings.Split(name, ":")

		rec.Tag = n[0]
		rec.Name = n[1]
		rec.Color = AppConfig.TagMap[rec.Tag]

		rec.Minus, _ = strconv.Atoi(minus)
		rec.Plus, _ = strconv.Atoi(plus)
		rec.Total = rec.Plus - rec.Minus

		if date == "" {
			currentTime := time.Now()
			rec.Date = currentTime.Format("2006-01-02")
		} else {
			rec.Date = date
		}

		log.Println("INFO: new record", rec)

		db.Insert(AppConfig.DB, rec)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func delRecordHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	log.Println("INFO: deleting record with ID =", id)

	db.Delete(AppConfig.DB, id)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
