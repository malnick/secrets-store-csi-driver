// +build !no_vault_provider
package register

import (
	"github.com/deislabs/secrets-store-csi-driver/pkg/providers"
	vault "github.com/hashicorp/secrets-store-csi-driver-provider-vault"
)

func init() {
	register("vault", initVault)
}

func initVault(cfg InitConfig) (providers.Provider, error) {
	return vault.NewProvider()
}
