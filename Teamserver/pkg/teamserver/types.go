package teamserver

import (
    "Havoc/pkg/agent"
    "Havoc/pkg/packager"
    "Havoc/pkg/profile"
    "Havoc/pkg/service"
    "sync"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

type Listener struct {
    Name   string
    Type   int
    Config any
}

type Client struct {
    ClientID      string
    Username      string
    GlobalIP      string
    ClientVersion string
    Connection    *websocket.Conn
    Packager      *packager.Packager
    Authenticated bool
    SessionID     string
    Mutex         sync.Mutex
}

type Users struct {
    Name     string
    Password string
    Hashed   bool
    Online   bool
}

type serverFlags struct {
    Host string
    Port string

    Profile  string
    Verbose  bool
    Debug    bool
    DebugDev bool
    Default  bool
}

type utilFlags struct {
    NoBanner bool
    Debug    bool
    Verbose  bool

    Test bool

    ListOperators bool
}

type TeamserverFlags struct {
    Server serverFlags
    Util   utilFlags
}

type Endpoint struct {
    Endpoint string
    Function func(ctx *gin.Context)
}

type Teamserver struct {
    Flags       TeamserverFlags
    Profile     *profile.Profile
    Clients     map[string]*Client
    Fingerprint string
    Users       []Users
    EventsList  []packager.Package
    Service     *service.Service

    Server struct {
        Path   string
        Engine *gin.Engine
    }

    Agents    agent.Agents
    Listeners []*Listener
    Endpoints []*Endpoint

    Settings struct {
        Compiler64 string
        Compiler32 string
        Nasm       string
    }
}
