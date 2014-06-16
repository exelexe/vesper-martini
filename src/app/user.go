package app

import (
    . "core/db"
    "log"
    "schemas"
    "github.com/go-martini/martini"
)

func FindUser(params martini.Params) string {
	var user schemas.User
    DB.First(&user, params["id"])
    log.Println(user)
    //return "Hello " + params["id"]
    return "name: " + user.Name
	//return Must(enc.Encode(toIface(db.GetAll())...))
}

//func GetAlbums(r *http.Request, enc Encoder, db DB) string {
//	// Get the query string arguments, if any
//	qs := r.URL.Query()
//	band, title, yrs := qs.Get("band"), qs.Get("title"), qs.Get("year")
//	yri, err := strconv.Atoi(yrs)
//	if err != nil {
//		// If year is not a valid integer, ignore it
//		yri = 0
//	}
//	if band != "" || title != "" || yri != 0 {
//		// At least one filter, use Find()
//		return Must(enc.Encode(toIface(db.Find(band, title, yri))...))
//	}
//	// Otherwise, return all albums
//	return Must(enc.Encode(toIface(db.GetAll())...))
//}
