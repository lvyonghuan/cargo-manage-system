package src

import (
	"cargo-manage-system/dao"
	"fmt"
)

func GetAutoOrderList() ([]string, error) {
	return dao.GetAutoOrderList()
}

func AddAutoOrder(upcCode string, storeID, limit, restockAmount int) error {
	//生成触发器sql语句
	triggerBody := fmt.Sprintf(
		"DECLARE existing_task INT; "+
			"SELECT COUNT(*) INTO existing_task FROM Restocking_Tasks WHERE upc_code = NEW.upc_code AND store_id = NEW.store_id; "+
			"IF NEW.store_id = %d AND NEW.upc_code = '%s' AND NEW.inventory_amount < %d AND existing_task = 0 THEN "+
			"INSERT INTO Restocking_Tasks (upc_code, store_id, restock_amount) "+
			"VALUES (NEW.upc_code, NEW.store_id, %d); "+
			"END IF;",
		storeID, upcCode, limit, restockAmount)

	triggerCreation := fmt.Sprintf(
		"CREATE TRIGGER ordering_%s_%d BEFORE UPDATE ON Product_Store FOR EACH ROW BEGIN %s END",
		upcCode, storeID, triggerBody)

	return dao.AddAutoOrder(triggerCreation)
}

func StopAutoOrder(upcCode string, storeID int) error {
	sql := fmt.Sprintf("DROP TRIGGER ordering_%s_%d", upcCode, storeID)
	return dao.StopAutoOrder(sql)
}
