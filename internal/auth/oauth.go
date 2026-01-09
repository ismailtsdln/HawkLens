package auth

import (
	"sync"
)

type Token struct {
	Platform string
	Value    string
}

type TokenPool struct {
	tokens map[string][]string
	index  map[string]int
	mu     sync.Mutex
}

func NewTokenPool() *TokenPool {
	return &TokenPool{
		tokens: make(map[string][]string),
		index:  make(map[string]int),
	}
}

func (p *TokenPool) AddToken(platform, token string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.tokens[platform] = append(p.tokens[platform], token)
}

func (p *TokenPool) GetNextToken(platform string) (string, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	tokens, ok := p.tokens[platform]
	if !ok || len(tokens) == 0 {
		return "", false
	}

	idx := p.index[platform]
	token := tokens[idx]
	p.index[platform] = (idx + 1) % len(tokens)

	return token, true
}
