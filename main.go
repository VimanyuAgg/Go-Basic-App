package main

import (
	"fmt"
	"net/http"

	"github.com/basicGo/webservice/controllers"

	"github.com/basicGo/webservice/models"
)

func main() {
	fmt.Println("hello go!")
	u := models.User{
		ID:        2,
		FirstName: "Vimanyu",
		LastName:  "Aggarwal",
	}

	fmt.Println(u)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
