package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/taybart/args"
)

var (
	app = args.App{
		Name:    "unix-timestamp",
		Version: "v0.0.1",
		Args: map[string]*args.Arg{
			"unix": {
				Short:   "u",
				Help:    "unix timestamp to convert",
				Default: -1,
			},
			"diff": {
				Short:   "d",
				Help:    "print difference to now (ex. 3min from now/2 days ago)",
				Default: -1,
			},
			"timestamp": {
				Short: "ts",
				Help:  "RFC3339 timestamp to convert to a unix timestamp",
			},
			// // TODO: add format string
			// "timezone": {
			// 	Short:    "tz",
			// 	Required: false,
			// 	Default:  time.UTC,
			// },
		},
	}
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
func run() error {

	if err := app.Parse(); err != nil {
		if errors.Is(err, args.ErrUsageRequested) {
			return nil
		}
		return err
	}

	// TODO: check timezone flag
	if app.UserSet("unix") {
		ts := time.Unix(int64(app.Int("unix")), 0)
		fmt.Println(ts.UTC())
		return nil
	}
	if app.UserSet("diff") {
		ts := time.Unix(int64(app.Int("diff")), 0)
		now := time.Now()
		if ts.After(now) {
			fmt.Println(time.Until(ts), "from now")
		} else {
			fmt.Println(time.Since(ts), "ago")
		}
		return nil
	}
	if app.UserSet("timestamp") {
		ts, err := time.Parse(time.RFC3339, app.String("timestamp"))
		if err != nil {
			return err
		}
		fmt.Println(ts.Unix())
		return nil
	}
	if len(os.Args) > 1 {
		i, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			return err
		}
		ts := time.Unix(i, 0)
		fmt.Println(ts.UTC())
		return nil
	}
	fmt.Println(time.Now().Unix())
	return nil
}
