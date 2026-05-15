package sensor

import (
	"context"
	"encoding/json"
	"log"

	mqttEntity "go-projects/envi-monitor/internal/adapter/inbound/mqtt/entity"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (s *Subscriber) SubSensorData() {
	token := s.client.Subscribe("sensor/data", 1, func(_ mqtt.Client, msg mqtt.Message) {
		var payload mqttEntity.SensorPayload
		if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
			return
		}

		ctx := context.Background()
		if err := s.service.Sensor.StoreSensorData(ctx, payload.ToServiceEntity()); err != nil {
			log.Printf("failed to store sensor data: %v", err)
		}
	})
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("failed to subscribe to sensor/data: %v", token.Error())
	}
	log.Println("subscribed to topic: sensor/data")
}
