package service

// NewPool 建立数据库连接
// func NewPool() *pgxpool.Pool {
// 	postgre := new(libs.Postgre)
// 	dbpool, err := postgre.Connect()
// 	if err != nil {
// 		fmt.Printf("connect postgre err:%s", err)
// 		os.Exit(1)
// 	}
// 	return dbpool
// }

// Service Service
type Service struct {
	Handle    *Handle
	Relaction *Relaction
}

// NewService NewService
func NewService() *Service {

	return &Service{
		Handle:    NewHandle(),
		Relaction: NewRelaction(),
	}
}
