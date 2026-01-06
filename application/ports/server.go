package ports

type HTTPServer interface {
	Run(port string)
	RegisterRoutes()
}
