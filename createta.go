package createta

import (
    "strings"
)

func CreateTA(name string) {
    name = cleanup(name)
    upperName := strings.ToUpper(name)
    lowerName := strings.ToUpper(name)
}

func cleanup(name string) {
    // replace all '-',' ' to _ because of C naming convention
    name = strings.Replace(name, "-", "_", -1)
    name = strings.Replace(name, " ", "_", -1)
}
