/*
	This file controls the algorithm of changing backgrounds after a certain time
*/

package utilities

import (
	"time"

	"github.com/BgeR/globals"
)

func Init(params *globals.CmdFlags, osType globals.OS_TYPE, deType globals.DE_TYPE) {

	for {

		files := GetAllImageFilePathsFromDirectory(params.Directory)
		randomRange := RandomRangeGenerator(len(files))

		for _, i := range randomRange {
			err := SetBackground(osType, deType, files[i])
			if err != nil {
				WriteKillReasonAndExit(err.Error())
			}
			time.Sleep(time.Duration(params.Duration) * time.Second)
		}

	}

}
