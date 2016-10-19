package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/armon/go-radix"
	"github.com/icrowley/fake"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

type IpParams struct {
	Allow       bool       `json:"allow"`
	Description string     `json:"description,omitempty"`
	ExpiredAt   *time.Time `json:"expired_at,omitempty"`
}

func (ip *IpParams) SetExpire(exp time.Time) *IpParams {
	ip.ExpiredAt = &exp
	return ip
}

func (ip IpParams) AllowIp() bool {
	return ip.Allow && (ip.ExpiredAt == nil || ip.ExpiredAt.After(time.Now()))
}

type Config struct {
	Title     string               `json:"title,omitempty"`
	Ips       map[string]*IpParams `json:"ips,omitempty"`
	CreatedAt *time.Time           `json:"created_at,omitempty"`
	UpdatedAt *time.Time           `json:"update_at,omitempty"`

	fileName string
	ipsRadix *radix.Tree
}

func NewConfig() *Config {
	c := new(Config)
	now := time.Now()
	c.CreatedAt = &now
	c.Ips = make(map[string]*IpParams)
	c.ipsRadix = radix.New()
	return c
}

func (c *Config) Load(fileName string) error {
	c.fileName = fileName
	if jsonRaw, err := ioutil.ReadFile(fileName); err == nil {
		if err := json.Unmarshal(jsonRaw, c); err != nil {
			return errors.Wrapf(err, "Configuration file `%s` load error", fileName)
		}
	}
	if c.Ips != nil {
		// Initiate radix tree
		for ip, desc := range c.Ips {
			c.ipsRadix.Insert(ip, desc)
		}
	}
	return nil
}

func (c *Config) SaveToFile(fileName string) error {
	now := time.Now()
	c.UpdatedAt = &now

	if jsonRaw, err := json.MarshalIndent(c, "", "  "); err == nil {
		err := ioutil.WriteFile(fileName, jsonRaw, 0640)
		if err != nil {
			return errors.Wrapf(err, "Configuration file `%s` save error", fileName)
		}
	} else {
		return errors.Wrap(err, "Marshaling error")
	}
	return nil
}

func (c *Config) Save() error {
	return c.SaveToFile(c.fileName)
}

func (c *Config) SetIp(ip string, allow bool) *IpParams {
	if ipParams, ok := c.Ips[ip]; ok {
		ipParams.Allow = allow
		return ipParams
	}
	ipParams := new(IpParams)
	ipParams.Allow = allow
	c.Ips[ip] = ipParams
	c.ipsRadix.Insert(ip, ipParams)
	return ipParams
}

func (c *Config) DeleteIp(ip string) {
	delete(c.Ips, ip)
	c.ipsRadix.Delete(ip)
}

func (c Config) FindIp(ip string) (string, *IpParams, bool) {
	if fip, ipParams, ok := c.ipsRadix.LongestPrefix(ip); ok {
		return fip, ipParams.(*IpParams), ok
	}
	return "", nil, false
}

func (c Config) AllowIp(ip string) bool {
	if _, p, ok := c.FindIp(ip); ok {
		return p.AllowIp()
	}
	return false
}

func (c Config) FileName() string {
	return c.fileName
}

func main() {
	cfg := NewConfig()

	if err := cfg.Load("config.json"); err != nil {
		log.Printf("[ERROR] %+v", err)
		return
	}
	defer func() {
		if err := cfg.Save(); err != nil {
			log.Printf("[ERROR] %+v", err)
			return
		}
	}()

	if len(cfg.Ips) > 100 {
		for k, v := range cfg.Ips {
			cfg.DeleteIp(k)
			log.Printf("Delete IP %s: %s", k, v)
			break
		}
	}

	fake.SetLang("ru")
	ipParams := cfg.SetIp(fake.IPv4(), true)
	ipParams.Description = fake.FullName()
	ipParams.SetExpire(time.Now().Add(time.Minute))

	pp.Println(cfg.AllowIp("181.113.250.128"))
	pp.Println(cfg.AllowIp("89.155.206.159"))
}
