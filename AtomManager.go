package nxReplicatorCommon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	bp "github.com/nexustix/boilerplate"
)

//Entry => a file
//Atom => thing with a specific ID

//AtomManager represents a way to acess and manage existing atoms
type AtomManager struct {
	Index      []string // index of Atom IDs for easy searching
	WorkingDir string   //
}

/*
func (am *AtomManager) HasAtom(atomID string) bool {
	filePath := path.Join(pdb.WorkingDir, fileName+pdb.FileEnding)
	return bp.FileExists(filePath)
	//return false
}
*/

//HasEntry checks if an atom of a given ID exists
func (am *AtomManager) HasEntry(provider, atomID string) bool {
	filePath := path.Join(am.WorkingDir, provider+"_"+atomID+".nxra")
	return bp.FileExists(filePath)
	//return false
}

/*
func (am *AtomManager) GetAtom(atomID string) NxAtom {
	return NxAtom{}
}
*/

//GetEntry returns atom of a given ID (or empty atom if none found)
func (am *AtomManager) GetEntry(provider, atomID string) Atom {
	filePath := path.Join(am.WorkingDir, provider+"_"+atomID+".nxra")

	if bp.FileExists(filePath) {
		dat, err := ioutil.ReadFile(filePath)
		bp.FailError(err)

		tmpAtom := Atom{}

		json.Unmarshal(dat, &tmpAtom)

		return tmpAtom

	}
	fmt.Printf("<!> ERROR entry >%s< not known\n", provider+"_"+atomID+".nxra")
	return Atom{}
}

/*
func (am *AtomManager) AddAtom(atom NxAtom) {
	if !am.HasAtom(atom.ID) {
		//create new Atom
	} else {
		//edit existing atom
	}
}
*/

//SetEntry sets Atom for a given ID
func (am *AtomManager) SetEntry(provider string, atom Atom) {
	filePath := path.Join(am.WorkingDir, provider+"_"+atom.ID+".nxra")

	//if bp.FileExists(filePath) {
	outFile, err := os.Create(filePath)
	bp.FailError(err)
	tmpJSON, _ := json.Marshal(atom)
	outFile.Write(tmpJSON)
	outFile.Close()
}

/*
func (am *AtomManager) RemoveAtom(atomID string) {
	if am.HasAtom(atomID) {
		//remove atom
	} else {
		//nothing to do
	}
}
*/

/*
func (am *AtomManager) RemoveEntry(provider, atomID string) {
	if am.HasEntry(provider, atomID) {
		//remove
	} else {
		//nothing to do
	}
}
*/

/*
func (am *AtomManager) LoadAtoms() {

}

func (am *AtomManager) SaveAtoms() {

}
*/
