package models

import (
	"ahutoj/web/dao"
	mysqldao "ahutoj/web/dao/mysqlDao"
	"ahutoj/web/utils"

	"github.com/gin-gonic/gin"
)

func IsListExistByLID(ctx *gin.Context, list *dao.List) bool {
	count, err := mysqldao.SelectListCountByLID(ctx, list.LID)
	if err != nil {
		return false
	}
	return count > 0
}

func CreateList(ctx *gin.Context, list *dao.List) error {
	logger := utils.GetLogInstance()
	if IsListExistByLID(ctx, list) {
		return nil
	}
	err := mysqldao.InsertTraning(ctx, *list)
	if err != nil {
		logger.Errorf("call InsertListTable failed,list= %+v, err=%s", utils.Sdump(list), err.Error())
	}
	return err
}

func CreateListProblem(ctx *gin.Context, listproblem *dao.ListProblem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.InsertListProblem(ctx, *listproblem)
	if err != nil {
		logger.Errorf("call InsertListProblemTable failed,listproblem= %+v,err=%s", utils.Sdump(listproblem), err.Error())
	}
	return err
}

func EditList(ctx *gin.Context, list *dao.List) error {
	logger := utils.GetLogInstance()
	err := mysqldao.UpdateTraning(ctx, *list)
	if err != nil {
		logger.Errorf("call EditListTable failed,list= %+v,err=%s", utils.Sdump(list), err.Error())
	}
	return err
}
func EditListProblem(ctx *gin.Context, listproblem *dao.ListProblem) error {
	logger := utils.GetLogInstance()
	err := mysqldao.UpdateListProblem(ctx, *listproblem)
	if err != nil {
		logger.Errorf("call EditListProblemTable failed,listproblem= %+v,err=%s", utils.Sdump(listproblem), err.Error())
	}
	return err
}
func DeleteTraining(ctx *gin.Context, list *dao.List) error {
	logger := utils.GetLogInstance()
	err := mysqldao.DeleteTraning(ctx, list.LID)
	if err != nil {
		logger.Errorf("call DeleteListTable failed,listproblem= %+v,err=%s", utils.Sdump(list), err.Error())
	}
	return err
}
func GetTrainingList(ctx *gin.Context, offset, pagesize int) ([]dao.List, error) {
	return mysqldao.GetTrainingList(ctx, offset, pagesize)
}

func GetCurrentLID(ctx *gin.Context, list dao.List) (int64, error) {
	return mysqldao.SelectListByUID(ctx, list.UID)
}
