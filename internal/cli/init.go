package cli

import (
	"github.com/ismailtsdln/HawkLens/internal/plugins/instagram"
	"github.com/ismailtsdln/HawkLens/internal/plugins/reddit"
	"github.com/ismailtsdln/HawkLens/internal/plugins/tiktok"
	"github.com/ismailtsdln/HawkLens/internal/plugins/twitter"
	"github.com/ismailtsdln/HawkLens/internal/plugins/youtube"
	"github.com/ismailtsdln/HawkLens/pkg/plugins"
)

func init() {
	// Register plugins
	plugins.Register(twitter.NewTwitterPlugin("mock_twitter_key"))
	plugins.Register(youtube.NewYouTubePlugin("mock_youtube_key"))
	plugins.Register(reddit.NewRedditPlugin())
	plugins.Register(instagram.NewInstagramPlugin())
	plugins.Register(tiktok.NewTikTokPlugin())
}
