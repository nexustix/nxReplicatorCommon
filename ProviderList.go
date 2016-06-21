package nxReplicatorCommon

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	bp "github.com/nexustix/boilerplate"
)

type ProviderEntry struct {
	ID   string
	Path string
}

type ProviderList struct {
	Dir       string
	Filename  string
	Providers []ProviderEntry
}

func (p *ProviderList) HasEntry(providerID string) bool {
	for _, v := range p.Providers {
		if v.ID == providerID {
			return true
		}
	}
	return false
}

func (p *ProviderList) GetEntry(providerID string) ProviderEntry {
	for _, v := range p.Providers {
		if v.ID == providerID {
			return v
		}
	}
	return ProviderEntry{}
}

func (p *ProviderList) AddEntry(providerID, providerPath string) {
	if !p.HasEntry(providerID) {
		p.Providers = append(p.Providers, ProviderEntry{ID: providerID, Path: providerPath})
	}

}

func (p *ProviderList) RemoveEntry(providerID string) {
	var tmpSlice []ProviderEntry
	for _, v := range p.Providers {
		if v.ID != providerID {
			tmpSlice = append(tmpSlice, v)
		}
	}
	p.Providers = tmpSlice
}

func (p *ProviderList) LoadEntries() {
	dat, err := ioutil.ReadFile(path.Join(p.Dir, p.Filename))
	//FIXME fails silently
	if !bp.GotError(err) {
		//bp.FailError(err)
		//Boilerplate.FailError(err)
		tmpProviderList := ProviderList{}

		json.Unmarshal(dat, &tmpProviderList)

		p.Providers = tmpProviderList.Providers
	}
}

func (p *ProviderList) SaveEntries() {
	filePath := path.Join(p.Dir, p.Filename)
	//if bp.FileExists(filePath) {
	outFile, err := os.Create(filePath)
	bp.FailError(err)
	tmpJSON, _ := json.Marshal(p)
	outFile.Write(tmpJSON)
	outFile.Close()
}
