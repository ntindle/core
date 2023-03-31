// Utility functions for DID Service - https://w3c.github.io/did-core/#services
// I.e. Service Endpoints for IPFS Cluster
package types

import (
	"encoding/base64"
	"encoding/json"
	fmt "fmt"
	"strings"

	"github.com/go-webauthn/webauthn/protocol"
	types "github.com/sonr-io/sonr/types/crypto"
)

const (
	VaultServiceType            = "EncryptedVault"
	LinkedDomainServiceType     = "LinkedDomains"
	DIDCommMessagingServiceType = "DIDCommMessaging"
)

// NewIPFSStoreService creates a new IPFS Store Service record for the given address.
// Addresses look like: /orbitdb/bafyreiepksmvjzvcbzsdqkf474hgfoqf3xj5t47olga5qnnhxggxssbcya/testKVStore
// The address is split into the CID and the DBName, and the CID is used to create the DID. Which results in:
// did:orbitdb:bafyreiepksmvjzvcbzsdqkf474hgfoqf3xj5t47olga5qnnhxggxssbcya
// The origin is the dbname and the type is "EncryptedVault"
func NewIPFSStoreService(address string, controllerDid string) *Service {
	parts := strings.Split(address, "/")
	if len(parts) < 4 {
		return nil
	}
	host := parts[1]
	cid := parts[2]
	dbName := parts[3]
	id := fmt.Sprintf("did:%s:%s", host, cid)
	return &Service{
		Id:         id,
		Type:       VaultServiceType,
		Origin:     dbName,
		Controller: controllerDid,
	}
}

func (d *DidDocument) AddService(s *Service) {
	d.Service = append(d.Service, s)
}

func (d *DidDocument) RemoveServiceByID(id string) bool {
	for i, s := range d.Service {
		if s.Id == id {
			d.Service = append(d.Service[:i], d.Service[i+1:]...)
			return true
		}
	}
	return false
}

func (d *DidDocument) GetVaultService() *Service {
	for _, s := range d.Service {
		if s.Type == "EncryptedVault" && s.CID() != "" {
			return s
		}
	}
	return nil
}

func (s *Service) CredentialEntity() protocol.CredentialEntity {
	return protocol.CredentialEntity{
		Name: s.Name,
	}
}

func (s *Service) GetUserEntity(id string) protocol.UserEntity {
	return protocol.UserEntity{
		ID:               []byte(id),
		DisplayName:      id,
		CredentialEntity: s.CredentialEntity(),
	}
}

// GetCredentialCreationOptions issues a challenge for the VerificationMethod to sign and return
func (vm *Service) GetCredentialCreationOptions(username string) (string, error) {
	hashString := base64.URLEncoding.EncodeToString([]byte(vm.Id))
	params := DefaultParams()
	chal :=  protocol.URLEncodedBase64(hashString)

	cco, err := params.NewWebauthnCreationOptions(vm, username, chal)
	if err != nil {
		return "", err
	}

	ccoJSON, err := json.Marshal(cco)
	if err != nil {
		return "", err
	}
	return string(ccoJSON), nil
}

// RelyingPartyEntity is a struct that represents a Relying Party entity.
func (s *Service) RelyingPartyEntity() protocol.RelyingPartyEntity {
	return protocol.RelyingPartyEntity{
		ID: s.Id,
		CredentialEntity: protocol.CredentialEntity{
			Name: s.Name,
		},
	}
}

// VerifyCreationChallenge verifies the challenge and a creation signature and returns an error if it fails to verify
func (vm *Service) VerifyCreationChallenge(resp string) (*types.WebauthnCredential, error) {
	pcc, err := parseCreationData(resp)
	if err != nil {
		return nil, err
	}
	return makeCredentialFromCreationData(pcc), nil
}

// VeriifyAssertionChallenge verifies the challenge and an assertion signature and returns an error if it fails to verify
func (vm *Service) VeriifyAssertionChallenge(resp string, cred *types.WebauthnCredential) error {
	// pca, err := parseAssertionData(resp)
	// if err != nil {
	// 	return err
	// }
	return nil
}

// ResolveEndpointURL finds the endpoint with the given type and unmarshalls it as single URL.
// It returns the endpoint ID and URL, or an error if anything went wrong;
// - holder document can't be resolved,
// - service with given type doesn't exist,
// - multiple services match,
// - serviceEndpoint isn't a string.
func (d *DidDocument) ResolveEndpointURL(serviceType string) (endpointID string, endpointURL string, err error) {
	var services []*Service
	for _, service := range d.Service {
		if service.Type == serviceType {
			services = append(services, service)
		}
	}
	if len(services) == 0 {
		return "", "", fmt.Errorf("service not found (did=%s, type=%s)", d.Id, serviceType)
	}
	if len(services) > 1 {
		return "", "", fmt.Errorf("multiple services found (did=%s, type=%s)", d.Id, serviceType)
	}
	err = services[0].UnmarshalServiceEndpoint(&endpointURL)
	if err != nil {
		return "", "", fmt.Errorf("unable to unmarshal single URL from service (id=%s): %w", services[0].Id, err)
	}
	return services[0].Id, endpointURL, nil
}

// Unmarshal unmarshalls the service endpoint into a domain-specific type.
func (s Service) UnmarshalServiceEndpoint(target interface{}) error {
	if asJSON, err := json.Marshal(s.Origin); err != nil {
		return err
	} else {
		return json.Unmarshal(asJSON, target)
	}
}

func (s *Service) CID() string {
	if strings.Contains(s.Origin, "ipfs") {
		return strings.TrimPrefix(s.Origin, "https://ipfs.sonr.network")
	}
	return ""
}