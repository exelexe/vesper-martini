package app

import (
    . "core/db"
    "schemas"
    "github.com/go-martini/martini"
    "github.com/codegangsta/martini-contrib/render"
)

func FindUser(params martini.Params, r render.Render) {
	var user schemas.User

    DB.First(&user, params["id"])
    if (user.Id != 0) {
        newmap := map[string]interface{}{"id":user.Id, "name": user.Name} 
        r.JSON(200, newmap)
    } else {
        newmap := map[string]interface{}{"error":"Bad Request"} 
        r.JSON(400, newmap)    	
    }
}
