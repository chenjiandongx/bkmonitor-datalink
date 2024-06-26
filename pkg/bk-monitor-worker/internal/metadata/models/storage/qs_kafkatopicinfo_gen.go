// Code generated by go-queryset. DO NOT EDIT.
package storage

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set KafkaTopicInfoQuerySet

// KafkaTopicInfoQuerySet is an queryset type for KafkaTopicInfo
type KafkaTopicInfoQuerySet struct {
	db *gorm.DB
}

// NewKafkaTopicInfoQuerySet constructs new KafkaTopicInfoQuerySet
func NewKafkaTopicInfoQuerySet(db *gorm.DB) KafkaTopicInfoQuerySet {
	return KafkaTopicInfoQuerySet{
		db: db.Model(&KafkaTopicInfo{}),
	}
}

func (qs KafkaTopicInfoQuerySet) w(db *gorm.DB) KafkaTopicInfoQuerySet {
	return NewKafkaTopicInfoQuerySet(db)
}

func (qs KafkaTopicInfoQuerySet) Select(fields ...KafkaTopicInfoDBSchemaField) KafkaTopicInfoQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *KafkaTopicInfo) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *KafkaTopicInfo) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) All(ret *[]KafkaTopicInfo) error {
	return qs.db.Find(ret).Error
}

// BatchSizeEq is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeEq(batchSize int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size = ?", batchSize))
}

// BatchSizeGt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeGt(batchSize int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size > ?", batchSize))
}

// BatchSizeGte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeGte(batchSize int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size >= ?", batchSize))
}

// BatchSizeIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeIn(batchSize ...int64) KafkaTopicInfoQuerySet {
	if len(batchSize) == 0 {
		qs.db.AddError(errors.New("must at least pass one batchSize in BatchSizeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("batch_size IN (?)", batchSize))
}

// BatchSizeIsNotNull is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeIsNotNull() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size IS NOT NULL"))
}

// BatchSizeIsNull is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeIsNull() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size IS NULL"))
}

// BatchSizeLt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeLt(batchSize int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size < ?", batchSize))
}

// BatchSizeLte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeLte(batchSize int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size <= ?", batchSize))
}

// BatchSizeNe is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeNe(batchSize int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("batch_size != ?", batchSize))
}

// BatchSizeNotIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BatchSizeNotIn(batchSize ...int64) KafkaTopicInfoQuerySet {
	if len(batchSize) == 0 {
		qs.db.AddError(errors.New("must at least pass one batchSize in BatchSizeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("batch_size NOT IN (?)", batchSize))
}

// BkDataIdEq is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdEq(bkDataId uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("bk_data_id = ?", bkDataId))
}

// BkDataIdGt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdGt(bkDataId uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("bk_data_id > ?", bkDataId))
}

// BkDataIdGte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdGte(bkDataId uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("bk_data_id >= ?", bkDataId))
}

// BkDataIdIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdIn(bkDataId ...uint) KafkaTopicInfoQuerySet {
	if len(bkDataId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkDataId in BkDataIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_data_id IN (?)", bkDataId))
}

// BkDataIdLt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdLt(bkDataId uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("bk_data_id < ?", bkDataId))
}

// BkDataIdLte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdLte(bkDataId uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("bk_data_id <= ?", bkDataId))
}

// BkDataIdNe is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdNe(bkDataId uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("bk_data_id != ?", bkDataId))
}

// BkDataIdNotIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) BkDataIdNotIn(bkDataId ...uint) KafkaTopicInfoQuerySet {
	if len(bkDataId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkDataId in BkDataIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_data_id NOT IN (?)", bkDataId))
}

// ConsumeRateEq is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateEq(consumeRate int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate = ?", consumeRate))
}

// ConsumeRateGt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateGt(consumeRate int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate > ?", consumeRate))
}

// ConsumeRateGte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateGte(consumeRate int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate >= ?", consumeRate))
}

// ConsumeRateIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateIn(consumeRate ...int64) KafkaTopicInfoQuerySet {
	if len(consumeRate) == 0 {
		qs.db.AddError(errors.New("must at least pass one consumeRate in ConsumeRateIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("consume_rate IN (?)", consumeRate))
}

// ConsumeRateIsNotNull is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateIsNotNull() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate IS NOT NULL"))
}

// ConsumeRateIsNull is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateIsNull() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate IS NULL"))
}

// ConsumeRateLt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateLt(consumeRate int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate < ?", consumeRate))
}

// ConsumeRateLte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateLte(consumeRate int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate <= ?", consumeRate))
}

// ConsumeRateNe is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateNe(consumeRate int64) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("consume_rate != ?", consumeRate))
}

// ConsumeRateNotIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) ConsumeRateNotIn(consumeRate ...int64) KafkaTopicInfoQuerySet {
	if len(consumeRate) == 0 {
		qs.db.AddError(errors.New("must at least pass one consumeRate in ConsumeRateNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("consume_rate NOT IN (?)", consumeRate))
}

// Count is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Delete is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) Delete() error {
	return qs.db.Delete(KafkaTopicInfo{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(KafkaTopicInfo{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(KafkaTopicInfo{})
	return db.RowsAffected, db.Error
}

// FlushIntervalEq is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalEq(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval = ?", flushInterval))
}

// FlushIntervalGt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalGt(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval > ?", flushInterval))
}

// FlushIntervalGte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalGte(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval >= ?", flushInterval))
}

// FlushIntervalIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalIn(flushInterval ...string) KafkaTopicInfoQuerySet {
	if len(flushInterval) == 0 {
		qs.db.AddError(errors.New("must at least pass one flushInterval in FlushIntervalIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("flush_interval IN (?)", flushInterval))
}

// FlushIntervalIsNotNull is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalIsNotNull() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval IS NOT NULL"))
}

// FlushIntervalIsNull is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalIsNull() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval IS NULL"))
}

// FlushIntervalLike is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalLike(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval LIKE ?", flushInterval))
}

// FlushIntervalLt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalLt(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval < ?", flushInterval))
}

// FlushIntervalLte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalLte(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval <= ?", flushInterval))
}

// FlushIntervalNe is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalNe(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval != ?", flushInterval))
}

// FlushIntervalNotIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalNotIn(flushInterval ...string) KafkaTopicInfoQuerySet {
	if len(flushInterval) == 0 {
		qs.db.AddError(errors.New("must at least pass one flushInterval in FlushIntervalNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("flush_interval NOT IN (?)", flushInterval))
}

// FlushIntervalNotlike is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) FlushIntervalNotlike(flushInterval string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("flush_interval NOT LIKE ?", flushInterval))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) GetUpdater() KafkaTopicInfoUpdater {
	return NewKafkaTopicInfoUpdater(qs.db)
}

// IdEq is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdEq(id uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("id = ?", id))
}

// IdGt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdGt(id uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("id > ?", id))
}

// IdGte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdGte(id uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("id >= ?", id))
}

// IdIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdIn(id ...uint) KafkaTopicInfoQuerySet {
	if len(id) == 0 {
		qs.db.AddError(errors.New("must at least pass one id in IdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", id))
}

// IdLt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdLt(id uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("id < ?", id))
}

// IdLte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdLte(id uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("id <= ?", id))
}

// IdNe is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdNe(id uint) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("id != ?", id))
}

// IdNotIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) IdNotIn(id ...uint) KafkaTopicInfoQuerySet {
	if len(id) == 0 {
		qs.db.AddError(errors.New("must at least pass one id in IdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", id))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) Limit(limit int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) Offset(offset int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs KafkaTopicInfoQuerySet) One(ret *KafkaTopicInfo) error {
	return qs.db.First(ret).Error
}

// OrderAscByBatchSize is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderAscByBatchSize() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("batch_size ASC"))
}

// OrderAscByBkDataId is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderAscByBkDataId() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("bk_data_id ASC"))
}

// OrderAscByConsumeRate is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderAscByConsumeRate() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("consume_rate ASC"))
}

// OrderAscByFlushInterval is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderAscByFlushInterval() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("flush_interval ASC"))
}

// OrderAscById is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderAscById() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByPartition is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderAscByPartition() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("partition ASC"))
}

// OrderAscByTopic is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderAscByTopic() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("topic ASC"))
}

// OrderDescByBatchSize is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderDescByBatchSize() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("batch_size DESC"))
}

// OrderDescByBkDataId is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderDescByBkDataId() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("bk_data_id DESC"))
}

// OrderDescByConsumeRate is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderDescByConsumeRate() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("consume_rate DESC"))
}

// OrderDescByFlushInterval is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderDescByFlushInterval() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("flush_interval DESC"))
}

// OrderDescById is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderDescById() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByPartition is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderDescByPartition() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("partition DESC"))
}

// OrderDescByTopic is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) OrderDescByTopic() KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Order("topic DESC"))
}

// PartitionEq is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionEq(partition int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("partition = ?", partition))
}

// PartitionGt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionGt(partition int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("partition > ?", partition))
}

// PartitionGte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionGte(partition int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("partition >= ?", partition))
}

// PartitionIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionIn(partition ...int) KafkaTopicInfoQuerySet {
	if len(partition) == 0 {
		qs.db.AddError(errors.New("must at least pass one partition in PartitionIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("partition IN (?)", partition))
}

// PartitionLt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionLt(partition int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("partition < ?", partition))
}

// PartitionLte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionLte(partition int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("partition <= ?", partition))
}

// PartitionNe is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionNe(partition int) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("partition != ?", partition))
}

// PartitionNotIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) PartitionNotIn(partition ...int) KafkaTopicInfoQuerySet {
	if len(partition) == 0 {
		qs.db.AddError(errors.New("must at least pass one partition in PartitionNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("partition NOT IN (?)", partition))
}

// TopicEq is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicEq(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic = ?", topic))
}

// TopicGt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicGt(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic > ?", topic))
}

