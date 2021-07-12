package hub

// AddKey implements KeyManager
func (ic IritaHubChain) AddKey(name string, passphrase string) (addr string, mnemonic string, err error) {
	return ic.IritaClient.Insert(name, passphrase)
}

// DeleteKey implements KeyManager
func (ic IritaHubChain) DeleteKey(name string, passphrase string) error {
	return ic.IritaClient.Delete(name, passphrase)
}

// ShowKey implements KeyManager
func (ic IritaHubChain) ShowKey(name string, passphrase string) (addr string, err error) {
	_, address, err := ic.IritaClient.Find(name, passphrase)
	return address.String(), err
}

// ImportKey implements KeyManager
func (ic IritaHubChain) ImportKey(name string, passphrase string, keyArmor string) (addr string, err error) {
	return ic.IritaClient.Import(name, passphrase, keyArmor)
}

// ExportKey implements KeyManager
func (ic IritaHubChain) ExportKey(name string, passphrase string) (keyArmor string, err error) {
	return ic.IritaClient.Export(name, passphrase)
}

// RecoverKey implements KeyManager
func (ic IritaHubChain) RecoverKey(name string, passphrase string, mnemonic string) (addr string, err error) {
	return ic.IritaClient.Recover(name, passphrase, mnemonic)
}
