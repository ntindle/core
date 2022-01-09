package builder

import (
	"strings"

	"github.com/sonr-io/core/did/parser"
)

type Network int

const (
	Method = "did:sonr"
)

type Option func(*DID)

// WithFragment adds a fragment to a DID
func WithFragment(fragment string) Option {
	return func(d *DID) {
		fragment := strings.SplitAfter(fragment, "#")
		d.Fragment = fragment[1]
	}
}

// WithNetwork adds a network to a DID
func WithNetwork(network string) Option {
	return func(d *DID) {
		if ok := parser.IsValidNetworkPrefix(network); ok {
			if network == "mainnet" {
				network = ""
			}
			d.Network = network
		} else {
			d.Network = "testnet"
		}
	}
}

// WithPath adds a path to a DID
func WithPath(p string) Option {
	return func(d *DID) {
		rawpaths := strings.SplitAfter(p, "/")
		paths := strings.Split(strings.Join(rawpaths[1:], ""), "/")
		d.Paths = paths
	}
}

// WithQuery adds a query to a DID
func WithQuery(query string) Option {
	return func(d *DID) {
		query := strings.SplitAfter(query, "?")
		d.Query = query[1]
	}
}

type DID struct {
	Method   string
	Network  string
	Id       string
	Paths    []string
	Query    string
	Fragment string
}

// NewDID creates a new DID
func NewDID(id string, options ...Option) (*DID, error) {
	var did DID
	did.Id = id
	did.Method = Method

	// Apply options
	for _, option := range options {
		option(&did)
	}

	// Check if the DID is valid
	if !did.IsValid() {
		return nil, ErrFragmentAndQuery
	}
	return &did, nil
}


// AddPath adds a path to the DID struct and returns the new DID
func (d *DID) AddPath(path string) *DID {
	// Check if the path contains a leading slash
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	// Split and join the path
	paths := strings.Split(path, "/")
	d.Paths = append(d.Paths, paths...)
	return d
}

// AddFragment adds a fragment to the DID struct and returns the new DID
func (d *DID) AddFragment(fragment string) *DID {
	// Check if the fragment contains a leading hash
	if !strings.HasPrefix(fragment, "#") {
		fragment = "#" + fragment
	}

	d.Fragment = fragment
	return d
}

// AddQuery adds a query to the DID struct and returns the new DID
func (d *DID) AddQuery(query string) *DID {
	// Check if the query contains a leading question mark
	if !strings.HasPrefix(query, "?") {
		query = "?" + query
	}

	d.Query = query
	return d
}

// GetBase returns the base DID string: Method + Network
func (d *DID) GetBase() string {
	if d.HasNetwork() {
		return d.Method + ":" + d.Network + ":"
	}
	return d.Method + ":"
}

// GetPath returns the path part of the DID
func (d *DID) GetPath() string {
	if d.HasPath() {
		return "/" + strings.Join(d.Paths, "/")
	}
	return ""
}

// GetQuery returns the query part of the DID
func (d *DID) GetQuery() string {
	if d.HasQuery() {
		return "?" + d.Query
	}
	return ""
}

// GetFragment returns the fragment part of the DID
func (d *DID) GetFragment() string {
	if d.HasFragment() {
		return "#" + d.Fragment
	}
	return ""
}

// HasNetwork returns true if the DID has a network
func (d *DID) HasNetwork() bool {
	return len(d.Network) > 0
}

// HasPath returns true if the DID has a path
func (d *DID) HasPath() bool {
	return len(d.Paths) > 0
}

// HasQuery returns true if the DID has a query
func (d *DID) HasQuery() bool {
	return len(d.Query) > 0
}

// HasFragment returns true if the DID has a fragment
func (d *DID) HasFragment() bool {
	return len(d.Fragment) > 0
}

// IsValid checks if a DID is valid and does not contain both a Fragment and a Query
func (d *DID) IsValid() bool {
	hq := d.HasQuery()
	hf := d.HasFragment()

	if hq && hf {
		return false
	}
	return true
}

// String combines all DID parts into a string
func (d *DID) String() string {
	return d.GetBase() + d.Id + d.GetPath() + d.GetQuery() + d.GetFragment() + d.GetFragment()
}
