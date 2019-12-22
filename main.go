package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/", addHandler)
	e.POST("/mul", mulHandler)
	e.Start(":8080")

	// m := http.NewServeMux()
	// m.Handle("/", http.HandlerFunc(indexHandler))
	// m.Handle("/about", http.HandlerFunc(aboutHandler))

	// :8080 => 0.0.0.0:8080
	// http.ListenAndServe(":8080", http.HandlerFunc(router))
	// http.ListenAndServe(":8080", http.HandlerFunc(handler))

	// http.ListenAndServe(":8080", chainMiddleware(
	// 	m1,
	// 	m2,
	// 	m3,
	// )(m))
}

type addRequest struct {
	A int `json:"x"`
	B int `json:"y"`
}

type addResponse struct {
	Result int `json:"res,omitempty"`
}

func addHandler(c echo.Context) error {
	//
	var reqBody addRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return err
	}
	//
	result := reqBody.A + reqBody.B
	//
	return c.JSON(200, addResponse{
		Result: result,
	})
}

func mulHandler(c echo.Context) error {
	//
	var reqBody addRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return err
	}
	//
	result := reqBody.A * reqBody.B
	//
	return c.JSON(200, addResponse{
		Result: result,
	})
}

// var indexTmpl = template.Must(template.New("").Parse(`
// 	<!doctype html>
// 	<title>Calculator</title>
// 	<h1>Calculator</h1>

// 	<form method=POST>
// 		<input name=a>
// 		<input name=b>
// 		<button>+</button>
// 	</form>
// 	{{if .}}
// 	<p>Result:{{.A}} + {{.B}} = {{.Result}}</p>
// 	{{end}}

// `))

// type IndexData struct {
// 	A      int
// 	B      int
// 	Result int
// }

// HTTP Handler

func handler(w http.ResponseWriter, r *http.Request) {
	// parse request
	var reqBody addRequest
	json.NewDecoder(r.Body).Decode(&reqBody)
	// json.Unmarshal()
	//calculate
	result := reqBody.A + reqBody.B
	//send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addResponse{
		Result: result,
	})
	// if r.Method == "POST" {
	// 	// form
	// 	a, _ := strconv.Atoi(r.PostFormValue("a"))
	// 	b, _ := strconv.Atoi(r.PostFormValue("b"))

	// 	result := a + b
	// 	w.Header().Set("Content-text", "text/html")
	// 	indexTmpl.Execute(w, IndexData{
	// 		A:      a,
	// 		B:      b,
	// 		Result: result,
	// 	})

	// 	fmt.Println(a + b)
	// 	// fmt.Println(b)

	// 	return
	// }
	// w.Header().Set("Content-text", "text/html")
	// indexTmpl.Execute(w, nil)
	// w.WriteHeader(200)
	// w.Write([]byte("<h1>Hello</h1>"))
}

// Router

// func router(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		//indexHandler()
// 		indexHandler(w, r)
// 	} else if r.URL.Path == "/about" {
// 		//aboutHandler
// 		aboutHandler(w, r)
// 	} else {
// 		//notFoundHandler 404
// 		notFoundHandler(w, r)
// 	}
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Index Page"))
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("About Page"))
// }

// func notFoundHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Not Found 404"))
// }

// middle ware

// func main() {
// 	// m := http.NewServeMux()
// 	// m.Handle("/", http.HandlerFunc(indexHandler))
// 	// m.Handle("/about", http.HandlerFunc(aboutHandler))
// 	// :8080 => 0.0.0.0:8080
// 	http.ListenAndServe(":8080",chainMidleware(
// 	m1,
// 	m2,
// 	m3,

// )(m))
// }

// func addContentType(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/html")
// 		h.ServeHTTP(w, r)
// 	})
// }

// func m1(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("m1 before")
// 		h.ServeHTTP(w, r)
// 		fmt.Println("m1 after")
// 	})
// }

// func m2(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("m2 before")
// 		h.ServeHTTP(w, r)
// 		fmt.Println("m2 after")
// 	})
// }

// func m3(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("m3 before")
// 		h.ServeHTTP(w, r)
// 		fmt.Println("m3 after")
// 	})
// }

// type middleware func(http.Handler) http.Handler

// func chainMiddleware(ms ...middleware) middleware {
// 	return func(h http.Handler) http.Handler {

// 		for i := len(ms) - 1; i >= 0; i-- {
// 			h = ms[i](h)
// 		}
// 		return h
// 	}
// }

// // func add(x int, y int) int {
// // 	return x + y
// // }

// // func main() {
// // 	// var name string = "bom"
// // 	name := "bomtaku"
// // 	r := add(1, 12)
// // 	fmt.Printf("Hello World, %s\n1 + 12 = %d", name, r)
// // }
