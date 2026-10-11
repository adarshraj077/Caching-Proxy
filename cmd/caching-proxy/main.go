package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
      "io"
	"github.com/adarshraj077/cachingProxy/internal/proxy"
	"github.com/adarshraj077/cachingProxy/internal/cache"
)


func  main (){
	
	port := flag.String("port","8080","port on which it runs")
	origin := flag.String("origin","","target origin url")
	clear_cache := flag.Bool("clear_cache",false,"used to clear cache")
	
	flag.Parse()
	if *origin=="" {
		fmt.Println("requires a target url")
		return 
	}
	Proxy := proxy.NewProxy(*origin)
	
	if *clear_cache {
		Proxy.ClearCache()
	}

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		CACHE_KEY := r.Method+":"+r.URL.String()
		originURL :=*origin+r.URL.String()


		Proxy.Mutex.RLock()

	    c,ok := Proxy.Cache[CACHE_KEY];

		Proxy.Mutex.RLocker().Unlock()

		
		if ok {
			proxy.ResponseWithHeader(w,*c.Response,"HIT",c.Body,CACHE_KEY)
			return
		}
		
		fmt.Printf("Cache Not Present for key : %s \n", CACHE_KEY)
		

	   fmt.Println(originURL)


Proxy.Mutex.RLock()
	   resp,err :=http.Get(originURL)
	   Proxy.Mutex.RLocker().Unlock()

	   defer resp.Body.Close()

	   if err != nil {
		http.Error(w,"error forwarding error",http.StatusInternalServerError)
		return 
	   }

	   body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error Forwarding Request Body", http.StatusInternalServerError)
		return
	}

	  Proxy.Cache[CACHE_KEY]= &cache.CacheEntry{
          Body: body,
		  Response: resp,
		  Created: time.Now(),
	   }

	   proxy.ResponseWithHeader(w,*resp,"MISS",body,CACHE_KEY)

	})


	fmt.Println(*port)
	fmt.Println(*origin)

	err := http.ListenAndServe(":"+*port,nil)
	if err !=nil {
		panic(err)
	}

}
