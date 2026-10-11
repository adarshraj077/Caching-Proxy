package cache

import (
	"net/http"
	"time"
	
)


type CacheEntry struct{
	
	Body   []byte
	Response *http.Response
	Created time.Time
}
