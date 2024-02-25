package web

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Zoomea/meal-planning-app/biz"
	"github.com/Zoomea/meal-planning-app/db"
	"github.com/Zoomea/meal-planning-app/db/fsdatabase"
)

type httpHandler func(http.ResponseWriter, *http.Request)

type State struct {
	recipeDB db.Crudler[biz.Recipe]
}

func Serve(dir string, port int, ready chan<- struct{}) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("unable to serve directory '%s' as it doesn't exist", dir)
	}

	if s := os.Getenv("GODEBUG"); !strings.Contains(s, "httpmuxgo121=0") {
		return fmt.Errorf("you must explicitly enabled the Go 1.22 HTTP server mux by setting the httpmuxgo121 GODEBUG variable to zero. Otherwise, it seemly like Go disabled the mux by default")
	}

	fileServer := http.StripPrefix(dir, http.FileServer(http.Dir(dir)))
	http.Handle("/", fileServer)

	state := initState()

	// Middleware
	wrap := func(f func(*http.Request, State) (any, int, error)) httpHandler {
		return log(sendJSON(f, state))
	}

	http.HandleFunc("GET /api/recipe/", wrap(getManyRecipes))
	http.HandleFunc("GET /api/recipe/{id}", wrap(getRecipe))
	http.HandleFunc("POST /api/recipe/", wrap(addRecipe))
	http.HandleFunc("DELETE /api/recipe/{id}", wrap(deleteRecipe))

	fmt.Printf("Serving directory '%s' at http://localhost:%d\n", dir, port)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	// Signal that the server is ready to start receiving requests
	ready <- struct{}{}

	return http.Serve(l, nil)
}

func initState() State {
	return State{
		recipeDB: fsdatabase.New[biz.Recipe](),
	}
}

func getManyRecipes(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	recipes, err := biz.ListRecipes(ctx, state.recipeDB)
	if err != nil {
		return nil, 500, err
	}
	return recipes, 200, nil
}

func getRecipe(req *http.Request, state State) (any, int, error) {
	idStr := req.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, 400, fmt.Errorf("Could not parse '%s' as integer", idStr)
	}

	ctx := req.Context()

	recipes, err := biz.GetRecipes(ctx, state.recipeDB, []int64{id})
	if err != nil {
		return nil, 500, err
	}
	if len(recipes) == 0 {
		return nil, 404, fmt.Errorf("no recipe with id %d", id)
	}

	return recipes[0], 200, nil
}

func addRecipe(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	var rec biz.Recipe
	if err := json.NewDecoder(req.Body).Decode(&rec); err != nil {
		return nil, 400, err
	}

	id, err := biz.AddRecipe(ctx, state.recipeDB, rec)
	if err != nil {
		return nil, 500, err
	}
	return id, 200, nil
}

func deleteRecipe(req *http.Request, state State) (any, int, error) {
	ctx := req.Context()

	idStr := req.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, 400, fmt.Errorf("Could not parse '%s' as integer", idStr)
	}

	err = biz.DeleteRecipe(ctx, state.recipeDB, id)
	if err != nil {
		return nil, 500, err
	}
	return nil, 200, nil
}

func log(f httpHandler) httpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("Incoming request: %s %s\n", req.Method, req.URL)
		f(w, req)
	}
}

type Response struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

func sendJSON(f func(*http.Request, State) (any, int, error), state State) httpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		var res Response
		data, code, err := f(req, state)
		if err != nil {
			res.Error = err.Error()
		} else {
			res.Data = data
		}

		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")

		output, _ := json.Marshal(res)

		fmt.Printf("Returned: code=%d body=%s\n", code, output)

		_, err = w.Write(output)
		if err != nil {
			fmt.Printf("err: %s", err)
		}
	}
}
