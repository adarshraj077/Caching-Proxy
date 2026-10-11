package proxy

import (
	"fmt"
	"net/http"
	"github.com/adarshraj077/cachingProxy/internal/cache"
	"sync"
)

type proxyObject struct {
	Origin string
	Cache map[string]*cache.CacheEntry
    Mutex sync.RWMutex
}

func (p *proxyObject) ClearCache() {
	p.Mutex.Lock()
	p.Cache = make(map[string]*cache.CacheEntry)
	p.Mutex.Unlock()
	fmt.Println("Cache Cleared Successfully")
}



func  NewProxy(originURL string )*proxyObject{
    return &proxyObject{
		Origin: originURL,
        Cache: make(map[string]*cache.CacheEntry),
        
	}
   
}




func  ResponseWithHeader(w http.ResponseWriter ,response http.Response,CACHE_HEADER string ,body []byte,key string){
	fmt.Printf("cache %s %s \n",CACHE_HEADER,key)
    w.Header().Set("X-Cache", CACHE_HEADER)
	for k, v := range response.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(response.StatusCode)
	w.Write(body)
}

