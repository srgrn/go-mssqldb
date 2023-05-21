package ntlm

import (
	"github.com/srgrn/go-mssqldb/integratedauth"
)

// AuthProvider handles NTLM SSPI Windows Authentication
var AuthProvider integratedauth.Provider = integratedauth.ProviderFunc(getAuth)

func init() {
	err := integratedauth.SetIntegratedAuthenticationProvider("ntlm", AuthProvider)
	if err != nil {
		panic(err)
	}
}
