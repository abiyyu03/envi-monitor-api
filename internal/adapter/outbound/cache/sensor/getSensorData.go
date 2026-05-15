package sensor

import (
	"context"
	"encoding/json"

	outboundEntity "go-projects/envi-monitor/internal/adapter/outbound/entity"
)

func (c sensorCache) GetSensorData(ctx context.Context) (outboundEntity.Sensor, error) {
	str, err := c.pkg.Cache.Client.Get(ctx, sensorKey).Result()
	if err != nil {
		return outboundEntity.Sensor{}, err
	}

	var data outboundEntity.Sensor
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		return outboundEntity.Sensor{}, err
	}

	return data, nil
}

type IGetSensorData interface {
	GetSensorData(ctx context.Context) (outboundEntity.Sensor, error)
}
