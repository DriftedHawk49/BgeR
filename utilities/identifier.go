package utilities

import (
	"errors"
	"os/exec"
	"runtime"
	"strings"

	"github.com/BgeR/globals"
)

/*
Validates the OS on which tool is being run. Currently only linux is supported
*/
func ValidateOSType() globals.OS_TYPE {
	switch runtime.GOOS {
	case "linux":
		return globals.OS_LINUX
	default:
		return globals.OS_TYPE(strings.ToUpper(runtime.GOOS))
	}
}

/*
Returns the type of Desktop Environment present.
*/
func GetDE() (globals.DE_TYPE, error) {
	de, err := exec.Command("echo", "$XDG_CURRENT_DESKTOP").Output()
	if err != nil {
		return globals.DE_INVALID, err
	}

	if strings.Contains(string(de), "gnome") {
		return globals.DE_GNOME, nil
	}
	if strings.Contains(string(de), "kde") {
		return globals.DE_KDE, nil
	}
	if strings.Contains(string(de), "xfce") {
		return globals.DE_XFCE, nil
	}

	return globals.DE_INVALID, errors.New("cannot detect desktop environment")
}
