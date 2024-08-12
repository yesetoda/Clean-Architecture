package controller

type GeneralControllerInterface interface {
	HandleGetTaskById(struct{})
	HandleCreateTask(struct{})
	HandleUpdateTask(struct{})
	HandleDeleteTask(struct{})
	HandleFilterTask(struct{})
	
	HandleGetAllUsers(struct{})
	HandleCreateUser(struct{})
	HandlePromote(struct{})
	HandleDeleteUser(struct{})
	HandleFilterUser(struct{})
	HandleLogin(struct{})
}
