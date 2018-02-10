package logger

import (
	"fmt"
	"time"

	//"github.com/aybabtme/color/brush"
	"github.com/go-redis/redis"
)

//redis

const (
	Log_Key = "log_key"
)

// This is the standard writer that prints to standard output.
type RedisLogWriter chan *LogRecord

func NewRedisLogWriter(redis *redis.Client) (RedisLogWriter, error) {
	records := make(RedisLogWriter, LogBufferLength)
	_, err := redis.Ping().Result()
	if err != nil {
		return nil, err
	}
	go records.run(redis)
	return records, nil
}

func NewRedisLogWriterFrom(redis_str, pass string) (RedisLogWriter, error) {
	records := make(RedisLogWriter, LogBufferLength)
	client := redis.NewClient(&redis.Options{
		Addr: redis_str,
		DB:   1, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	go records.run(client)
	return records, nil
}

func (w RedisLogWriter) run(redis *redis.Client) {
	var timestr string
	var timestrAt int64

	for rec := range w {
		if at := rec.Created.UnixNano() / 1e9; at != timestrAt {
			timestr, timestrAt = rec.Created.Format("2006/01/02 15:04:05"), at
		}
		message := fmt.Sprintf("[%s] [%s] (%s) %s\n",
			timestr,
			levelStrings[rec.Level],
			/*fmt.Sprintf("%s", brush.LightGray(rec.Source)),
			fmt.Sprintf("%s", brush.DarkBlue(rec.Message)),*/
			fmt.Sprintf("%s", rec.Source),
			fmt.Sprintf("%s", rec.Message),
		)
		//TODO
		redis.RPush(Log_Key, message)
	}
}

// This is the ConsoleLogWriter's output method. This will block if the output
// buffer is full.
func (w RedisLogWriter) LogWrite(rec *LogRecord) {
	w <- rec
}

// Close stops the logger from sending messages to standard output. Attempts to
// send log messages to this logger after a Close have undefined behavior.
func (w RedisLogWriter) Close() {
	close(w)
	time.Sleep(50 * time.Millisecond) // Try to give console I/O time to complete
}
