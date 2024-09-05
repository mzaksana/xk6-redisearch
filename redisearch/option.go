package redisearch

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/redis/rueidis"
)

type socketOptions struct {
	Host               string      `json:"host,omitempty"`
	Port               int         `json:"port,omitempty"`
	TLS                *tlsOptions `json:"tls,omitempty"`
	DialTimeout        int64       `json:"dialTimeout,omitempty"`
	ReadTimeout        int64       `json:"readTimeout,omitempty"`
	WriteTimeout       int64       `json:"writeTimeout,omitempty"`
	PoolSize           int         `json:"poolSize,omitempty"`
	MinIdleConns       int         `json:"minIdleConns,omitempty"`
	MaxConnAge         int64       `json:"maxConnAge,omitempty"`
	PoolTimeout        int64       `json:"poolTimeout,omitempty"`
	IdleTimeout        int64       `json:"idleTimeout,omitempty"`
	IdleCheckFrequency int64       `json:"idleCheckFrequency,omitempty"`
}

type tlsOptions struct {
	// TODO: Handle binary data (ArrayBuffer) for all these as well.
	CA   []string `json:"ca,omitempty"`
	Cert string   `json:"cert,omitempty"`
	Key  string   `json:"key,omitempty"`
}

type singleNodeOptions struct {
	Socket          *socketOptions `json:"socket,omitempty"`
	Username        string         `json:"username,omitempty"`
	Password        string         `json:"password,omitempty"`
	ClientName      string         `json:"clientName,omitempty"`
	Database        int            `json:"database,omitempty"`
	MaxRetries      int            `json:"maxRetries,omitempty"`
	MinRetryBackoff int64          `json:"minRetryBackoff,omitempty"`
	MaxRetryBackoff int64          `json:"maxRetryBackoff,omitempty"`
}


// readOptions parses the input options (either a string or object) and converts them to rueidis.ClientOption.
func readOptions(input interface{}) (*rueidis.ClientOption, error) {
	switch val := input.(type) {
	case string:
		// If the input is a string, treat it as a Redis URL
		return parseRedisURL(val)
	case map[string]interface{}:
		// If it's an object, assume it's a map of connection options
		return parseOptionsFromMap(val)
	default:
		return nil, fmt.Errorf("invalid options type: %T; expected string or object", input)
	}
}

// parseRedisURL parses a Redis URL into rueidis.ClientOption
func parseRedisURL(url string) (*rueidis.ClientOption, error) {
	// Split the URL into components
	// Example URL: redis://username:password@localhost:6379/0
	parts := strings.Split(url, "@")
	if len(parts) != 2 {
		return nil, errors.New("invalid Redis URL format")
	}

	// Extract username and password from the first part
	authParts := strings.Split(parts[0], ":")
	if len(authParts) != 2 {
		return nil, errors.New("invalid auth credentials in Redis URL")
	}
	username := authParts[0]
	password := authParts[1]

	// Extract host and port from the second part
	host := parts[1]

	return &rueidis.ClientOption{
		InitAddress:  []string{host},
		Username:     username,
		Password:     password,
		DisableCache: true,
	}, nil
}

// parseOptionsFromMap parses a map of options into rueidis.ClientOption
func parseOptionsFromMap(options map[string]interface{}) (*rueidis.ClientOption, error) {
	var sopts singleNodeOptions

	// Convert the map to a struct (this assumes the options follow the correct schema)
	optionsJSON, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize options: %w", err)
	}

	err = json.Unmarshal(optionsJSON, &sopts)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize options: %w", err)
	}

	// Convert singleNodeOptions to rueidis.ClientOption
	clientOption, err := sopts.toRueidisOptions()
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}

	return &clientOption, nil
}


// toRueidisOptions converts singleNodeOptions into rueidis.ClientOption
func (opts singleNodeOptions) toRueidisOptions() (rueidis.ClientOption, error) {
	if opts.Socket == nil {
		return rueidis.ClientOption{}, fmt.Errorf("socket options are required")
	}

	clientOption := rueidis.ClientOption{
		InitAddress:  []string{fmt.Sprintf("%s:%d", opts.Socket.Host, opts.Socket.Port)},
		Username:     opts.Username,
		Password:     opts.Password,
		ClientName:   opts.ClientName,
		DisableCache: true, // Default to disable caching unless specified
	}

	// Set TLS configuration if provided
	if opts.Socket.TLS != nil {
		tlsCfg := &tls.Config{}
		if len(opts.Socket.TLS.CA) > 0 {
			caCertPool := x509.NewCertPool()
			for _, cert := range opts.Socket.TLS.CA {
				caCertPool.AppendCertsFromPEM([]byte(cert))
			}
			tlsCfg.RootCAs = caCertPool
		}

		if opts.Socket.TLS.Cert != "" && opts.Socket.TLS.Key != "" {
			clientCertPair, err := tls.X509KeyPair([]byte(opts.Socket.TLS.Cert), []byte(opts.Socket.TLS.Key))
			if err != nil {
				return rueidis.ClientOption{}, fmt.Errorf("failed to load TLS certificates: %w", err)
			}
			tlsCfg.Certificates = []tls.Certificate{clientCertPair}
		}

		clientOption.TLSConfig = tlsCfg
	}

	// Handle timeouts (if they exist in the socket options)
	// if opts.Socket.DialTimeout > 0 {
	// 	clientOption.DialTimeout = time.Duration(opts.Socket.DialTimeout) * time.Millisecond
	// }
	// if opts.Socket.ReadTimeout > 0 {
	// 	clientOption.ReadTimeout = time.Duration(opts.Socket.ReadTimeout) * time.Millisecond
	// }
	// if opts.Socket.WriteTimeout > 0 {
	// 	clientOption.WriteTimeout = time.Duration(opts.Socket.WriteTimeout) * time.Millisecond
	// }

	return clientOption, nil
}
