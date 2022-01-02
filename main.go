package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

type Activity struct {
	ID        int         `json:"id"`
	Email     string      `json:"email"`
	Title     string      `json:"title"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	DeletedAt interface{} `json:"deleted_at"`
}

type Todo struct {
	ID              int         `json:"id"`
	ActivityGroupId interface{} `json:"activity_group_id"`
	Title           interface{} `json:"title"`
	IsActive        interface{} `json:"is_active"`
	Priority        interface{} `json:"priority"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	DeletedAt       interface{} `json:"deleted_at"`
}

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Kosong struct {
}

var kosong Kosong
var activities = []Activity{}
var todos = []Todo{}
var resp response

// var currentActivity int
// var currentTodo int

// var db *gorm.DB
var err error
var db *sql.DB
var req int

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ":3306)/" + os.Getenv("MYSQL_DBNAME")
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("Database Not Connected")
	// } else {
	// db.AutoMigrate(&Activity{})
	// db.AutoMigrate(&Todo{})
	// }

	db, err = sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/"+os.Getenv("MYSQL_DBNAME"))
	if err != nil {
		panic(err)
	}

	req = 1
	defer db.Close()

	db.Query(`CREATE TABLE IF NOT EXISTS activities (
		id bigint(20) NOT NULL,
		email varchar(255) DEFAULT NULL,
		title varchar(255) DEFAULT NULL,
		created_at varchar(255) DEFAULT NULL,
		updated_at varchar(255) DEFAULT NULL,
		deleted_at varchar(255) DEFAULT NULL
	  ) ENGINE=MyIsam DEFAULT CHARSET=latin1;`)

	db.Query(`CREATE TABLE IF NOT EXISTS todos (
		id bigint(20) NOT NULL,
		activity_group_id varchar(255) DEFAULT NULL,
		title varchar(255) DEFAULT NULL,
		is_active varchar(255) DEFAULT NULL,
		priority varchar(255) DEFAULT NULL,
		created_at varchar(255) DEFAULT NULL,
		updated_at varchar(255) DEFAULT NULL,
		deleted_at varchar(255) DEFAULT NULL
	  ) ENGINE=MyIsam DEFAULT CHARSET=latin1;
	  `)

	db.Query(`INSERT INTO activities (id, email, title, created_at, updated_at, deleted_at) VALUES
	(1, 'e9c4786a-66a3-48e4-9478-320fc834d1f1@test.com', 'testing112', NULL, NULL, NULL),
	(2, 'performance@test.com', 'performanceTesting', NULL, NULL, NULL);`)

	db.Query(`INSERT INTO todos (id, activity_group_id, title, is_active, priority, created_at, updated_at, deleted_at) VALUES
	(1, '1', 'todoTesting', '1', 'very-high', NULL, NULL, NULL),
	(2, '2', 'performanceTesting1', '1', 'very-high', NULL, NULL, NULL),
	(3, '2', 'performanceTesting2', '1', 'very-high', NULL, NULL, NULL),
	(4, '2', 'performanceTesting3', '1', 'very-high', NULL, NULL, NULL),
	(5, '2', 'performanceTesting4', '1', 'very-high', NULL, NULL, NULL),
	(6, '2', 'performanceTesting5', '1', 'very-high', NULL, NULL, NULL),
	(7, '2', 'performanceTesting6', '1', 'very-high', NULL, NULL, NULL),
	(8, '2', 'performanceTesting7', '1', 'very-high', NULL, NULL, NULL),
	(9, '2', 'performanceTesting8', '1', 'very-high', NULL, NULL, NULL),
	(10, '2', 'performanceTesting9', '1', 'very-high', NULL, NULL, NULL),
	(11, '2', 'performanceTesting10', '1', 'very-high', NULL, NULL, NULL),
	(12, '2', 'performanceTesting11', '1', 'very-high', NULL, NULL, NULL),
	(13, '2', 'performanceTesting12', '1', 'very-high', NULL, NULL, NULL),
	(14, '2', 'performanceTesting13', '1', 'very-high', NULL, NULL, NULL),
	(15, '2', 'performanceTesting14', '1', 'very-high', NULL, NULL, NULL),
	(16, '2', 'performanceTesting15', '1', 'very-high', NULL, NULL, NULL),
	(17, '2', 'performanceTesting16', '1', 'very-high', NULL, NULL, NULL),
	(18, '2', 'performanceTesting17', '1', 'very-high', NULL, NULL, NULL),
	(19, '2', 'performanceTesting18', '1', 'very-high', NULL, NULL, NULL),
	(20, '2', 'performanceTesting19', '1', 'very-high', NULL, NULL, NULL),
	(21, '2', 'performanceTesting20', '1', 'very-high', NULL, NULL, NULL),
	(22, '2', 'performanceTesting21', '1', 'very-high', NULL, NULL, NULL),
	(23, '2', 'performanceTesting22', '1', 'very-high', NULL, NULL, NULL),
	(24, '2', 'performanceTesting23', '1', 'very-high', NULL, NULL, NULL),
	(25, '2', 'performanceTesting24', '1', 'very-high', NULL, NULL, NULL),
	(26, '2', 'performanceTesting25', '1', 'very-high', NULL, NULL, NULL),
	(27, '2', 'performanceTesting26', '1', 'very-high', NULL, NULL, NULL),
	(28, '2', 'performanceTesting27', '1', 'very-high', NULL, NULL, NULL),
	(29, '2', 'performanceTesting28', '1', 'very-high', NULL, NULL, NULL),
	(30, '2', 'performanceTesting29', '1', 'very-high', NULL, NULL, NULL),
	(31, '2', 'performanceTesting30', '1', 'very-high', NULL, NULL, NULL),
	(32, '2', 'performanceTesting31', '1', 'very-high', NULL, NULL, NULL),
	(33, '2', 'performanceTesting32', '1', 'very-high', NULL, NULL, NULL),
	(34, '2', 'performanceTesting33', '1', 'very-high', NULL, NULL, NULL),
	(35, '2', 'performanceTesting34', '1', 'very-high', NULL, NULL, NULL),
	(36, '2', 'performanceTesting35', '1', 'very-high', NULL, NULL, NULL),
	(37, '2', 'performanceTesting36', '1', 'very-high', NULL, NULL, NULL),
	(38, '2', 'performanceTesting37', '1', 'very-high', NULL, NULL, NULL),
	(39, '2', 'performanceTesting38', '1', 'very-high', NULL, NULL, NULL),
	(40, '2', 'performanceTesting39', '1', 'very-high', NULL, NULL, NULL),
	(41, '2', 'performanceTesting40', '1', 'very-high', NULL, NULL, NULL),
	(42, '2', 'performanceTesting41', '1', 'very-high', NULL, NULL, NULL),
	(43, '2', 'performanceTesting42', '1', 'very-high', NULL, NULL, NULL),
	(44, '2', 'performanceTesting43', '1', 'very-high', NULL, NULL, NULL),
	(45, '2', 'performanceTesting44', '1', 'very-high', NULL, NULL, NULL),
	(46, '2', 'performanceTesting45', '1', 'very-high', NULL, NULL, NULL),
	(47, '2', 'performanceTesting46', '1', 'very-high', NULL, NULL, NULL),
	(48, '2', 'performanceTesting47', '1', 'very-high', NULL, NULL, NULL),
	(49, '2', 'performanceTesting48', '1', 'very-high', NULL, NULL, NULL),
	(50, '2', 'performanceTesting49', '1', 'very-high', NULL, NULL, NULL),
	(51, '2', 'performanceTesting50', '1', 'very-high', NULL, NULL, NULL),
	(52, '2', 'performanceTesting51', '1', 'very-high', NULL, NULL, NULL),
	(53, '2', 'performanceTesting52', '1', 'very-high', NULL, NULL, NULL),
	(54, '2', 'performanceTesting53', '1', 'very-high', NULL, NULL, NULL),
	(55, '2', 'performanceTesting54', '1', 'very-high', NULL, NULL, NULL),
	(56, '2', 'performanceTesting55', '1', 'very-high', NULL, NULL, NULL),
	(57, '2', 'performanceTesting56', '1', 'very-high', NULL, NULL, NULL),
	(58, '2', 'performanceTesting57', '1', 'very-high', NULL, NULL, NULL),
	(59, '2', 'performanceTesting58', '1', 'very-high', NULL, NULL, NULL),
	(60, '2', 'performanceTesting59', '1', 'very-high', NULL, NULL, NULL),
	(61, '2', 'performanceTesting60', '1', 'very-high', NULL, NULL, NULL),
	(62, '2', 'performanceTesting61', '1', 'very-high', NULL, NULL, NULL),
	(63, '2', 'performanceTesting62', '1', 'very-high', NULL, NULL, NULL),
	(64, '2', 'performanceTesting63', '1', 'very-high', NULL, NULL, NULL),
	(65, '2', 'performanceTesting64', '1', 'very-high', NULL, NULL, NULL),
	(66, '2', 'performanceTesting65', '1', 'very-high', NULL, NULL, NULL),
	(67, '2', 'performanceTesting66', '1', 'very-high', NULL, NULL, NULL),
	(68, '2', 'performanceTesting67', '1', 'very-high', NULL, NULL, NULL),
	(69, '2', 'performanceTesting68', '1', 'very-high', NULL, NULL, NULL),
	(70, '2', 'performanceTesting69', '1', 'very-high', NULL, NULL, NULL),
	(71, '2', 'performanceTesting70', '1', 'very-high', NULL, NULL, NULL),
	(72, '2', 'performanceTesting71', '1', 'very-high', NULL, NULL, NULL),
	(73, '2', 'performanceTesting72', '1', 'very-high', NULL, NULL, NULL),
	(74, '2', 'performanceTesting73', '1', 'very-high', NULL, NULL, NULL),
	(75, '2', 'performanceTesting74', '1', 'very-high', NULL, NULL, NULL),
	(76, '2', 'performanceTesting75', '1', 'very-high', NULL, NULL, NULL),
	(77, '2', 'performanceTesting76', '1', 'very-high', NULL, NULL, NULL),
	(78, '2', 'performanceTesting77', '1', 'very-high', NULL, NULL, NULL),
	(79, '2', 'performanceTesting78', '1', 'very-high', NULL, NULL, NULL),
	(80, '2', 'performanceTesting79', '1', 'very-high', NULL, NULL, NULL),
	(81, '2', 'performanceTesting80', '1', 'very-high', NULL, NULL, NULL),
	(82, '2', 'performanceTesting81', '1', 'very-high', NULL, NULL, NULL),
	(83, '2', 'performanceTesting82', '1', 'very-high', NULL, NULL, NULL),
	(84, '2', 'performanceTesting83', '1', 'very-high', NULL, NULL, NULL),
	(85, '2', 'performanceTesting84', '1', 'very-high', NULL, NULL, NULL),
	(86, '2', 'performanceTesting85', '1', 'very-high', NULL, NULL, NULL),
	(87, '2', 'performanceTesting86', '1', 'very-high', NULL, NULL, NULL),
	(88, '2', 'performanceTesting87', '1', 'very-high', NULL, NULL, NULL),
	(89, '2', 'performanceTesting88', '1', 'very-high', NULL, NULL, NULL),
	(90, '2', 'performanceTesting89', '1', 'very-high', NULL, NULL, NULL),
	(91, '2', 'performanceTesting90', '1', 'very-high', NULL, NULL, NULL),
	(92, '2', 'performanceTesting91', '1', 'very-high', NULL, NULL, NULL),
	(93, '2', 'performanceTesting92', '1', 'very-high', NULL, NULL, NULL),
	(94, '2', 'performanceTesting93', '1', 'very-high', NULL, NULL, NULL),
	(95, '2', 'performanceTesting94', '1', 'very-high', NULL, NULL, NULL),
	(96, '2', 'performanceTesting95', '1', 'very-high', NULL, NULL, NULL),
	(97, '2', 'performanceTesting96', '1', 'very-high', NULL, NULL, NULL),
	(98, '2', 'performanceTesting97', '1', 'very-high', NULL, NULL, NULL),
	(99, '2', 'performanceTesting98', '1', 'very-high', NULL, NULL, NULL),
	(100, '2', 'performanceTesting99', '1', 'very-high', NULL, NULL, NULL),
	(101, '2', 'performanceTesting100', '1', 'very-high', NULL, NULL, NULL),
	(102, '2', 'performanceTesting101', '1', 'very-high', NULL, NULL, NULL),
	(103, '2', 'performanceTesting102', '1', 'very-high', NULL, NULL, NULL),
	(104, '2', 'performanceTesting103', '1', 'very-high', NULL, NULL, NULL),
	(105, '2', 'performanceTesting104', '1', 'very-high', NULL, NULL, NULL),
	(106, '2', 'performanceTesting105', '1', 'very-high', NULL, NULL, NULL),
	(107, '2', 'performanceTesting106', '1', 'very-high', NULL, NULL, NULL),
	(108, '2', 'performanceTesting107', '1', 'very-high', NULL, NULL, NULL),
	(109, '2', 'performanceTesting108', '1', 'very-high', NULL, NULL, NULL),
	(110, '2', 'performanceTesting109', '1', 'very-high', NULL, NULL, NULL),
	(111, '2', 'performanceTesting110', '1', 'very-high', NULL, NULL, NULL),
	(112, '2', 'performanceTesting111', '1', 'very-high', NULL, NULL, NULL),
	(113, '2', 'performanceTesting112', '1', 'very-high', NULL, NULL, NULL),
	(114, '2', 'performanceTesting113', '1', 'very-high', NULL, NULL, NULL),
	(115, '2', 'performanceTesting114', '1', 'very-high', NULL, NULL, NULL),
	(116, '2', 'performanceTesting115', '1', 'very-high', NULL, NULL, NULL),
	(117, '2', 'performanceTesting116', '1', 'very-high', NULL, NULL, NULL),
	(118, '2', 'performanceTesting117', '1', 'very-high', NULL, NULL, NULL),
	(119, '2', 'performanceTesting118', '1', 'very-high', NULL, NULL, NULL),
	(120, '2', 'performanceTesting119', '1', 'very-high', NULL, NULL, NULL),
	(121, '2', 'performanceTesting120', '1', 'very-high', NULL, NULL, NULL),
	(122, '2', 'performanceTesting121', '1', 'very-high', NULL, NULL, NULL),
	(123, '2', 'performanceTesting122', '1', 'very-high', NULL, NULL, NULL),
	(124, '2', 'performanceTesting123', '1', 'very-high', NULL, NULL, NULL),
	(125, '2', 'performanceTesting124', '1', 'very-high', NULL, NULL, NULL),
	(126, '2', 'performanceTesting125', '1', 'very-high', NULL, NULL, NULL),
	(127, '2', 'performanceTesting126', '1', 'very-high', NULL, NULL, NULL),
	(128, '2', 'performanceTesting127', '1', 'very-high', NULL, NULL, NULL),
	(129, '2', 'performanceTesting128', '1', 'very-high', NULL, NULL, NULL),
	(130, '2', 'performanceTesting129', '1', 'very-high', NULL, NULL, NULL),
	(131, '2', 'performanceTesting130', '1', 'very-high', NULL, NULL, NULL),
	(132, '2', 'performanceTesting131', '1', 'very-high', NULL, NULL, NULL),
	(133, '2', 'performanceTesting132', '1', 'very-high', NULL, NULL, NULL),
	(134, '2', 'performanceTesting133', '1', 'very-high', NULL, NULL, NULL),
	(135, '2', 'performanceTesting134', '1', 'very-high', NULL, NULL, NULL),
	(136, '2', 'performanceTesting135', '1', 'very-high', NULL, NULL, NULL),
	(137, '2', 'performanceTesting136', '1', 'very-high', NULL, NULL, NULL),
	(138, '2', 'performanceTesting137', '1', 'very-high', NULL, NULL, NULL),
	(139, '2', 'performanceTesting138', '1', 'very-high', NULL, NULL, NULL),
	(140, '2', 'performanceTesting139', '1', 'very-high', NULL, NULL, NULL),
	(141, '2', 'performanceTesting140', '1', 'very-high', NULL, NULL, NULL),
	(142, '2', 'performanceTesting141', '1', 'very-high', NULL, NULL, NULL),
	(143, '2', 'performanceTesting142', '1', 'very-high', NULL, NULL, NULL),
	(144, '2', 'performanceTesting143', '1', 'very-high', NULL, NULL, NULL),
	(145, '2', 'performanceTesting144', '1', 'very-high', NULL, NULL, NULL),
	(146, '2', 'performanceTesting145', '1', 'very-high', NULL, NULL, NULL),
	(147, '2', 'performanceTesting146', '1', 'very-high', NULL, NULL, NULL),
	(148, '2', 'performanceTesting147', '1', 'very-high', NULL, NULL, NULL),
	(149, '2', 'performanceTesting148', '1', 'very-high', NULL, NULL, NULL),
	(150, '2', 'performanceTesting149', '1', 'very-high', NULL, NULL, NULL),
	(151, '2', 'performanceTesting150', '1', 'very-high', NULL, NULL, NULL),
	(152, '2', 'performanceTesting151', '1', 'very-high', NULL, NULL, NULL),
	(153, '2', 'performanceTesting152', '1', 'very-high', NULL, NULL, NULL),
	(154, '2', 'performanceTesting153', '1', 'very-high', NULL, NULL, NULL),
	(155, '2', 'performanceTesting154', '1', 'very-high', NULL, NULL, NULL),
	(156, '2', 'performanceTesting155', '1', 'very-high', NULL, NULL, NULL),
	(157, '2', 'performanceTesting156', '1', 'very-high', NULL, NULL, NULL),
	(158, '2', 'performanceTesting157', '1', 'very-high', NULL, NULL, NULL),
	(159, '2', 'performanceTesting158', '1', 'very-high', NULL, NULL, NULL),
	(160, '2', 'performanceTesting159', '1', 'very-high', NULL, NULL, NULL),
	(161, '2', 'performanceTesting160', '1', 'very-high', NULL, NULL, NULL),
	(162, '2', 'performanceTesting161', '1', 'very-high', NULL, NULL, NULL),
	(163, '2', 'performanceTesting162', '1', 'very-high', NULL, NULL, NULL),
	(164, '2', 'performanceTesting163', '1', 'very-high', NULL, NULL, NULL),
	(165, '2', 'performanceTesting164', '1', 'very-high', NULL, NULL, NULL),
	(166, '2', 'performanceTesting165', '1', 'very-high', NULL, NULL, NULL),
	(167, '2', 'performanceTesting166', '1', 'very-high', NULL, NULL, NULL),
	(168, '2', 'performanceTesting167', '1', 'very-high', NULL, NULL, NULL),
	(169, '2', 'performanceTesting168', '1', 'very-high', NULL, NULL, NULL),
	(170, '2', 'performanceTesting169', '1', 'very-high', NULL, NULL, NULL),
	(171, '2', 'performanceTesting170', '1', 'very-high', NULL, NULL, NULL),
	(172, '2', 'performanceTesting171', '1', 'very-high', NULL, NULL, NULL),
	(173, '2', 'performanceTesting172', '1', 'very-high', NULL, NULL, NULL),
	(174, '2', 'performanceTesting173', '1', 'very-high', NULL, NULL, NULL),
	(175, '2', 'performanceTesting174', '1', 'very-high', NULL, NULL, NULL),
	(176, '2', 'performanceTesting175', '1', 'very-high', NULL, NULL, NULL),
	(177, '2', 'performanceTesting176', '1', 'very-high', NULL, NULL, NULL),
	(178, '2', 'performanceTesting177', '1', 'very-high', NULL, NULL, NULL),
	(179, '2', 'performanceTesting178', '1', 'very-high', NULL, NULL, NULL),
	(180, '2', 'performanceTesting179', '1', 'very-high', NULL, NULL, NULL),
	(181, '2', 'performanceTesting180', '1', 'very-high', NULL, NULL, NULL),
	(182, '2', 'performanceTesting181', '1', 'very-high', NULL, NULL, NULL),
	(183, '2', 'performanceTesting182', '1', 'very-high', NULL, NULL, NULL),
	(184, '2', 'performanceTesting183', '1', 'very-high', NULL, NULL, NULL),
	(185, '2', 'performanceTesting184', '1', 'very-high', NULL, NULL, NULL),
	(186, '2', 'performanceTesting185', '1', 'very-high', NULL, NULL, NULL),
	(187, '2', 'performanceTesting186', '1', 'very-high', NULL, NULL, NULL),
	(188, '2', 'performanceTesting187', '1', 'very-high', NULL, NULL, NULL),
	(189, '2', 'performanceTesting188', '1', 'very-high', NULL, NULL, NULL),
	(190, '2', 'performanceTesting189', '1', 'very-high', NULL, NULL, NULL),
	(191, '2', 'performanceTesting190', '1', 'very-high', NULL, NULL, NULL),
	(192, '2', 'performanceTesting191', '1', 'very-high', NULL, NULL, NULL),
	(193, '2', 'performanceTesting192', '1', 'very-high', NULL, NULL, NULL),
	(194, '2', 'performanceTesting193', '1', 'very-high', NULL, NULL, NULL),
	(195, '2', 'performanceTesting194', '1', 'very-high', NULL, NULL, NULL),
	(196, '2', 'performanceTesting195', '1', 'very-high', NULL, NULL, NULL),
	(197, '2', 'performanceTesting196', '1', 'very-high', NULL, NULL, NULL),
	(198, '2', 'performanceTesting197', '1', 'very-high', NULL, NULL, NULL),
	(199, '2', 'performanceTesting198', '1', 'very-high', NULL, NULL, NULL),
	(200, '2', 'performanceTesting199', '1', 'very-high', NULL, NULL, NULL),
	(201, '2', 'performanceTesting200', '1', 'very-high', NULL, NULL, NULL),
	(202, '2', 'performanceTesting201', '1', 'very-high', NULL, NULL, NULL),
	(203, '2', 'performanceTesting202', '1', 'very-high', NULL, NULL, NULL),
	(204, '2', 'performanceTesting203', '1', 'very-high', NULL, NULL, NULL),
	(205, '2', 'performanceTesting204', '1', 'very-high', NULL, NULL, NULL),
	(206, '2', 'performanceTesting205', '1', 'very-high', NULL, NULL, NULL),
	(207, '2', 'performanceTesting206', '1', 'very-high', NULL, NULL, NULL),
	(208, '2', 'performanceTesting207', '1', 'very-high', NULL, NULL, NULL),
	(209, '2', 'performanceTesting208', '1', 'very-high', NULL, NULL, NULL),
	(210, '2', 'performanceTesting209', '1', 'very-high', NULL, NULL, NULL),
	(211, '2', 'performanceTesting210', '1', 'very-high', NULL, NULL, NULL),
	(212, '2', 'performanceTesting211', '1', 'very-high', NULL, NULL, NULL),
	(213, '2', 'performanceTesting212', '1', 'very-high', NULL, NULL, NULL),
	(214, '2', 'performanceTesting213', '1', 'very-high', NULL, NULL, NULL),
	(215, '2', 'performanceTesting214', '1', 'very-high', NULL, NULL, NULL),
	(216, '2', 'performanceTesting215', '1', 'very-high', NULL, NULL, NULL),
	(217, '2', 'performanceTesting216', '1', 'very-high', NULL, NULL, NULL),
	(218, '2', 'performanceTesting217', '1', 'very-high', NULL, NULL, NULL),
	(219, '2', 'performanceTesting218', '1', 'very-high', NULL, NULL, NULL),
	(220, '2', 'performanceTesting219', '1', 'very-high', NULL, NULL, NULL),
	(221, '2', 'performanceTesting220', '1', 'very-high', NULL, NULL, NULL),
	(222, '2', 'performanceTesting221', '1', 'very-high', NULL, NULL, NULL),
	(223, '2', 'performanceTesting222', '1', 'very-high', NULL, NULL, NULL),
	(224, '2', 'performanceTesting223', '1', 'very-high', NULL, NULL, NULL),
	(225, '2', 'performanceTesting224', '1', 'very-high', NULL, NULL, NULL),
	(226, '2', 'performanceTesting225', '1', 'very-high', NULL, NULL, NULL),
	(227, '2', 'performanceTesting226', '1', 'very-high', NULL, NULL, NULL),
	(228, '2', 'performanceTesting227', '1', 'very-high', NULL, NULL, NULL),
	(229, '2', 'performanceTesting228', '1', 'very-high', NULL, NULL, NULL),
	(230, '2', 'performanceTesting229', '1', 'very-high', NULL, NULL, NULL),
	(231, '2', 'performanceTesting230', '1', 'very-high', NULL, NULL, NULL),
	(232, '2', 'performanceTesting231', '1', 'very-high', NULL, NULL, NULL),
	(233, '2', 'performanceTesting232', '1', 'very-high', NULL, NULL, NULL),
	(234, '2', 'performanceTesting233', '1', 'very-high', NULL, NULL, NULL),
	(235, '2', 'performanceTesting234', '1', 'very-high', NULL, NULL, NULL),
	(236, '2', 'performanceTesting235', '1', 'very-high', NULL, NULL, NULL),
	(237, '2', 'performanceTesting236', '1', 'very-high', NULL, NULL, NULL),
	(238, '2', 'performanceTesting237', '1', 'very-high', NULL, NULL, NULL),
	(239, '2', 'performanceTesting238', '1', 'very-high', NULL, NULL, NULL),
	(240, '2', 'performanceTesting239', '1', 'very-high', NULL, NULL, NULL),
	(241, '2', 'performanceTesting240', '1', 'very-high', NULL, NULL, NULL),
	(242, '2', 'performanceTesting241', '1', 'very-high', NULL, NULL, NULL),
	(243, '2', 'performanceTesting242', '1', 'very-high', NULL, NULL, NULL),
	(244, '2', 'performanceTesting243', '1', 'very-high', NULL, NULL, NULL),
	(245, '2', 'performanceTesting244', '1', 'very-high', NULL, NULL, NULL),
	(246, '2', 'performanceTesting245', '1', 'very-high', NULL, NULL, NULL),
	(247, '2', 'performanceTesting246', '1', 'very-high', NULL, NULL, NULL),
	(248, '2', 'performanceTesting247', '1', 'very-high', NULL, NULL, NULL),
	(249, '2', 'performanceTesting248', '1', 'very-high', NULL, NULL, NULL),
	(250, '2', 'performanceTesting249', '1', 'very-high', NULL, NULL, NULL),
	(251, '2', 'performanceTesting250', '1', 'very-high', NULL, NULL, NULL),
	(252, '2', 'performanceTesting251', '1', 'very-high', NULL, NULL, NULL),
	(253, '2', 'performanceTesting252', '1', 'very-high', NULL, NULL, NULL),
	(254, '2', 'performanceTesting253', '1', 'very-high', NULL, NULL, NULL),
	(255, '2', 'performanceTesting254', '1', 'very-high', NULL, NULL, NULL),
	(256, '2', 'performanceTesting255', '1', 'very-high', NULL, NULL, NULL),
	(257, '2', 'performanceTesting256', '1', 'very-high', NULL, NULL, NULL),
	(258, '2', 'performanceTesting257', '1', 'very-high', NULL, NULL, NULL),
	(259, '2', 'performanceTesting258', '1', 'very-high', NULL, NULL, NULL),
	(260, '2', 'performanceTesting259', '1', 'very-high', NULL, NULL, NULL),
	(261, '2', 'performanceTesting260', '1', 'very-high', NULL, NULL, NULL),
	(262, '2', 'performanceTesting261', '1', 'very-high', NULL, NULL, NULL),
	(263, '2', 'performanceTesting262', '1', 'very-high', NULL, NULL, NULL),
	(264, '2', 'performanceTesting263', '1', 'very-high', NULL, NULL, NULL),
	(265, '2', 'performanceTesting264', '1', 'very-high', NULL, NULL, NULL),
	(266, '2', 'performanceTesting265', '1', 'very-high', NULL, NULL, NULL),
	(267, '2', 'performanceTesting266', '1', 'very-high', NULL, NULL, NULL),
	(268, '2', 'performanceTesting267', '1', 'very-high', NULL, NULL, NULL),
	(269, '2', 'performanceTesting268', '1', 'very-high', NULL, NULL, NULL),
	(270, '2', 'performanceTesting269', '1', 'very-high', NULL, NULL, NULL),
	(271, '2', 'performanceTesting270', '1', 'very-high', NULL, NULL, NULL),
	(272, '2', 'performanceTesting271', '1', 'very-high', NULL, NULL, NULL),
	(273, '2', 'performanceTesting272', '1', 'very-high', NULL, NULL, NULL),
	(274, '2', 'performanceTesting273', '1', 'very-high', NULL, NULL, NULL),
	(275, '2', 'performanceTesting274', '1', 'very-high', NULL, NULL, NULL),
	(276, '2', 'performanceTesting275', '1', 'very-high', NULL, NULL, NULL),
	(277, '2', 'performanceTesting276', '1', 'very-high', NULL, NULL, NULL),
	(278, '2', 'performanceTesting277', '1', 'very-high', NULL, NULL, NULL),
	(279, '2', 'performanceTesting278', '1', 'very-high', NULL, NULL, NULL),
	(280, '2', 'performanceTesting279', '1', 'very-high', NULL, NULL, NULL),
	(281, '2', 'performanceTesting280', '1', 'very-high', NULL, NULL, NULL),
	(282, '2', 'performanceTesting281', '1', 'very-high', NULL, NULL, NULL),
	(283, '2', 'performanceTesting282', '1', 'very-high', NULL, NULL, NULL),
	(284, '2', 'performanceTesting283', '1', 'very-high', NULL, NULL, NULL),
	(285, '2', 'performanceTesting284', '1', 'very-high', NULL, NULL, NULL),
	(286, '2', 'performanceTesting285', '1', 'very-high', NULL, NULL, NULL),
	(287, '2', 'performanceTesting286', '1', 'very-high', NULL, NULL, NULL),
	(288, '2', 'performanceTesting287', '1', 'very-high', NULL, NULL, NULL),
	(289, '2', 'performanceTesting288', '1', 'very-high', NULL, NULL, NULL),
	(290, '2', 'performanceTesting289', '1', 'very-high', NULL, NULL, NULL),
	(291, '2', 'performanceTesting290', '1', 'very-high', NULL, NULL, NULL),
	(292, '2', 'performanceTesting291', '1', 'very-high', NULL, NULL, NULL),
	(293, '2', 'performanceTesting292', '1', 'very-high', NULL, NULL, NULL),
	(294, '2', 'performanceTesting293', '1', 'very-high', NULL, NULL, NULL),
	(295, '2', 'performanceTesting294', '1', 'very-high', NULL, NULL, NULL),
	(296, '2', 'performanceTesting295', '1', 'very-high', NULL, NULL, NULL),
	(297, '2', 'performanceTesting296', '1', 'very-high', NULL, NULL, NULL),
	(298, '2', 'performanceTesting297', '1', 'very-high', NULL, NULL, NULL),
	(299, '2', 'performanceTesting298', '1', 'very-high', NULL, NULL, NULL),
	(300, '2', 'performanceTesting299', '1', 'very-high', NULL, NULL, NULL),
	(301, '2', 'performanceTesting300', '1', 'very-high', NULL, NULL, NULL),
	(302, '2', 'performanceTesting301', '1', 'very-high', NULL, NULL, NULL),
	(303, '2', 'performanceTesting302', '1', 'very-high', NULL, NULL, NULL),
	(304, '2', 'performanceTesting303', '1', 'very-high', NULL, NULL, NULL),
	(305, '2', 'performanceTesting304', '1', 'very-high', NULL, NULL, NULL),
	(306, '2', 'performanceTesting305', '1', 'very-high', NULL, NULL, NULL),
	(307, '2', 'performanceTesting306', '1', 'very-high', NULL, NULL, NULL),
	(308, '2', 'performanceTesting307', '1', 'very-high', NULL, NULL, NULL),
	(309, '2', 'performanceTesting308', '1', 'very-high', NULL, NULL, NULL),
	(310, '2', 'performanceTesting309', '1', 'very-high', NULL, NULL, NULL),
	(311, '2', 'performanceTesting310', '1', 'very-high', NULL, NULL, NULL),
	(312, '2', 'performanceTesting311', '1', 'very-high', NULL, NULL, NULL),
	(313, '2', 'performanceTesting312', '1', 'very-high', NULL, NULL, NULL),
	(314, '2', 'performanceTesting313', '1', 'very-high', NULL, NULL, NULL),
	(315, '2', 'performanceTesting314', '1', 'very-high', NULL, NULL, NULL),
	(316, '2', 'performanceTesting315', '1', 'very-high', NULL, NULL, NULL),
	(317, '2', 'performanceTesting316', '1', 'very-high', NULL, NULL, NULL),
	(318, '2', 'performanceTesting317', '1', 'very-high', NULL, NULL, NULL),
	(319, '2', 'performanceTesting318', '1', 'very-high', NULL, NULL, NULL),
	(320, '2', 'performanceTesting319', '1', 'very-high', NULL, NULL, NULL),
	(321, '2', 'performanceTesting320', '1', 'very-high', NULL, NULL, NULL),
	(322, '2', 'performanceTesting321', '1', 'very-high', NULL, NULL, NULL),
	(323, '2', 'performanceTesting322', '1', 'very-high', NULL, NULL, NULL),
	(324, '2', 'performanceTesting323', '1', 'very-high', NULL, NULL, NULL),
	(325, '2', 'performanceTesting324', '1', 'very-high', NULL, NULL, NULL),
	(326, '2', 'performanceTesting325', '1', 'very-high', NULL, NULL, NULL),
	(327, '2', 'performanceTesting326', '1', 'very-high', NULL, NULL, NULL),
	(328, '2', 'performanceTesting327', '1', 'very-high', NULL, NULL, NULL),
	(329, '2', 'performanceTesting328', '1', 'very-high', NULL, NULL, NULL),
	(330, '2', 'performanceTesting329', '1', 'very-high', NULL, NULL, NULL),
	(331, '2', 'performanceTesting330', '1', 'very-high', NULL, NULL, NULL),
	(332, '2', 'performanceTesting331', '1', 'very-high', NULL, NULL, NULL),
	(333, '2', 'performanceTesting332', '1', 'very-high', NULL, NULL, NULL),
	(334, '2', 'performanceTesting333', '1', 'very-high', NULL, NULL, NULL),
	(335, '2', 'performanceTesting334', '1', 'very-high', NULL, NULL, NULL),
	(336, '2', 'performanceTesting335', '1', 'very-high', NULL, NULL, NULL),
	(337, '2', 'performanceTesting336', '1', 'very-high', NULL, NULL, NULL),
	(338, '2', 'performanceTesting337', '1', 'very-high', NULL, NULL, NULL),
	(339, '2', 'performanceTesting338', '1', 'very-high', NULL, NULL, NULL),
	(340, '2', 'performanceTesting339', '1', 'very-high', NULL, NULL, NULL),
	(341, '2', 'performanceTesting340', '1', 'very-high', NULL, NULL, NULL),
	(342, '2', 'performanceTesting341', '1', 'very-high', NULL, NULL, NULL),
	(343, '2', 'performanceTesting342', '1', 'very-high', NULL, NULL, NULL),
	(344, '2', 'performanceTesting343', '1', 'very-high', NULL, NULL, NULL),
	(345, '2', 'performanceTesting344', '1', 'very-high', NULL, NULL, NULL),
	(346, '2', 'performanceTesting345', '1', 'very-high', NULL, NULL, NULL),
	(347, '2', 'performanceTesting346', '1', 'very-high', NULL, NULL, NULL),
	(348, '2', 'performanceTesting347', '1', 'very-high', NULL, NULL, NULL),
	(349, '2', 'performanceTesting348', '1', 'very-high', NULL, NULL, NULL),
	(350, '2', 'performanceTesting349', '1', 'very-high', NULL, NULL, NULL),
	(351, '2', 'performanceTesting350', '1', 'very-high', NULL, NULL, NULL),
	(352, '2', 'performanceTesting351', '1', 'very-high', NULL, NULL, NULL),
	(353, '2', 'performanceTesting352', '1', 'very-high', NULL, NULL, NULL),
	(354, '2', 'performanceTesting353', '1', 'very-high', NULL, NULL, NULL),
	(355, '2', 'performanceTesting354', '1', 'very-high', NULL, NULL, NULL),
	(356, '2', 'performanceTesting355', '1', 'very-high', NULL, NULL, NULL),
	(357, '2', 'performanceTesting356', '1', 'very-high', NULL, NULL, NULL),
	(358, '2', 'performanceTesting357', '1', 'very-high', NULL, NULL, NULL),
	(359, '2', 'performanceTesting358', '1', 'very-high', NULL, NULL, NULL),
	(360, '2', 'performanceTesting359', '1', 'very-high', NULL, NULL, NULL),
	(361, '2', 'performanceTesting360', '1', 'very-high', NULL, NULL, NULL),
	(362, '2', 'performanceTesting361', '1', 'very-high', NULL, NULL, NULL),
	(363, '2', 'performanceTesting362', '1', 'very-high', NULL, NULL, NULL),
	(364, '2', 'performanceTesting363', '1', 'very-high', NULL, NULL, NULL),
	(365, '2', 'performanceTesting364', '1', 'very-high', NULL, NULL, NULL),
	(366, '2', 'performanceTesting365', '1', 'very-high', NULL, NULL, NULL),
	(367, '2', 'performanceTesting366', '1', 'very-high', NULL, NULL, NULL),
	(368, '2', 'performanceTesting367', '1', 'very-high', NULL, NULL, NULL),
	(369, '2', 'performanceTesting368', '1', 'very-high', NULL, NULL, NULL),
	(370, '2', 'performanceTesting369', '1', 'very-high', NULL, NULL, NULL),
	(371, '2', 'performanceTesting370', '1', 'very-high', NULL, NULL, NULL),
	(372, '2', 'performanceTesting371', '1', 'very-high', NULL, NULL, NULL),
	(373, '2', 'performanceTesting372', '1', 'very-high', NULL, NULL, NULL),
	(374, '2', 'performanceTesting373', '1', 'very-high', NULL, NULL, NULL),
	(375, '2', 'performanceTesting374', '1', 'very-high', NULL, NULL, NULL),
	(376, '2', 'performanceTesting375', '1', 'very-high', NULL, NULL, NULL),
	(377, '2', 'performanceTesting376', '1', 'very-high', NULL, NULL, NULL),
	(378, '2', 'performanceTesting377', '1', 'very-high', NULL, NULL, NULL),
	(379, '2', 'performanceTesting378', '1', 'very-high', NULL, NULL, NULL),
	(380, '2', 'performanceTesting379', '1', 'very-high', NULL, NULL, NULL),
	(381, '2', 'performanceTesting380', '1', 'very-high', NULL, NULL, NULL),
	(382, '2', 'performanceTesting381', '1', 'very-high', NULL, NULL, NULL),
	(383, '2', 'performanceTesting382', '1', 'very-high', NULL, NULL, NULL),
	(384, '2', 'performanceTesting383', '1', 'very-high', NULL, NULL, NULL),
	(385, '2', 'performanceTesting384', '1', 'very-high', NULL, NULL, NULL),
	(386, '2', 'performanceTesting385', '1', 'very-high', NULL, NULL, NULL),
	(387, '2', 'performanceTesting386', '1', 'very-high', NULL, NULL, NULL),
	(388, '2', 'performanceTesting387', '1', 'very-high', NULL, NULL, NULL),
	(389, '2', 'performanceTesting388', '1', 'very-high', NULL, NULL, NULL),
	(390, '2', 'performanceTesting389', '1', 'very-high', NULL, NULL, NULL),
	(391, '2', 'performanceTesting390', '1', 'very-high', NULL, NULL, NULL),
	(392, '2', 'performanceTesting391', '1', 'very-high', NULL, NULL, NULL),
	(393, '2', 'performanceTesting392', '1', 'very-high', NULL, NULL, NULL),
	(394, '2', 'performanceTesting393', '1', 'very-high', NULL, NULL, NULL),
	(395, '2', 'performanceTesting394', '1', 'very-high', NULL, NULL, NULL),
	(396, '2', 'performanceTesting395', '1', 'very-high', NULL, NULL, NULL),
	(397, '2', 'performanceTesting396', '1', 'very-high', NULL, NULL, NULL),
	(398, '2', 'performanceTesting397', '1', 'very-high', NULL, NULL, NULL),
	(399, '2', 'performanceTesting398', '1', 'very-high', NULL, NULL, NULL),
	(400, '2', 'performanceTesting399', '1', 'very-high', NULL, NULL, NULL),
	(401, '2', 'performanceTesting400', '1', 'very-high', NULL, NULL, NULL),
	(402, '2', 'performanceTesting401', '1', 'very-high', NULL, NULL, NULL),
	(403, '2', 'performanceTesting402', '1', 'very-high', NULL, NULL, NULL),
	(404, '2', 'performanceTesting403', '1', 'very-high', NULL, NULL, NULL),
	(405, '2', 'performanceTesting404', '1', 'very-high', NULL, NULL, NULL),
	(406, '2', 'performanceTesting405', '1', 'very-high', NULL, NULL, NULL),
	(407, '2', 'performanceTesting406', '1', 'very-high', NULL, NULL, NULL),
	(408, '2', 'performanceTesting407', '1', 'very-high', NULL, NULL, NULL),
	(409, '2', 'performanceTesting408', '1', 'very-high', NULL, NULL, NULL),
	(410, '2', 'performanceTesting409', '1', 'very-high', NULL, NULL, NULL),
	(411, '2', 'performanceTesting410', '1', 'very-high', NULL, NULL, NULL),
	(412, '2', 'performanceTesting411', '1', 'very-high', NULL, NULL, NULL),
	(413, '2', 'performanceTesting412', '1', 'very-high', NULL, NULL, NULL),
	(414, '2', 'performanceTesting413', '1', 'very-high', NULL, NULL, NULL),
	(415, '2', 'performanceTesting414', '1', 'very-high', NULL, NULL, NULL),
	(416, '2', 'performanceTesting415', '1', 'very-high', NULL, NULL, NULL),
	(417, '2', 'performanceTesting416', '1', 'very-high', NULL, NULL, NULL),
	(418, '2', 'performanceTesting417', '1', 'very-high', NULL, NULL, NULL),
	(419, '2', 'performanceTesting418', '1', 'very-high', NULL, NULL, NULL),
	(420, '2', 'performanceTesting419', '1', 'very-high', NULL, NULL, NULL),
	(421, '2', 'performanceTesting420', '1', 'very-high', NULL, NULL, NULL),
	(422, '2', 'performanceTesting421', '1', 'very-high', NULL, NULL, NULL),
	(423, '2', 'performanceTesting422', '1', 'very-high', NULL, NULL, NULL),
	(424, '2', 'performanceTesting423', '1', 'very-high', NULL, NULL, NULL),
	(425, '2', 'performanceTesting424', '1', 'very-high', NULL, NULL, NULL),
	(426, '2', 'performanceTesting425', '1', 'very-high', NULL, NULL, NULL),
	(427, '2', 'performanceTesting426', '1', 'very-high', NULL, NULL, NULL),
	(428, '2', 'performanceTesting427', '1', 'very-high', NULL, NULL, NULL),
	(429, '2', 'performanceTesting428', '1', 'very-high', NULL, NULL, NULL),
	(430, '2', 'performanceTesting429', '1', 'very-high', NULL, NULL, NULL),
	(431, '2', 'performanceTesting430', '1', 'very-high', NULL, NULL, NULL),
	(432, '2', 'performanceTesting431', '1', 'very-high', NULL, NULL, NULL),
	(433, '2', 'performanceTesting432', '1', 'very-high', NULL, NULL, NULL),
	(434, '2', 'performanceTesting433', '1', 'very-high', NULL, NULL, NULL),
	(435, '2', 'performanceTesting434', '1', 'very-high', NULL, NULL, NULL),
	(436, '2', 'performanceTesting435', '1', 'very-high', NULL, NULL, NULL),
	(437, '2', 'performanceTesting436', '1', 'very-high', NULL, NULL, NULL),
	(438, '2', 'performanceTesting437', '1', 'very-high', NULL, NULL, NULL),
	(439, '2', 'performanceTesting438', '1', 'very-high', NULL, NULL, NULL),
	(440, '2', 'performanceTesting439', '1', 'very-high', NULL, NULL, NULL),
	(441, '2', 'performanceTesting440', '1', 'very-high', NULL, NULL, NULL),
	(442, '2', 'performanceTesting441', '1', 'very-high', NULL, NULL, NULL),
	(443, '2', 'performanceTesting442', '1', 'very-high', NULL, NULL, NULL),
	(444, '2', 'performanceTesting443', '1', 'very-high', NULL, NULL, NULL),
	(445, '2', 'performanceTesting444', '1', 'very-high', NULL, NULL, NULL),
	(446, '2', 'performanceTesting445', '1', 'very-high', NULL, NULL, NULL),
	(447, '2', 'performanceTesting446', '1', 'very-high', NULL, NULL, NULL),
	(448, '2', 'performanceTesting447', '1', 'very-high', NULL, NULL, NULL),
	(449, '2', 'performanceTesting448', '1', 'very-high', NULL, NULL, NULL),
	(450, '2', 'performanceTesting449', '1', 'very-high', NULL, NULL, NULL),
	(451, '2', 'performanceTesting450', '1', 'very-high', NULL, NULL, NULL),
	(452, '2', 'performanceTesting451', '1', 'very-high', NULL, NULL, NULL),
	(453, '2', 'performanceTesting452', '1', 'very-high', NULL, NULL, NULL),
	(454, '2', 'performanceTesting453', '1', 'very-high', NULL, NULL, NULL),
	(455, '2', 'performanceTesting454', '1', 'very-high', NULL, NULL, NULL),
	(456, '2', 'performanceTesting455', '1', 'very-high', NULL, NULL, NULL),
	(457, '2', 'performanceTesting456', '1', 'very-high', NULL, NULL, NULL),
	(458, '2', 'performanceTesting457', '1', 'very-high', NULL, NULL, NULL),
	(459, '2', 'performanceTesting458', '1', 'very-high', NULL, NULL, NULL),
	(460, '2', 'performanceTesting459', '1', 'very-high', NULL, NULL, NULL),
	(461, '2', 'performanceTesting460', '1', 'very-high', NULL, NULL, NULL),
	(462, '2', 'performanceTesting461', '1', 'very-high', NULL, NULL, NULL),
	(463, '2', 'performanceTesting462', '1', 'very-high', NULL, NULL, NULL),
	(464, '2', 'performanceTesting463', '1', 'very-high', NULL, NULL, NULL),
	(465, '2', 'performanceTesting464', '1', 'very-high', NULL, NULL, NULL),
	(466, '2', 'performanceTesting465', '1', 'very-high', NULL, NULL, NULL),
	(467, '2', 'performanceTesting466', '1', 'very-high', NULL, NULL, NULL),
	(468, '2', 'performanceTesting467', '1', 'very-high', NULL, NULL, NULL),
	(469, '2', 'performanceTesting468', '1', 'very-high', NULL, NULL, NULL),
	(470, '2', 'performanceTesting469', '1', 'very-high', NULL, NULL, NULL),
	(471, '2', 'performanceTesting470', '1', 'very-high', NULL, NULL, NULL),
	(472, '2', 'performanceTesting471', '1', 'very-high', NULL, NULL, NULL),
	(473, '2', 'performanceTesting472', '1', 'very-high', NULL, NULL, NULL),
	(474, '2', 'performanceTesting473', '1', 'very-high', NULL, NULL, NULL),
	(475, '2', 'performanceTesting474', '1', 'very-high', NULL, NULL, NULL),
	(476, '2', 'performanceTesting475', '1', 'very-high', NULL, NULL, NULL),
	(477, '2', 'performanceTesting476', '1', 'very-high', NULL, NULL, NULL),
	(478, '2', 'performanceTesting477', '1', 'very-high', NULL, NULL, NULL),
	(479, '2', 'performanceTesting478', '1', 'very-high', NULL, NULL, NULL),
	(480, '2', 'performanceTesting479', '1', 'very-high', NULL, NULL, NULL),
	(481, '2', 'performanceTesting480', '1', 'very-high', NULL, NULL, NULL),
	(482, '2', 'performanceTesting481', '1', 'very-high', NULL, NULL, NULL),
	(483, '2', 'performanceTesting482', '1', 'very-high', NULL, NULL, NULL),
	(484, '2', 'performanceTesting483', '1', 'very-high', NULL, NULL, NULL),
	(485, '2', 'performanceTesting484', '1', 'very-high', NULL, NULL, NULL),
	(486, '2', 'performanceTesting485', '1', 'very-high', NULL, NULL, NULL),
	(487, '2', 'performanceTesting486', '1', 'very-high', NULL, NULL, NULL),
	(488, '2', 'performanceTesting487', '1', 'very-high', NULL, NULL, NULL),
	(489, '2', 'performanceTesting488', '1', 'very-high', NULL, NULL, NULL),
	(490, '2', 'performanceTesting489', '1', 'very-high', NULL, NULL, NULL),
	(491, '2', 'performanceTesting490', '1', 'very-high', NULL, NULL, NULL),
	(492, '2', 'performanceTesting491', '1', 'very-high', NULL, NULL, NULL),
	(493, '2', 'performanceTesting492', '1', 'very-high', NULL, NULL, NULL),
	(494, '2', 'performanceTesting493', '1', 'very-high', NULL, NULL, NULL),
	(495, '2', 'performanceTesting494', '1', 'very-high', NULL, NULL, NULL),
	(496, '2', 'performanceTesting495', '1', 'very-high', NULL, NULL, NULL),
	(497, '2', 'performanceTesting496', '1', 'very-high', NULL, NULL, NULL),
	(498, '2', 'performanceTesting497', '1', 'very-high', NULL, NULL, NULL),
	(499, '2', 'performanceTesting498', '1', 'very-high', NULL, NULL, NULL),
	(500, '2', 'performanceTesting499', '1', 'very-high', NULL, NULL, NULL),
	(501, '2', 'performanceTesting500', '1', 'very-high', NULL, NULL, NULL),
	(502, '2', 'performanceTesting501', '1', 'very-high', NULL, NULL, NULL),
	(503, '2', 'performanceTesting502', '1', 'very-high', NULL, NULL, NULL),
	(504, '2', 'performanceTesting503', '1', 'very-high', NULL, NULL, NULL),
	(505, '2', 'performanceTesting504', '1', 'very-high', NULL, NULL, NULL),
	(506, '2', 'performanceTesting505', '1', 'very-high', NULL, NULL, NULL),
	(507, '2', 'performanceTesting506', '1', 'very-high', NULL, NULL, NULL),
	(508, '2', 'performanceTesting507', '1', 'very-high', NULL, NULL, NULL),
	(509, '2', 'performanceTesting508', '1', 'very-high', NULL, NULL, NULL),
	(510, '2', 'performanceTesting509', '1', 'very-high', NULL, NULL, NULL),
	(511, '2', 'performanceTesting510', '1', 'very-high', NULL, NULL, NULL),
	(512, '2', 'performanceTesting511', '1', 'very-high', NULL, NULL, NULL),
	(513, '2', 'performanceTesting512', '1', 'very-high', NULL, NULL, NULL),
	(514, '2', 'performanceTesting513', '1', 'very-high', NULL, NULL, NULL),
	(515, '2', 'performanceTesting514', '1', 'very-high', NULL, NULL, NULL),
	(516, '2', 'performanceTesting515', '1', 'very-high', NULL, NULL, NULL),
	(517, '2', 'performanceTesting516', '1', 'very-high', NULL, NULL, NULL),
	(518, '2', 'performanceTesting517', '1', 'very-high', NULL, NULL, NULL),
	(519, '2', 'performanceTesting518', '1', 'very-high', NULL, NULL, NULL),
	(520, '2', 'performanceTesting519', '1', 'very-high', NULL, NULL, NULL),
	(521, '2', 'performanceTesting520', '1', 'very-high', NULL, NULL, NULL),
	(522, '2', 'performanceTesting521', '1', 'very-high', NULL, NULL, NULL),
	(523, '2', 'performanceTesting522', '1', 'very-high', NULL, NULL, NULL),
	(524, '2', 'performanceTesting523', '1', 'very-high', NULL, NULL, NULL),
	(525, '2', 'performanceTesting524', '1', 'very-high', NULL, NULL, NULL),
	(526, '2', 'performanceTesting525', '1', 'very-high', NULL, NULL, NULL),
	(527, '2', 'performanceTesting526', '1', 'very-high', NULL, NULL, NULL),
	(528, '2', 'performanceTesting527', '1', 'very-high', NULL, NULL, NULL),
	(529, '2', 'performanceTesting528', '1', 'very-high', NULL, NULL, NULL),
	(530, '2', 'performanceTesting529', '1', 'very-high', NULL, NULL, NULL),
	(531, '2', 'performanceTesting530', '1', 'very-high', NULL, NULL, NULL),
	(532, '2', 'performanceTesting531', '1', 'very-high', NULL, NULL, NULL),
	(533, '2', 'performanceTesting532', '1', 'very-high', NULL, NULL, NULL),
	(534, '2', 'performanceTesting533', '1', 'very-high', NULL, NULL, NULL),
	(535, '2', 'performanceTesting534', '1', 'very-high', NULL, NULL, NULL),
	(536, '2', 'performanceTesting535', '1', 'very-high', NULL, NULL, NULL),
	(537, '2', 'performanceTesting536', '1', 'very-high', NULL, NULL, NULL),
	(538, '2', 'performanceTesting537', '1', 'very-high', NULL, NULL, NULL),
	(539, '2', 'performanceTesting538', '1', 'very-high', NULL, NULL, NULL),
	(540, '2', 'performanceTesting539', '1', 'very-high', NULL, NULL, NULL),
	(541, '2', 'performanceTesting540', '1', 'very-high', NULL, NULL, NULL),
	(542, '2', 'performanceTesting541', '1', 'very-high', NULL, NULL, NULL),
	(543, '2', 'performanceTesting542', '1', 'very-high', NULL, NULL, NULL),
	(544, '2', 'performanceTesting543', '1', 'very-high', NULL, NULL, NULL),
	(545, '2', 'performanceTesting544', '1', 'very-high', NULL, NULL, NULL),
	(546, '2', 'performanceTesting545', '1', 'very-high', NULL, NULL, NULL),
	(547, '2', 'performanceTesting546', '1', 'very-high', NULL, NULL, NULL),
	(548, '2', 'performanceTesting547', '1', 'very-high', NULL, NULL, NULL),
	(549, '2', 'performanceTesting548', '1', 'very-high', NULL, NULL, NULL),
	(550, '2', 'performanceTesting549', '1', 'very-high', NULL, NULL, NULL),
	(551, '2', 'performanceTesting550', '1', 'very-high', NULL, NULL, NULL),
	(552, '2', 'performanceTesting551', '1', 'very-high', NULL, NULL, NULL),
	(553, '2', 'performanceTesting552', '1', 'very-high', NULL, NULL, NULL),
	(554, '2', 'performanceTesting553', '1', 'very-high', NULL, NULL, NULL),
	(555, '2', 'performanceTesting554', '1', 'very-high', NULL, NULL, NULL),
	(556, '2', 'performanceTesting555', '1', 'very-high', NULL, NULL, NULL),
	(557, '2', 'performanceTesting556', '1', 'very-high', NULL, NULL, NULL),
	(558, '2', 'performanceTesting557', '1', 'very-high', NULL, NULL, NULL),
	(559, '2', 'performanceTesting558', '1', 'very-high', NULL, NULL, NULL),
	(560, '2', 'performanceTesting559', '1', 'very-high', NULL, NULL, NULL),
	(561, '2', 'performanceTesting560', '1', 'very-high', NULL, NULL, NULL),
	(562, '2', 'performanceTesting561', '1', 'very-high', NULL, NULL, NULL),
	(563, '2', 'performanceTesting562', '1', 'very-high', NULL, NULL, NULL),
	(564, '2', 'performanceTesting563', '1', 'very-high', NULL, NULL, NULL),
	(565, '2', 'performanceTesting564', '1', 'very-high', NULL, NULL, NULL),
	(566, '2', 'performanceTesting565', '1', 'very-high', NULL, NULL, NULL),
	(567, '2', 'performanceTesting566', '1', 'very-high', NULL, NULL, NULL),
	(568, '2', 'performanceTesting567', '1', 'very-high', NULL, NULL, NULL),
	(569, '2', 'performanceTesting568', '1', 'very-high', NULL, NULL, NULL),
	(570, '2', 'performanceTesting569', '1', 'very-high', NULL, NULL, NULL),
	(571, '2', 'performanceTesting570', '1', 'very-high', NULL, NULL, NULL),
	(572, '2', 'performanceTesting571', '1', 'very-high', NULL, NULL, NULL),
	(573, '2', 'performanceTesting572', '1', 'very-high', NULL, NULL, NULL),
	(574, '2', 'performanceTesting573', '1', 'very-high', NULL, NULL, NULL),
	(575, '2', 'performanceTesting574', '1', 'very-high', NULL, NULL, NULL),
	(576, '2', 'performanceTesting575', '1', 'very-high', NULL, NULL, NULL),
	(577, '2', 'performanceTesting576', '1', 'very-high', NULL, NULL, NULL),
	(578, '2', 'performanceTesting577', '1', 'very-high', NULL, NULL, NULL),
	(579, '2', 'performanceTesting578', '1', 'very-high', NULL, NULL, NULL),
	(580, '2', 'performanceTesting579', '1', 'very-high', NULL, NULL, NULL),
	(581, '2', 'performanceTesting580', '1', 'very-high', NULL, NULL, NULL),
	(582, '2', 'performanceTesting581', '1', 'very-high', NULL, NULL, NULL),
	(583, '2', 'performanceTesting582', '1', 'very-high', NULL, NULL, NULL),
	(584, '2', 'performanceTesting583', '1', 'very-high', NULL, NULL, NULL),
	(585, '2', 'performanceTesting584', '1', 'very-high', NULL, NULL, NULL),
	(586, '2', 'performanceTesting585', '1', 'very-high', NULL, NULL, NULL),
	(587, '2', 'performanceTesting586', '1', 'very-high', NULL, NULL, NULL),
	(588, '2', 'performanceTesting587', '1', 'very-high', NULL, NULL, NULL),
	(589, '2', 'performanceTesting588', '1', 'very-high', NULL, NULL, NULL),
	(590, '2', 'performanceTesting589', '1', 'very-high', NULL, NULL, NULL),
	(591, '2', 'performanceTesting590', '1', 'very-high', NULL, NULL, NULL),
	(592, '2', 'performanceTesting591', '1', 'very-high', NULL, NULL, NULL),
	(593, '2', 'performanceTesting592', '1', 'very-high', NULL, NULL, NULL),
	(594, '2', 'performanceTesting593', '1', 'very-high', NULL, NULL, NULL),
	(595, '2', 'performanceTesting594', '1', 'very-high', NULL, NULL, NULL),
	(596, '2', 'performanceTesting595', '1', 'very-high', NULL, NULL, NULL),
	(597, '2', 'performanceTesting596', '1', 'very-high', NULL, NULL, NULL),
	(598, '2', 'performanceTesting597', '1', 'very-high', NULL, NULL, NULL),
	(599, '2', 'performanceTesting598', '1', 'very-high', NULL, NULL, NULL),
	(600, '2', 'performanceTesting599', '1', 'very-high', NULL, NULL, NULL),
	(601, '2', 'performanceTesting600', '1', 'very-high', NULL, NULL, NULL),
	(602, '2', 'performanceTesting601', '1', 'very-high', NULL, NULL, NULL),
	(603, '2', 'performanceTesting602', '1', 'very-high', NULL, NULL, NULL),
	(604, '2', 'performanceTesting603', '1', 'very-high', NULL, NULL, NULL),
	(605, '2', 'performanceTesting604', '1', 'very-high', NULL, NULL, NULL),
	(606, '2', 'performanceTesting605', '1', 'very-high', NULL, NULL, NULL),
	(607, '2', 'performanceTesting606', '1', 'very-high', NULL, NULL, NULL),
	(608, '2', 'performanceTesting607', '1', 'very-high', NULL, NULL, NULL),
	(609, '2', 'performanceTesting608', '1', 'very-high', NULL, NULL, NULL),
	(610, '2', 'performanceTesting609', '1', 'very-high', NULL, NULL, NULL),
	(611, '2', 'performanceTesting610', '1', 'very-high', NULL, NULL, NULL),
	(612, '2', 'performanceTesting611', '1', 'very-high', NULL, NULL, NULL),
	(613, '2', 'performanceTesting612', '1', 'very-high', NULL, NULL, NULL),
	(614, '2', 'performanceTesting613', '1', 'very-high', NULL, NULL, NULL),
	(615, '2', 'performanceTesting614', '1', 'very-high', NULL, NULL, NULL),
	(616, '2', 'performanceTesting615', '1', 'very-high', NULL, NULL, NULL),
	(617, '2', 'performanceTesting616', '1', 'very-high', NULL, NULL, NULL),
	(618, '2', 'performanceTesting617', '1', 'very-high', NULL, NULL, NULL),
	(619, '2', 'performanceTesting618', '1', 'very-high', NULL, NULL, NULL),
	(620, '2', 'performanceTesting619', '1', 'very-high', NULL, NULL, NULL),
	(621, '2', 'performanceTesting620', '1', 'very-high', NULL, NULL, NULL),
	(622, '2', 'performanceTesting621', '1', 'very-high', NULL, NULL, NULL),
	(623, '2', 'performanceTesting622', '1', 'very-high', NULL, NULL, NULL),
	(624, '2', 'performanceTesting623', '1', 'very-high', NULL, NULL, NULL),
	(625, '2', 'performanceTesting624', '1', 'very-high', NULL, NULL, NULL),
	(626, '2', 'performanceTesting625', '1', 'very-high', NULL, NULL, NULL),
	(627, '2', 'performanceTesting626', '1', 'very-high', NULL, NULL, NULL),
	(628, '2', 'performanceTesting627', '1', 'very-high', NULL, NULL, NULL),
	(629, '2', 'performanceTesting628', '1', 'very-high', NULL, NULL, NULL),
	(630, '2', 'performanceTesting629', '1', 'very-high', NULL, NULL, NULL),
	(631, '2', 'performanceTesting630', '1', 'very-high', NULL, NULL, NULL),
	(632, '2', 'performanceTesting631', '1', 'very-high', NULL, NULL, NULL),
	(633, '2', 'performanceTesting632', '1', 'very-high', NULL, NULL, NULL),
	(634, '2', 'performanceTesting633', '1', 'very-high', NULL, NULL, NULL),
	(635, '2', 'performanceTesting634', '1', 'very-high', NULL, NULL, NULL),
	(636, '2', 'performanceTesting635', '1', 'very-high', NULL, NULL, NULL),
	(637, '2', 'performanceTesting636', '1', 'very-high', NULL, NULL, NULL),
	(638, '2', 'performanceTesting637', '1', 'very-high', NULL, NULL, NULL),
	(639, '2', 'performanceTesting638', '1', 'very-high', NULL, NULL, NULL),
	(640, '2', 'performanceTesting639', '1', 'very-high', NULL, NULL, NULL),
	(641, '2', 'performanceTesting640', '1', 'very-high', NULL, NULL, NULL),
	(642, '2', 'performanceTesting641', '1', 'very-high', NULL, NULL, NULL),
	(643, '2', 'performanceTesting642', '1', 'very-high', NULL, NULL, NULL),
	(644, '2', 'performanceTesting643', '1', 'very-high', NULL, NULL, NULL),
	(645, '2', 'performanceTesting644', '1', 'very-high', NULL, NULL, NULL),
	(646, '2', 'performanceTesting645', '1', 'very-high', NULL, NULL, NULL),
	(647, '2', 'performanceTesting646', '1', 'very-high', NULL, NULL, NULL),
	(648, '2', 'performanceTesting647', '1', 'very-high', NULL, NULL, NULL),
	(649, '2', 'performanceTesting648', '1', 'very-high', NULL, NULL, NULL),
	(650, '2', 'performanceTesting649', '1', 'very-high', NULL, NULL, NULL),
	(651, '2', 'performanceTesting650', '1', 'very-high', NULL, NULL, NULL),
	(652, '2', 'performanceTesting651', '1', 'very-high', NULL, NULL, NULL),
	(653, '2', 'performanceTesting652', '1', 'very-high', NULL, NULL, NULL),
	(654, '2', 'performanceTesting653', '1', 'very-high', NULL, NULL, NULL),
	(655, '2', 'performanceTesting654', '1', 'very-high', NULL, NULL, NULL),
	(656, '2', 'performanceTesting655', '1', 'very-high', NULL, NULL, NULL),
	(657, '2', 'performanceTesting656', '1', 'very-high', NULL, NULL, NULL),
	(658, '2', 'performanceTesting657', '1', 'very-high', NULL, NULL, NULL),
	(659, '2', 'performanceTesting658', '1', 'very-high', NULL, NULL, NULL),
	(660, '2', 'performanceTesting659', '1', 'very-high', NULL, NULL, NULL),
	(661, '2', 'performanceTesting660', '1', 'very-high', NULL, NULL, NULL),
	(662, '2', 'performanceTesting661', '1', 'very-high', NULL, NULL, NULL),
	(663, '2', 'performanceTesting662', '1', 'very-high', NULL, NULL, NULL),
	(664, '2', 'performanceTesting663', '1', 'very-high', NULL, NULL, NULL),
	(665, '2', 'performanceTesting664', '1', 'very-high', NULL, NULL, NULL),
	(666, '2', 'performanceTesting665', '1', 'very-high', NULL, NULL, NULL),
	(667, '2', 'performanceTesting666', '1', 'very-high', NULL, NULL, NULL),
	(668, '2', 'performanceTesting667', '1', 'very-high', NULL, NULL, NULL),
	(669, '2', 'performanceTesting668', '1', 'very-high', NULL, NULL, NULL),
	(670, '2', 'performanceTesting669', '1', 'very-high', NULL, NULL, NULL),
	(671, '2', 'performanceTesting670', '1', 'very-high', NULL, NULL, NULL),
	(672, '2', 'performanceTesting671', '1', 'very-high', NULL, NULL, NULL),
	(673, '2', 'performanceTesting672', '1', 'very-high', NULL, NULL, NULL),
	(674, '2', 'performanceTesting673', '1', 'very-high', NULL, NULL, NULL),
	(675, '2', 'performanceTesting674', '1', 'very-high', NULL, NULL, NULL),
	(676, '2', 'performanceTesting675', '1', 'very-high', NULL, NULL, NULL),
	(677, '2', 'performanceTesting676', '1', 'very-high', NULL, NULL, NULL),
	(678, '2', 'performanceTesting677', '1', 'very-high', NULL, NULL, NULL),
	(679, '2', 'performanceTesting678', '1', 'very-high', NULL, NULL, NULL),
	(680, '2', 'performanceTesting679', '1', 'very-high', NULL, NULL, NULL),
	(681, '2', 'performanceTesting680', '1', 'very-high', NULL, NULL, NULL),
	(682, '2', 'performanceTesting681', '1', 'very-high', NULL, NULL, NULL),
	(683, '2', 'performanceTesting682', '1', 'very-high', NULL, NULL, NULL),
	(684, '2', 'performanceTesting683', '1', 'very-high', NULL, NULL, NULL),
	(685, '2', 'performanceTesting684', '1', 'very-high', NULL, NULL, NULL),
	(686, '2', 'performanceTesting685', '1', 'very-high', NULL, NULL, NULL),
	(687, '2', 'performanceTesting686', '1', 'very-high', NULL, NULL, NULL),
	(688, '2', 'performanceTesting687', '1', 'very-high', NULL, NULL, NULL),
	(689, '2', 'performanceTesting688', '1', 'very-high', NULL, NULL, NULL),
	(690, '2', 'performanceTesting689', '1', 'very-high', NULL, NULL, NULL),
	(691, '2', 'performanceTesting690', '1', 'very-high', NULL, NULL, NULL),
	(692, '2', 'performanceTesting691', '1', 'very-high', NULL, NULL, NULL),
	(693, '2', 'performanceTesting692', '1', 'very-high', NULL, NULL, NULL),
	(694, '2', 'performanceTesting693', '1', 'very-high', NULL, NULL, NULL),
	(695, '2', 'performanceTesting694', '1', 'very-high', NULL, NULL, NULL),
	(696, '2', 'performanceTesting695', '1', 'very-high', NULL, NULL, NULL),
	(697, '2', 'performanceTesting696', '1', 'very-high', NULL, NULL, NULL),
	(698, '2', 'performanceTesting697', '1', 'very-high', NULL, NULL, NULL),
	(699, '2', 'performanceTesting698', '1', 'very-high', NULL, NULL, NULL),
	(700, '2', 'performanceTesting699', '1', 'very-high', NULL, NULL, NULL),
	(701, '2', 'performanceTesting700', '1', 'very-high', NULL, NULL, NULL),
	(702, '2', 'performanceTesting701', '1', 'very-high', NULL, NULL, NULL),
	(703, '2', 'performanceTesting702', '1', 'very-high', NULL, NULL, NULL),
	(704, '2', 'performanceTesting703', '1', 'very-high', NULL, NULL, NULL),
	(705, '2', 'performanceTesting704', '1', 'very-high', NULL, NULL, NULL),
	(706, '2', 'performanceTesting705', '1', 'very-high', NULL, NULL, NULL),
	(707, '2', 'performanceTesting706', '1', 'very-high', NULL, NULL, NULL),
	(708, '2', 'performanceTesting707', '1', 'very-high', NULL, NULL, NULL),
	(709, '2', 'performanceTesting708', '1', 'very-high', NULL, NULL, NULL),
	(710, '2', 'performanceTesting709', '1', 'very-high', NULL, NULL, NULL),
	(711, '2', 'performanceTesting710', '1', 'very-high', NULL, NULL, NULL),
	(712, '2', 'performanceTesting711', '1', 'very-high', NULL, NULL, NULL),
	(713, '2', 'performanceTesting712', '1', 'very-high', NULL, NULL, NULL),
	(714, '2', 'performanceTesting713', '1', 'very-high', NULL, NULL, NULL),
	(715, '2', 'performanceTesting714', '1', 'very-high', NULL, NULL, NULL),
	(716, '2', 'performanceTesting715', '1', 'very-high', NULL, NULL, NULL),
	(717, '2', 'performanceTesting716', '1', 'very-high', NULL, NULL, NULL),
	(718, '2', 'performanceTesting717', '1', 'very-high', NULL, NULL, NULL),
	(719, '2', 'performanceTesting718', '1', 'very-high', NULL, NULL, NULL),
	(720, '2', 'performanceTesting719', '1', 'very-high', NULL, NULL, NULL),
	(721, '2', 'performanceTesting720', '1', 'very-high', NULL, NULL, NULL),
	(722, '2', 'performanceTesting721', '1', 'very-high', NULL, NULL, NULL),
	(723, '2', 'performanceTesting722', '1', 'very-high', NULL, NULL, NULL),
	(724, '2', 'performanceTesting723', '1', 'very-high', NULL, NULL, NULL),
	(725, '2', 'performanceTesting724', '1', 'very-high', NULL, NULL, NULL),
	(726, '2', 'performanceTesting725', '1', 'very-high', NULL, NULL, NULL),
	(727, '2', 'performanceTesting726', '1', 'very-high', NULL, NULL, NULL),
	(728, '2', 'performanceTesting727', '1', 'very-high', NULL, NULL, NULL),
	(729, '2', 'performanceTesting728', '1', 'very-high', NULL, NULL, NULL),
	(730, '2', 'performanceTesting729', '1', 'very-high', NULL, NULL, NULL),
	(731, '2', 'performanceTesting730', '1', 'very-high', NULL, NULL, NULL),
	(732, '2', 'performanceTesting731', '1', 'very-high', NULL, NULL, NULL),
	(733, '2', 'performanceTesting732', '1', 'very-high', NULL, NULL, NULL),
	(734, '2', 'performanceTesting733', '1', 'very-high', NULL, NULL, NULL),
	(735, '2', 'performanceTesting734', '1', 'very-high', NULL, NULL, NULL),
	(736, '2', 'performanceTesting735', '1', 'very-high', NULL, NULL, NULL),
	(737, '2', 'performanceTesting736', '1', 'very-high', NULL, NULL, NULL),
	(738, '2', 'performanceTesting737', '1', 'very-high', NULL, NULL, NULL),
	(739, '2', 'performanceTesting738', '1', 'very-high', NULL, NULL, NULL),
	(740, '2', 'performanceTesting739', '1', 'very-high', NULL, NULL, NULL),
	(741, '2', 'performanceTesting740', '1', 'very-high', NULL, NULL, NULL),
	(742, '2', 'performanceTesting741', '1', 'very-high', NULL, NULL, NULL),
	(743, '2', 'performanceTesting742', '1', 'very-high', NULL, NULL, NULL),
	(744, '2', 'performanceTesting743', '1', 'very-high', NULL, NULL, NULL),
	(745, '2', 'performanceTesting744', '1', 'very-high', NULL, NULL, NULL),
	(746, '2', 'performanceTesting745', '1', 'very-high', NULL, NULL, NULL),
	(747, '2', 'performanceTesting746', '1', 'very-high', NULL, NULL, NULL),
	(748, '2', 'performanceTesting747', '1', 'very-high', NULL, NULL, NULL),
	(749, '2', 'performanceTesting748', '1', 'very-high', NULL, NULL, NULL),
	(750, '2', 'performanceTesting749', '1', 'very-high', NULL, NULL, NULL),
	(751, '2', 'performanceTesting750', '1', 'very-high', NULL, NULL, NULL),
	(752, '2', 'performanceTesting751', '1', 'very-high', NULL, NULL, NULL),
	(753, '2', 'performanceTesting752', '1', 'very-high', NULL, NULL, NULL),
	(754, '2', 'performanceTesting753', '1', 'very-high', NULL, NULL, NULL),
	(755, '2', 'performanceTesting754', '1', 'very-high', NULL, NULL, NULL),
	(756, '2', 'performanceTesting755', '1', 'very-high', NULL, NULL, NULL),
	(757, '2', 'performanceTesting756', '1', 'very-high', NULL, NULL, NULL),
	(758, '2', 'performanceTesting757', '1', 'very-high', NULL, NULL, NULL),
	(759, '2', 'performanceTesting758', '1', 'very-high', NULL, NULL, NULL),
	(760, '2', 'performanceTesting759', '1', 'very-high', NULL, NULL, NULL),
	(761, '2', 'performanceTesting760', '1', 'very-high', NULL, NULL, NULL),
	(762, '2', 'performanceTesting761', '1', 'very-high', NULL, NULL, NULL),
	(763, '2', 'performanceTesting762', '1', 'very-high', NULL, NULL, NULL),
	(764, '2', 'performanceTesting763', '1', 'very-high', NULL, NULL, NULL),
	(765, '2', 'performanceTesting764', '1', 'very-high', NULL, NULL, NULL),
	(766, '2', 'performanceTesting765', '1', 'very-high', NULL, NULL, NULL),
	(767, '2', 'performanceTesting766', '1', 'very-high', NULL, NULL, NULL),
	(768, '2', 'performanceTesting767', '1', 'very-high', NULL, NULL, NULL),
	(769, '2', 'performanceTesting768', '1', 'very-high', NULL, NULL, NULL),
	(770, '2', 'performanceTesting769', '1', 'very-high', NULL, NULL, NULL),
	(771, '2', 'performanceTesting770', '1', 'very-high', NULL, NULL, NULL),
	(772, '2', 'performanceTesting771', '1', 'very-high', NULL, NULL, NULL),
	(773, '2', 'performanceTesting772', '1', 'very-high', NULL, NULL, NULL),
	(774, '2', 'performanceTesting773', '1', 'very-high', NULL, NULL, NULL),
	(775, '2', 'performanceTesting774', '1', 'very-high', NULL, NULL, NULL),
	(776, '2', 'performanceTesting775', '1', 'very-high', NULL, NULL, NULL),
	(777, '2', 'performanceTesting776', '1', 'very-high', NULL, NULL, NULL),
	(778, '2', 'performanceTesting777', '1', 'very-high', NULL, NULL, NULL),
	(779, '2', 'performanceTesting778', '1', 'very-high', NULL, NULL, NULL),
	(780, '2', 'performanceTesting779', '1', 'very-high', NULL, NULL, NULL),
	(781, '2', 'performanceTesting780', '1', 'very-high', NULL, NULL, NULL),
	(782, '2', 'performanceTesting781', '1', 'very-high', NULL, NULL, NULL),
	(783, '2', 'performanceTesting782', '1', 'very-high', NULL, NULL, NULL),
	(784, '2', 'performanceTesting783', '1', 'very-high', NULL, NULL, NULL),
	(785, '2', 'performanceTesting784', '1', 'very-high', NULL, NULL, NULL),
	(786, '2', 'performanceTesting785', '1', 'very-high', NULL, NULL, NULL),
	(787, '2', 'performanceTesting786', '1', 'very-high', NULL, NULL, NULL),
	(788, '2', 'performanceTesting787', '1', 'very-high', NULL, NULL, NULL),
	(789, '2', 'performanceTesting788', '1', 'very-high', NULL, NULL, NULL),
	(790, '2', 'performanceTesting789', '1', 'very-high', NULL, NULL, NULL),
	(791, '2', 'performanceTesting790', '1', 'very-high', NULL, NULL, NULL),
	(792, '2', 'performanceTesting791', '1', 'very-high', NULL, NULL, NULL),
	(793, '2', 'performanceTesting792', '1', 'very-high', NULL, NULL, NULL),
	(794, '2', 'performanceTesting793', '1', 'very-high', NULL, NULL, NULL),
	(795, '2', 'performanceTesting794', '1', 'very-high', NULL, NULL, NULL),
	(796, '2', 'performanceTesting795', '1', 'very-high', NULL, NULL, NULL),
	(797, '2', 'performanceTesting796', '1', 'very-high', NULL, NULL, NULL),
	(798, '2', 'performanceTesting797', '1', 'very-high', NULL, NULL, NULL),
	(799, '2', 'performanceTesting798', '1', 'very-high', NULL, NULL, NULL),
	(800, '2', 'performanceTesting799', '1', 'very-high', NULL, NULL, NULL),
	(801, '2', 'performanceTesting800', '1', 'very-high', NULL, NULL, NULL),
	(802, '2', 'performanceTesting801', '1', 'very-high', NULL, NULL, NULL),
	(803, '2', 'performanceTesting802', '1', 'very-high', NULL, NULL, NULL),
	(804, '2', 'performanceTesting803', '1', 'very-high', NULL, NULL, NULL),
	(805, '2', 'performanceTesting804', '1', 'very-high', NULL, NULL, NULL),
	(806, '2', 'performanceTesting805', '1', 'very-high', NULL, NULL, NULL),
	(807, '2', 'performanceTesting806', '1', 'very-high', NULL, NULL, NULL),
	(808, '2', 'performanceTesting807', '1', 'very-high', NULL, NULL, NULL),
	(809, '2', 'performanceTesting808', '1', 'very-high', NULL, NULL, NULL),
	(810, '2', 'performanceTesting809', '1', 'very-high', NULL, NULL, NULL),
	(811, '2', 'performanceTesting810', '1', 'very-high', NULL, NULL, NULL),
	(812, '2', 'performanceTesting811', '1', 'very-high', NULL, NULL, NULL),
	(813, '2', 'performanceTesting812', '1', 'very-high', NULL, NULL, NULL),
	(814, '2', 'performanceTesting813', '1', 'very-high', NULL, NULL, NULL),
	(815, '2', 'performanceTesting814', '1', 'very-high', NULL, NULL, NULL),
	(816, '2', 'performanceTesting815', '1', 'very-high', NULL, NULL, NULL),
	(817, '2', 'performanceTesting816', '1', 'very-high', NULL, NULL, NULL),
	(818, '2', 'performanceTesting817', '1', 'very-high', NULL, NULL, NULL),
	(819, '2', 'performanceTesting818', '1', 'very-high', NULL, NULL, NULL),
	(820, '2', 'performanceTesting819', '1', 'very-high', NULL, NULL, NULL),
	(821, '2', 'performanceTesting820', '1', 'very-high', NULL, NULL, NULL),
	(822, '2', 'performanceTesting821', '1', 'very-high', NULL, NULL, NULL),
	(823, '2', 'performanceTesting822', '1', 'very-high', NULL, NULL, NULL),
	(824, '2', 'performanceTesting823', '1', 'very-high', NULL, NULL, NULL),
	(825, '2', 'performanceTesting824', '1', 'very-high', NULL, NULL, NULL),
	(826, '2', 'performanceTesting825', '1', 'very-high', NULL, NULL, NULL),
	(827, '2', 'performanceTesting826', '1', 'very-high', NULL, NULL, NULL),
	(828, '2', 'performanceTesting827', '1', 'very-high', NULL, NULL, NULL),
	(829, '2', 'performanceTesting828', '1', 'very-high', NULL, NULL, NULL),
	(830, '2', 'performanceTesting829', '1', 'very-high', NULL, NULL, NULL),
	(831, '2', 'performanceTesting830', '1', 'very-high', NULL, NULL, NULL),
	(832, '2', 'performanceTesting831', '1', 'very-high', NULL, NULL, NULL),
	(833, '2', 'performanceTesting832', '1', 'very-high', NULL, NULL, NULL),
	(834, '2', 'performanceTesting833', '1', 'very-high', NULL, NULL, NULL),
	(835, '2', 'performanceTesting834', '1', 'very-high', NULL, NULL, NULL),
	(836, '2', 'performanceTesting835', '1', 'very-high', NULL, NULL, NULL),
	(837, '2', 'performanceTesting836', '1', 'very-high', NULL, NULL, NULL),
	(838, '2', 'performanceTesting837', '1', 'very-high', NULL, NULL, NULL),
	(839, '2', 'performanceTesting838', '1', 'very-high', NULL, NULL, NULL),
	(840, '2', 'performanceTesting839', '1', 'very-high', NULL, NULL, NULL),
	(841, '2', 'performanceTesting840', '1', 'very-high', NULL, NULL, NULL),
	(842, '2', 'performanceTesting841', '1', 'very-high', NULL, NULL, NULL),
	(843, '2', 'performanceTesting842', '1', 'very-high', NULL, NULL, NULL),
	(844, '2', 'performanceTesting843', '1', 'very-high', NULL, NULL, NULL),
	(845, '2', 'performanceTesting844', '1', 'very-high', NULL, NULL, NULL),
	(846, '2', 'performanceTesting845', '1', 'very-high', NULL, NULL, NULL),
	(847, '2', 'performanceTesting846', '1', 'very-high', NULL, NULL, NULL),
	(848, '2', 'performanceTesting847', '1', 'very-high', NULL, NULL, NULL),
	(849, '2', 'performanceTesting848', '1', 'very-high', NULL, NULL, NULL),
	(850, '2', 'performanceTesting849', '1', 'very-high', NULL, NULL, NULL),
	(851, '2', 'performanceTesting850', '1', 'very-high', NULL, NULL, NULL),
	(852, '2', 'performanceTesting851', '1', 'very-high', NULL, NULL, NULL),
	(853, '2', 'performanceTesting852', '1', 'very-high', NULL, NULL, NULL),
	(854, '2', 'performanceTesting853', '1', 'very-high', NULL, NULL, NULL),
	(855, '2', 'performanceTesting854', '1', 'very-high', NULL, NULL, NULL),
	(856, '2', 'performanceTesting855', '1', 'very-high', NULL, NULL, NULL),
	(857, '2', 'performanceTesting856', '1', 'very-high', NULL, NULL, NULL),
	(858, '2', 'performanceTesting857', '1', 'very-high', NULL, NULL, NULL),
	(859, '2', 'performanceTesting858', '1', 'very-high', NULL, NULL, NULL),
	(860, '2', 'performanceTesting859', '1', 'very-high', NULL, NULL, NULL),
	(861, '2', 'performanceTesting860', '1', 'very-high', NULL, NULL, NULL),
	(862, '2', 'performanceTesting861', '1', 'very-high', NULL, NULL, NULL),
	(863, '2', 'performanceTesting862', '1', 'very-high', NULL, NULL, NULL),
	(864, '2', 'performanceTesting863', '1', 'very-high', NULL, NULL, NULL),
	(865, '2', 'performanceTesting864', '1', 'very-high', NULL, NULL, NULL),
	(866, '2', 'performanceTesting865', '1', 'very-high', NULL, NULL, NULL),
	(867, '2', 'performanceTesting866', '1', 'very-high', NULL, NULL, NULL),
	(868, '2', 'performanceTesting867', '1', 'very-high', NULL, NULL, NULL),
	(869, '2', 'performanceTesting868', '1', 'very-high', NULL, NULL, NULL),
	(870, '2', 'performanceTesting869', '1', 'very-high', NULL, NULL, NULL),
	(871, '2', 'performanceTesting870', '1', 'very-high', NULL, NULL, NULL),
	(872, '2', 'performanceTesting871', '1', 'very-high', NULL, NULL, NULL),
	(873, '2', 'performanceTesting872', '1', 'very-high', NULL, NULL, NULL),
	(874, '2', 'performanceTesting873', '1', 'very-high', NULL, NULL, NULL),
	(875, '2', 'performanceTesting874', '1', 'very-high', NULL, NULL, NULL),
	(876, '2', 'performanceTesting875', '1', 'very-high', NULL, NULL, NULL),
	(877, '2', 'performanceTesting876', '1', 'very-high', NULL, NULL, NULL),
	(878, '2', 'performanceTesting877', '1', 'very-high', NULL, NULL, NULL),
	(879, '2', 'performanceTesting878', '1', 'very-high', NULL, NULL, NULL),
	(880, '2', 'performanceTesting879', '1', 'very-high', NULL, NULL, NULL),
	(881, '2', 'performanceTesting880', '1', 'very-high', NULL, NULL, NULL),
	(882, '2', 'performanceTesting881', '1', 'very-high', NULL, NULL, NULL),
	(883, '2', 'performanceTesting882', '1', 'very-high', NULL, NULL, NULL),
	(884, '2', 'performanceTesting883', '1', 'very-high', NULL, NULL, NULL),
	(885, '2', 'performanceTesting884', '1', 'very-high', NULL, NULL, NULL),
	(886, '2', 'performanceTesting885', '1', 'very-high', NULL, NULL, NULL),
	(887, '2', 'performanceTesting886', '1', 'very-high', NULL, NULL, NULL),
	(888, '2', 'performanceTesting887', '1', 'very-high', NULL, NULL, NULL),
	(889, '2', 'performanceTesting888', '1', 'very-high', NULL, NULL, NULL),
	(890, '2', 'performanceTesting889', '1', 'very-high', NULL, NULL, NULL),
	(891, '2', 'performanceTesting890', '1', 'very-high', NULL, NULL, NULL),
	(892, '2', 'performanceTesting891', '1', 'very-high', NULL, NULL, NULL),
	(893, '2', 'performanceTesting892', '1', 'very-high', NULL, NULL, NULL),
	(894, '2', 'performanceTesting893', '1', 'very-high', NULL, NULL, NULL),
	(895, '2', 'performanceTesting894', '1', 'very-high', NULL, NULL, NULL),
	(896, '2', 'performanceTesting895', '1', 'very-high', NULL, NULL, NULL),
	(897, '2', 'performanceTesting896', '1', 'very-high', NULL, NULL, NULL),
	(898, '2', 'performanceTesting897', '1', 'very-high', NULL, NULL, NULL),
	(899, '2', 'performanceTesting898', '1', 'very-high', NULL, NULL, NULL),
	(900, '2', 'performanceTesting899', '1', 'very-high', NULL, NULL, NULL),
	(901, '2', 'performanceTesting900', '1', 'very-high', NULL, NULL, NULL),
	(902, '2', 'performanceTesting901', '1', 'very-high', NULL, NULL, NULL),
	(903, '2', 'performanceTesting902', '1', 'very-high', NULL, NULL, NULL),
	(904, '2', 'performanceTesting903', '1', 'very-high', NULL, NULL, NULL),
	(905, '2', 'performanceTesting904', '1', 'very-high', NULL, NULL, NULL),
	(906, '2', 'performanceTesting905', '1', 'very-high', NULL, NULL, NULL),
	(907, '2', 'performanceTesting906', '1', 'very-high', NULL, NULL, NULL),
	(908, '2', 'performanceTesting907', '1', 'very-high', NULL, NULL, NULL),
	(909, '2', 'performanceTesting908', '1', 'very-high', NULL, NULL, NULL),
	(910, '2', 'performanceTesting909', '1', 'very-high', NULL, NULL, NULL),
	(911, '2', 'performanceTesting910', '1', 'very-high', NULL, NULL, NULL),
	(912, '2', 'performanceTesting911', '1', 'very-high', NULL, NULL, NULL),
	(913, '2', 'performanceTesting912', '1', 'very-high', NULL, NULL, NULL),
	(914, '2', 'performanceTesting913', '1', 'very-high', NULL, NULL, NULL),
	(915, '2', 'performanceTesting914', '1', 'very-high', NULL, NULL, NULL),
	(916, '2', 'performanceTesting915', '1', 'very-high', NULL, NULL, NULL),
	(917, '2', 'performanceTesting916', '1', 'very-high', NULL, NULL, NULL),
	(918, '2', 'performanceTesting917', '1', 'very-high', NULL, NULL, NULL),
	(919, '2', 'performanceTesting918', '1', 'very-high', NULL, NULL, NULL),
	(920, '2', 'performanceTesting919', '1', 'very-high', NULL, NULL, NULL),
	(921, '2', 'performanceTesting920', '1', 'very-high', NULL, NULL, NULL),
	(922, '2', 'performanceTesting921', '1', 'very-high', NULL, NULL, NULL),
	(923, '2', 'performanceTesting922', '1', 'very-high', NULL, NULL, NULL),
	(924, '2', 'performanceTesting923', '1', 'very-high', NULL, NULL, NULL),
	(925, '2', 'performanceTesting924', '1', 'very-high', NULL, NULL, NULL),
	(926, '2', 'performanceTesting925', '1', 'very-high', NULL, NULL, NULL),
	(927, '2', 'performanceTesting926', '1', 'very-high', NULL, NULL, NULL),
	(928, '2', 'performanceTesting927', '1', 'very-high', NULL, NULL, NULL),
	(929, '2', 'performanceTesting928', '1', 'very-high', NULL, NULL, NULL),
	(930, '2', 'performanceTesting929', '1', 'very-high', NULL, NULL, NULL),
	(931, '2', 'performanceTesting930', '1', 'very-high', NULL, NULL, NULL),
	(932, '2', 'performanceTesting931', '1', 'very-high', NULL, NULL, NULL),
	(933, '2', 'performanceTesting932', '1', 'very-high', NULL, NULL, NULL),
	(934, '2', 'performanceTesting933', '1', 'very-high', NULL, NULL, NULL),
	(935, '2', 'performanceTesting934', '1', 'very-high', NULL, NULL, NULL),
	(936, '2', 'performanceTesting935', '1', 'very-high', NULL, NULL, NULL),
	(937, '2', 'performanceTesting936', '1', 'very-high', NULL, NULL, NULL),
	(938, '2', 'performanceTesting937', '1', 'very-high', NULL, NULL, NULL),
	(939, '2', 'performanceTesting938', '1', 'very-high', NULL, NULL, NULL),
	(940, '2', 'performanceTesting939', '1', 'very-high', NULL, NULL, NULL),
	(941, '2', 'performanceTesting940', '1', 'very-high', NULL, NULL, NULL),
	(942, '2', 'performanceTesting941', '1', 'very-high', NULL, NULL, NULL),
	(943, '2', 'performanceTesting942', '1', 'very-high', NULL, NULL, NULL),
	(944, '2', 'performanceTesting943', '1', 'very-high', NULL, NULL, NULL),
	(945, '2', 'performanceTesting944', '1', 'very-high', NULL, NULL, NULL),
	(946, '2', 'performanceTesting945', '1', 'very-high', NULL, NULL, NULL),
	(947, '2', 'performanceTesting946', '1', 'very-high', NULL, NULL, NULL),
	(948, '2', 'performanceTesting947', '1', 'very-high', NULL, NULL, NULL),
	(949, '2', 'performanceTesting948', '1', 'very-high', NULL, NULL, NULL),
	(950, '2', 'performanceTesting949', '1', 'very-high', NULL, NULL, NULL),
	(951, '2', 'performanceTesting950', '1', 'very-high', NULL, NULL, NULL),
	(952, '2', 'performanceTesting951', '1', 'very-high', NULL, NULL, NULL),
	(953, '2', 'performanceTesting952', '1', 'very-high', NULL, NULL, NULL),
	(954, '2', 'performanceTesting953', '1', 'very-high', NULL, NULL, NULL),
	(955, '2', 'performanceTesting954', '1', 'very-high', NULL, NULL, NULL),
	(956, '2', 'performanceTesting955', '1', 'very-high', NULL, NULL, NULL),
	(957, '2', 'performanceTesting956', '1', 'very-high', NULL, NULL, NULL),
	(958, '2', 'performanceTesting957', '1', 'very-high', NULL, NULL, NULL),
	(959, '2', 'performanceTesting958', '1', 'very-high', NULL, NULL, NULL),
	(960, '2', 'performanceTesting959', '1', 'very-high', NULL, NULL, NULL),
	(961, '2', 'performanceTesting960', '1', 'very-high', NULL, NULL, NULL),
	(962, '2', 'performanceTesting961', '1', 'very-high', NULL, NULL, NULL),
	(963, '2', 'performanceTesting962', '1', 'very-high', NULL, NULL, NULL),
	(964, '2', 'performanceTesting963', '1', 'very-high', NULL, NULL, NULL),
	(965, '2', 'performanceTesting964', '1', 'very-high', NULL, NULL, NULL),
	(966, '2', 'performanceTesting965', '1', 'very-high', NULL, NULL, NULL),
	(967, '2', 'performanceTesting966', '1', 'very-high', NULL, NULL, NULL),
	(968, '2', 'performanceTesting967', '1', 'very-high', NULL, NULL, NULL),
	(969, '2', 'performanceTesting968', '1', 'very-high', NULL, NULL, NULL),
	(970, '2', 'performanceTesting969', '1', 'very-high', NULL, NULL, NULL),
	(971, '2', 'performanceTesting970', '1', 'very-high', NULL, NULL, NULL),
	(972, '2', 'performanceTesting971', '1', 'very-high', NULL, NULL, NULL),
	(973, '2', 'performanceTesting972', '1', 'very-high', NULL, NULL, NULL),
	(974, '2', 'performanceTesting973', '1', 'very-high', NULL, NULL, NULL),
	(975, '2', 'performanceTesting974', '1', 'very-high', NULL, NULL, NULL),
	(976, '2', 'performanceTesting975', '1', 'very-high', NULL, NULL, NULL),
	(977, '2', 'performanceTesting976', '1', 'very-high', NULL, NULL, NULL),
	(978, '2', 'performanceTesting977', '1', 'very-high', NULL, NULL, NULL),
	(979, '2', 'performanceTesting978', '1', 'very-high', NULL, NULL, NULL),
	(980, '2', 'performanceTesting979', '1', 'very-high', NULL, NULL, NULL),
	(981, '2', 'performanceTesting980', '1', 'very-high', NULL, NULL, NULL),
	(982, '2', 'performanceTesting981', '1', 'very-high', NULL, NULL, NULL),
	(983, '2', 'performanceTesting982', '1', 'very-high', NULL, NULL, NULL),
	(984, '2', 'performanceTesting983', '1', 'very-high', NULL, NULL, NULL),
	(985, '2', 'performanceTesting984', '1', 'very-high', NULL, NULL, NULL),
	(986, '2', 'performanceTesting985', '1', 'very-high', NULL, NULL, NULL),
	(987, '2', 'performanceTesting986', '1', 'very-high', NULL, NULL, NULL),
	(988, '2', 'performanceTesting987', '1', 'very-high', NULL, NULL, NULL),
	(989, '2', 'performanceTesting988', '1', 'very-high', NULL, NULL, NULL),
	(990, '2', 'performanceTesting989', '1', 'very-high', NULL, NULL, NULL),
	(991, '2', 'performanceTesting990', '1', 'very-high', NULL, NULL, NULL),
	(992, '2', 'performanceTesting991', '1', 'very-high', NULL, NULL, NULL),
	(993, '2', 'performanceTesting992', '1', 'very-high', NULL, NULL, NULL),
	(994, '2', 'performanceTesting993', '1', 'very-high', NULL, NULL, NULL),
	(995, '2', 'performanceTesting994', '1', 'very-high', NULL, NULL, NULL),
	(996, '2', 'performanceTesting995', '1', 'very-high', NULL, NULL, NULL),
	(997, '2', 'performanceTesting996', '1', 'very-high', NULL, NULL, NULL),
	(998, '2', 'performanceTesting997', '1', 'very-high', NULL, NULL, NULL),
	(999, '2', 'performanceTesting998', '1', 'very-high', NULL, NULL, NULL),
	(1000, '2', 'performanceTesting999', '1', 'very-high', NULL, NULL, NULL),
	(1001, '2', 'performanceTesting1000', '1', 'very-high', NULL, NULL, NULL);
	`)

	// currentActivity = 1
	// currentTodo = 1
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3030", nil)
}

