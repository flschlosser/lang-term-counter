package wiki

import "sync"

type TermFrequency struct {
    sync.RWMutex
    internal map[string]int
}

func NewTermFrequency() *TermFrequency {
    return &TermFrequency{internal: make(map[string]int)}
}

func (tf* TermFrequency) Get(key string) int {
    tf.RLock()
    result := tf.internal[key]
    tf.RUnlock()
    return result
}

func (tf* TermFrequency) Inc(key string) {
    tf.Lock()
    tf.internal[key]++
    tf.Unlock()
}
