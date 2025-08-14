package main
import(
	"log"
	"server/config"
)
func main(){
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}
	config.InitMySQLDB()
	config.InitRedis()
	config.CloseDB()
	config.CloseRedis()
}