package pingpong

import (
	"log"
	"time"
)

func WatchKevin(callPort string, callAddress string, streamerReqBody string, delaySeconds int) {

	for {
		go func() {
			_, err := CallPingPong(callPort, callAddress)
			if err != nil {
				WatchPongHealth.Set(0)
				log.Println(err)
				log.Println("Pong service is unhealthy")
				return
			}
			WatchPongHealth.Set(1)
			log.Println("Pong service is healthy")
		}()

		go func() {
			err := CallStreamPong(callPort, callAddress, streamerReqBody)
			if err != nil {
				WatchStreamHealth.Set(0)
				log.Println(err)
				log.Println("Stream service is unhealthy")
				return
			}
			WatchStreamHealth.Set(1)
			log.Println("Stream service is healthy")
		}()
		time.Sleep(time.Duration(delaySeconds) * time.Second)
	}

}
