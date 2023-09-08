package err

import (
	"fmt"
	launcher "mtgolauncher/backend/Launcher"
)

type AppError struct {
	Err      error
	Message  string
	Location string
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.Location, e.Message, e.Err)
}

func HandleError(location, message string) {
	app := launcher.NewLauncher()

	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		appErr := AppError{
			Err:      err,
			Message:  message,
			Location: location,
		}
		app.UI.Error("Error!", fmt.Sprintf("We encountered an error! Oh no!\n %s", appErr))
		fmt.Printf("Error: %s", appErr)
	}
}