func ActivityRest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var activity Activity
		err := decoder.Decode(&activity)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}

		if activity.Title == "" {

			resp.Status = "Bad Request"
			resp.Message = "title cannot be null"
			resp.Data = kosong
			w.WriteHeader(http.StatusBadRequest)

			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		}

		activity.CreatedAt = "2021-12-01T09:23:05.825Z"
		activity.UpdatedAt = "2021-12-01T09:23:05.825Z"
		activity.DeletedAt = nil
		activity.ID = len(activities) + 1

		w.WriteHeader(http.StatusCreated)
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = activity
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}

		w.Write(jData)
		go func() {
			// currentActivity = currentActivity + 1
			activities = append(activities, activity)

			// db.Create(&activity)
		}()

	case "GET":

		resp.Status = "Success"
		resp.Message = "Success"
		// data := []Activity{}
		// for i := range activities {
		// 	if activities[i].DeletedAt == nil {
		// 		data = append(data, activities[i])
		// 	}
		// }
		resp.Data = activities
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.Write(jData)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}

func TodoRest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t Todo
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}

		if t.ActivityGroupId == "" || t.ActivityGroupId == nil {

			resp.Status = "Bad Request"
			resp.Message = "activity_group_id cannot be null"
			resp.Data = kosong
			w.WriteHeader(http.StatusBadRequest)

			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		}

		if t.Title == "" || t.Title == nil {

			resp.Status = "Bad Request"
			resp.Message = "title cannot be null"
			resp.Data = kosong
			w.WriteHeader(http.StatusBadRequest)

			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		}

		if t.Priority == "" || t.Priority == nil {
			t.Priority = "very-high"
		}

		t.IsActive = "1"
		t.CreatedAt = "2021-12-01T09:23:05.825Z"
		t.UpdatedAt = "2021-12-01T09:23:05.825Z"
		t.DeletedAt = nil
		t.ID = len(todos) + 1
		t.IsActive = true

		w.WriteHeader(http.StatusCreated)
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = t
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		todos = append(todos, t)

		w.Write(jData)

	case "GET":

		resp.Status = "Success"
		resp.Message = "Success"

		param1 := r.URL.Query().Get("activity_group_id")

		if param1 != "" {
			resp.Data = []Todo{}
		} else {
			resp.Data = todos
		}
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.Write(jData)

	default:
		http.Error(w, "", http.StatusBadRequest)
	}

}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func HandleParamActivity(w http.ResponseWriter, r *http.Request, ids string) {

	id, err := strconv.Atoi(ids)
	if err != nil {
		return
	}

	if id > len(activities) {
		resp.Status = "Not Found"
		resp.Message = "Activity with ID " + ids + " Not Found"
		resp.Data = kosong
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(jData)

	} else {

		if r.Method == "GET" {
			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = activities[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		} else if r.Method == "PATCH" {
			var activity Activity
			decoder := json.NewDecoder(r.Body)

			err := decoder.Decode(&activity)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}

			if activity.Title == "" {

				resp.Status = "Bad Request"
				resp.Message = "title cannot be null"
				resp.Data = kosong
				w.WriteHeader(http.StatusBadRequest)

				jData, err := json.Marshal(resp)
				if err != nil {
					fmt.Fprint(w, "Test Error")
					return
				}
				w.Write(jData)
				return
			}

			activities[id-1].Title = activity.Title
			if activity.Email != "" {
				activities[id-1].Email = activity.Email
			}

			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = activities[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		} else if r.Method == "DELETE" {
			resp.Status = "Success"
			resp.Message = "Success"
			activities[id-1].DeletedAt = "2021-12-01T09:23:05.825Z"
			resp.Data = kosong
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		}
	}
	return
}

func HandleParamTodo(w http.ResponseWriter, r *http.Request, ids string) {

	id, err := strconv.Atoi(ids)
	if err != nil {
		return
	}

	if id > len(todos) {
		resp.Status = "Not Found"
		resp.Message = "Todo with ID " + ids + " Not Found"
		resp.Data = kosong
		jData, err := json.Marshal(resp)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(jData)

	} else {

		if r.Method == "GET" {
			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = todos[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return
		} else if r.Method == "PATCH" {
			decoder := json.NewDecoder(r.Body)
			var t Todo
			err := decoder.Decode(&t)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}

			todos[id-1].IsActive = t.IsActive

			if t.Title != "" && t.Title != nil {
				todos[id-1].Title = t.Title
			}

			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = todos[id-1]
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		} else if r.Method == "DELETE" {
			resp.Status = "Success"
			resp.Message = "Success"
			activities[id-1].DeletedAt = "2021-12-01T09:23:05.825Z"
			resp.Data = kosong
			jData, err := json.Marshal(resp)
			if err != nil {
				fmt.Fprint(w, "Test Error")
				return
			}
			w.Write(jData)
			return

		}
	}
	return
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var path = r.URL.Path
	lenPath := len(path)

	if path == "/todo-items" {
		TodoRest(w, r)
		return
	}

	if path == "/activity-groups" {
		ActivityRest(w, r)
		return
	}

	if lenPath > 12 {
		if path[0:12] == "/todo-items/" && isInt(path[12:lenPath]) {
			ids := path[12:lenPath]
			HandleParamTodo(w, r, ids)
			return
		}
	}

	if lenPath > 17 {
		if path[0:17] == "/activity-groups/" && isInt(path[17:lenPath]) {
			ids := path[17:lenPath]
			HandleParamActivity(w, r, ids)
			return
		}
	}

	names := "Oke"
	jData, err := json.Marshal(names)
	if err != nil {
		fmt.Fprint(w, "Internal Server Error")
	} else {
		w.Write(jData)
	}

}
