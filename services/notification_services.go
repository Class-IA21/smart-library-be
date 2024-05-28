package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/entity"
	"github.com/dimassfeb-09/smart-library-be/helper"
	"net/http"
	"time"
)

type NotificationServicesInterface interface {
	GetNotificationByAccountID(ctx context.Context, accountID int) ([]*entity.Notification, *entity.ErrorResponse)
}

type NotificationServices struct {
	*sql.DB
	*StudentServices
	*AccountServices
	*BorrowServices
}

func NewNotificationServices(db *sql.DB, ss *StudentServices, as *AccountServices, bs *BorrowServices) *NotificationServices {
	return &NotificationServices{
		DB:              db,
		StudentServices: ss,
		AccountServices: as,
		BorrowServices:  bs,
	}
}

func (ns *NotificationServices) GetNotificationByAccountID(ctx context.Context, accountID int) ([]*entity.Notification, *entity.ErrorResponse) {
	_, errorResponse := ns.AccountServices.GetAccountByID(ctx, accountID)
	if errorResponse != nil {
		return nil, errorResponse
	}

	student, errorResponse := ns.StudentServices.GetStudentByAccountID(ctx, accountID)
	if errorResponse != nil {
		return nil, errorResponse
	}

	borrows, errorResponse := ns.BorrowServices.GetBorrowsByStudentID(ctx, student.ID)
	if errorResponse != nil {
		return nil, errorResponse
	}

	var notifications []*entity.Notification
	dueDateLayout := "2006-01-02 15:04:05"
	for _, trx := range borrows.Transactions {

		t, err := time.Parse(dueDateLayout, trx.DueDate)
		if err != nil {
			return nil, helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
		}
		dueDateUnix := t.Unix()
		currentTimeUnix := time.Now().Unix()

		if dueDateUnix-currentTimeUnix < 24*60*60 && dueDateUnix-currentTimeUnix > 0 {
			notification := entity.Notification{
				TransactionID: trx.ID,
				DueDate:       trx.DueDate,
				BookIDS:       trx.BookIDS,
				Message:       fmt.Sprintf("Please return the book to the library before %s", trx.DueDate),
			}
			notifications = append(notifications, &notification)
		}
	}

	return notifications, nil
}
