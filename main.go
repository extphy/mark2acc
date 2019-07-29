package main

import (
   "flag"

   "github.com/sirupsen/logrus"
   "github.com/extphy/mark2acc/app"
   "github.com/extphy/mark2acc/env"
)

func main() {

   var configPath = flag.String("config", "./config.json", "configuration path")
   var debug = flag.Bool("debug", false, "debug")
   flag.Parse()

   log := logrus.New()

   if *debug {
      log.Level = logrus.DebugLevel
   } else {
      log.Level = logrus.WarnLevel
   }
   log.Formatter = &logrus.JSONFormatter{}

   config, err := env.LoadConfig(*configPath)
   if err != nil {
      log.Fatal(err)
   }

   application := app.NewApp(log, config)
   application.Run()
}
