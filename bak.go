package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/linqiurong2021/go-postgis/conf"
	"github.com/linqiurong2021/go-postgis/db"
)

func main11() {
	//
	err := conf.InitConfig("./conf/conf.ini")
	if err != nil {
		log.Fatal("init config error: ", err)
	}
	// user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10
	postgre := db.NewPostgreDB(conf.Conf)
	dbpool, err := postgre.Connect()
	if err != nil {
		fmt.Printf("connect postgre err:%s", err)
		os.Exit(1)
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select st_centroid('01060000208A11000001000000010300000001000000A90000005A798801B3835D40AF87E447587438401ACC9EE5B2835D4063D9572D587438408A4243C8B2835D400235F8245874384003796879AE835D40E35397615874384072223696AB835D401781981B57743840E0EE8901A1835D40630E2E2049743840DAF53DE5A0835D40773C200549743840D5857EC8A0835D4010E7B9FC4874384041E9C8ABA0835D40FC76BB0649743840BB6FD3A2A0835D407A0FB80D497438401DFEDE279D835D409AB289234C74384062BC864398835D4061A25512497438409294A8BA94835D404E29A73C407438404A14474A93835D405D8744A435743840B98E16FF91835D40F8E115DD0F74384053A454B491835D40930C7E1CF273384036FD6AB291835D401B60E3D3F173384048789F858C835D40FAA222AA79733840AFAFE7848C835D40F75BC19A79733840C98FA8838C835D40E7325B84797338407E09CD858A835D40EF5EB7E358733840C0145F7F8A835D405D735494587338407420FC728A835D4072B1F83458733840CB9223718A835D407A65BF29587338406E5C2D8F87835D409C98860E4773384043724D8087835D40C82DB7C3467338404F83C16B87835D40BBBA9E7946733840E6D1FB6887835D40F2746A7146733840BE4D014D82835D40F1801FA737733840146F073882835D40DEE46B733773384038F3CC1D82835D400696864737733840A549C21382835D40A2CBC33B37733840E64A52527C835D40CFA1B93931733840D78123117A835D403550B40B2E733840C0486B5F75835D4065EBAC67027338404E0B729775835D402A0667CCFD72384045B0657376835D40312ED083FB7238404502857D76835D407012A466FB7238406E59C19176835D401557681BFB72384029A01DA276835D409A5773C4FA72384047941AAE76835D40BF776964FA7238404460E6B176835D4021DEF237FA72384092ABD09777835D40BE8B0514EE723840C440459B77835D40616C48DAED723840DCA7909D77835D4058D30571ED723840773CD59A77835D40E31CE707ED7238401B7A969777835D40198F35D4EC723840607799C776835D4002DB9211E2723840617C637777835D400DBC89C2DB723840BDCE2DE080835D40D28DEB41C0723840180046E480835D402BAB8D35C0723840BF4782F880835D4019DB51EABF723840EC7FDE0881835D40FBCA5C93BF723840DE66DB1481835D40F8DE5233BF72384063BD1B1C81835D402A1F1FCDBE723840BE18481D81835D409B7108AFBE723840A5A5CBEA81835D40AD9EE59DA572384074A8EAEB81835D4034B1B952A5723840B3362FE981835D40B6FD9AE9A4723840013682E181835D40D3DCD183A4723840D04D12D981835D4059C6843EA47238406F5910E67F835D40C752019A967238406BF08D9180835D40AFADC6168C723840CC31959380835D40E631C5F18B7238407890E09580835D409B9682888B7238409620259380835D401CE2631F8B723840BA22788B80835D407EBF9AB98A7238407C4B157F80835D403BEA3E5A8A7238400EF35C6E80835D400C2036048A723840CF27D15980835D4019971DBA89723840B2BA114280835D40D4A8357E897238401664D72780835D40844E505289723840F726EE0B80835D401AF9C23789723840081E2FEF7F835D4039315C2F89723840FAE279D27F835D406E505D39897238403BC3ADB67F835D4063857855897238400AF7A29C7F835D40A130D38289723840930F24857F835D4069890CC0897238405ACFE7707F835D40E655480B8A723840979C8B607F835D4013633D628A723840EDB88E547F835D40064D47C28A72384043A5554F7F835D406E9079038B7238401F9CE1957E835D4087B2D06196723840085ADA937E835D40F02DD2869672384095F98E917E835D40D5C814F09672384029684A947E835D40EF7D3359977238405165F79B7E835D401DA2FCBE97723840E04A67A47E835D40FFBB4904987238404856399C80835D4044837DCAA572384017D994DC7F835D401DA7772ABD723840B872B88676835D40B4D3C073D87238408F3FA08276835D40FEB11E80D8723840AFED636E76835D40DB6C5ACBD8723840F6AA075E76835D40C06B4F22D9723840A6553F5476835D4082A8EA6CD9723840B323A17F75835D40B0014F0EE172384048876C7D75835D403CA4BD23E1723840A1252C7675835D40465CF189E1723840DBBCE07375835D4019F533F3E172384025269C7675835D409FAC525CE272384054E7DA7975835D40593B0490E2723840469A9B4C76835D40E9394577ED72384096C6CA8175835D40F074DE2CF872384006CFC59E74835D40F30C3E88FA723840A57CA69474835D40C1266AA5FA72384056246A8074835D400BDEA5F0FA72384014DC0D7074835D40FED99A47FB72384029E6106474835D403BB7A4A7FB723840FE89F65C74835D408A26AB0AFC7238402257661174835D40EC072F41027338400BE9070F74835D40194408AC027338402652C31174835D40A4FB261503733840B5ABFC1974835D40A9642D80037338402CD1E9F778835D4065EE56BF30733840DBF8330379835D40DD4F3814317338409563EC1379835D400923416A31733840EB46782879835D400AB759B43173384030D1374079835D400DB241F031733840867D615679835D404C49A41632733840D06557E27B835D4041DB1FAE357338409A3168E67B835D40E9AAA2B3357338408EDA72F07B835D40E97B65BF35733840F1BC849981835D401F1100A83B733840EA8DE68486835D40D54F9EE54973384099FAF24889835D40B8FB544F5A733840BD74983F8B835D406C69DB797A733840EE63036A90835D40BAC4846CF27338408ADAACB490835D40990B802310743840D099DAB590835D40B17B1C59107438401097740592835D4050442CA136743840034AE80692835D4051C68BC3367438403173950E92835D40A4CF542937743840A5DB791592835D40543D8A6337743840D06D98A593835D40846D9EE542743840C61717AB93835D4031B0C40A437438402BBACFBB93835D40C149CD604374384093DA5BD093835D400299E5AA437438405B1F75DC93835D40680C35CC437438409219C4A697835D40CBA874454D7438404EA16AB297835D40A8DD0C604D743840965EA5CC97835D4035E6F18B4D7438404D971EE897835D4007FA38A64D74384091A0E7119D835D40B688F6E250743840EA48DF2E9D835D40E4DE7FEB50743840A7E7944B9D835D404C597EE150743840F3618A549D835D40DDC381DA5074384044BA5CC0A0835D40DF801AD24D743840B905F737AB835D406E8A1AA75B743840D2044354AB835D40763C28C25B743840370BF65BAB835D401C273AC65B743840151AEF5CAE835D4055AF5A195D7438401C9E9772AE835D404C79A81D5D743840C0F79ECCB2835D40F5146CE05C743840DC6B1CE8B2835D40AB4077D65C74384043E4E803B3835D40726A5BBA5C743840F8FAF31DB3835D40BE29008D5C7438407E1C7335B3835D40F54BC64F5C7438408684AF49B3835D403F0F8A045C743840A4CB0B5AB3835D4015AA94AD5B743840A0AF0866B3835D4050838A4D5B7438405DF1486DB3835D40A9A456E75A743840222A946FB3835D40F407147E5A7438405482D86CB3835D40936FF5145A743840353C2B65B3835D40C3852CAF59743840910EC858B3835D404104D14F5974384034540F48B3835D403CA6C8F958743840821E8333B3835D40909EB0AF58743840A141C31BB3835D40C742C973587438405A798801B3835D40AF87E44758743840') as Center").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	fmt.Println(greeting)
	fmt.Println("Main")
}
