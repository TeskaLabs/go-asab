package asab

type Service struct {
	App *Application
}

type ServiceI interface {
	Finalize()
}

func (svc *Service) Initialize(app *Application) {
	svc.App = app
}

func (svc *Service) Finalize() {
	svc.App = nil
}
