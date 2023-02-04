package repository

import (
	"errors"
)

func (r *Repository) Restore() (err error) {
	if r.BackupStorage == nil {
		return errors.New("неподдерживаемый тип хранилища")
	}

	r.Mu.Lock()
	defer r.Mu.Unlock()

	metrics, err := r.BackupStorage.GetAllMetrics()
	if err != nil {
		return err
	}

	for _, metric := range metrics {
		r.Store[metric.ID] = metric.Value
	}

	return nil
}
