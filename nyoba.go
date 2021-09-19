// package main

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo/v4"
// )

// type User struct {
// 	Id      int    `json:"id"`
// 	Name    string `json:"name"`
// 	Age     int    `json:"age"`
// 	Email   string `json:"email"`
// 	Address string `json:"address"`
// }

// type BaseResponse struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// type UserLogin struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func main() {
// 	//endpoint
// 	// http.HandleFunc("/v1/users", GetUserController)
// 	// fmt.Println("Server REST API telah berjalan di port 8000")
// 	// http.ListenAndServe(":8000", nil)
// 	e := echo.New()
// 	ev1 := e.Group("v1/")
// 	ev1.GET("users", GetUserController)
// 	ev1.POST("users/login", LoginController)
// 	ev1.GET("users/:userId", DetailUserController)
// 	e.Start(":8000")

// }

// func LoginController(c echo.Context) error {
// 	userLogin := UserLogin{}
// 	c.Bind(&userLogin)
// 	// email := c.FormValue("email")
// 	// password := c.FormValue("password")
// 	return c.JSON(http.StatusOK, BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    userLogin,
// 	})
// }

// func DetailUserController(c echo.Context) error {
// 	userId, err := strconv.Atoi(c.Param("userId"))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, BaseResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Gagal koneksi userId",
// 			Data:    nil,
// 		})
// 	}
// 	return c.JSON(http.StatusOK, BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    User{Id: userId},
// 	})
// }

// func GetUserController(c echo.Context) error {
// 	user := User{}
// 	name := c.QueryParam("name")
// 	if name == "" {
// 		user = User{1, "Diah", 21, "diahaufaarini5@gmail.com", "Bekasi"}
// 	} else {
// 		user = User{1, name, 21, "diahaufaarini5@gmail.com", "Bekasi"}
// 	}
// 	return c.JSON(http.StatusOK, BaseResponse{
// 		Code:    http.StatusOK,
// 		Message: "Berhasil",
// 		Data:    user,
// 	})

// 	// if r.Method == "GET" {
// 	// 	response, _ := http.Get("https://swapi.dev/api/people/1")
// 	// 	responseBody, _ := ioutil.ReadAll(response.Body)
// 	// 	defer response.Body.Close()

// 	// 	peopleSwapi := PeopleSwapi{}
// 	// 	json.Unmarshal(responseBody, &peopleSwapi)

// 	// 	user := User{peopleSwapi.Name, peopleSwapi.Mass, peopleSwapi.Name, peopleSwapi.Hair_Color}

// 	// 	// user := []User{
// 	// 	// 	{"Diah", 21, "diahaufaarini5@gmail.com", "Bekasi"},
// 	// 	// 	{"Vernon", 23, "vernonchwe@gmail.com", "Seoul"},
// 	// 	// }
// 	// 	resultJSON, err := json.Marshal(user)
// 	// 	if err != nil {
// 	// 		http.Error(w, "Gagal convert", http.StatusInternalServerError) //untuk code bisa pake code 500
// 	// 		return
// 	// 	}
// 	// 	w.Write(resultJSON)
// 	// 	return
// 	// }
// 	// http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// }
