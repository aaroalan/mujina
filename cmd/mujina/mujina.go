package main

import (
	"flag"
	"fmt"
	"github.com/aaroalan/mujina/internal/config"
	"github.com/aaroalan/mujina/internal/handler"
	"github.com/aaroalan/mujina/internal/help"
	"github.com/aaroalan/mujina/internal/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
)

const (
	configOptsName  = "configPath"
	configOptsUsage = "Path to the configuration file that defines the endpoints."
	portOptsName    = "port"
	portOptsUsage   = "Port where http server will run"
	versionOptsName = "version"
	versionUsage    = "Prints current app version"
	defaultFileName = "/mujina.json"
	defaultPort     = 8080
	allDomains      = "*"
	appName         = "Mujina"
	appVersion      = "0.1.0"
)

func main() {
	cfgPath := flag.String(configOptsName, defaultPath(), configOptsUsage)
	port := flag.Int(portOptsName, defaultPort, portOptsUsage)
	versionOpt := flag.Bool(versionOptsName, false, versionUsage)
	flag.Parse()
	if *versionOpt {
		version()
		return
	}
	// Generate global config using path.
	cfg, err := config.NewConfig(*cfgPath)
	help.PanicIfError(&err)
	// Start Gin server
	fmt.Println("Running server config: " + *cfgPath)
	startServer(&cfg, *port)
}

func version() {
	fmt.Println(appName + " v" + appVersion)
}

// defaultPath : By default it will look for a file called "mujina.json" in the same folder as binary.
func defaultPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	help.PanicIfError(&err)
	return dir + defaultFileName
}

// configMiddleware : Sets the pointer to the global config so can be use in the Handler method.
func configMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	}
}

// startServer : Starts the Gin server, server will run in port 8080 unless specify, all domains are allowed in CORS
// config.
func startServer(cfg *config.Config, port int) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowOrigins = []string{allDomains}
	r.Use(cors.New(corsCfg))
	r.Use(configMiddleware(cfg))
	route.AddRoutes(r, handler.Handler, cfg.Endpoints)
	strPort := ":" + strconv.Itoa(port)
	fmt.Println("Starting server: http://localhost" + strPort + "...")
	_ = r.Run(strPort)
}
