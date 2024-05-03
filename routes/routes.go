package routes

import (
	"audiohub/controller"

	"github.com/gin-gonic/gin"
)

// Router creates and configures the Gin router.
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(controller.CorsMiddleware())
	
	router.POST("/tokenvalidity",controller.ValidateToken)
	router.POST("/login", controller.Login)
	router.POST("/signupwithgoogle", controller.SignUpWithGoogle)
	router.POST("/register", controller.CreateProfile)
	router.POST("/verify", controller.VerifyEmail)
	router.POST("/forgetpassword", controller.ForgetPassword)
	router.POST("/passwordchange", controller.PasswordChange)
	router.POST("/getuserbyid", controller.UserDetails)

	// History Routes
	router.POST("/savehistory",controller.SaveToHistory)
	router.POST("/listhistory",controller.DisplayHistory)
	router.POST("/deletehistory",controller.DeleteHistory)
	router.POST("/viewhistory",controller.ViewHistory)
	router.POST("/listhistoryforadmin",controller.ListHistoryForAdmin)


	// Admin routes
	// router.Static("/admin","./admin")
	router.POST("/adminlogin", controller.AdminLogin)
	router.POST("/createadmin", controller.CreateAdmin)
	router.POST("/getdataforadmin", controller.GetData)
	router.POST("/getallcustomerforadmin", controller.GetallCustomerdata)
	router.POST("/getdetailsforadmin", controller.GetAllDetailsForAdmin)
	router.POST("/update", controller.Update)
	router.POST("/deleteuserbyadmin", controller.DeleteUser)
	router.POST("/deletehistorybyadmin", controller.DeleteHistorybyadmin)
	router.POST("/block", controller.Block)
	router.POST("/shutdown", controller.ShutDown)
	router.POST("/cleardb", controller.ClearDB)

	// FeedBack routes
	router.POST("/listfeedback", controller.GetFeedbacks)
	router.POST("/insertfeedback", controller.InsertFeedback)
	router.POST("/deletefeedback", controller.Deletefeedback)
	
	return router

}
