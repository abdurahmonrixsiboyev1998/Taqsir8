// package main

// import (
// 	"fmt"
// 	"sync"
// )

// type SafeMap struct {
// 	num  sync.Mutex
// 	data map[string]int
// }

// func NewSafeMap() *SafeMap {
// 	return &SafeMap{
// 		data: make(map[string]int),
// 	}
// }

// func (sm *SafeMap) Get(key string) int {
// 	sm.num.Lock()
// 	defer sm.num.Unlock()
// 	return sm.data[key]
// }

// func (sm *SafeMap) Set(key string, value int) {
// 	sm.num.Lock()
// 	defer sm.num.Unlock()
// 	sm.data[key] = value
// }

// func (sm *SafeMap) Delete(key string) {
// 	sm.num.Lock()
// 	defer sm.num.Unlock()
// 	delete(sm.data, key)
// }

// func main() {
// 	newMap := NewSafeMap()

// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		newMap.Set("key1", 1)
// 		newMap.Set("key2", 2)
// 		newMap.Set("key3", 3)
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Value of key1:", newMap.Get("key1"))
// 		fmt.Println("Value of key2:", newMap.Get("key2"))
// 		fmt.Println("Value of key3:", newMap.Get("key3"))
// 	}()

// 	wg.Wait()

// 	newMap.Delete("key2")
// 	fmt.Println("Value of key2 after deletion:", newMap.Get("key2"))
// }



package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Map struct {
	sync.Mutex
	gorot map[int]int
}

func (n *Map) read(key int) (int, bool) {
	n.Lock()
	defer n.Unlock()
	val, ok := n.gorot[key]
	return val, ok
}

func (n *Map) write(key, val int) {
	n.Lock()
	defer n.Unlock()
	n.gorot[key] = val
}

func (n *Map) delete(key int) {
	n.Lock()
	defer n.Unlock()
	delete(n.gorot, key)
}

func main() {
	var wg sync.WaitGroup
	n := &Map{gorot: make(map[int]int)}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			n.write(i, rand.Intn(100))
		}(i)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			val, ok := n.read(i)
			if ok {
				fmt.Printf("Key: %d, Value: %d\n", i, val)
			} else {
				fmt.Printf("Key: %d, Value: mavjud emas\n", i)
			}
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			n.delete(i)
		}(i)
	}

	wg.Wait()
}
