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
	SendEmailNotification(ctx context.Context) ([]*entity.Notification, *entity.ErrorResponse)
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

func (ns *NotificationServices) SendEmailNotification(ctx context.Context) *entity.ErrorResponse {
	borrows, errorResponse := ns.BorrowRepository.GetBorrows(ctx, ns.DB)
	if errorResponse != nil {
		return errorResponse
	}

	dueDateLayout := "2006-01-02 15:04:05"
	var notifications []entity.Notification
	for _, trx := range borrows.Transactions {
		location, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
		}

		t, err := time.ParseInLocation(dueDateLayout, trx.DueDate, location)
		if err != nil {
			return helper.ErrorResponse(http.StatusInternalServerError, "Internal Server Error")
		}
		timeDifference := t.Sub(time.Now())

		if trx.Status == "borrowed" && trx.ReturnDate == "" {
			if timeDifference.Hours() >= 72 && timeDifference.Hours() <= 72.06 {
				var notification = entity.Notification{
					TransactionID: trx.ID,
					BookIDS:       trx.BookIDS,
					DueDate:       trx.DueDate,
					Message:       fmt.Sprintf("Peringatan! Sudah lebih dari 2 hari sejak batas pengembalian buku pada tanggal dan waktu %s. Segera kembalikan buku.", trx.DueDate),
				}
				notifications = append(notifications, notification)
			}

			if timeDifference.Hours() >= 48 && timeDifference.Hours() <= 48.06 {
				var notification = entity.Notification{
					TransactionID: trx.ID,
					BookIDS:       trx.BookIDS,
					DueDate:       trx.DueDate,
					Message:       "Harap kembalikan buku, karena telah melebihi 1 hari dari batas pengembalian buku yang telah ditentukan.",
				}
				notifications = append(notifications, notification)
			}

			if timeDifference.Hours() >= 24 && timeDifference.Hours() <= 24.06 {
				fmt.Println(timeDifference.Hours())
				var notification = entity.Notification{
					TransactionID: trx.ID,
					BookIDS:       trx.BookIDS,
					DueDate:       trx.DueDate,
					Message:       fmt.Sprintf("Harap kembalikan buku dalam waktu 24 jam. Batas akhir pengembalian buku pada tanggal dan waktu %s.", trx.DueDate),
				}
				notifications = append(notifications, notification)
			}

		}
	}

	for _, notif := range notifications {
		fmt.Println(notif)
		fmt.Println()
	}

	return nil
}
