
package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path"
	"sync"
)

type Service struct {
	Server      Server