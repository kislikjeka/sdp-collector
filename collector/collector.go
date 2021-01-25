package collector

import (
	"database/sql"
	"github.com/kislikjeka/sdp-collector/models"
	"gorm.io/gorm"
	"time"
)

type Collector struct {
	Name            string
	MaxTimeDuration time.Duration
	errorDB         *gorm.DB
}

func (cl Collector) GetTimeProceed() (*time.Time, *time.Time, error) {

	var state models.CollectorsState
	state = models.CollectorsState{
		CollectorName: cl.Name,
		LastTimeProceed: sql.NullTime{
			Time:  time.Date(2020, 11, 24, 14, 0, 0, 0, time.UTC),
			Valid: true,
		},
	}

	lastTime := state.LastTimeProceed.Time
	nextTime := cl.calculateNextTimeProceed(lastTime)

	return &lastTime, nextTime, nil
}

//Расчитывает
func (cl Collector) calculateNextTimeProceed(t time.Time) *time.Time {
	now := time.Now()
	var nextTime time.Time
	if now.Sub(t) > cl.MaxTimeDuration {
		nextTime = t.Add(cl.MaxTimeDuration)
	} else {
		nextTime = now
	}

	return &nextTime
}

func (cl *Collector) WriteError(t string, level string, event string, message string, shardGuid string) error {
	lvlError := models.LevelError{
		Type:      t,
		Level:     level,
		Event:     event,
		Message:   message,
		ShardGuid: shardGuid,
	}
	err := cl.errorDB.Create(&lvlError).Error
	if err != nil {
		return err
	}
	return nil
}

//func initCollectorsDB()  {
//	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", //TODO SetConfig
//		"localhost",
//		"root",
//		"pass",
//		"db",
//		"port")
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return nil, nil, err
//	}
//	var state models.CollectorsState
//	err = db.Where(&models.CollectorsState{CollectorName: coll.Name}).First(&state).Error
//}
