package router

import (
	"log"
	"net/http"

	"github.com/satriowibowo1701/e-commorce-api/controller"
	"github.com/satriowibowo1701/e-commorce-api/execption"
)

type Handler struct {
	HandlerFunc http.HandlerFunc
	Path        string
	Method      string
}

type Method struct {
	Mux     *http.ServeMux
	Handler map[string]map[string]*Handler
	Cache   string
}

func (m *Method) Set(path string, method string, handler http.HandlerFunc) *Method {
	if m.Handler[path] == nil {
		m.Handler[path] = make(map[string]*Handler)
	}
	if m.Handler[path][method] == nil {
		m.Handler[path][method] = &Handler{}
		m.Mux.HandleFunc(path, handler)
	}
	m.Handler[path][method].HandlerFunc = handler
	m.Handler[path][method].Method = method
	m.Handler[path][method].Path = path
	return m

}
func (m *Method) Get(w http.ResponseWriter, r *http.Request) bool {
	if m.Handler[r.URL.Path] == nil {
		http.Error(w, "Not Path Found ", http.StatusNotFound)
		log.Println("error3")
		return false
	}
	if m.Handler[r.URL.Path][r.Method] == nil {
		http.Error(w, "Method Not Allowed ", http.StatusMethodNotAllowed)
		log.Println("error2")
		return false
	}

	return true
}

func (m *Method) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if m.Get(w, r) {
		log.Println("error1")
		m.Mux.ServeHTTP(w, r)
		return
	}
	execption.NotAllowed(w)

}

func NewRouter(ProdukController controller.ProdukController, TransactionController controller.TransactionController, UserConttoller controller.UserController) *Method {
	router := &http.ServeMux{}
	newRouter := &Method{
		Mux:     router,
		Handler: make(map[string]map[string]*Handler),
	}
	newRouter.Set("/api/v1/user/register", "POST", UserConttoller.RegisterUser)
	newRouter.Set("/api/v1/user/login", "POST", UserConttoller.LoginUser)
	newRouter.Set("/api/v1/user/updateuser", "POST", UserConttoller.UpdateUser)
	newRouter.Set("/api/v1/user/getallusers", "GET", UserConttoller.FindAllUsers)
	newRouter.Set("/api/v1/user/gettrxcusbyid", "POST", TransactionController.GetAllTransactionsByIdCus)
	newRouter.Set("/api/v1/user/gettrxcus", "GET", TransactionController.GetAllTransactionsCus)
	newRouter.Set("/api/v1/user/createtrx", "GET", TransactionController.CreateTransaction)
	newRouter.Set("/api/v1/transaction/gettrxstatus", "GET", TransactionController.GetAllTransactionsByStatusCus)
	newRouter.Set("/api/v1/transaction/inserttmptrx", "POST", TransactionController.InsertTransactionsTmp)
	newRouter.Set("/api/v1/transaction/deletetmptrx", "GET", TransactionController.DeleteTransactionsTmp)
	newRouter.Set("/api/v1/transaction/updatetmptrx", "POST", TransactionController.UpdateTransactionsTmp)
	return newRouter
}
