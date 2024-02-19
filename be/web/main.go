package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type httpHandler func(http.ResponseWriter, *http.Request)

func Serve(dir string, port int) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("unable to serve directory '%s' as it doesn't exist", dir)
	}

	if s := os.Getenv("GODEBUG"); strings.Contains(s, "httpmuxgo121=1") {
		return fmt.Errorf("Go 1.22 Mux was disabled with the httpmuxgo121 GODEBUG variable. Remove this option or explicitly set httpmuxgo121=0 in order to run this server")
	}

	fileServer := http.StripPrefix(dir, http.FileServer(http.Dir(dir)))
	http.Handle("/", fileServer)

	// Middleware
	wrap := func(f func(*http.Request) (any, int, error)) httpHandler {
		return log(sendJSON(f))
	}

	http.HandleFunc("GET /recipe/", wrap(getRecipes))
	http.HandleFunc("GET /recipe/{id}", wrap(getRecipe))

	fmt.Printf("Serving directory '%s' at http://localhost:%d\n", dir, port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getRecipes(req *http.Request) (any, int, error) {
	return "all", 200, nil
}

func getRecipe(req *http.Request) (any, int, error) {
	idStr := req.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, 400, fmt.Errorf("Could not parse '%s' as integer", idStr)
	}

	return fmt.Sprintf("wew %d", id), 200, nil
}

func log(f httpHandler) httpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("Incoming request: %s %s\n", req.Method, req.URL)
		f(w, req)
	}
}

func sendJSON(f func(*http.Request) (any, int, error)) httpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		data, code, err := f(req)
		if err != nil {
			w.WriteHeader(code)
			data = fmt.Sprintf(`{"error":"%v"}`, err)
			fmt.Printf("err: %+v", err)
		}

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Printf("err: %s", err)
		}
	}
}
