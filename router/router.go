package router

import (
	"fmt"
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
	Mux        *http.ServeMux
	Handler    map[string]map[string]*Handler
	Middleware map[string]string
	Cache      string
}

func (m *Method) Set(path string, method string, handler http.HandlerFunc, middleware string) *Method {
	if m.Handler[path] == nil {
		m.Handler[path] = make(map[string]*Handler)
	}
	if m.Handler[path][method] == nil {
		m.Handler[path][method] = &Handler{}
		m.Mux.HandleFunc(path, handler)
	}
	m.Middleware[path] = middleware
	m.Handler[path][method].HandlerFunc = handler
	m.Handler[path][method].Method = method
	m.Handler[path][method].Path = path
	return m

}
func (m *Method) Get(w http.ResponseWriter, r *http.Request) bool {
	if m.Handler[r.URL.Path] == nil {
		http.Error(w, "Not Path Found ", http.StatusNotFound)

		return false
	}
	if m.Handler[r.URL.Path][r.Method] == nil {
		http.Error(w, "Method Not Allowed ", http.StatusMethodNotAllowed)
		
		return false
	}

	return true
}

func (m *Method) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Get(w, r) {
		m.Mux.ServeHTTP(w, r)
		return
	}
	execption.NotAllowed(w)

}

func NewRouter(ProdukController controller.ProdukController, TransactionController controller.TransactionController, UserConttoller controller.UserController) *Method {
	router := &http.ServeMux{}
	newRouter := &Method{
		Mux:        router,
		Handler:    make(map[string]map[string]*Handler),
		Middleware: make(map[string]string),
	}
	//noauth
	newRouter.Set("/api/v1/user/register", "POST", UserConttoller.RegisterUser, "noauth")
	newRouter.Set("/api/v1/user/login", "POST", UserConttoller.LoginUser, "noauth")
	//general
	newRouter.Set("/api/v1/user/logout", "GET", UserConttoller.Logout, "general")
	newRouter.Set("/api/v1/user/updateuser", "PUT", UserConttoller.UpdateUser, "general")
	newRouter.Set("/api/v1/user/getuserbyid", "GET", UserConttoller.FindAllUsers, "general")
	newRouter.Set("/api/v1/user/profile", "GET", UserConttoller.FindByUserid, "general")
	newRouter.Set("/api/v1/produk/all", "GET", ProdukController.FindAll, "general")
	//admin
	newRouter.Set("/api/v1/user/getallusers", "GET", UserConttoller.FindAllUsers, "admin")
	newRouter.Set("/api/v1/transaction/getalltrxcus", "GET", TransactionController.GetAllTransactionsCus, "admin")
	newRouter.Set("/api/v1/product/create", "POST", ProdukController.CreateProduct, "admin")
	newRouter.Set("/api/v1/product/update", "PUT", ProdukController.UpdateProduct, "admin")
	newRouter.Set("/api/v1/product/delete", "DELETE", ProdukController.DeleteProduct, "admin")
	//cus
	newRouter.Set("/api/v1/transaction/createtrx", "POST", TransactionController.CreateTransaction, "customer")
	newRouter.Set("/api/v1/transaction/gettrxcusbyid", "POST", TransactionController.GetAllTransactionsByIdCus, "customer")
	newRouter.Set("/api/v1/transaction/gettrxstatus", "GET", TransactionController.GetAllTransactionsByStatusCus, "customer")
	newRouter.Set("/api/v1/transaction/inserttmptrx", "POST", TransactionController.InsertTransactionsTmp, "customer")
	newRouter.Set("/api/v1/transaction/deletetmptrx", "DELETE", TransactionController.DeleteTransactionsTmp, "customer")
	newRouter.Set("/api/v1/transaction/updatetmptrx", "PUT", TransactionController.UpdateTransactionsTmp, "customer")
	fmt.Println(newRouter.Middleware)

	return newRouter
}
