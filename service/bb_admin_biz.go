package service

import (
	"github.com/bensema/goadmin/model"
	"github.com/bensema/library/biz"
	"github.com/gin-gonic/gin"
)

func (s *Service) FindAdvertisePage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertisePages)
}

func (s *Service) AdvertiseAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseAdd)
}

func (s *Service) AdvertiseDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseDel)
}

func (s *Service) AdvertiseQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseQuery)
}

func (s *Service) AdvertiseUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AdvertiseUpdate)
}

func (s *Service) FindAnnouncementsPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementsPages)
}

func (s *Service) AnnouncementsAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementsAdd)
}

func (s *Service) AnnouncementsDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementsDel)
}

func (s *Service) AnnouncementsQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementsQuery)
}

func (s *Service) AnnouncementsUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AnnouncementsUpdate)
}

func (s *Service) FindAuthPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AuthPages)
}

func (s *Service) AuthAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AuthAdd)
}

func (s *Service) AuthDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AuthDel)
}

func (s *Service) AuthQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AuthQuery)
}

func (s *Service) AuthUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.AuthUpdate)
}

func (s *Service) FindCoinPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinPages)
}

func (s *Service) CoinAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinAdd)
}

func (s *Service) CoinDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinDel)
}

func (s *Service) CoinQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinQuery)
}

func (s *Service) CoinUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinUpdate)
}

func (s *Service) FindCoinChainPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinChainPages)
}

func (s *Service) CoinChainAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinChainAdd)
}

func (s *Service) CoinChainDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinChainDel)
}

func (s *Service) CoinChainQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinChainQuery)
}

func (s *Service) CoinChainUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.CoinChainUpdate)
}

func (s *Service) FindGamePage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GamePages)
}

func (s *Service) GameAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameAdd)
}

func (s *Service) GameDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameDel)
}

func (s *Service) GameQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameQuery)
}

func (s *Service) GameUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameUpdate)
}

func (s *Service) FindGameGroupPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameGroupPages)
}

func (s *Service) GameGroupAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameGroupAdd)
}

func (s *Service) GameGroupDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameGroupDel)
}

func (s *Service) GameGroupQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameGroupQuery)
}

func (s *Service) GameGroupUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameGroupUpdate)
}

func (s *Service) FindGameTypePage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameTypePages)
}

func (s *Service) GameTypeAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameTypeAdd)
}

func (s *Service) GameTypeDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameTypeDel)
}

func (s *Service) GameTypeQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameTypeQuery)
}

func (s *Service) GameTypeUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GameTypeUpdate)
}

func (s *Service) FindGradePage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GradePages)
}

func (s *Service) GradeAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GradeAdd)
}

func (s *Service) GradeDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GradeDel)
}

func (s *Service) GradeQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GradeQuery)
}

func (s *Service) GradeUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.GradeUpdate)
}

func (s *Service) FindLogOperatePage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogOperatePages)
}

func (s *Service) LogOperateAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogOperateAdd)
}

func (s *Service) LogOperateDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogOperateDel)
}

func (s *Service) LogOperateQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogOperateQuery)
}

func (s *Service) LogOperateUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogOperateUpdate)
}

func (s *Service) FindLogUserLoginPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogUserLoginPages)
}

func (s *Service) LogUserLoginAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogUserLoginAdd)
}

func (s *Service) LogUserLoginDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogUserLoginDel)
}

func (s *Service) LogUserLoginQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogUserLoginQuery)
}

func (s *Service) LogUserLoginUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.LogUserLoginUpdate)
}

func (s *Service) FindPowerPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.PowerPages)
}

func (s *Service) PowerAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.PowerAdd)
}

func (s *Service) PowerDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.PowerDel)
}

func (s *Service) PowerQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.PowerQuery)
}

func (s *Service) PowerUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.PowerUpdate)
}

func (s *Service) FindSecurityQuestionsPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.SecurityQuestionsPages)
}

func (s *Service) SecurityQuestionsAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.SecurityQuestionsAdd)
}

func (s *Service) SecurityQuestionsDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.SecurityQuestionsDel)
}

func (s *Service) SecurityQuestionsQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.SecurityQuestionsQuery)
}

func (s *Service) SecurityQuestionsUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.SecurityQuestionsUpdate)
}

func (s *Service) FindUserInfoPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserInfoPages)
}

func (s *Service) UserInfoAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserInfoAdd)
}

func (s *Service) UserInfoDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserInfoDel)
}

func (s *Service) UserInfoQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserInfoQuery)
}

func (s *Service) UserInfoUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserInfoUpdate)
}

func (s *Service) FindUserPowerPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserPowerPages)
}

func (s *Service) UserPowerAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserPowerAdd)
}

func (s *Service) UserPowerDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserPowerDel)
}

func (s *Service) UserPowerQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserPowerQuery)
}

func (s *Service) UserPowerUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserPowerUpdate)
}

func (s *Service) FindUserSecurityQuestionsPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserSecurityQuestionsPages)
}

func (s *Service) UserSecurityQuestionsAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserSecurityQuestionsAdd)
}

func (s *Service) UserSecurityQuestionsDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserSecurityQuestionsDel)
}

func (s *Service) UserSecurityQuestionsQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserSecurityQuestionsQuery)
}

func (s *Service) UserSecurityQuestionsUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserSecurityQuestionsUpdate)
}

func (s *Service) FindUserTokenPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserTokenPages)
}

func (s *Service) UserTokenAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserTokenAdd)
}

func (s *Service) UserTokenDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserTokenDel)
}

func (s *Service) UserTokenQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserTokenQuery)
}

func (s *Service) UserTokenUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserTokenUpdate)
}

func (s *Service) FindUserWalletPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserWalletPages)
}

func (s *Service) UserWalletAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserWalletAdd)
}

func (s *Service) UserWalletDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserWalletDel)
}

func (s *Service) UserWalletQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserWalletQuery)
}

func (s *Service) UserWalletUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UserWalletUpdate)
}

func (s *Service) FindUsersPage(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UsersPages)
}

func (s *Service) UsersAdd(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UsersAdd)
}

func (s *Service) UsersDel(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UsersDel)
}

func (s *Service) UsersQuery(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UsersQuery)
}

func (s *Service) UsersUpdate(c *gin.Context) (reply *model.AdminApiReply, err error) {
	return s.doRequest(c, biz.UsersUpdate)
}
