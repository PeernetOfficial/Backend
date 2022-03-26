/*
File Name:  Update.go
Copyright:  2021 Peernet Foundation s.r.o.
Author:     Peter Kleissner
*/

package main

import (
	"github.com/PeernetOfficial/core"
	"github.com/PeernetOfficial/core/system"
)

func update(backend *core.Backend) {
	if config.UpdateFolder == "" {
		return
	}

	files, err := system.ParseUpdateFiles(config.UpdateFolder)
	if err != nil {
		backend.LogError("update", "read update files: %v\n", err)
		return
	}

	for _, file := range files {
		if file.Err != nil {
			backend.LogError("update", "parse update package '%s': %v\n", file.Filename, err)
			continue
		}

		// run the update file
		if err = file.Execute(backend.Config.DataFolder, config.PluginFolder); err != nil {
			file.Reader.Close()
			backend.LogError("update", "process package '%s' error: %v\n", file.Filename, err)
			continue
		}

		file.Reader.Close()

		// successful run, delete the update file
		if err = file.Delete(); err != nil {
			backend.LogError("update", "delete package '%s' error: %v\n", file.Filename, err)
			continue
		}

		backend.LogError("update", "success package '%s'\n", file.Filename)
	}
}
