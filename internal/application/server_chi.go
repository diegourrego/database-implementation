package application

import (
	"database/sql"
	"database_implementation/internal/handler"
	"database_implementation/internal/repository"
	"database_implementation/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ConfigServerChi struct {
	Addr     string
	MySQLDSN string
}

func NewServerChi(cfg ConfigServerChi) *ServerChi {
	defaultCfg := ConfigServerChi{
		Addr:     ":8080",
		MySQLDSN: "",
	}
	if cfg.Addr != "" {
		defaultCfg.Addr = cfg.Addr
	}
	if cfg.MySQLDSN != "" {
		defaultCfg.MySQLDSN = cfg.MySQLDSN
	}

	return &ServerChi{
		addr:     defaultCfg.Addr,
		mysqlDSN: defaultCfg.MySQLDSN,
	}
}

type ServerChi struct {
	addr     string
	mysqlDSN string
}

func (sc *ServerChi) Run() (err error) {
	db, err := sql.Open("mysql", sc.mysqlDSN)
	if err != nil {
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		return
	}

	// repositories
	rp := repository.NewProductRepository(db)
	// services
	sv := service.NewProductDefault(rp)
	// handlers
	hd := handler.NewProductDefault(sv)
	// router
	router := chi.NewRouter()

	router.Route("/products", func(rt chi.Router) {
		rt.Get("/", hd.GetAll())
		rt.Get("/{id}", hd.GetOne())
		rt.Post("/", hd.Save())
	})

	err = http.ListenAndServe(sc.addr, router)
	return
}
