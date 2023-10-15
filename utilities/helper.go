package utilities

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
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

	de := os.Getenv("XDG_CURRENT_DESKTOP")

	deStr := strings.ToLower(de)
	log.Println(deStr)

	if strings.Contains(string(deStr), "gnome") {
		return globals.DE_GNOME, nil
	}
	if strings.Contains(string(deStr), "kde") {
		return globals.DE_KDE, nil
	}
	if strings.Contains(string(deStr), "xfce") {
		return globals.DE_XFCE, nil
	}

	return globals.DE_INVALID, errors.New("cannot detect desktop environment")
}

/*
Gets runtime flags provided so that configuration options can be set
*/
func GetFlags() *globals.CmdFlags {
	directory := flag.String("dir", "", "directory where images are present for desktop background.")
	duration := flag.Int("dur", 60, "Shuffle Duration in seconds")
	flag.Parse()

	if directory == nil || *directory == "" {
		WriteKillReasonAndExit("dir flag is missing, supply -h flag to see options")
	}

	if duration == nil {
		duration = getAddress(globals.DUR_DEFAULT)
	}

	result := globals.CmdFlags{
		Duration:  *duration,
		Directory: *directory,
	}

	return &result
}

func GetAllImageFilePathsFromDirectory(dirPath string) []string {
	result := make([]string, 0)

	d, err := os.ReadDir(dirPath)
	if err != nil {
		WriteKillReasonAndExit(fmt.Sprintf("error while getting list of items in given directory, err : %s", err.Error()))
	}

	for _, f := range d {
		if f.IsDir() {
			continue
		}

		nameArr := strings.Split(f.Name(), ".")
		if !(strings.EqualFold(nameArr[len(nameArr)-1], "jpg") || strings.Contains(nameArr[len(nameArr)-1], "jpeg")) {
			continue
		}

		result = append(result, filepath.Join(dirPath, f.Name()))
	}

	return result
}

// This function is called in case program cannot continue anymore, and writes the reason to quit to a file for debugging purposes
func WriteKillReasonAndExit(reason string) {
	fp := filepath.Join(os.TempDir(), globals.TEMP_FAILURE_FILE)
	f, err := os.Create(fp)
	if err != nil {
		log.Fatalln("error while creating temp file for failure reason", "err", err.Error())
	}
	defer f.Close()
	log.Println(reason)
	f.WriteString(reason)
	os.Exit(1)
}

func getAddress[T any](t T) *T {
	return &t
}

func SetBackground(osType globals.OS_TYPE, deType globals.DE_TYPE, fp string) error {
	switch osType {
	case globals.OS_LINUX:
		switch deType {
		case globals.DE_GNOME:
			err := setBackgroundGNOME(fp)
			if err != nil {
				return err
			}
			return nil

		default:
			return errors.New("current desktop environment is not supported")
		}

	default:
		return errors.New("current os type is not supported")
	}
}

func setBackgroundGNOME(filepath string) error {

	cs, err := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "color-scheme").Output()

	if err != nil {
		return fmt.Errorf("error while getting desktop interface, err : %s", err.Error())
	}

	if strings.Contains(string(cs), "prefer-dark") {
		output, err := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", fmt.Sprintf("'file://%s'", filepath)).Output()
		if err != nil {
			return fmt.Errorf("error while setting background for GNOME, err : %s", err.Error())
		}

		log.Println(string(output))
		return nil
	}

	output, err := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", fmt.Sprintf("'file://%s'", filepath)).Output()
	if err != nil {
		return fmt.Errorf("error while setting background for GNOME, err : %s", err.Error())
	}

	log.Println(string(output))
	return nil

}

func RandomRangeGenerator(end int) []int {

	ur := make(map[int]bool)
	result := make([]int, 0)

	for len(ur) != end-1 {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(end)))
		if err != nil {
			log.Panicln("failed to generate random number", "err", err.Error())
		}

		if !ur[int(n.Int64())] {
			fmt.Println(n.Int64())
			ur[int(n.Int64())] = true
			result = append(result, int(n.Int64()))
		}
	}

	return result
}
