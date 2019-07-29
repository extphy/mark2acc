package app

import (
   "time"

   "github.com/sirupsen/logrus"

   "github.com/extphy/mark2acc/env"
   "github.com/extphy/mark2acc/sensor"
   "github.com/extphy/mark2acc/regulator"
)

type App struct {
   log         *logrus.Entry
   config      *env.Config
   sensor      *sensor.TempSensor
   regulator   *regulator.ACCRegulator
}

func NewApp(log *logrus.Logger, c *env.Config) *App {

   a := &App {
      log: log.WithField("module", "app"),
      config: c,
      sensor: sensor.NewTempSensor(log, c),
      regulator: regulator.NewACCRegulator(log, c),
   }

   return a
}

func (a *App) Run() {
   a.log.Info("running")

   go a.sensor.Run()
   go a.regulator.Run()

   for {
      select {
      case temp := <-a.sensor.RecvChan:
         a.regulator.SendChan <- temp
         break
      case <-time.After(10 * time.Second):
         break
      }
   }
}
