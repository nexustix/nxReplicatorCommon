package nxReplicatorCommon

//NxProvider represents an interface for providing "Atoms"
type NxProvider interface {
	Search(searchTerm, modVersion, gameVersion string) []NxAtom
	GetDependencies(modName, modID string) []NxAtom
}
