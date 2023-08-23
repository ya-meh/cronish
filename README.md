# Cronish

The main reason of this micro-library is to determine the closest time, when then given restrictions apply.

Goal: have similar syntax to `time.Parse`. `cronish.Parse("Mon 15:04", "Fri 23:59")` gets you closest `time.Time`, which is
*Friday 23:59* after now. Realisation is a little different from a dream.

As a bonus, supports Russian days of the week. And simply extendable in this regard.

## Core functions

> ### TD;DR
> Just use `Parse(Time("15:04", 19:00), Day("Mon", "Fri"))` to get the closest `time.Time`, which is Friday 19:00.

* `Time(layout, value string) `
* `Date(layout, value string)`
* `Day(layout, value string)` - to specify the day of the week
* `Parse(options...)`

To simplify the use even further the library comes with `...Try` alternative functions, that allow to specify just the
restriction and let the library guess which layout you meant.

## Example

```golang
package main

import (
	"github.com/ya-meh/cronish"
	"time"
)

func main() {
	println(cronish.Parse(cronish.DayTry("Monday"), cronish.TimeTry("23:59")).String())
	
	certain := cronish.New(
		cronish.Time("15:04:05", "00:10:05"),
		cronish.Date("02.01", "29.02"))

	println(certain.Time().Format(time.DateTime))
}

```

## Performance

The library is written in an evening. Keep it stupid simple. Thanks GOd it performs.

* Time only - `4045178	       277.9 ns/op` (this and below were run on Apple M1 PRO)
* Date only - `3828038	       300.2 ns/op`
* Time and date - `2276107	       508.6 ns/op`
