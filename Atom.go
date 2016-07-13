package nxReplicatorCommon

import (
	"fmt"
	"net/url"
	"strings"

	bp "github.com/nexustix/boilerplate"
)

//NxAtom represents an "Atom" (an analog for a single downloadable file)
type Atom struct {
	Provider     string   // Name/ID of provider
	Name         string   // Name of Atom
	ID           string   // ID of Atom
	Filename     string   // Filename to save Atom as
	URL          string   // URL use depens on Provider used (Download URL ?)
	RelativePath string   // Relative path to download location //TODO exploitable be aware of "../"
	Groups       []string // Groups for cataloging (hashtag-style since trees are evil)

	Priority     int      // TBD
	DoDepCheck   bool     // do a check for dependencies
	Dependencies []string // Atom dependencies
}

// URL encoded string seperated by spaces followed by a Pipe and all Grops seperated by spaces
// <Provider> <Name> <ID> <Filename> <URL> <RelativePath>|<Group1> <Group2> <Group3>...

func StringToAtom(stringAtom string, depCheck bool) Atom {

	var decodedAtomData []string

	mainSegments := strings.Split(stringAtom, "|")
	atomData := strings.Split(bp.StringAtIndex(0, mainSegments), " ")
	//FIXME may cause error
	atomGroups := strings.Split(bp.StringAtIndex(1, mainSegments), " ")

	/*
		for _, v := range atomData {
			fmt.Printf(">%s<\n", v)
		}
		fmt.Printf("---\n")
	*/

	for _, v := range atomData {
		data, err := url.QueryUnescape(v)
		bp.FailError(err)

		decodedAtomData = append(decodedAtomData, data)
	}

	for _, v := range decodedAtomData {
		fmt.Printf(">%s<\n", v)
	}
	fmt.Printf("---\n")

	return Atom{
		Provider:     bp.StringAtIndex(0, decodedAtomData),
		Name:         bp.StringAtIndex(1, decodedAtomData),
		ID:           bp.StringAtIndex(2, decodedAtomData),
		Filename:     bp.StringAtIndex(3, decodedAtomData),
		URL:          bp.StringAtIndex(4, decodedAtomData),
		RelativePath: bp.StringAtIndex(5, decodedAtomData),
		Groups:       atomGroups,
		DoDepCheck:   depCheck}

}

func OutputToAtoms(stringAtoms string, depCheck bool) []Atom {
	lines := strings.Split(stringAtoms, "\n")
	var tmpAtoms []Atom
	for _, v := range lines {
		if !strings.HasPrefix(v, "<") && !strings.HasPrefix(v, "#") && v != "" {
			fmt.Printf("==>%s<\n", v)
			tmpAtom := StringToAtom(v, depCheck)
			if tmpAtom.ID != "" {
				//tmpAtoms = append(tmpAtoms, StringToAtom(v))
				tmpAtoms = append(tmpAtoms, tmpAtom)
			}
		}
	}
	for _, v := range tmpAtoms {
		fmt.Printf(">>%s|\n", v.ID)
	}

	return tmpAtoms
}

func OutputToAtomsAndAdd(provider, stringAtoms string, manager *AtomManager, depCheck bool) {
	tmpAtoms := OutputToAtoms(stringAtoms, depCheck)
	for _, v := range tmpAtoms {
		if !manager.HasEntry(provider, v.ID) {
			manager.SetEntry(provider, v)
		}

		//AddAtom(provider, v, manager)
	}
}

/*
func AddAtom(provider string, atom Atom, manager *AtomManager) {
	manager.SetEntry(provider, atom)
}
*/
