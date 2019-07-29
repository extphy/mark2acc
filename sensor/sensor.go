package sensor

import (
   "time"

   "github.com/sirupsen/logrus"
   "github.com/extphy/mark2acc/env"
   "github.com/extphy/mark2acc/common"
)

type TempSensor struct {
   log         *logrus.Entry
   config      *env.Config
   lastTemp    common.TempData
   RecvChan    <-chan common.TempData
   recvChan    chan<- common.TempData
}

func NewTempSensor(log *logrus.Logger, c *env.Config) *TempSensor {

   recvChan := make(chan common.TempData)

   s := &TempSensor {
      log: log.WithField("module", "sensor"),
      config: c,
      recvChan: recvChan,
      RecvChan: recvChan,
   }

   return s
}

func (s *TempSensor) Run() {

   s.log.Info("running temp sensor")

   for {
      select {
      case <-time.After(time.Duration(s.config.TempScanIntervalSec) * time.Second):
         s.log.Info("read temp data from sensor")
         td := common.TempData{}
         td[0] = 1
         td[1] = 2
         s.recvChan <- td
      }
   }

}
