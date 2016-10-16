package playground

import "sync"

type KeyValue struct {
	store map[string]string
	mu    sync.RWMutex //排他制御用のMutex
}

func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)}
}

func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()
	defer kv.mu.Unlock() //メソッドを抜けるときは必ずアンロック
	kv.store[key] = val
}

func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.store[key]
	return val, ok
}
