package createta

import (
	"io/fs"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// TAName holds the name in different formats
type TAName struct {
	Uppercase    string
	Lowercase    string
	Name         string
	NewUppercase string
	NewLowercase string
	NewName      string
}

func Create(name, newName string) *TAName {
	return &TAName{
		Name:         name,
		Uppercase:    strings.ToUpper(name),
		Lowercase:    strings.ToLower(name),
		NewName:      newName,
		NewUppercase: strings.ToUpper(newName),
		NewLowercase: strings.ToLower(newName),
	}
}

// replaceName replace the name of all the variables with the ta-name in it with the new ta-name.
// All uppercase names are replaced with the upper case representative names. All lowercase names
// respectively. Implements the [fs.WalkDirFunc] interface.
func (n *TAName) replaceName(path string, dirEntry fs.DirEntry, err error) error {
	if err != nil {
		log.Warnf("Error reading directory %s: %v", path, err)
		return err
	}
	// replace the name of all the variables with the ta-name in it with the new ta-name
	// upper case
	if !dirEntry.Type().IsRegular() {
		log.Infof("Skipping directory %s", path)
		return nil
	}
	// warn if file is very big
	info, err := dirEntry.Info()
	if err != nil {
		log.Warnf("Error reading file info %s: %v", path, err)
		return err
	}
	if info.Size() > 1000000 {
		log.Warnf("File %s is very big: %d", path, info.Size())
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Warnf("Error reading file %s: %v", path, err)
		return err
	}

	sdata := string(data)

	sdata = strings.Replace(sdata, n.Uppercase, n.NewUppercase, -1)
	sdata = strings.Replace(sdata, n.Lowercase, n.NewLowercase, -1)
	if strings.Index(sdata, "UUID") != -1 {
		log.Infof("The UUID is in file %s.", path)
	}

	// write back to the file
	err = os.WriteFile(path, []byte(sdata), 0644)
	if err != nil {
		log.Warnf("Error writing file %s: %v", path, err)
		return err
	}

	return err
}

func (n *TAName) GenerateNew(path string) {
	// walk the directory and replace the name
	err := fs.WalkDir(os.DirFS(path), path, n.replaceName)
	if err != nil {
		log.Warnf("Error walking directory %s: %v", path, err)
	}
}
