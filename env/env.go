package env

import (
   "io/ioutil"
   "encoding/json"
)

type Config struct {
   TempScanIntervalSec   uint32  `json:"TempScanIntervalSec"`
}

func (c *Config) ParseJSON(b []byte) error {
   return json.Unmarshal(b, &c)
}

func New(path string) *Config {
   return &Config{}
}

func LoadConfig(configPath string) (*Config, error) {

   config := New(configPath)

   jsonBytes, err := ioutil.ReadFile(configPath)
   if err != nil {
      return nil, err
   }

   if err := config.ParseJSON(jsonBytes); err != nil {
      return nil, err
   }

   return config, err
}
