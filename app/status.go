package main

var statuses map[string]bool

func InitStatuses() {
	cfg := GetConfig()
	statuses = make(map[string]bool)

	for _, v := range cfg.General.Target {
		statuses[v] = true
	}
}

func IsTargetUp(target string) bool {
	return statuses[target]
}

func UpTarget(target string) {
	statuses[target] = true
}

func DownTarget(target string) {
	statuses[target] = false
}
