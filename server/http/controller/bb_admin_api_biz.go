package controller

import (
	"github.com/bensema/goadmin/global"
	"github.com/bensema/goadmin/server/http/internal"
	"github.com/gin-gonic/gin"
)

func (_this *ApiAuth) RegisterBBAdminRoute(g *gin.RouterGroup) {

	g.GET("/api/v1/advertise/pages", _this.advertisePages)
	g.POST("/api/v1/advertise/add", _this.advertiseAdd)
	g.POST("/api/v1/advertise/del", _this.advertiseDel)
	g.GET("/api/v1/advertise/query", _this.advertiseQuery)
	g.POST("/api/v1/advertise/update", _this.advertiseUpdate)

	g.GET("/api/v1/announcements/pages", _this.announcementsPages)
	g.POST("/api/v1/announcements/add", _this.announcementsAdd)
	g.POST("/api/v1/announcements/del", _this.announcementsDel)
	g.GET("/api/v1/announcements/query", _this.announcementsQuery)
	g.POST("/api/v1/announcements/update", _this.announcementsUpdate)

	g.GET("/api/v1/auth/pages", _this.authPages)
	g.POST("/api/v1/auth/add", _this.authAdd)
	g.POST("/api/v1/auth/del", _this.authDel)
	g.GET("/api/v1/auth/query", _this.authQuery)
	g.POST("/api/v1/auth/update", _this.authUpdate)

	g.GET("/api/v1/coin/pages", _this.coinPages)
	g.POST("/api/v1/coin/add", _this.coinAdd)
	g.POST("/api/v1/coin/del", _this.coinDel)
	g.GET("/api/v1/coin/query", _this.coinQuery)
	g.POST("/api/v1/coin/update", _this.coinUpdate)

	g.GET("/api/v1/coin_chain/pages", _this.coinChainPages)
	g.POST("/api/v1/coin_chain/add", _this.coinChainAdd)
	g.POST("/api/v1/coin_chain/del", _this.coinChainDel)
	g.GET("/api/v1/coin_chain/query", _this.coinChainQuery)
	g.POST("/api/v1/coin_chain/update", _this.coinChainUpdate)

	g.GET("/api/v1/game/pages", _this.gamePages)
	g.POST("/api/v1/game/add", _this.gameAdd)
	g.POST("/api/v1/game/del", _this.gameDel)
	g.GET("/api/v1/game/query", _this.gameQuery)
	g.POST("/api/v1/game/update", _this.gameUpdate)

	g.GET("/api/v1/game_group/pages", _this.gameGroupPages)
	g.POST("/api/v1/game_group/add", _this.gameGroupAdd)
	g.POST("/api/v1/game_group/del", _this.gameGroupDel)
	g.GET("/api/v1/game_group/query", _this.gameGroupQuery)
	g.POST("/api/v1/game_group/update", _this.gameGroupUpdate)

	g.GET("/api/v1/game_type/pages", _this.gameTypePages)
	g.POST("/api/v1/game_type/add", _this.gameTypeAdd)
	g.POST("/api/v1/game_type/del", _this.gameTypeDel)
	g.GET("/api/v1/game_type/query", _this.gameTypeQuery)
	g.POST("/api/v1/game_type/update", _this.gameTypeUpdate)

	g.GET("/api/v1/grade/pages", _this.gradePages)
	g.POST("/api/v1/grade/add", _this.gradeAdd)
	g.POST("/api/v1/grade/del", _this.gradeDel)
	g.GET("/api/v1/grade/query", _this.gradeQuery)
	g.POST("/api/v1/grade/update", _this.gradeUpdate)

	g.GET("/api/v1/log_operate/pages", _this.logOperatePages)
	g.POST("/api/v1/log_operate/add", _this.logOperateAdd)
	g.POST("/api/v1/log_operate/del", _this.logOperateDel)
	g.GET("/api/v1/log_operate/query", _this.logOperateQuery)
	g.POST("/api/v1/log_operate/update", _this.logOperateUpdate)

	g.GET("/api/v1/log_user_login/pages", _this.logUserLoginPages)
	g.POST("/api/v1/log_user_login/add", _this.logUserLoginAdd)
	g.POST("/api/v1/log_user_login/del", _this.logUserLoginDel)
	g.GET("/api/v1/log_user_login/query", _this.logUserLoginQuery)
	g.POST("/api/v1/log_user_login/update", _this.logUserLoginUpdate)

	g.GET("/api/v1/power/pages", _this.powerPages)
	g.POST("/api/v1/power/add", _this.powerAdd)
	g.POST("/api/v1/power/del", _this.powerDel)
	g.GET("/api/v1/power/query", _this.powerQuery)
	g.POST("/api/v1/power/update", _this.powerUpdate)

	g.GET("/api/v1/security_questions/pages", _this.securityQuestionsPages)
	g.POST("/api/v1/security_questions/add", _this.securityQuestionsAdd)
	g.POST("/api/v1/security_questions/del", _this.securityQuestionsDel)
	g.GET("/api/v1/security_questions/query", _this.securityQuestionsQuery)
	g.POST("/api/v1/security_questions/update", _this.securityQuestionsUpdate)

	g.GET("/api/v1/user_info/pages", _this.userInfoPages)
	g.POST("/api/v1/user_info/add", _this.userInfoAdd)
	g.POST("/api/v1/user_info/del", _this.userInfoDel)
	g.GET("/api/v1/user_info/query", _this.userInfoQuery)
	g.POST("/api/v1/user_info/update", _this.userInfoUpdate)

	g.GET("/api/v1/user_power/pages", _this.userPowerPages)
	g.POST("/api/v1/user_power/add", _this.userPowerAdd)
	g.POST("/api/v1/user_power/del", _this.userPowerDel)
	g.GET("/api/v1/user_power/query", _this.userPowerQuery)
	g.POST("/api/v1/user_power/update", _this.userPowerUpdate)

	g.GET("/api/v1/user_security_questions/pages", _this.userSecurityQuestionsPages)
	g.POST("/api/v1/user_security_questions/add", _this.userSecurityQuestionsAdd)
	g.POST("/api/v1/user_security_questions/del", _this.userSecurityQuestionsDel)
	g.GET("/api/v1/user_security_questions/query", _this.userSecurityQuestionsQuery)
	g.POST("/api/v1/user_security_questions/update", _this.userSecurityQuestionsUpdate)

	g.GET("/api/v1/user_token/pages", _this.userTokenPages)
	g.POST("/api/v1/user_token/add", _this.userTokenAdd)
	g.POST("/api/v1/user_token/del", _this.userTokenDel)
	g.GET("/api/v1/user_token/query", _this.userTokenQuery)
	g.POST("/api/v1/user_token/update", _this.userTokenUpdate)

	g.GET("/api/v1/user_wallet/pages", _this.userWalletPages)
	g.POST("/api/v1/user_wallet/add", _this.userWalletAdd)
	g.POST("/api/v1/user_wallet/del", _this.userWalletDel)
	g.GET("/api/v1/user_wallet/query", _this.userWalletQuery)
	g.POST("/api/v1/user_wallet/update", _this.userWalletUpdate)

	g.GET("/api/v1/users/pages", _this.usersPages)
	g.POST("/api/v1/users/add", _this.usersAdd)
	g.POST("/api/v1/users/del", _this.usersDel)
	g.GET("/api/v1/users/query", _this.usersQuery)
	g.POST("/api/v1/users/update", _this.usersUpdate)

}

