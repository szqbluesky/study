package study

func main(){

	engine := gin.Default()
	engine.Use(middleware.time)
}