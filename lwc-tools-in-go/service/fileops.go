package service

import (
	"errors"
	"log"
	"lwc-tools-in-go/model"
	"os"
	"path/filepath"
	"strings"
)

func ListAllFiles(rootDir string, skipFn func(path *string, info *os.FileInfo) bool) ([]string, error) {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if skipFn != nil && skipFn(&path, &info) {
			return nil
		}
		if !info.IsDir() {
			files = append(files, path[len(rootDir)+1:])
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func ParseLwcComps(lwcFiles []string) (map[string]*model.LwcComp, error) {
	lwcComps := make(map[string]*model.LwcComp)
	for _, path := range lwcFiles {
		lastSlash := strings.LastIndex(path, "/")
		if lastSlash == -1 {
			log.Println("Invalid path: ", path)
			continue
		}
		file := path[lastSlash+1:]
		compDir := path[:lastSlash]
		compName := compDir
		lastSlash = strings.LastIndex(compDir, "/")
		if lastSlash != -1 {
			compName = compDir[lastSlash+1:]
		}
		lwcComp, ok := lwcComps[compName]
		if !ok {
			lwcComp = &model.LwcComp{
				Name: compName,
				Dir:  compDir,
			}
			lwcComps[compName] = lwcComp
		}
		if strings.HasSuffix(file, ".html") {
			lwcComp.SetHtml(&model.LwcHtml{
				LwcFile: model.LwcFile{
					Name: file,
				},
			})
		} else if strings.HasSuffix(file, ".js") {
			lwcComp.SetJs(&model.LwcJs{
				LwcFile: model.LwcFile{
					Name: file,
				},
			})
		} else if strings.HasSuffix(file, ".css") {
			lwcComp.SetCss(&model.LwcCss{
				LwcFile: model.LwcFile{
					Name: file,
				},
			})
		} else if strings.HasSuffix(file, ".js-meta.xml") {
			lwcComp.SetMeta(&model.LwcMeta{
				LwcFile: model.LwcFile{
					Name: file,
				},
				Exposed: "unknown",
			})
		} else {
			return nil, errors.New("Invalid lwc file type: " + file)
		}
	}
	return lwcComps, nil
}