func (_this *ApiAuth) advertisePages(c *gin.Context) {
	reply, err := global.Srv.FindAdvertisePage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) advertiseAdd(c *gin.Context) {
	reply, err := global.Srv.AdvertiseAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) advertiseDel(c *gin.Context) {
	reply, err := global.Srv.AdvertiseDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) advertiseQuery(c *gin.Context) {
	reply, err := global.Srv.AdvertiseQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) advertiseUpdate(c *gin.Context) {
	reply, err := global.Srv.AdvertiseUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) announcementsPages(c *gin.Context) {
	reply, err := global.Srv.FindAnnouncementsPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) announcementsAdd(c *gin.Context) {
	reply, err := global.Srv.AnnouncementsAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) announcementsDel(c *gin.Context) {
	reply, err := global.Srv.AnnouncementsDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) announcementsQuery(c *gin.Context) {
	reply, err := global.Srv.AnnouncementsQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) announcementsUpdate(c *gin.Context) {
	reply, err := global.Srv.AnnouncementsUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) authPages(c *gin.Context) {
	reply, err := global.Srv.FindAuthPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) authAdd(c *gin.Context) {
	reply, err := global.Srv.AuthAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) authDel(c *gin.Context) {
	reply, err := global.Srv.AuthDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) authQuery(c *gin.Context) {
	reply, err := global.Srv.AuthQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) authUpdate(c *gin.Context) {
	reply, err := global.Srv.AuthUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinPages(c *gin.Context) {
	reply, err := global.Srv.FindCoinPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinAdd(c *gin.Context) {
	reply, err := global.Srv.CoinAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinDel(c *gin.Context) {
	reply, err := global.Srv.CoinDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinQuery(c *gin.Context) {
	reply, err := global.Srv.CoinQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinUpdate(c *gin.Context) {
	reply, err := global.Srv.CoinUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinChainPages(c *gin.Context) {
	reply, err := global.Srv.FindCoinChainPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinChainAdd(c *gin.Context) {
	reply, err := global.Srv.CoinChainAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinChainDel(c *gin.Context) {
	reply, err := global.Srv.CoinChainDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinChainQuery(c *gin.Context) {
	reply, err := global.Srv.CoinChainQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) coinChainUpdate(c *gin.Context) {
	reply, err := global.Srv.CoinChainUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gamePages(c *gin.Context) {
	reply, err := global.Srv.FindGamePage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameAdd(c *gin.Context) {
	reply, err := global.Srv.GameAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameDel(c *gin.Context) {
	reply, err := global.Srv.GameDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameQuery(c *gin.Context) {
	reply, err := global.Srv.GameQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameUpdate(c *gin.Context) {
	reply, err := global.Srv.GameUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameGroupPages(c *gin.Context) {
	reply, err := global.Srv.FindGameGroupPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameGroupAdd(c *gin.Context) {
	reply, err := global.Srv.GameGroupAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameGroupDel(c *gin.Context) {
	reply, err := global.Srv.GameGroupDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameGroupQuery(c *gin.Context) {
	reply, err := global.Srv.GameGroupQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameGroupUpdate(c *gin.Context) {
	reply, err := global.Srv.GameGroupUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameTypePages(c *gin.Context) {
	reply, err := global.Srv.FindGameTypePage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameTypeAdd(c *gin.Context) {
	reply, err := global.Srv.GameTypeAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameTypeDel(c *gin.Context) {
	reply, err := global.Srv.GameTypeDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameTypeQuery(c *gin.Context) {
	reply, err := global.Srv.GameTypeQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gameTypeUpdate(c *gin.Context) {
	reply, err := global.Srv.GameTypeUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gradePages(c *gin.Context) {
	reply, err := global.Srv.FindGradePage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gradeAdd(c *gin.Context) {
	reply, err := global.Srv.GradeAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gradeDel(c *gin.Context) {
	reply, err := global.Srv.GradeDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gradeQuery(c *gin.Context) {
	reply, err := global.Srv.GradeQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) gradeUpdate(c *gin.Context) {
	reply, err := global.Srv.GradeUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logOperatePages(c *gin.Context) {
	reply, err := global.Srv.FindLogOperatePage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logOperateAdd(c *gin.Context) {
	reply, err := global.Srv.LogOperateAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logOperateDel(c *gin.Context) {
	reply, err := global.Srv.LogOperateDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logOperateQuery(c *gin.Context) {
	reply, err := global.Srv.LogOperateQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logOperateUpdate(c *gin.Context) {
	reply, err := global.Srv.LogOperateUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logUserLoginPages(c *gin.Context) {
	reply, err := global.Srv.FindLogUserLoginPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logUserLoginAdd(c *gin.Context) {
	reply, err := global.Srv.LogUserLoginAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logUserLoginDel(c *gin.Context) {
	reply, err := global.Srv.LogUserLoginDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logUserLoginQuery(c *gin.Context) {
	reply, err := global.Srv.LogUserLoginQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) logUserLoginUpdate(c *gin.Context) {
	reply, err := global.Srv.LogUserLoginUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) powerPages(c *gin.Context) {
	reply, err := global.Srv.FindPowerPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) powerAdd(c *gin.Context) {
	reply, err := global.Srv.PowerAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) powerDel(c *gin.Context) {
	reply, err := global.Srv.PowerDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) powerQuery(c *gin.Context) {
	reply, err := global.Srv.PowerQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) powerUpdate(c *gin.Context) {
	reply, err := global.Srv.PowerUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) securityQuestionsPages(c *gin.Context) {
	reply, err := global.Srv.FindSecurityQuestionsPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) securityQuestionsAdd(c *gin.Context) {
	reply, err := global.Srv.SecurityQuestionsAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) securityQuestionsDel(c *gin.Context) {
	reply, err := global.Srv.SecurityQuestionsDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) securityQuestionsQuery(c *gin.Context) {
	reply, err := global.Srv.SecurityQuestionsQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) securityQuestionsUpdate(c *gin.Context) {
	reply, err := global.Srv.SecurityQuestionsUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userInfoPages(c *gin.Context) {
	reply, err := global.Srv.FindUserInfoPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userInfoAdd(c *gin.Context) {
	reply, err := global.Srv.UserInfoAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userInfoDel(c *gin.Context) {
	reply, err := global.Srv.UserInfoDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userInfoQuery(c *gin.Context) {
	reply, err := global.Srv.UserInfoQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userInfoUpdate(c *gin.Context) {
	reply, err := global.Srv.UserInfoUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userPowerPages(c *gin.Context) {
	reply, err := global.Srv.FindUserPowerPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userPowerAdd(c *gin.Context) {
	reply, err := global.Srv.UserPowerAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userPowerDel(c *gin.Context) {
	reply, err := global.Srv.UserPowerDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userPowerQuery(c *gin.Context) {
	reply, err := global.Srv.UserPowerQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userPowerUpdate(c *gin.Context) {
	reply, err := global.Srv.UserPowerUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userSecurityQuestionsPages(c *gin.Context) {
	reply, err := global.Srv.FindUserSecurityQuestionsPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userSecurityQuestionsAdd(c *gin.Context) {
	reply, err := global.Srv.UserSecurityQuestionsAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userSecurityQuestionsDel(c *gin.Context) {
	reply, err := global.Srv.UserSecurityQuestionsDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userSecurityQuestionsQuery(c *gin.Context) {
	reply, err := global.Srv.UserSecurityQuestionsQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userSecurityQuestionsUpdate(c *gin.Context) {
	reply, err := global.Srv.UserSecurityQuestionsUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userTokenPages(c *gin.Context) {
	reply, err := global.Srv.FindUserTokenPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userTokenAdd(c *gin.Context) {
	reply, err := global.Srv.UserTokenAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userTokenDel(c *gin.Context) {
	reply, err := global.Srv.UserTokenDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userTokenQuery(c *gin.Context) {
	reply, err := global.Srv.UserTokenQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userTokenUpdate(c *gin.Context) {
	reply, err := global.Srv.UserTokenUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userWalletPages(c *gin.Context) {
	reply, err := global.Srv.FindUserWalletPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userWalletAdd(c *gin.Context) {
	reply, err := global.Srv.UserWalletAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userWalletDel(c *gin.Context) {
	reply, err := global.Srv.UserWalletDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userWalletQuery(c *gin.Context) {
	reply, err := global.Srv.UserWalletQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) userWalletUpdate(c *gin.Context) {
	reply, err := global.Srv.UserWalletUpdate(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) usersPages(c *gin.Context) {
	reply, err := global.Srv.FindUsersPage(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) usersAdd(c *gin.Context) {
	reply, err := global.Srv.UsersAdd(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) usersDel(c *gin.Context) {
	reply, err := global.Srv.UsersDel(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) usersQuery(c *gin.Context) {
	reply, err := global.Srv.UsersQuery(c)
	internal.AdminJSON(c, reply, err)
}

func (_this *ApiAuth) usersUpdate(c *gin.Context) {
	reply, err := global.Srv.UsersUpdate(c)
	internal.AdminJSON(c, reply, err)
}
