package currtime

import (
	"fmt"
	"sync"
	"time"
)

var locMutex sync.RWMutex
var locCache map[string]*time.Location

const DefaultTimezone string = "UTC"
const packageName string = "currtime"

func init() {
	locMutex = sync.RWMutex{}
	locCache = map[string]*time.Location{}

	loc, err := time.LoadLocation(DefaultTimezone)
	if err != nil {
		panic(packageName + ": cannot load timezone \"" + DefaultTimezone + "\" with package \"time\": " + fmt.Sprint(err))
	}

	locCache[DefaultTimezone] = loc
}

func getLocation(timezone string) (*time.Location, error) {
	locMutex.RLock()
	loc, exist := locCache[timezone]
	locMutex.RUnlock()

	if exist {
		return loc, nil
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}

	locMutex.Lock()
	locCache[timezone] = loc
	locMutex.Unlock()

	return loc, nil
}

func GetTime(timezone string) (time.Time, error) {
	loc, err := getLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().In(loc), nil
}
