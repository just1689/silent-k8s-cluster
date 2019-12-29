package disk

import "os"

func DeleteDir(p string) {
	os.RemoveAll(p)

}

func CreateDir(p string) {
	os.Mkdir(p, os.ModeDir)

}
