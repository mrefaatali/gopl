package storage1

var usage = make(map[string]int64)

func bytesInUse(username string) int64 {
  return usage[username]
}

const sender = "notification@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

