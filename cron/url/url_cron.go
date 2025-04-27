package urlcron

import (
	"fmt"
	"time"
	"urlshortener/api/url"

	"github.com/robfig/cron/v3"
)

func DeleteExpiredUrls() {
	fmt.Println(time.Now().Format("2006-01-02T15:04:05") + " - Cron to delete expired URLs")
	urlController := url.MakeUrlController()

	cron := cron.New()
	id, _ := cron.AddFunc("@daily", func() {
		msg, err := urlController.DeleteExpiredUrlsWithCron()
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02T15:04:05") + " - Error deleting expired URLs: " + err.Error())
			return
		}

		fmt.Println(time.Now().Format("2006-01-02T15:04:05") + " - " + msg)
	})

	cron.Entry(id).Job.Run() // This will run the job immediately

	go cron.Start()
}
