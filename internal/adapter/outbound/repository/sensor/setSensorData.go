package sensor

import (
	"context"

	outboundEntity "go-projects/envi-monitor/internal/adapter/outbound/entity"
)

func (r sensor) SetSensorData(ctx context.Context, data outboundEntity.Sensor) error {
	if err := r.pkg.DB.WithContext(ctx).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

type IStoreSensorData interface {
	SetSensorData(ctx context.Context, data outboundEntity.Sensor) error
}
