package nxReplicatorCommon

//NxAtom represents "Atom"s (an analog for single downloadable files)
type NxAtom struct {
	Groups       []string //Groups for cataloging (hashtag-style since trees are evil)
	Provider     string   // Name/ID of provider
	Name         string   // Name of Atom
	ID           string   // ID of Atom
	Filename     string   // Filename to save Atom as
	URL          string   // URL use depens on Provider used (Download URL ?)
	RelativePath string   // Relative path to download location //TODO exploitable be aware of "../"
}
