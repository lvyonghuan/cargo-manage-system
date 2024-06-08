package dao

func GetAutoOrderList() ([]string, error) {
	var triggerNames []string
	err := db.Raw("SELECT TRIGGER_NAME FROM information_schema.TRIGGERS WHERE TRIGGER_NAME LIKE 'ordering%';").Scan(&triggerNames).Error
	return triggerNames, err
}

func AddAutoOrder(trigger string) error {
	return db.Exec(trigger).Error
}

func StopAutoOrder(sql string) error {
	return db.Exec(sql).Error
}
