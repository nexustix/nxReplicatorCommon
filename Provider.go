package nxReplicatorCommon

//NxProvider represents an interface for providing "Atoms"
type Provider interface {
	Search(searchTerm, modVersion, gameVersion string) []Atom
	GetDependencies(modName, modID string) []Atom
}
