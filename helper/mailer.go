package helper

import (
	"fmt"
	"github.com/dimassfeb-09/smart-library-be/entity"
	"gopkg.in/mail.v2"
	"log"
)

func SendMail(notification *entity.Notification) {
	err := sendMailNotifications("ia21.gunadarma@gmail.com", "dimassfeb@gmail.com", "hello world", notification)
	if err != nil {
		log.Println(err)
	}
}

func sendMailNotifications(fromFromEmail string, sendToMail string, subject string, notification *entity.Notification) error {
	m := mail.NewMessage()

	m.SetHeader("From", fromFromEmail)
	m.SetHeader("To", sendToMail)

	m.SetHeader("Subject", "Hello from Go!")
	m.SetBody("text/html", htmlBody(notification))

	d := mail.NewDialer("smtp.gmail.com", 587, "ia21.gunadarma@gmail.com", "ifqj szmd gclu orqi")

	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Error Sending Mail", err)
	}

	return nil
}

func htmlBody(notification *entity.Notification) string {
	var body = fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>Simple Transactional Email</title>
    <style media="all" type="text/css">
        @media all {
            .btn-primary table td:hover {
                background-color: #ec0867 !important;
            }
            .btn-primary a:hover {
                background-color: #ec0867 !important;
                border-color: #ec0867 !important;
            }
        }
        @media only screen and (max-width: 640px) {
            .main p,
            .main td,
            .main span {
                font-size: 16px !important;
            }
            .wrapper {
                padding: 8px !important;
            }
            .content {
                padding: 0 !important;
            }
            .container {
                padding: 0 !important;
                padding-top: 8px !important;
                width: 100% !important;
            }
            .main {
                border-left-width: 0 !important;
                border-radius: 0 !important;
                border-right-width: 0 !important;
            }
            .btn table {
                max-width: 100% !important;
                width: 100% !important;
            }
            .btn a {
                font-size: 16px !important;
                max-width: 100% !important;
                width: 100% !important;
            }
        }
        @media all {
            .ExternalClass {
                width: 100%;
            }
            .ExternalClass,
            .ExternalClass p,
            .ExternalClass span,
            .ExternalClass font,
            .ExternalClass td,
            .ExternalClass div {
                line-height: 100%;
            }
            .apple-link a {
                color: inherit !important;
                font-family: inherit !important;
                font-size: inherit !important;
                font-weight: inherit !important;
                line-height: inherit !important;
                text-decoration: none !important;
            }
            #MessageViewBody a {
                color: inherit;
                text-decoration: none;
                font-size: inherit;
                font-family: inherit;
                font-weight: inherit;
                line-height: inherit;
            }
        }
    </style>
</head>
<body style="font-family: Helvetica, sans-serif; -webkit-font-smoothing: antialiased; font-size: 16px; line-height: 1.3; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; background-color: #f4f5f6; margin: 0; padding: 0;">
<table role="presentation" border="0" cellpadding="0" cellspacing="0" class="body" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #f4f5f6; width: 100%;" width="100%" bgcolor="#f4f5f6">
    <tr>
        <td style="font-family: Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">&nbsp;</td>
        <td class="container" style="font-family: Helvetica, sans-serif; font-size: 16px; vertical-align: top; max-width: 600px; padding: 0; padding-top: 24px; width: 600px; margin: 0 auto;" width="600" valign="top">
            <div class="content" style="box-sizing: border-box; display: block; margin: 0 auto; max-width: 600px; padding: 0;">

                <!-- START CENTERED WHITE CONTAINER -->
                <span class="preheader" style="color: transparent; display: none; height: 0; max-height: 0; max-width: 0; opacity: 0; overflow: hidden; mso-hide: all; visibility: hidden; width: 0;">This is preheader text. Some clients will show this text as a preview.</span>
                <table role="presentation" border="0" cellpadding="0" cellspacing="0" class="main" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background: #ffffff; border: 1px solid #eaebed; border-radius: 16px; width: 100%;" width="100%">

                    <!-- START MAIN CONTENT AREA -->
                    <tr>
                        <td class="wrapper" style="font-family: Helvetica, sans-serif; font-size: 16px; vertical-align: top; box-sizing: border-box; padding: 24px;" valign="top">
                            <p style="font-family: Helvetica, sans-serif; font-size: 16px; font-weight: normal; margin: 0; margin-bottom: 16px;">Halo!</p>
                            <p style="font-family: Helvetica, sans-serif; font-size: 16px; font-weight: normal; margin: 0; margin-bottom: 16px;">Batas akhir peminjaman buku kamu sampai dengan tanggal %s, harap dikembalikan sebelum batas pengembalian yang sudah ditentukan.</p>
                            <p style="font-family: Helvetica, sans-serif; font-size: 16px; font-weight: normal; margin: 0; margin-bottom: 20px;"><b>Batas Pengembalian</b>: %s.</p>
                            <a href="" style="height: 30px; width: 100px; background-color: blue; padding: 10px; color: white; text-align: center; text-decoration: none; border-radius: 5px; font-family: Arial, sans-serif; outline: none;">Lihat Detail Transaksi</a>
                        </td>
                    </tr>

                    <!-- END MAIN CONTENT AREA -->
                </table>

                <!-- START FOOTER -->
                <div class="footer" style="clear: both; padding-top: 24px; text-align: center; width: 100%;">
                    <table role="presentation" border="0" cellpadding="0" cellspacing="0" style="border-collapse: separate; mso-table-lspace: 0pt; mso-table-rspace: 0pt; width: 100%;" width="100%">
                        <tr>
                            <td class="content-block" style="font-family: Helvetica, sans-serif; vertical-align: top; color: #9a9ea6; font-size: 16px; text-align: center;" valign="top" align="center">
                                <span class="apple-link" style="color: #9a9ea6; font-size: 16px; text-align: center;">Smart Library - Pengagum Rahasia, 2024.</span>
                            </td>
                        </tr>
                    </table>
                </div>

                <!-- END FOOTER -->

                <!-- END CENTERED WHITE CONTAINER -->
            </div>
        </td>
        <td style="font-family: Helvetica, sans-serif; font-size: 16px; vertical-align: top;" valign="top">&nbsp;</td>
    </tr>
</table>
</body>
</html>
`, notification.DueDate, notification.DueDate)

	return body
}
