package super_rate

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/throttled/throttled/v2"
	"github.com/throttled/throttled/v2/store/memstore"
)

type RateLimitConfig struct {
	Rate  int `json:"rate_per_min,omitempty"`
	Burst int `json:"burst,omitempty"`
}

// Create the plugin configuration.
type Config struct {
	Default RateLimitConfig `json:"default,omitempty"`
	Headers []string        `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// GraphQLRateLimiter holds the necessary components of a Traefik plugin
type RateLimiter struct {
	next   http.Handler
	name   string
	config Config
	//	rl     *throttled.GCRARateLimiter
}

// New instantiates and returns the required components used to handle a HTTP request
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	return &RateLimiter{
		next:   next,
		config: *config,
		name:   name,
		//	rl:     rateLimiter,
	}, nil
}

// Iterate over every headers to match the ones specified in the config and
// return nothing if regexp failed.
func (r *RateLimiter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	store, err := memstore.New(65536)
	if err != nil {
		fmt.Println("error creating store")
		fmt.Println(err)
		r.next.ServeHTTP(rw, req)
		//		return nil, fmt.Errorf("Error creating store %w", err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(10),
		MaxBurst: 1,
		//MaxBurst: config.Default.Burst,
	}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		fmt.Println("error creating newgcraratelimiter")
		fmt.Println(err)
		r.next.ServeHTTP(rw, req)
		//		return nil, fmt.Errorf("Error creating rl %w", err)
	}

	keys := []string{}
	for _, header := range r.config.Headers {
		val := req.Header.Get(header)
		if val == "" {
			val = "nil"
		}
		keys = append(keys, val)
	}
	uKey := strings.Join(keys, "::")
	fmt.Println("Header values are", uKey)

	//allowed, result, err := r.rl.RateLimit(uKey, 1)
	allowed, result, err := rateLimiter.RateLimit(uKey, 1)
	if err != nil {
		fmt.Println("Problem with rate limit")
		fmt.Println(err)
		rw.WriteHeader(http.StatusTeapot)
		return
	}

	fmt.Println("Rate limit allowed:", allowed)
	fmt.Println("Rate limit result limit:", result.Limit)
	fmt.Println("Rate limit result remaining:", result.Remaining)

	if hval := req.Header.Get("Super-Man"); hval == "yes" {
		fmt.Println("Rejecting request because of super-man")
		rw.WriteHeader(http.StatusTeapot)
		_, err := rw.Write([]byte(http.StatusText(http.StatusTeapot)))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	r.next.ServeHTTP(rw, req)
}
