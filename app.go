package app
type App struct {
	logger bool
	err error
}

func (App App) Error()string{
	return v.Error()
}

func (App *App) IsErr(errs ...error) bool {
	if App.err != nil {
		return true
	}
	for _, err := range errs {
		if err != nil {
			App.err = err
			return true
		}
	}
	return false
}

func (App App) Println(v ...interface{}) {
	if App.logger {
		log.Println(v...)
	}
}

func (App App) Printf(format string, v ...interface{}) {
	if App.logger {
		log.Printf(format, v...)
	}
}

func (App App) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
