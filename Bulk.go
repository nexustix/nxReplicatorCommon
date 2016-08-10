package nxReplicatorCommon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	bp "github.com/nexustix/boilerplate"
)

type BulkItem struct {
	RelativePath string
	Download     bp.Download
}

//Bulk represents a collection of downloads
type Bulk struct {
	BulkItems []BulkItem
}

/*
//AddDownload adds download to bulk
func (b *Bulk) AddDownload(download bp.Download) {
	b.Downloads = append(b.Downloads, download)
}
*/

func (b *Bulk) AddDownload(bulkItem BulkItem) {
	for _, v := range b.BulkItems {
		if v.Download.Filename == bulkItem.Download.Filename {
			return
		}
	}
	b.BulkItems = append(b.BulkItems, bulkItem)
}

/*
//RemoveDownload removes download from bulk
func (b *Bulk) RemoveDownload(providerID, filename string) {
	var tmpDownloads []bp.Download
	for _, v := range b.Downloads {
		if v.Filename == filename {

		} else {
			tmpDownloads = append(tmpDownloads, v)
		}
	}
	b.Downloads = tmpDownloads
}
*/

//RemoveDownload removes download from bulk
func (b *Bulk) RemoveDownload(providerID, filename string) {
	var tmpItems []BulkItem
	for _, v := range b.BulkItems {
		if v.Download.Filename == filename {

		} else {
			tmpItems = append(tmpItems, v)
		}
	}
	b.BulkItems = tmpItems
}

//LoadFromFile loads Items into molecule from file
func (b *Bulk) LoadFromFile(filePath string) {
	if bp.FileExists(filePath) {
		dat, err := ioutil.ReadFile(filePath)
		bp.FailError(err)

		json.Unmarshal(dat, &b)

		return

	}
	fmt.Printf("<!> INFO Bulk >%s< not found\n", filePath)

}

//SaveToFile saves molecule to file
func (b *Bulk) SaveToFile(filePath string) {
	//if bp.FileExists(filePath) {
	outFile, err := os.Create(filePath)
	bp.FailError(err)
	tmpJSON, _ := json.Marshal(b)
	outFile.Write(tmpJSON)
	outFile.Close()
	//}
}
