package app

import "errors"

// ProduceTitle produces title with the set mode
func ProduceTitle(mode string) (string, error) {
	switch mode {
	case "chile":
		return "Trabajadores al poder!", nil
	case "brazil":
		return "Oi sim sim sim, oi não não não", nil
	default:
		return "", errors.New("Mode is not set")
	}
}
