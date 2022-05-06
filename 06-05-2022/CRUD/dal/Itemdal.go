package dal

import (
	"go-store/CRUD/model"

)

//InsertItem to product
func InsertItem(item model.Item) (int64, int64, error) {
	GetConnection()
	sqlQuery := "INSERT items SET item_name=?, unit_price=?, amount=?, item_status=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(item.ItemName, item.UnitPrice, item.Amount, item.ItemStatus)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertId, err := res.LastInsertId()
	return rowsAffected, lastInsertId, err
}

//UpdateItem in product
func UpdateItem(item model.Item) (int64, error) {
	GetConnection()
	sqlQuery := "UPDATE items SET item_name=?, unit_price=?, amount=?, item_status=? WHERE item_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(item.ItemName, item.UnitPrice, item.Amount, item.ItemStatus, item.ItemId)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//DeleteItem in product
func DeleteItem(itemId int64) (int64, error) {
	GetConnection()
	sqlQuery := "DELETE FROM items WHERE item_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(itemId)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

//GetItem from itemId
func GetItem(itemId int64) (model.Item, error) {
	sqlQuery := "SELECT item_id, item_name, unit_price, amount, item_status, item_description FROM items WHERE item_id = ?"
	GetConnection()
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var item model.Item
	if err != nil {
		return item, err
	}
	res, err := stmt.Query(itemId)
	defer CloseRows(res)
	if err != nil {
		return item, err
	}
	if res.Next() {
		res.Scan(&item.ItemId, &item.ItemName, &item.UnitPrice, &item.Amount, &item.ItemStatus, &item.ItemDescription)
	}
	return item, err
}