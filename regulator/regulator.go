package regulator

import (
   "time"

   "github.com/sirupsen/logrus"
   "github.com/extphy/mark2acc/env"
   "github.com/extphy/mark2acc/common"
)

type ACCRegulator struct {
   log         *logrus.Entry
   config      *env.Config
   temp        common.TempData
   SendChan    chan<- common.TempData
   sendChan    <-chan common.TempData
}

func NewACCRegulator(log *logrus.Logger, c *env.Config) *ACCRegulator {

   sendChan := make(chan common.TempData)

   r := &ACCRegulator {
      log: log.WithField("module", "regulator"),
      config: c,
      sendChan: sendChan,
      SendChan: sendChan,
   }

   return r
}

func (r *ACCRegulator) Run() {

   r.log.Info("running acc regulator")

   for {
      select {
      case <-time.After(2 * time.Second):
         r.log.Info("modulate signal")
      case temp := <-r.sendChan:
         r.log.Info("change temp:", temp)
      }
   }
}

