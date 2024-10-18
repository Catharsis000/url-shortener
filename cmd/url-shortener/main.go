package main

import (
	"fmt"

	"github.com/Catharsis000/url-shortener.git/internal/config"
)

func main() {

	cfg := config.MustLoad()
	fmt.Println(cfg)

	//TODO: init config: cleanenv библиотека

	//TODO: init logger: slog библиотека (import "log/slog")

	//TODO: init storage: sqlite

	//TODO: init router: chi, "chi render" ("net/http")

	//TODO: run server

}
