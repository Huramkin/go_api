package dao

import "go_api/domain"

func FindAll() []domain.ServerInstance {
	db := CreateDbcore()
	var nodes []domain.ServerInstance
	db.Find(&nodes)
	return nodes
}

func FindById(id int) domain.ServerInstance {
	db := CreateDbcore()
	var node domain.ServerInstance
	db.Find(&node, "id =?", id)
	return node
}
