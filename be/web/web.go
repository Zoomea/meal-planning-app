package web

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/Zoomea/meal-planning-app/biz"
	"github.com/Zoomea/meal-planning-app/db"
	"github.com/Zoomea/meal-planning-app/db/fsdatabase"
)

type httpHandler func(http.ResponseWriter, *http.Request)

type State struct {
	recipeDB   db.Crudler[biz.Recipe]
	scheduleDB db.ScheduleStore
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

	http.HandleFunc("GET /api/schedule/", wrap(getManySchedules))
	http.HandleFunc("POST /api/schedule/", wrap(addSchedule))

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
		recipeDB:   fsdatabase.New[biz.Recipe](),
		scheduleDB: fsdatabase.NewScheduleDB(),
	}
}

func log(f httpHandler) httpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("Incoming request: %s %s\n", req.Method, req.URL)
		f(w, req)
	}
}


func sendJSON(f func(*http.Request, State) (any, int, error), state State) httpHandler {
	return func(w http.ResponseWriter, req *http.Request) {
		var res Response[any]
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
