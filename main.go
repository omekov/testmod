package testmod

import "fmt"

// Hi - it is hello method
func Hi(name, lang string) string {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name)
	case "kz":
		return fmt.Sprintf("Калайсын, %s!", name)
	case "ru":
		return fmt.Sprintf("Привет, %s!", name)
	default:
		return fmt.Sprintf("#$%^&*(), %s!", name)
	}

}
