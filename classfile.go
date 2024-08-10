package amisgo

const GopPackage = true

type (
	amis interface {
		init()
		HandleRouter()
	}
	IMain interface {
		amis
	}
	IWorker interface {
		amis
	}
)

type App = Amis

var (
	_ IMain   = (*App)(nil)
	_ IWorker = (*Amis)(nil)
)

func Gopt_App_Main(app IMain, workers ...IWorker) {
	for _, worker := range workers {
		worker.init()
		worker.(interface{ Main() }).Main()
		worker.HandleRouter()
	}
	app.init()
	app.HandleRouter()
	app.(interface{ MainEntry() }).MainEntry()
}
