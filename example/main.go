package main

import (
	"cache"
	"log"
	"time"
)

func main() {
	cache := cache.NewCache(100 * time.Microsecond)
	cache.Set("test", "hello", 1*time.Second)
	value, err := cache.Get("test")
	log.Print(value, err)
	time.Sleep(3 * time.Second)
	value, err = cache.Get("test")
	log.Print(value, err)
	cache.Set("delete", "delete", 3*time.Second)
	valueDelete, err := cache.Get("delete")
	log.Print(valueDelete, err)
	cache.Delete("test")
	valueDelete, err = cache.Get("test")
	log.Print(valueDelete, err)
	cache.Set("updateA", "A", 3*time.Second)
	valueUpdate, err := cache.Get("updateA")
	log.Print(valueUpdate, err)
	cache.Set("updateA", "B", 3*time.Second)
	valueUpdate, err = cache.Get("updateA")
	log.Print(valueUpdate, err)

}
