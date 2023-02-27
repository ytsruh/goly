package server

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var Monitor = monitor.New(monitor.Config{
	Title:      "Goly Monitor",
	Refresh:    3 * time.Second,
	APIOnly:    false,
	Next:       nil,
	CustomHead: "",
	FontURL:    "https://fonts.googleapis.com/css2?family=Almendra+Display&display=swap",
})

var ApiMonitor = monitor.New(monitor.Config{
	APIOnly: true,
})

var RateLimiter = limiter.New(limiter.Config{
	Max:               20,
	Expiration:        30 * time.Second,
	LimiterMiddleware: limiter.SlidingWindow{},
})

var Recover = recover.New()

// Compression levels
const (
	LevelDisabled        = -1
	LevelDefault         = 0
	LevelBestSpeed       = 1
	LevelBestCompression = 2
)

var Compression = compress.New(compress.Config{
	Level: LevelBestSpeed,
})

func RedirectMiddleware(c *fiber.Ctx) error {
	requestIp := c.IP()
	log.Println("A redirect has been triggered from IP: " + requestIp)
	// Go to next middleware:
	return c.Next()
}
