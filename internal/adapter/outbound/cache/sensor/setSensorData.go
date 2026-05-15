package sensor

import (
	"context"
	"encoding/json"
	"time"

	outboundEntity "go-projects/envi-monitor/internal/adapter/outbound/entity"
)

const sensorKey = "sensor:data"

func (c sensorCache) SetSensorData(ctx context.Context, data outboundEntity.Sensor) error {
	val, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	_, err = c.pkg.Cache.Client.Set(ctx, sensorKey, val, time.Minute*10).Result()
	return err
}

type ISetSensorData interface {
	SetSensorData(ctx context.Context, data outboundEntity.Sensor) error
}