// TopicGte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicGte(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic >= ?", topic))
}

// TopicIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicIn(topic ...string) KafkaTopicInfoQuerySet {
	if len(topic) == 0 {
		qs.db.AddError(errors.New("must at least pass one topic in TopicIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("topic IN (?)", topic))
}

// TopicLike is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicLike(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic LIKE ?", topic))
}

// TopicLt is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicLt(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic < ?", topic))
}

// TopicLte is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicLte(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic <= ?", topic))
}

// TopicNe is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicNe(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic != ?", topic))
}

// TopicNotIn is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicNotIn(topic ...string) KafkaTopicInfoQuerySet {
	if len(topic) == 0 {
		qs.db.AddError(errors.New("must at least pass one topic in TopicNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("topic NOT IN (?)", topic))
}

// TopicNotlike is an autogenerated method
// nolint: dupl
func (qs KafkaTopicInfoQuerySet) TopicNotlike(topic string) KafkaTopicInfoQuerySet {
	return qs.w(qs.db.Where("topic NOT LIKE ?", topic))
}

// SetBatchSize is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) SetBatchSize(batchSize *int64) KafkaTopicInfoUpdater {
	u.fields[string(KafkaTopicInfoDBSchema.BatchSize)] = batchSize
	return u
}

// SetBkDataId is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) SetBkDataId(bkDataId uint) KafkaTopicInfoUpdater {
	u.fields[string(KafkaTopicInfoDBSchema.BkDataId)] = bkDataId
	return u
}

// SetConsumeRate is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) SetConsumeRate(consumeRate *int64) KafkaTopicInfoUpdater {
	u.fields[string(KafkaTopicInfoDBSchema.ConsumeRate)] = consumeRate
	return u
}

// SetFlushInterval is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) SetFlushInterval(flushInterval *string) KafkaTopicInfoUpdater {
	u.fields[string(KafkaTopicInfoDBSchema.FlushInterval)] = flushInterval
	return u
}

// SetId is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) SetId(id uint) KafkaTopicInfoUpdater {
	u.fields[string(KafkaTopicInfoDBSchema.Id)] = id
	return u
}

// SetPartition is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) SetPartition(partition int) KafkaTopicInfoUpdater {
	u.fields[string(KafkaTopicInfoDBSchema.Partition)] = partition
	return u
}

// SetTopic is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) SetTopic(topic string) KafkaTopicInfoUpdater {
	u.fields[string(KafkaTopicInfoDBSchema.Topic)] = topic
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u KafkaTopicInfoUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set KafkaTopicInfoQuerySet

// ===== BEGIN of KafkaTopicInfo modifiers

// KafkaTopicInfoDBSchemaField describes database schema field. It requires for method 'Update'
type KafkaTopicInfoDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f KafkaTopicInfoDBSchemaField) String() string {
	return string(f)
}

// KafkaTopicInfoDBSchema stores db field names of KafkaTopicInfo
var KafkaTopicInfoDBSchema = struct {
	Id            KafkaTopicInfoDBSchemaField
	BkDataId      KafkaTopicInfoDBSchemaField
	Topic         KafkaTopicInfoDBSchemaField
	Partition     KafkaTopicInfoDBSchemaField
	BatchSize     KafkaTopicInfoDBSchemaField
	FlushInterval KafkaTopicInfoDBSchemaField
	ConsumeRate   KafkaTopicInfoDBSchemaField
}{

	Id:            KafkaTopicInfoDBSchemaField("id"),
	BkDataId:      KafkaTopicInfoDBSchemaField("bk_data_id"),
	Topic:         KafkaTopicInfoDBSchemaField("topic"),
	Partition:     KafkaTopicInfoDBSchemaField("partition"),
	BatchSize:     KafkaTopicInfoDBSchemaField("batch_size"),
	FlushInterval: KafkaTopicInfoDBSchemaField("flush_interval"),
	ConsumeRate:   KafkaTopicInfoDBSchemaField("consume_rate"),
}

// Update updates KafkaTopicInfo fields by primary key
// nolint: dupl
func (o *KafkaTopicInfo) Update(db *gorm.DB, fields ...KafkaTopicInfoDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":             o.Id,
		"bk_data_id":     o.BkDataId,
		"topic":          o.Topic,
		"partition":      o.Partition,
		"batch_size":     o.BatchSize,
		"flush_interval": o.FlushInterval,
		"consume_rate":   o.ConsumeRate,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update KafkaTopicInfo %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// KafkaTopicInfoUpdater is an KafkaTopicInfo updates manager
type KafkaTopicInfoUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewKafkaTopicInfoUpdater creates new KafkaTopicInfo updater
// nolint: dupl
func NewKafkaTopicInfoUpdater(db *gorm.DB) KafkaTopicInfoUpdater {
	return KafkaTopicInfoUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&KafkaTopicInfo{}),
	}
}

// ===== END of KafkaTopicInfo modifiers

// ===== END of all query sets
