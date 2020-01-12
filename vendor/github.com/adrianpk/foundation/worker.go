package kabestan

type (
	Worker struct {
		Cfg  *Config
		Log  Logger
		Name string
	}
)

func NewWorker(cfg *Config, log Logger, name string) *Worker {
	name = genName(name, "worker")

	return &Worker{
		Cfg:  cfg,
		Log:  log,
		Name: name,
	}
}
