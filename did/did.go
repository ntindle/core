package did

import (
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/mr-tron/base58"
	b "github.com/sonr-io/core/did/builder"
	p "github.com/sonr-io/core/did/parser"
)

// Parse parses a DID string into a DID struct
func Parse(s string) (*b.DID, error) {
	var did b.DID

	if !p.IsValidDid(b.Method, s) {
		return nil, b.ErrParseInvalid
	}
	// Parse Items from string into DID struct
	did.Method = b.Method

	return &did, nil
}

// ResolveId resolves a DID into a string
func ResolveId(did string, methodId string) string {
	result := methodId

	methodDid, methodFragment := p.SplitDidUrlIntoDidAndFragment(methodId)
	if len(methodDid) == 0 {
		result = did + "#" + methodFragment
	}

	return result
}

// CreateBaseDID creates a base DID with a given users libp2p public key
func CreateBaseDID(pubKey crypto.PubKey) (*b.DID, string, error) {
	// Marshal the public key into bytes
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, "", err
	}

	pubStr := base58.Encode(pubBuf)

	// Encode the public key into base58 and return the result
	did, err := b.NewDID(pubStr)
	if err != nil {
		return nil, "", err
	}
	return did, pubStr, nil
}

// CreateServiceDID creates a service DID with a developers libp2p public key
func CreateServiceDID(pubKey crypto.PubKey, name string) (*b.DID, string, error) {
	// Marshal the public key into bytes
	pubBuf, err := crypto.MarshalPublicKey(pubKey)
	if err != nil {
		return nil, "", err
	}

	pubStr := base58.Encode(pubBuf)

	// Encode the public key into base58 and return the result
	did, err := b.NewDID(pubStr, b.WithPath(name))
	if err != nil {
		return nil, "", err
	}
	return did, pubStr, nil
}
