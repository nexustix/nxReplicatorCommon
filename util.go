package nxReplicatorCommon

import (
	"os"
	"path"
)

//InitWorkFolder creates a nested folder wihin a subfolder of "workingDir"
func InitWorkFolder(workingDir, subfolder, folderName string) string {

	//XXX could be combined into one variable
	tmpDir := path.Join(workingDir, subfolder, folderName)
	os.MkdirAll(tmpDir, 0711)
	return tmpDir
}
