package crypto

import (
	"sync"
)

var signingMethods = map[string]func() SigningMethod{}
var signingMethodLock = new(sync.RWMutex)

// SigningMethod can be used add new methods for signing or verifying tokens.
type SigningMethod interface {
	Verify(signingString, signature []byte, key interface{}) error   // Returns nil if signature is valid
	VerifyHash(signingHash, signature []byte, key interface{}) error // Returns nil if signature is valid
	Sign(signingString []byte, key interface{}) ([]byte, error)      // Returns encoded signature or error
	SignHash(signingHash []byte, key interface{}) ([]byte, error)    // Returns encoded signature or error
	Alg() string                                                     // returns the alg identifier for this method (example: 'HS256')
}

// RegisterSigningMethod registers the "alg" name and a factory function for signing method.
// This is typically done during init() in the method's implementation
func RegisterSigningMethod(alg string, f func() SigningMethod) {
	signingMethodLock.Lock()
	defer signingMethodLock.Unlock()

	signingMethods[alg] = f
}

// GetSigningMethod retrieves a signing method from an "alg" string
func GetSigningMethod(alg string) (method SigningMethod) {
	signingMethodLock.RLock()
	defer signingMethodLock.RUnlock()

	if methodF, ok := signingMethods[alg]; ok {
		method = methodF()
	}
	return
}

// GetAlgorithms returns a list of registered "alg" names
func GetAlgorithms() (algs []string) {
	signingMethodLock.RLock()
	defer signingMethodLock.RUnlock()

	for alg := range signingMethods {
		algs = append(algs, alg)
	}
	return
}
