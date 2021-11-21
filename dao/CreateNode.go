package dao

import "go_api/domain"

func AddNode(node domain.ServerInstance) {
	db := CreateDbcore()
	db.Create(&node)
}

func UpdateNodeById(node domain.ServerInstance) {
	db := CreateDbcore()
	if node.Id != 0 {
		db.Model(&node).Where("1=1").Updates(node)
	}
}
