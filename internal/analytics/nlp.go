package analytics

import (
	"strings"
)

type AnalysisResult struct {
	Sentiment string   `json:"sentiment"`
	Score     float64  `json:"score"`
	Topics    []string `json:"topics"`
	Clusters  []string `json:"clusters"`
}

// AnalyzeText performs basic sentiment and topic analysis
func AnalyzeText(text string) AnalysisResult {
	text = strings.ToLower(text)

	// Basic sentiment mock
	sentiment := "neutral"
	score := 0.0

	positiveWords := []string{"good", "great", "excellent", "powerful", "amazing", "success", "resolved"}
	negativeWords := []string{"bad", "error", "fail", "poor", "issue", "vulnerab", "breach", "attack"}

	posCount := 0
	for _, word := range positiveWords {
		if strings.Contains(text, word) {
			posCount++
		}
	}

	negCount := 0
	for _, word := range negativeWords {
		if strings.Contains(text, word) {
			negCount++
		}
	}

	if posCount > negCount {
		sentiment = "positive"
		score = 0.8
	} else if negCount > posCount {
		sentiment = "negative"
		score = -0.8
	}

	// Enhanced Topic Modeling
	topics := []string{}
	topicKeywords := map[string][]string{
		"cybersecurity": {"osint", "security", "exploit", "vulnerab", "threat"},
		"social_media":  {"twitter", "tweet", "post", "hashtag", "viral"},
		"development":   {"code", "github", "programming", "api", "framework"},
		"finance":       {"market", "crypto", "stock", "price", "wallet"},
	}

	for topic, keywords := range topicKeywords {
		for _, kw := range keywords {
			if strings.Contains(text, kw) {
				topics = append(topics, topic)
				break
			}
		}
	}

	// Mock Network Clustering logic
	clusters := []string{}
	if len(topics) > 0 {
		clusters = append(clusters, "community_"+topics[0])
	} else {
		clusters = append(clusters, "general_community")
	}

	return AnalysisResult{
		Sentiment: sentiment,
		Score:     score,
		Topics:    topics,
		Clusters:  clusters,
	}
}
