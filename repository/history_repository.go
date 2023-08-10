package repository

import (
	"database/sql"
	"employeeleave/model"
)

type HistoryRepository interface {
	Create(payload model.HistoryLeave) error
	GetHistoryById(id string) (model.HistoryLeave, error)
	List() ([]model.HistoryLeave, error)
}

type historyRepository struct {
	db *sql.DB
}

func (h *historyRepository) Create(payload model.HistoryLeave) error {
	_, err := h.db.Exec("INSERT INTO history_leave (id, employee_id, transaction_id, date_start, date_end, leave_duration, status_leave) VALUES ($1, $2, $3, $4, $5, $6, $7)", payload.Id, payload.Employee.ID, payload.Transaction.ID, payload.DateStart, payload.DateEnd, payload.LeaveDuration, payload.StatusLeave)
	if err != nil {
		return err
	}
	return nil
}

func (h *historyRepository) GetHistoryById(id string) (model.HistoryLeave, error) {
	var history model.HistoryLeave
	err := h.db.QueryRow("SELECT id, employee_id, transaction_id, date_start, date_end, leave_duration, status_leave FROM history_leave WHERE id = $1", id).Scan(&history.Id, &history.Employee.ID, &history.Transaction.ID, &history.DateStart, &history.DateEnd, &history.LeaveDuration, &history.StatusLeave)
	if err != nil {
		return model.HistoryLeave{}, err
	}
	return history, nil
}

func (h *historyRepository) List() ([]model.HistoryLeave, error) {
	var histories []model.HistoryLeave

	rows, err := h.db.Query("SELECT h.id, e.id, t.id, h.date_start, h.date_end, h.leave_duration, h.status_leave FROM history_leave h INNER JOIN employee e ON e.id = h.employee_id INNER JOIN transaction_leave t ON t.id = h.transaction_id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var history model.HistoryLeave
		err := rows.Scan(&history.Id, &history.Employee.ID, &history.Transaction.ID, &history.DateStart, &history.DateEnd, &history.LeaveDuration, &history.StatusLeave)
		if err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}

func NewHistoryRepository(db *sql.DB) HistoryRepository {
	return &historyRepository{db: db}
}