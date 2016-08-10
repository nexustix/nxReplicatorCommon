package nxReplicatorCommon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	bp "github.com/nexustix/boilerplate"
)

//MoleculeItem is an item of a molecule representig a single atom
type MoleculeItem struct {
	Explicit   bool // if an atom was added explicitly (or implicitly as a dependency) (for re-fetching or purging dependencies)
	ProviderID string
	AtomID     string
	Dir        string
	//Filename string // only needed for Bulks
	//URL      string // only needed for bulks
}

//Molecule represents a colletion of "Atom"s only containing "Atom" IDs
type Molecule struct {
	MoleculeItems []MoleculeItem
}

//AddItem adds item to molecule
func (m *Molecule) AddItem(moleculeItem MoleculeItem) {
	for _, v := range m.MoleculeItems {
		if v.AtomID == moleculeItem.AtomID {
			return
		}
	}
	m.MoleculeItems = append(m.MoleculeItems, moleculeItem)
}

//RemoveItem removes item from molecule
func (m *Molecule) RemoveItem(providerID, atomID string) {
	var tmpItems []MoleculeItem
	for _, v := range m.MoleculeItems {
		if v.AtomID == atomID && v.ProviderID == providerID {

		} else {
			tmpItems = append(tmpItems, v)
		}
	}
	m.MoleculeItems = tmpItems
}

//AddAtom adds atom to molecule
//func (m *Molecule) AddAtom(atom Atom) {

//}

//LoadFromFile loads Items into molecule from file
func (m *Molecule) LoadFromFile(filePath string) {
	/*
		file, err := os.Open(filepath)
		bp.FailError(err)
		reader := bufio.NewReader(file)
		line := reader.ReadString('\n')
		segments := strings.Split(line, "|||")

		filePath := path.Join(am.WorkingDir, provider+"_"+atom.ID)
	*/
	//filePath := path.Join(am.WorkingDir, provider+"_"+atomID)

	if bp.FileExists(filePath) {
		dat, err := ioutil.ReadFile(filePath)
		bp.FailError(err)

		//tmpAtom := Atom{}

		json.Unmarshal(dat, &m)

		//return tmpAtom
		return

	}
	fmt.Printf("<!> INFO molecule >%s< not found\n", filePath)
	//return Atom{}

}

//SaveToFile saves molecule to file
func (m *Molecule) SaveToFile(filePath string) {
	//if bp.FileExists(filePath) {
	outFile, err := os.Create(filePath)
	bp.FailError(err)
	tmpJSON, _ := json.Marshal(m)
	outFile.Write(tmpJSON)
	outFile.Close()
	//}
}
