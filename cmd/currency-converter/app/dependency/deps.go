package dependency

type Dependencies struct {
	gateways    *gateways
	useCases    *useCases
	Entrypoints *Entrypoints
}

func ResolveDependencies() *Dependencies {
	d := &Dependencies{}

	return d.
		setupGateways().
		setupUseCases().
		setupEntrypoints()
}
