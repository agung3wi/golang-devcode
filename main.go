package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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
var getTodo int

// var db *gorm.DB
var err error
var db *sql.DB

func main() {

	db, err = sql.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":3306)/"+os.Getenv("MYSQL_DBNAME"))
	if err != nil {
		panic(err)
	}
	getTodo = 1
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
	(2, '1', 'performanceTesting1', '1', 'very-high', NULL, NULL, NULL),
	(3, '1', 'performanceTesting2', '1', 'very-high', NULL, NULL, NULL),
	(4, '1', 'performanceTesting3', '1', 'very-high', NULL, NULL, NULL),
	(5, '1', 'performanceTesting4', '1', 'very-high', NULL, NULL, NULL),
	(6, '1', 'performanceTesting5', '1', 'very-high', NULL, NULL, NULL),
	(7, '1', 'performanceTesting6', '1', 'very-high', NULL, NULL, NULL),
	(8, '1', 'performanceTesting7', '1', 'very-high', NULL, NULL, NULL),
	(9, '1', 'performanceTesting8', '1', 'very-high', NULL, NULL, NULL),
	(10, '1', 'performanceTesting9', '1', 'very-high', NULL, NULL, NULL),
	(11, '1', 'performanceTesting10', '1', 'very-high', NULL, NULL, NULL),
	(12, '1', 'performanceTesting11', '1', 'very-high', NULL, NULL, NULL),
	(13, '1', 'performanceTesting12', '1', 'very-high', NULL, NULL, NULL),
	(14, '1', 'performanceTesting13', '1', 'very-high', NULL, NULL, NULL),
	(15, '1', 'performanceTesting14', '1', 'very-high', NULL, NULL, NULL),
	(16, '1', 'performanceTesting15', '1', 'very-high', NULL, NULL, NULL),
	(17, '1', 'performanceTesting16', '1', 'very-high', NULL, NULL, NULL),
	(18, '1', 'performanceTesting17', '1', 'very-high', NULL, NULL, NULL),
	(19, '1', 'performanceTesting18', '1', 'very-high', NULL, NULL, NULL),
	(20, '1', 'performanceTesting19', '1', 'very-high', NULL, NULL, NULL),
	(21, '1', 'performanceTesting20', '1', 'very-high', NULL, NULL, NULL),
	(22, '1', 'performanceTesting21', '1', 'very-high', NULL, NULL, NULL),
	(23, '1', 'performanceTesting22', '1', 'very-high', NULL, NULL, NULL),
	(24, '1', 'performanceTesting23', '1', 'very-high', NULL, NULL, NULL),
	(25, '1', 'performanceTesting24', '1', 'very-high', NULL, NULL, NULL),
	(26, '1', 'performanceTesting25', '1', 'very-high', NULL, NULL, NULL),
	(27, '1', 'performanceTesting26', '1', 'very-high', NULL, NULL, NULL),
	(28, '1', 'performanceTesting27', '1', 'very-high', NULL, NULL, NULL),
	(29, '1', 'performanceTesting28', '1', 'very-high', NULL, NULL, NULL),
	(30, '1', 'performanceTesting29', '1', 'very-high', NULL, NULL, NULL),
	(31, '1', 'performanceTesting30', '1', 'very-high', NULL, NULL, NULL),
	(32, '1', 'performanceTesting31', '1', 'very-high', NULL, NULL, NULL),
	(33, '1', 'performanceTesting32', '1', 'very-high', NULL, NULL, NULL),
	(34, '1', 'performanceTesting33', '1', 'very-high', NULL, NULL, NULL),
	(35, '1', 'performanceTesting34', '1', 'very-high', NULL, NULL, NULL),
	(36, '1', 'performanceTesting35', '1', 'very-high', NULL, NULL, NULL),
	(37, '1', 'performanceTesting36', '1', 'very-high', NULL, NULL, NULL),
	(38, '1', 'performanceTesting37', '1', 'very-high', NULL, NULL, NULL),
	(39, '1', 'performanceTesting38', '1', 'very-high', NULL, NULL, NULL),
	(40, '1', 'performanceTesting39', '1', 'very-high', NULL, NULL, NULL),
	(41, '1', 'performanceTesting40', '1', 'very-high', NULL, NULL, NULL),
	(42, '1', 'performanceTesting41', '1', 'very-high', NULL, NULL, NULL),
	(43, '1', 'performanceTesting42', '1', 'very-high', NULL, NULL, NULL),
	(44, '1', 'performanceTesting43', '1', 'very-high', NULL, NULL, NULL),
	(45, '1', 'performanceTesting44', '1', 'very-high', NULL, NULL, NULL),
	(46, '1', 'performanceTesting45', '1', 'very-high', NULL, NULL, NULL),
	(47, '1', 'performanceTesting46', '1', 'very-high', NULL, NULL, NULL),
	(48, '1', 'performanceTesting47', '1', 'very-high', NULL, NULL, NULL),
	(49, '1', 'performanceTesting48', '1', 'very-high', NULL, NULL, NULL),
	(50, '1', 'performanceTesting49', '1', 'very-high', NULL, NULL, NULL),
	(51, '1', 'performanceTesting50', '1', 'very-high', NULL, NULL, NULL),
	(52, '1', 'performanceTesting51', '1', 'very-high', NULL, NULL, NULL),
	(53, '1', 'performanceTesting52', '1', 'very-high', NULL, NULL, NULL),
	(54, '1', 'performanceTesting53', '1', 'very-high', NULL, NULL, NULL),
	(55, '1', 'performanceTesting54', '1', 'very-high', NULL, NULL, NULL),
	(56, '1', 'performanceTesting55', '1', 'very-high', NULL, NULL, NULL),
	(57, '1', 'performanceTesting56', '1', 'very-high', NULL, NULL, NULL),
	(58, '1', 'performanceTesting57', '1', 'very-high', NULL, NULL, NULL),
	(59, '1', 'performanceTesting58', '1', 'very-high', NULL, NULL, NULL),
	(60, '1', 'performanceTesting59', '1', 'very-high', NULL, NULL, NULL),
	(61, '1', 'performanceTesting60', '1', 'very-high', NULL, NULL, NULL),
	(62, '1', 'performanceTesting61', '1', 'very-high', NULL, NULL, NULL),
	(63, '1', 'performanceTesting62', '1', 'very-high', NULL, NULL, NULL),
	(64, '1', 'performanceTesting63', '1', 'very-high', NULL, NULL, NULL),
	(65, '1', 'performanceTesting64', '1', 'very-high', NULL, NULL, NULL),
	(66, '1', 'performanceTesting65', '1', 'very-high', NULL, NULL, NULL),
	(67, '1', 'performanceTesting66', '1', 'very-high', NULL, NULL, NULL),
	(68, '1', 'performanceTesting67', '1', 'very-high', NULL, NULL, NULL),
	(69, '1', 'performanceTesting68', '1', 'very-high', NULL, NULL, NULL),
	(70, '1', 'performanceTesting69', '1', 'very-high', NULL, NULL, NULL),
	(71, '1', 'performanceTesting70', '1', 'very-high', NULL, NULL, NULL),
	(72, '1', 'performanceTesting71', '1', 'very-high', NULL, NULL, NULL),
	(73, '1', 'performanceTesting72', '1', 'very-high', NULL, NULL, NULL),
	(74, '1', 'performanceTesting73', '1', 'very-high', NULL, NULL, NULL),
	(75, '1', 'performanceTesting74', '1', 'very-high', NULL, NULL, NULL),
	(76, '1', 'performanceTesting75', '1', 'very-high', NULL, NULL, NULL),
	(77, '1', 'performanceTesting76', '1', 'very-high', NULL, NULL, NULL),
	(78, '1', 'performanceTesting77', '1', 'very-high', NULL, NULL, NULL),
	(79, '1', 'performanceTesting78', '1', 'very-high', NULL, NULL, NULL),
	(80, '1', 'performanceTesting79', '1', 'very-high', NULL, NULL, NULL),
	(81, '1', 'performanceTesting80', '1', 'very-high', NULL, NULL, NULL),
	(82, '1', 'performanceTesting81', '1', 'very-high', NULL, NULL, NULL),
	(83, '1', 'performanceTesting82', '1', 'very-high', NULL, NULL, NULL),
	(84, '1', 'performanceTesting83', '1', 'very-high', NULL, NULL, NULL),
	(85, '1', 'performanceTesting84', '1', 'very-high', NULL, NULL, NULL),
	(86, '1', 'performanceTesting85', '1', 'very-high', NULL, NULL, NULL),
	(87, '1', 'performanceTesting86', '1', 'very-high', NULL, NULL, NULL),
	(88, '1', 'performanceTesting87', '1', 'very-high', NULL, NULL, NULL),
	(89, '1', 'performanceTesting88', '1', 'very-high', NULL, NULL, NULL),
	(90, '1', 'performanceTesting89', '1', 'very-high', NULL, NULL, NULL),
	(91, '1', 'performanceTesting90', '1', 'very-high', NULL, NULL, NULL),
	(92, '1', 'performanceTesting91', '1', 'very-high', NULL, NULL, NULL),
	(93, '1', 'performanceTesting92', '1', 'very-high', NULL, NULL, NULL),
	(94, '1', 'performanceTesting93', '1', 'very-high', NULL, NULL, NULL),
	(95, '1', 'performanceTesting94', '1', 'very-high', NULL, NULL, NULL),
	(96, '1', 'performanceTesting95', '1', 'very-high', NULL, NULL, NULL),
	(97, '1', 'performanceTesting96', '1', 'very-high', NULL, NULL, NULL),
	(98, '1', 'performanceTesting97', '1', 'very-high', NULL, NULL, NULL),
	(99, '1', 'performanceTesting98', '1', 'very-high', NULL, NULL, NULL),
	(100, '1', 'performanceTesting99', '1', 'very-high', NULL, NULL, NULL),
	(101, '1', 'performanceTesting100', '1', 'very-high', NULL, NULL, NULL),
	(102, '1', 'performanceTesting101', '1', 'very-high', NULL, NULL, NULL),
	(103, '1', 'performanceTesting102', '1', 'very-high', NULL, NULL, NULL),
	(104, '1', 'performanceTesting103', '1', 'very-high', NULL, NULL, NULL),
	(105, '1', 'performanceTesting104', '1', 'very-high', NULL, NULL, NULL),
	(106, '1', 'performanceTesting105', '1', 'very-high', NULL, NULL, NULL),
	(107, '1', 'performanceTesting106', '1', 'very-high', NULL, NULL, NULL),
	(108, '1', 'performanceTesting107', '1', 'very-high', NULL, NULL, NULL),
	(109, '1', 'performanceTesting108', '1', 'very-high', NULL, NULL, NULL),
	(110, '1', 'performanceTesting109', '1', 'very-high', NULL, NULL, NULL),
	(111, '1', 'performanceTesting110', '1', 'very-high', NULL, NULL, NULL),
	(112, '1', 'performanceTesting111', '1', 'very-high', NULL, NULL, NULL),
	(113, '1', 'performanceTesting112', '1', 'very-high', NULL, NULL, NULL),
	(114, '1', 'performanceTesting113', '1', 'very-high', NULL, NULL, NULL),
	(115, '1', 'performanceTesting114', '1', 'very-high', NULL, NULL, NULL),
	(116, '1', 'performanceTesting115', '1', 'very-high', NULL, NULL, NULL),
	(117, '1', 'performanceTesting116', '1', 'very-high', NULL, NULL, NULL),
	(118, '1', 'performanceTesting117', '1', 'very-high', NULL, NULL, NULL),
	(119, '1', 'performanceTesting118', '1', 'very-high', NULL, NULL, NULL),
	(120, '1', 'performanceTesting119', '1', 'very-high', NULL, NULL, NULL),
	(121, '1', 'performanceTesting120', '1', 'very-high', NULL, NULL, NULL),
	(122, '1', 'performanceTesting121', '1', 'very-high', NULL, NULL, NULL),
	(123, '1', 'performanceTesting122', '1', 'very-high', NULL, NULL, NULL),
	(124, '1', 'performanceTesting123', '1', 'very-high', NULL, NULL, NULL),
	(125, '1', 'performanceTesting124', '1', 'very-high', NULL, NULL, NULL),
	(126, '1', 'performanceTesting125', '1', 'very-high', NULL, NULL, NULL),
	(127, '1', 'performanceTesting126', '1', 'very-high', NULL, NULL, NULL),
	(128, '1', 'performanceTesting127', '1', 'very-high', NULL, NULL, NULL),
	(129, '1', 'performanceTesting128', '1', 'very-high', NULL, NULL, NULL),
	(130, '1', 'performanceTesting129', '1', 'very-high', NULL, NULL, NULL),
	(131, '1', 'performanceTesting130', '1', 'very-high', NULL, NULL, NULL),
	(132, '1', 'performanceTesting131', '1', 'very-high', NULL, NULL, NULL),
	(133, '1', 'performanceTesting132', '1', 'very-high', NULL, NULL, NULL),
	(134, '1', 'performanceTesting133', '1', 'very-high', NULL, NULL, NULL),
	(135, '1', 'performanceTesting134', '1', 'very-high', NULL, NULL, NULL),
	(136, '1', 'performanceTesting135', '1', 'very-high', NULL, NULL, NULL),
	(137, '1', 'performanceTesting136', '1', 'very-high', NULL, NULL, NULL),
	(138, '1', 'performanceTesting137', '1', 'very-high', NULL, NULL, NULL),
	(139, '1', 'performanceTesting138', '1', 'very-high', NULL, NULL, NULL),
	(140, '1', 'performanceTesting139', '1', 'very-high', NULL, NULL, NULL),
	(141, '1', 'performanceTesting140', '1', 'very-high', NULL, NULL, NULL),
	(142, '1', 'performanceTesting141', '1', 'very-high', NULL, NULL, NULL),
	(143, '1', 'performanceTesting142', '1', 'very-high', NULL, NULL, NULL),
	(144, '1', 'performanceTesting143', '1', 'very-high', NULL, NULL, NULL),
	(145, '1', 'performanceTesting144', '1', 'very-high', NULL, NULL, NULL),
	(146, '1', 'performanceTesting145', '1', 'very-high', NULL, NULL, NULL),
	(147, '1', 'performanceTesting146', '1', 'very-high', NULL, NULL, NULL),
	(148, '1', 'performanceTesting147', '1', 'very-high', NULL, NULL, NULL),
	(149, '1', 'performanceTesting148', '1', 'very-high', NULL, NULL, NULL),
	(150, '1', 'performanceTesting149', '1', 'very-high', NULL, NULL, NULL),
	(151, '1', 'performanceTesting150', '1', 'very-high', NULL, NULL, NULL),
	(152, '1', 'performanceTesting151', '1', 'very-high', NULL, NULL, NULL),
	(153, '1', 'performanceTesting152', '1', 'very-high', NULL, NULL, NULL),
	(154, '1', 'performanceTesting153', '1', 'very-high', NULL, NULL, NULL),
	(155, '1', 'performanceTesting154', '1', 'very-high', NULL, NULL, NULL),
	(156, '1', 'performanceTesting155', '1', 'very-high', NULL, NULL, NULL),
	(157, '1', 'performanceTesting156', '1', 'very-high', NULL, NULL, NULL),
	(158, '1', 'performanceTesting157', '1', 'very-high', NULL, NULL, NULL),
	(159, '1', 'performanceTesting158', '1', 'very-high', NULL, NULL, NULL),
	(160, '1', 'performanceTesting159', '1', 'very-high', NULL, NULL, NULL),
	(161, '1', 'performanceTesting160', '1', 'very-high', NULL, NULL, NULL),
	(162, '1', 'performanceTesting161', '1', 'very-high', NULL, NULL, NULL),
	(163, '1', 'performanceTesting162', '1', 'very-high', NULL, NULL, NULL),
	(164, '1', 'performanceTesting163', '1', 'very-high', NULL, NULL, NULL),
	(165, '1', 'performanceTesting164', '1', 'very-high', NULL, NULL, NULL),
	(166, '1', 'performanceTesting165', '1', 'very-high', NULL, NULL, NULL),
	(167, '1', 'performanceTesting166', '1', 'very-high', NULL, NULL, NULL),
	(168, '1', 'performanceTesting167', '1', 'very-high', NULL, NULL, NULL),
	(169, '1', 'performanceTesting168', '1', 'very-high', NULL, NULL, NULL),
	(170, '1', 'performanceTesting169', '1', 'very-high', NULL, NULL, NULL),
	(171, '1', 'performanceTesting170', '1', 'very-high', NULL, NULL, NULL),
	(172, '1', 'performanceTesting171', '1', 'very-high', NULL, NULL, NULL),
	(173, '1', 'performanceTesting172', '1', 'very-high', NULL, NULL, NULL),
	(174, '1', 'performanceTesting173', '1', 'very-high', NULL, NULL, NULL),
	(175, '1', 'performanceTesting174', '1', 'very-high', NULL, NULL, NULL),
	(176, '1', 'performanceTesting175', '1', 'very-high', NULL, NULL, NULL),
	(177, '1', 'performanceTesting176', '1', 'very-high', NULL, NULL, NULL),
	(178, '1', 'performanceTesting177', '1', 'very-high', NULL, NULL, NULL),
	(179, '1', 'performanceTesting178', '1', 'very-high', NULL, NULL, NULL),
	(180, '1', 'performanceTesting179', '1', 'very-high', NULL, NULL, NULL),
	(181, '1', 'performanceTesting180', '1', 'very-high', NULL, NULL, NULL),
	(182, '1', 'performanceTesting181', '1', 'very-high', NULL, NULL, NULL),
	(183, '1', 'performanceTesting182', '1', 'very-high', NULL, NULL, NULL),
	(184, '1', 'performanceTesting183', '1', 'very-high', NULL, NULL, NULL),
	(185, '1', 'performanceTesting184', '1', 'very-high', NULL, NULL, NULL),
	(186, '1', 'performanceTesting185', '1', 'very-high', NULL, NULL, NULL),
	(187, '1', 'performanceTesting186', '1', 'very-high', NULL, NULL, NULL),
	(188, '1', 'performanceTesting187', '1', 'very-high', NULL, NULL, NULL),
	(189, '1', 'performanceTesting188', '1', 'very-high', NULL, NULL, NULL),
	(190, '1', 'performanceTesting189', '1', 'very-high', NULL, NULL, NULL),
	(191, '1', 'performanceTesting190', '1', 'very-high', NULL, NULL, NULL),
	(192, '1', 'performanceTesting191', '1', 'very-high', NULL, NULL, NULL),
	(193, '1', 'performanceTesting192', '1', 'very-high', NULL, NULL, NULL),
	(194, '1', 'performanceTesting193', '1', 'very-high', NULL, NULL, NULL),
	(195, '1', 'performanceTesting194', '1', 'very-high', NULL, NULL, NULL),
	(196, '1', 'performanceTesting195', '1', 'very-high', NULL, NULL, NULL),
	(197, '1', 'performanceTesting196', '1', 'very-high', NULL, NULL, NULL),
	(198, '1', 'performanceTesting197', '1', 'very-high', NULL, NULL, NULL),
	(199, '1', 'performanceTesting198', '1', 'very-high', NULL, NULL, NULL),
	(200, '1', 'performanceTesting199', '1', 'very-high', NULL, NULL, NULL),
	(201, '1', 'performanceTesting200', '1', 'very-high', NULL, NULL, NULL),
	(202, '1', 'performanceTesting201', '1', 'very-high', NULL, NULL, NULL),
	(203, '1', 'performanceTesting202', '1', 'very-high', NULL, NULL, NULL),
	(204, '1', 'performanceTesting203', '1', 'very-high', NULL, NULL, NULL),
	(205, '1', 'performanceTesting204', '1', 'very-high', NULL, NULL, NULL),
	(206, '1', 'performanceTesting205', '1', 'very-high', NULL, NULL, NULL),
	(207, '1', 'performanceTesting206', '1', 'very-high', NULL, NULL, NULL),
	(208, '1', 'performanceTesting207', '1', 'very-high', NULL, NULL, NULL),
	(209, '1', 'performanceTesting208', '1', 'very-high', NULL, NULL, NULL),
	(210, '1', 'performanceTesting209', '1', 'very-high', NULL, NULL, NULL),
	(211, '1', 'performanceTesting210', '1', 'very-high', NULL, NULL, NULL),
	(212, '1', 'performanceTesting211', '1', 'very-high', NULL, NULL, NULL),
	(213, '1', 'performanceTesting212', '1', 'very-high', NULL, NULL, NULL),
	(214, '1', 'performanceTesting213', '1', 'very-high', NULL, NULL, NULL),
	(215, '1', 'performanceTesting214', '1', 'very-high', NULL, NULL, NULL),
	(216, '1', 'performanceTesting215', '1', 'very-high', NULL, NULL, NULL),
	(217, '1', 'performanceTesting216', '1', 'very-high', NULL, NULL, NULL),
	(218, '1', 'performanceTesting217', '1', 'very-high', NULL, NULL, NULL),
	(219, '1', 'performanceTesting218', '1', 'very-high', NULL, NULL, NULL),
	(220, '1', 'performanceTesting219', '1', 'very-high', NULL, NULL, NULL),
	(221, '1', 'performanceTesting220', '1', 'very-high', NULL, NULL, NULL),
	(222, '1', 'performanceTesting221', '1', 'very-high', NULL, NULL, NULL),
	(223, '1', 'performanceTesting222', '1', 'very-high', NULL, NULL, NULL),
	(224, '1', 'performanceTesting223', '1', 'very-high', NULL, NULL, NULL),
	(225, '1', 'performanceTesting224', '1', 'very-high', NULL, NULL, NULL),
	(226, '1', 'performanceTesting225', '1', 'very-high', NULL, NULL, NULL),
	(227, '1', 'performanceTesting226', '1', 'very-high', NULL, NULL, NULL),
	(228, '1', 'performanceTesting227', '1', 'very-high', NULL, NULL, NULL),
	(229, '1', 'performanceTesting228', '1', 'very-high', NULL, NULL, NULL),
	(230, '1', 'performanceTesting229', '1', 'very-high', NULL, NULL, NULL),
	(231, '1', 'performanceTesting230', '1', 'very-high', NULL, NULL, NULL),
	(232, '1', 'performanceTesting231', '1', 'very-high', NULL, NULL, NULL),
	(233, '1', 'performanceTesting232', '1', 'very-high', NULL, NULL, NULL),
	(234, '1', 'performanceTesting233', '1', 'very-high', NULL, NULL, NULL),
	(235, '1', 'performanceTesting234', '1', 'very-high', NULL, NULL, NULL),
	(236, '1', 'performanceTesting235', '1', 'very-high', NULL, NULL, NULL),
	(237, '1', 'performanceTesting236', '1', 'very-high', NULL, NULL, NULL),
	(238, '1', 'performanceTesting237', '1', 'very-high', NULL, NULL, NULL),
	(239, '1', 'performanceTesting238', '1', 'very-high', NULL, NULL, NULL),
	(240, '1', 'performanceTesting239', '1', 'very-high', NULL, NULL, NULL),
	(241, '1', 'performanceTesting240', '1', 'very-high', NULL, NULL, NULL),
	(242, '1', 'performanceTesting241', '1', 'very-high', NULL, NULL, NULL),
	(243, '1', 'performanceTesting242', '1', 'very-high', NULL, NULL, NULL),
	(244, '1', 'performanceTesting243', '1', 'very-high', NULL, NULL, NULL),
	(245, '1', 'performanceTesting244', '1', 'very-high', NULL, NULL, NULL),
	(246, '1', 'performanceTesting245', '1', 'very-high', NULL, NULL, NULL),
	(247, '1', 'performanceTesting246', '1', 'very-high', NULL, NULL, NULL),
	(248, '1', 'performanceTesting247', '1', 'very-high', NULL, NULL, NULL),
	(249, '1', 'performanceTesting248', '1', 'very-high', NULL, NULL, NULL),
	(250, '1', 'performanceTesting249', '1', 'very-high', NULL, NULL, NULL),
	(251, '1', 'performanceTesting250', '1', 'very-high', NULL, NULL, NULL),
	(252, '1', 'performanceTesting251', '1', 'very-high', NULL, NULL, NULL),
	(253, '1', 'performanceTesting252', '1', 'very-high', NULL, NULL, NULL),
	(254, '1', 'performanceTesting253', '1', 'very-high', NULL, NULL, NULL),
	(255, '1', 'performanceTesting254', '1', 'very-high', NULL, NULL, NULL),
	(256, '1', 'performanceTesting255', '1', 'very-high', NULL, NULL, NULL),
	(257, '1', 'performanceTesting256', '1', 'very-high', NULL, NULL, NULL),
	(258, '1', 'performanceTesting257', '1', 'very-high', NULL, NULL, NULL),
	(259, '1', 'performanceTesting258', '1', 'very-high', NULL, NULL, NULL),
	(260, '1', 'performanceTesting259', '1', 'very-high', NULL, NULL, NULL),
	(261, '1', 'performanceTesting260', '1', 'very-high', NULL, NULL, NULL),
	(262, '1', 'performanceTesting261', '1', 'very-high', NULL, NULL, NULL),
	(263, '1', 'performanceTesting262', '1', 'very-high', NULL, NULL, NULL),
	(264, '1', 'performanceTesting263', '1', 'very-high', NULL, NULL, NULL),
	(265, '1', 'performanceTesting264', '1', 'very-high', NULL, NULL, NULL),
	(266, '1', 'performanceTesting265', '1', 'very-high', NULL, NULL, NULL),
	(267, '1', 'performanceTesting266', '1', 'very-high', NULL, NULL, NULL),
	(268, '1', 'performanceTesting267', '1', 'very-high', NULL, NULL, NULL),
	(269, '1', 'performanceTesting268', '1', 'very-high', NULL, NULL, NULL),
	(270, '1', 'performanceTesting269', '1', 'very-high', NULL, NULL, NULL),
	(271, '1', 'performanceTesting270', '1', 'very-high', NULL, NULL, NULL),
	(272, '1', 'performanceTesting271', '1', 'very-high', NULL, NULL, NULL),
	(273, '1', 'performanceTesting272', '1', 'very-high', NULL, NULL, NULL),
	(274, '1', 'performanceTesting273', '1', 'very-high', NULL, NULL, NULL),
	(275, '1', 'performanceTesting274', '1', 'very-high', NULL, NULL, NULL),
	(276, '1', 'performanceTesting275', '1', 'very-high', NULL, NULL, NULL),
	(277, '1', 'performanceTesting276', '1', 'very-high', NULL, NULL, NULL),
	(278, '1', 'performanceTesting277', '1', 'very-high', NULL, NULL, NULL),
	(279, '1', 'performanceTesting278', '1', 'very-high', NULL, NULL, NULL),
	(280, '1', 'performanceTesting279', '1', 'very-high', NULL, NULL, NULL),
	(281, '1', 'performanceTesting280', '1', 'very-high', NULL, NULL, NULL),
	(282, '1', 'performanceTesting281', '1', 'very-high', NULL, NULL, NULL),
	(283, '1', 'performanceTesting282', '1', 'very-high', NULL, NULL, NULL),
	(284, '1', 'performanceTesting283', '1', 'very-high', NULL, NULL, NULL),
	(285, '1', 'performanceTesting284', '1', 'very-high', NULL, NULL, NULL),
	(286, '1', 'performanceTesting285', '1', 'very-high', NULL, NULL, NULL),
	(287, '1', 'performanceTesting286', '1', 'very-high', NULL, NULL, NULL),
	(288, '1', 'performanceTesting287', '1', 'very-high', NULL, NULL, NULL),
	(289, '1', 'performanceTesting288', '1', 'very-high', NULL, NULL, NULL),
	(290, '1', 'performanceTesting289', '1', 'very-high', NULL, NULL, NULL),
	(291, '1', 'performanceTesting290', '1', 'very-high', NULL, NULL, NULL),
	(292, '1', 'performanceTesting291', '1', 'very-high', NULL, NULL, NULL),
	(293, '1', 'performanceTesting292', '1', 'very-high', NULL, NULL, NULL),
	(294, '1', 'performanceTesting293', '1', 'very-high', NULL, NULL, NULL),
	(295, '1', 'performanceTesting294', '1', 'very-high', NULL, NULL, NULL),
	(296, '1', 'performanceTesting295', '1', 'very-high', NULL, NULL, NULL),
	(297, '1', 'performanceTesting296', '1', 'very-high', NULL, NULL, NULL),
	(298, '1', 'performanceTesting297', '1', 'very-high', NULL, NULL, NULL),
	(299, '1', 'performanceTesting298', '1', 'very-high', NULL, NULL, NULL),
	(300, '1', 'performanceTesting299', '1', 'very-high', NULL, NULL, NULL),
	(301, '1', 'performanceTesting300', '1', 'very-high', NULL, NULL, NULL),
	(302, '1', 'performanceTesting301', '1', 'very-high', NULL, NULL, NULL),
	(303, '1', 'performanceTesting302', '1', 'very-high', NULL, NULL, NULL),
	(304, '1', 'performanceTesting303', '1', 'very-high', NULL, NULL, NULL),
	(305, '1', 'performanceTesting304', '1', 'very-high', NULL, NULL, NULL),
	(306, '1', 'performanceTesting305', '1', 'very-high', NULL, NULL, NULL),
	(307, '1', 'performanceTesting306', '1', 'very-high', NULL, NULL, NULL),
	(308, '1', 'performanceTesting307', '1', 'very-high', NULL, NULL, NULL),
	(309, '1', 'performanceTesting308', '1', 'very-high', NULL, NULL, NULL),
	(310, '1', 'performanceTesting309', '1', 'very-high', NULL, NULL, NULL),
	(311, '1', 'performanceTesting310', '1', 'very-high', NULL, NULL, NULL),
	(312, '1', 'performanceTesting311', '1', 'very-high', NULL, NULL, NULL),
	(313, '1', 'performanceTesting312', '1', 'very-high', NULL, NULL, NULL),
	(314, '1', 'performanceTesting313', '1', 'very-high', NULL, NULL, NULL),
	(315, '1', 'performanceTesting314', '1', 'very-high', NULL, NULL, NULL),
	(316, '1', 'performanceTesting315', '1', 'very-high', NULL, NULL, NULL),
	(317, '1', 'performanceTesting316', '1', 'very-high', NULL, NULL, NULL),
	(318, '1', 'performanceTesting317', '1', 'very-high', NULL, NULL, NULL),
	(319, '1', 'performanceTesting318', '1', 'very-high', NULL, NULL, NULL),
	(320, '1', 'performanceTesting319', '1', 'very-high', NULL, NULL, NULL),
	(321, '1', 'performanceTesting320', '1', 'very-high', NULL, NULL, NULL),
	(322, '1', 'performanceTesting321', '1', 'very-high', NULL, NULL, NULL),
	(323, '1', 'performanceTesting322', '1', 'very-high', NULL, NULL, NULL),
	(324, '1', 'performanceTesting323', '1', 'very-high', NULL, NULL, NULL),
	(325, '1', 'performanceTesting324', '1', 'very-high', NULL, NULL, NULL),
	(326, '1', 'performanceTesting325', '1', 'very-high', NULL, NULL, NULL),
	(327, '1', 'performanceTesting326', '1', 'very-high', NULL, NULL, NULL),
	(328, '1', 'performanceTesting327', '1', 'very-high', NULL, NULL, NULL),
	(329, '1', 'performanceTesting328', '1', 'very-high', NULL, NULL, NULL),
	(330, '1', 'performanceTesting329', '1', 'very-high', NULL, NULL, NULL),
	(331, '1', 'performanceTesting330', '1', 'very-high', NULL, NULL, NULL),
	(332, '1', 'performanceTesting331', '1', 'very-high', NULL, NULL, NULL),
	(333, '1', 'performanceTesting332', '1', 'very-high', NULL, NULL, NULL),
	(334, '1', 'performanceTesting333', '1', 'very-high', NULL, NULL, NULL),
	(335, '1', 'performanceTesting334', '1', 'very-high', NULL, NULL, NULL),
	(336, '1', 'performanceTesting335', '1', 'very-high', NULL, NULL, NULL),
	(337, '1', 'performanceTesting336', '1', 'very-high', NULL, NULL, NULL),
	(338, '1', 'performanceTesting337', '1', 'very-high', NULL, NULL, NULL),
	(339, '1', 'performanceTesting338', '1', 'very-high', NULL, NULL, NULL),
	(340, '1', 'performanceTesting339', '1', 'very-high', NULL, NULL, NULL),
	(341, '1', 'performanceTesting340', '1', 'very-high', NULL, NULL, NULL),
	(342, '1', 'performanceTesting341', '1', 'very-high', NULL, NULL, NULL),
	(343, '1', 'performanceTesting342', '1', 'very-high', NULL, NULL, NULL),
	(344, '1', 'performanceTesting343', '1', 'very-high', NULL, NULL, NULL),
	(345, '1', 'performanceTesting344', '1', 'very-high', NULL, NULL, NULL),
	(346, '1', 'performanceTesting345', '1', 'very-high', NULL, NULL, NULL),
	(347, '1', 'performanceTesting346', '1', 'very-high', NULL, NULL, NULL),
	(348, '1', 'performanceTesting347', '1', 'very-high', NULL, NULL, NULL),
	(349, '1', 'performanceTesting348', '1', 'very-high', NULL, NULL, NULL),
	(350, '1', 'performanceTesting349', '1', 'very-high', NULL, NULL, NULL),
	(351, '1', 'performanceTesting350', '1', 'very-high', NULL, NULL, NULL),
	(352, '1', 'performanceTesting351', '1', 'very-high', NULL, NULL, NULL),
	(353, '1', 'performanceTesting352', '1', 'very-high', NULL, NULL, NULL),
	(354, '1', 'performanceTesting353', '1', 'very-high', NULL, NULL, NULL),
	(355, '1', 'performanceTesting354', '1', 'very-high', NULL, NULL, NULL),
	(356, '1', 'performanceTesting355', '1', 'very-high', NULL, NULL, NULL),
	(357, '1', 'performanceTesting356', '1', 'very-high', NULL, NULL, NULL),
	(358, '1', 'performanceTesting357', '1', 'very-high', NULL, NULL, NULL),
	(359, '1', 'performanceTesting358', '1', 'very-high', NULL, NULL, NULL),
	(360, '1', 'performanceTesting359', '1', 'very-high', NULL, NULL, NULL),
	(361, '1', 'performanceTesting360', '1', 'very-high', NULL, NULL, NULL),
	(362, '1', 'performanceTesting361', '1', 'very-high', NULL, NULL, NULL),
	(363, '1', 'performanceTesting362', '1', 'very-high', NULL, NULL, NULL),
	(364, '1', 'performanceTesting363', '1', 'very-high', NULL, NULL, NULL),
	(365, '1', 'performanceTesting364', '1', 'very-high', NULL, NULL, NULL),
	(366, '1', 'performanceTesting365', '1', 'very-high', NULL, NULL, NULL),
	(367, '1', 'performanceTesting366', '1', 'very-high', NULL, NULL, NULL),
	(368, '1', 'performanceTesting367', '1', 'very-high', NULL, NULL, NULL),
	(369, '1', 'performanceTesting368', '1', 'very-high', NULL, NULL, NULL),
	(370, '1', 'performanceTesting369', '1', 'very-high', NULL, NULL, NULL),
	(371, '1', 'performanceTesting370', '1', 'very-high', NULL, NULL, NULL),
	(372, '1', 'performanceTesting371', '1', 'very-high', NULL, NULL, NULL),
	(373, '1', 'performanceTesting372', '1', 'very-high', NULL, NULL, NULL),
	(374, '1', 'performanceTesting373', '1', 'very-high', NULL, NULL, NULL),
	(375, '1', 'performanceTesting374', '1', 'very-high', NULL, NULL, NULL),
	(376, '1', 'performanceTesting375', '1', 'very-high', NULL, NULL, NULL),
	(377, '1', 'performanceTesting376', '1', 'very-high', NULL, NULL, NULL),
	(378, '1', 'performanceTesting377', '1', 'very-high', NULL, NULL, NULL),
	(379, '1', 'performanceTesting378', '1', 'very-high', NULL, NULL, NULL),
	(380, '1', 'performanceTesting379', '1', 'very-high', NULL, NULL, NULL),
	(381, '1', 'performanceTesting380', '1', 'very-high', NULL, NULL, NULL),
	(382, '1', 'performanceTesting381', '1', 'very-high', NULL, NULL, NULL),
	(383, '1', 'performanceTesting382', '1', 'very-high', NULL, NULL, NULL),
	(384, '1', 'performanceTesting383', '1', 'very-high', NULL, NULL, NULL),
	(385, '1', 'performanceTesting384', '1', 'very-high', NULL, NULL, NULL),
	(386, '1', 'performanceTesting385', '1', 'very-high', NULL, NULL, NULL),
	(387, '1', 'performanceTesting386', '1', 'very-high', NULL, NULL, NULL),
	(388, '1', 'performanceTesting387', '1', 'very-high', NULL, NULL, NULL),
	(389, '1', 'performanceTesting388', '1', 'very-high', NULL, NULL, NULL),
	(390, '1', 'performanceTesting389', '1', 'very-high', NULL, NULL, NULL),
	(391, '1', 'performanceTesting390', '1', 'very-high', NULL, NULL, NULL),
	(392, '1', 'performanceTesting391', '1', 'very-high', NULL, NULL, NULL),
	(393, '1', 'performanceTesting392', '1', 'very-high', NULL, NULL, NULL),
	(394, '1', 'performanceTesting393', '1', 'very-high', NULL, NULL, NULL),
	(395, '1', 'performanceTesting394', '1', 'very-high', NULL, NULL, NULL),
	(396, '1', 'performanceTesting395', '1', 'very-high', NULL, NULL, NULL),
	(397, '1', 'performanceTesting396', '1', 'very-high', NULL, NULL, NULL),
	(398, '1', 'performanceTesting397', '1', 'very-high', NULL, NULL, NULL),
	(399, '1', 'performanceTesting398', '1', 'very-high', NULL, NULL, NULL),
	(400, '1', 'performanceTesting399', '1', 'very-high', NULL, NULL, NULL),
	(401, '1', 'performanceTesting400', '1', 'very-high', NULL, NULL, NULL),
	(402, '1', 'performanceTesting401', '1', 'very-high', NULL, NULL, NULL),
	(403, '1', 'performanceTesting402', '1', 'very-high', NULL, NULL, NULL),
	(404, '1', 'performanceTesting403', '1', 'very-high', NULL, NULL, NULL),
	(405, '1', 'performanceTesting404', '1', 'very-high', NULL, NULL, NULL),
	(406, '1', 'performanceTesting405', '1', 'very-high', NULL, NULL, NULL),
	(407, '1', 'performanceTesting406', '1', 'very-high', NULL, NULL, NULL),
	(408, '1', 'performanceTesting407', '1', 'very-high', NULL, NULL, NULL),
	(409, '1', 'performanceTesting408', '1', 'very-high', NULL, NULL, NULL),
	(410, '1', 'performanceTesting409', '1', 'very-high', NULL, NULL, NULL),
	(411, '1', 'performanceTesting410', '1', 'very-high', NULL, NULL, NULL),
	(412, '1', 'performanceTesting411', '1', 'very-high', NULL, NULL, NULL),
	(413, '1', 'performanceTesting412', '1', 'very-high', NULL, NULL, NULL),
	(414, '1', 'performanceTesting413', '1', 'very-high', NULL, NULL, NULL),
	(415, '1', 'performanceTesting414', '1', 'very-high', NULL, NULL, NULL),
	(416, '1', 'performanceTesting415', '1', 'very-high', NULL, NULL, NULL),
	(417, '1', 'performanceTesting416', '1', 'very-high', NULL, NULL, NULL),
	(418, '1', 'performanceTesting417', '1', 'very-high', NULL, NULL, NULL),
	(419, '1', 'performanceTesting418', '1', 'very-high', NULL, NULL, NULL),
	(420, '1', 'performanceTesting419', '1', 'very-high', NULL, NULL, NULL),
	(421, '1', 'performanceTesting420', '1', 'very-high', NULL, NULL, NULL),
	(422, '1', 'performanceTesting421', '1', 'very-high', NULL, NULL, NULL),
	(423, '1', 'performanceTesting422', '1', 'very-high', NULL, NULL, NULL),
	(424, '1', 'performanceTesting423', '1', 'very-high', NULL, NULL, NULL),
	(425, '1', 'performanceTesting424', '1', 'very-high', NULL, NULL, NULL),
	(426, '1', 'performanceTesting425', '1', 'very-high', NULL, NULL, NULL),
	(427, '1', 'performanceTesting426', '1', 'very-high', NULL, NULL, NULL),
	(428, '1', 'performanceTesting427', '1', 'very-high', NULL, NULL, NULL),
	(429, '1', 'performanceTesting428', '1', 'very-high', NULL, NULL, NULL),
	(430, '1', 'performanceTesting429', '1', 'very-high', NULL, NULL, NULL),
	(431, '1', 'performanceTesting430', '1', 'very-high', NULL, NULL, NULL),
	(432, '1', 'performanceTesting431', '1', 'very-high', NULL, NULL, NULL),
	(433, '1', 'performanceTesting432', '1', 'very-high', NULL, NULL, NULL),
	(434, '1', 'performanceTesting433', '1', 'very-high', NULL, NULL, NULL),
	(435, '1', 'performanceTesting434', '1', 'very-high', NULL, NULL, NULL),
	(436, '1', 'performanceTesting435', '1', 'very-high', NULL, NULL, NULL),
	(437, '1', 'performanceTesting436', '1', 'very-high', NULL, NULL, NULL),
	(438, '1', 'performanceTesting437', '1', 'very-high', NULL, NULL, NULL),
	(439, '1', 'performanceTesting438', '1', 'very-high', NULL, NULL, NULL),
	(440, '1', 'performanceTesting439', '1', 'very-high', NULL, NULL, NULL),
	(441, '1', 'performanceTesting440', '1', 'very-high', NULL, NULL, NULL),
	(442, '1', 'performanceTesting441', '1', 'very-high', NULL, NULL, NULL),
	(443, '1', 'performanceTesting442', '1', 'very-high', NULL, NULL, NULL),
	(444, '1', 'performanceTesting443', '1', 'very-high', NULL, NULL, NULL),
	(445, '1', 'performanceTesting444', '1', 'very-high', NULL, NULL, NULL),
	(446, '1', 'performanceTesting445', '1', 'very-high', NULL, NULL, NULL),
	(447, '1', 'performanceTesting446', '1', 'very-high', NULL, NULL, NULL),
	(448, '1', 'performanceTesting447', '1', 'very-high', NULL, NULL, NULL),
	(449, '1', 'performanceTesting448', '1', 'very-high', NULL, NULL, NULL),
	(450, '1', 'performanceTesting449', '1', 'very-high', NULL, NULL, NULL),
	(451, '1', 'performanceTesting450', '1', 'very-high', NULL, NULL, NULL),
	(452, '1', 'performanceTesting451', '1', 'very-high', NULL, NULL, NULL),
	(453, '1', 'performanceTesting452', '1', 'very-high', NULL, NULL, NULL),
	(454, '1', 'performanceTesting453', '1', 'very-high', NULL, NULL, NULL),
	(455, '1', 'performanceTesting454', '1', 'very-high', NULL, NULL, NULL),
	(456, '1', 'performanceTesting455', '1', 'very-high', NULL, NULL, NULL),
	(457, '1', 'performanceTesting456', '1', 'very-high', NULL, NULL, NULL),
	(458, '1', 'performanceTesting457', '1', 'very-high', NULL, NULL, NULL),
	(459, '1', 'performanceTesting458', '1', 'very-high', NULL, NULL, NULL),
	(460, '1', 'performanceTesting459', '1', 'very-high', NULL, NULL, NULL),
	(461, '1', 'performanceTesting460', '1', 'very-high', NULL, NULL, NULL),
	(462, '1', 'performanceTesting461', '1', 'very-high', NULL, NULL, NULL),
	(463, '1', 'performanceTesting462', '1', 'very-high', NULL, NULL, NULL),
	(464, '1', 'performanceTesting463', '1', 'very-high', NULL, NULL, NULL),
	(465, '1', 'performanceTesting464', '1', 'very-high', NULL, NULL, NULL),
	(466, '1', 'performanceTesting465', '1', 'very-high', NULL, NULL, NULL),
	(467, '1', 'performanceTesting466', '1', 'very-high', NULL, NULL, NULL),
	(468, '1', 'performanceTesting467', '1', 'very-high', NULL, NULL, NULL),
	(469, '1', 'performanceTesting468', '1', 'very-high', NULL, NULL, NULL),
	(470, '1', 'performanceTesting469', '1', 'very-high', NULL, NULL, NULL),
	(471, '1', 'performanceTesting470', '1', 'very-high', NULL, NULL, NULL),
	(472, '1', 'performanceTesting471', '1', 'very-high', NULL, NULL, NULL),
	(473, '1', 'performanceTesting472', '1', 'very-high', NULL, NULL, NULL),
	(474, '1', 'performanceTesting473', '1', 'very-high', NULL, NULL, NULL),
	(475, '1', 'performanceTesting474', '1', 'very-high', NULL, NULL, NULL),
	(476, '1', 'performanceTesting475', '1', 'very-high', NULL, NULL, NULL),
	(477, '1', 'performanceTesting476', '1', 'very-high', NULL, NULL, NULL),
	(478, '1', 'performanceTesting477', '1', 'very-high', NULL, NULL, NULL),
	(479, '1', 'performanceTesting478', '1', 'very-high', NULL, NULL, NULL),
	(480, '1', 'performanceTesting479', '1', 'very-high', NULL, NULL, NULL),
	(481, '1', 'performanceTesting480', '1', 'very-high', NULL, NULL, NULL),
	(482, '1', 'performanceTesting481', '1', 'very-high', NULL, NULL, NULL),
	(483, '1', 'performanceTesting482', '1', 'very-high', NULL, NULL, NULL),
	(484, '1', 'performanceTesting483', '1', 'very-high', NULL, NULL, NULL),
	(485, '1', 'performanceTesting484', '1', 'very-high', NULL, NULL, NULL),
	(486, '1', 'performanceTesting485', '1', 'very-high', NULL, NULL, NULL),
	(487, '1', 'performanceTesting486', '1', 'very-high', NULL, NULL, NULL),
	(488, '1', 'performanceTesting487', '1', 'very-high', NULL, NULL, NULL),
	(489, '1', 'performanceTesting488', '1', 'very-high', NULL, NULL, NULL),
	(490, '1', 'performanceTesting489', '1', 'very-high', NULL, NULL, NULL),
	(491, '1', 'performanceTesting490', '1', 'very-high', NULL, NULL, NULL),
	(492, '1', 'performanceTesting491', '1', 'very-high', NULL, NULL, NULL),
	(493, '1', 'performanceTesting492', '1', 'very-high', NULL, NULL, NULL),
	(494, '1', 'performanceTesting493', '1', 'very-high', NULL, NULL, NULL),
	(495, '1', 'performanceTesting494', '1', 'very-high', NULL, NULL, NULL),
	(496, '1', 'performanceTesting495', '1', 'very-high', NULL, NULL, NULL),
	(497, '1', 'performanceTesting496', '1', 'very-high', NULL, NULL, NULL),
	(498, '1', 'performanceTesting497', '1', 'very-high', NULL, NULL, NULL),
	(499, '1', 'performanceTesting498', '1', 'very-high', NULL, NULL, NULL),
	(500, '1', 'performanceTesting499', '1', 'very-high', NULL, NULL, NULL),
	(501, '1', 'performanceTesting500', '1', 'very-high', NULL, NULL, NULL),
	(502, '1', 'performanceTesting501', '1', 'very-high', NULL, NULL, NULL),
	(503, '1', 'performanceTesting502', '1', 'very-high', NULL, NULL, NULL),
	(504, '1', 'performanceTesting503', '1', 'very-high', NULL, NULL, NULL),
	(505, '1', 'performanceTesting504', '1', 'very-high', NULL, NULL, NULL),
	(506, '1', 'performanceTesting505', '1', 'very-high', NULL, NULL, NULL),
	(507, '1', 'performanceTesting506', '1', 'very-high', NULL, NULL, NULL),
	(508, '1', 'performanceTesting507', '1', 'very-high', NULL, NULL, NULL),
	(509, '1', 'performanceTesting508', '1', 'very-high', NULL, NULL, NULL),
	(510, '1', 'performanceTesting509', '1', 'very-high', NULL, NULL, NULL),
	(511, '1', 'performanceTesting510', '1', 'very-high', NULL, NULL, NULL),
	(512, '1', 'performanceTesting511', '1', 'very-high', NULL, NULL, NULL),
	(513, '1', 'performanceTesting512', '1', 'very-high', NULL, NULL, NULL),
	(514, '1', 'performanceTesting513', '1', 'very-high', NULL, NULL, NULL),
	(515, '1', 'performanceTesting514', '1', 'very-high', NULL, NULL, NULL),
	(516, '1', 'performanceTesting515', '1', 'very-high', NULL, NULL, NULL),
	(517, '1', 'performanceTesting516', '1', 'very-high', NULL, NULL, NULL),
	(518, '1', 'performanceTesting517', '1', 'very-high', NULL, NULL, NULL),
	(519, '1', 'performanceTesting518', '1', 'very-high', NULL, NULL, NULL),
	(520, '1', 'performanceTesting519', '1', 'very-high', NULL, NULL, NULL),
	(521, '1', 'performanceTesting520', '1', 'very-high', NULL, NULL, NULL),
	(522, '1', 'performanceTesting521', '1', 'very-high', NULL, NULL, NULL),
	(523, '1', 'performanceTesting522', '1', 'very-high', NULL, NULL, NULL),
	(524, '1', 'performanceTesting523', '1', 'very-high', NULL, NULL, NULL),
	(525, '1', 'performanceTesting524', '1', 'very-high', NULL, NULL, NULL),
	(526, '1', 'performanceTesting525', '1', 'very-high', NULL, NULL, NULL),
	(527, '1', 'performanceTesting526', '1', 'very-high', NULL, NULL, NULL),
	(528, '1', 'performanceTesting527', '1', 'very-high', NULL, NULL, NULL),
	(529, '1', 'performanceTesting528', '1', 'very-high', NULL, NULL, NULL),
	(530, '1', 'performanceTesting529', '1', 'very-high', NULL, NULL, NULL),
	(531, '1', 'performanceTesting530', '1', 'very-high', NULL, NULL, NULL),
	(532, '1', 'performanceTesting531', '1', 'very-high', NULL, NULL, NULL),
	(533, '1', 'performanceTesting532', '1', 'very-high', NULL, NULL, NULL),
	(534, '1', 'performanceTesting533', '1', 'very-high', NULL, NULL, NULL),
	(535, '1', 'performanceTesting534', '1', 'very-high', NULL, NULL, NULL),
	(536, '1', 'performanceTesting535', '1', 'very-high', NULL, NULL, NULL),
	(537, '1', 'performanceTesting536', '1', 'very-high', NULL, NULL, NULL),
	(538, '1', 'performanceTesting537', '1', 'very-high', NULL, NULL, NULL),
	(539, '1', 'performanceTesting538', '1', 'very-high', NULL, NULL, NULL),
	(540, '1', 'performanceTesting539', '1', 'very-high', NULL, NULL, NULL),
	(541, '1', 'performanceTesting540', '1', 'very-high', NULL, NULL, NULL),
	(542, '1', 'performanceTesting541', '1', 'very-high', NULL, NULL, NULL),
	(543, '1', 'performanceTesting542', '1', 'very-high', NULL, NULL, NULL),
	(544, '1', 'performanceTesting543', '1', 'very-high', NULL, NULL, NULL),
	(545, '1', 'performanceTesting544', '1', 'very-high', NULL, NULL, NULL),
	(546, '1', 'performanceTesting545', '1', 'very-high', NULL, NULL, NULL),
	(547, '1', 'performanceTesting546', '1', 'very-high', NULL, NULL, NULL),
	(548, '1', 'performanceTesting547', '1', 'very-high', NULL, NULL, NULL),
	(549, '1', 'performanceTesting548', '1', 'very-high', NULL, NULL, NULL),
	(550, '1', 'performanceTesting549', '1', 'very-high', NULL, NULL, NULL),
	(551, '1', 'performanceTesting550', '1', 'very-high', NULL, NULL, NULL),
	(552, '1', 'performanceTesting551', '1', 'very-high', NULL, NULL, NULL),
	(553, '1', 'performanceTesting552', '1', 'very-high', NULL, NULL, NULL),
	(554, '1', 'performanceTesting553', '1', 'very-high', NULL, NULL, NULL),
	(555, '1', 'performanceTesting554', '1', 'very-high', NULL, NULL, NULL),
	(556, '1', 'performanceTesting555', '1', 'very-high', NULL, NULL, NULL),
	(557, '1', 'performanceTesting556', '1', 'very-high', NULL, NULL, NULL),
	(558, '1', 'performanceTesting557', '1', 'very-high', NULL, NULL, NULL),
	(559, '1', 'performanceTesting558', '1', 'very-high', NULL, NULL, NULL),
	(560, '1', 'performanceTesting559', '1', 'very-high', NULL, NULL, NULL),
	(561, '1', 'performanceTesting560', '1', 'very-high', NULL, NULL, NULL),
	(562, '1', 'performanceTesting561', '1', 'very-high', NULL, NULL, NULL),
	(563, '1', 'performanceTesting562', '1', 'very-high', NULL, NULL, NULL),
	(564, '1', 'performanceTesting563', '1', 'very-high', NULL, NULL, NULL),
	(565, '1', 'performanceTesting564', '1', 'very-high', NULL, NULL, NULL),
	(566, '1', 'performanceTesting565', '1', 'very-high', NULL, NULL, NULL),
	(567, '1', 'performanceTesting566', '1', 'very-high', NULL, NULL, NULL),
	(568, '1', 'performanceTesting567', '1', 'very-high', NULL, NULL, NULL),
	(569, '1', 'performanceTesting568', '1', 'very-high', NULL, NULL, NULL),
	(570, '1', 'performanceTesting569', '1', 'very-high', NULL, NULL, NULL),
	(571, '1', 'performanceTesting570', '1', 'very-high', NULL, NULL, NULL),
	(572, '1', 'performanceTesting571', '1', 'very-high', NULL, NULL, NULL),
	(573, '1', 'performanceTesting572', '1', 'very-high', NULL, NULL, NULL),
	(574, '1', 'performanceTesting573', '1', 'very-high', NULL, NULL, NULL),
	(575, '1', 'performanceTesting574', '1', 'very-high', NULL, NULL, NULL),
	(576, '1', 'performanceTesting575', '1', 'very-high', NULL, NULL, NULL),
	(577, '1', 'performanceTesting576', '1', 'very-high', NULL, NULL, NULL),
	(578, '1', 'performanceTesting577', '1', 'very-high', NULL, NULL, NULL),
	(579, '1', 'performanceTesting578', '1', 'very-high', NULL, NULL, NULL),
	(580, '1', 'performanceTesting579', '1', 'very-high', NULL, NULL, NULL),
	(581, '1', 'performanceTesting580', '1', 'very-high', NULL, NULL, NULL),
	(582, '1', 'performanceTesting581', '1', 'very-high', NULL, NULL, NULL),
	(583, '1', 'performanceTesting582', '1', 'very-high', NULL, NULL, NULL),
	(584, '1', 'performanceTesting583', '1', 'very-high', NULL, NULL, NULL),
	(585, '1', 'performanceTesting584', '1', 'very-high', NULL, NULL, NULL),
	(586, '1', 'performanceTesting585', '1', 'very-high', NULL, NULL, NULL),
	(587, '1', 'performanceTesting586', '1', 'very-high', NULL, NULL, NULL),
	(588, '1', 'performanceTesting587', '1', 'very-high', NULL, NULL, NULL),
	(589, '1', 'performanceTesting588', '1', 'very-high', NULL, NULL, NULL),
	(590, '1', 'performanceTesting589', '1', 'very-high', NULL, NULL, NULL),
	(591, '1', 'performanceTesting590', '1', 'very-high', NULL, NULL, NULL),
	(592, '1', 'performanceTesting591', '1', 'very-high', NULL, NULL, NULL),
	(593, '1', 'performanceTesting592', '1', 'very-high', NULL, NULL, NULL),
	(594, '1', 'performanceTesting593', '1', 'very-high', NULL, NULL, NULL),
	(595, '1', 'performanceTesting594', '1', 'very-high', NULL, NULL, NULL),
	(596, '1', 'performanceTesting595', '1', 'very-high', NULL, NULL, NULL),
	(597, '1', 'performanceTesting596', '1', 'very-high', NULL, NULL, NULL),
	(598, '1', 'performanceTesting597', '1', 'very-high', NULL, NULL, NULL),
	(599, '1', 'performanceTesting598', '1', 'very-high', NULL, NULL, NULL),
	(600, '1', 'performanceTesting599', '1', 'very-high', NULL, NULL, NULL),
	(601, '1', 'performanceTesting600', '1', 'very-high', NULL, NULL, NULL),
	(602, '1', 'performanceTesting601', '1', 'very-high', NULL, NULL, NULL),
	(603, '1', 'performanceTesting602', '1', 'very-high', NULL, NULL, NULL),
	(604, '1', 'performanceTesting603', '1', 'very-high', NULL, NULL, NULL),
	(605, '1', 'performanceTesting604', '1', 'very-high', NULL, NULL, NULL),
	(606, '1', 'performanceTesting605', '1', 'very-high', NULL, NULL, NULL),
	(607, '1', 'performanceTesting606', '1', 'very-high', NULL, NULL, NULL),
	(608, '1', 'performanceTesting607', '1', 'very-high', NULL, NULL, NULL),
	(609, '1', 'performanceTesting608', '1', 'very-high', NULL, NULL, NULL),
	(610, '1', 'performanceTesting609', '1', 'very-high', NULL, NULL, NULL),
	(611, '1', 'performanceTesting610', '1', 'very-high', NULL, NULL, NULL),
	(612, '1', 'performanceTesting611', '1', 'very-high', NULL, NULL, NULL),
	(613, '1', 'performanceTesting612', '1', 'very-high', NULL, NULL, NULL),
	(614, '1', 'performanceTesting613', '1', 'very-high', NULL, NULL, NULL),
	(615, '1', 'performanceTesting614', '1', 'very-high', NULL, NULL, NULL),
	(616, '1', 'performanceTesting615', '1', 'very-high', NULL, NULL, NULL),
	(617, '1', 'performanceTesting616', '1', 'very-high', NULL, NULL, NULL),
	(618, '1', 'performanceTesting617', '1', 'very-high', NULL, NULL, NULL),
	(619, '1', 'performanceTesting618', '1', 'very-high', NULL, NULL, NULL),
	(620, '1', 'performanceTesting619', '1', 'very-high', NULL, NULL, NULL),
	(621, '1', 'performanceTesting620', '1', 'very-high', NULL, NULL, NULL),
	(622, '1', 'performanceTesting621', '1', 'very-high', NULL, NULL, NULL),
	(623, '1', 'performanceTesting622', '1', 'very-high', NULL, NULL, NULL),
	(624, '1', 'performanceTesting623', '1', 'very-high', NULL, NULL, NULL),
	(625, '1', 'performanceTesting624', '1', 'very-high', NULL, NULL, NULL),
	(626, '1', 'performanceTesting625', '1', 'very-high', NULL, NULL, NULL),
	(627, '1', 'performanceTesting626', '1', 'very-high', NULL, NULL, NULL),
	(628, '1', 'performanceTesting627', '1', 'very-high', NULL, NULL, NULL),
	(629, '1', 'performanceTesting628', '1', 'very-high', NULL, NULL, NULL),
	(630, '1', 'performanceTesting629', '1', 'very-high', NULL, NULL, NULL),
	(631, '1', 'performanceTesting630', '1', 'very-high', NULL, NULL, NULL),
	(632, '1', 'performanceTesting631', '1', 'very-high', NULL, NULL, NULL),
	(633, '1', 'performanceTesting632', '1', 'very-high', NULL, NULL, NULL),
	(634, '1', 'performanceTesting633', '1', 'very-high', NULL, NULL, NULL),
	(635, '1', 'performanceTesting634', '1', 'very-high', NULL, NULL, NULL),
	(636, '1', 'performanceTesting635', '1', 'very-high', NULL, NULL, NULL),
	(637, '1', 'performanceTesting636', '1', 'very-high', NULL, NULL, NULL),
	(638, '1', 'performanceTesting637', '1', 'very-high', NULL, NULL, NULL),
	(639, '1', 'performanceTesting638', '1', 'very-high', NULL, NULL, NULL),
	(640, '1', 'performanceTesting639', '1', 'very-high', NULL, NULL, NULL),
	(641, '1', 'performanceTesting640', '1', 'very-high', NULL, NULL, NULL),
	(642, '1', 'performanceTesting641', '1', 'very-high', NULL, NULL, NULL),
	(643, '1', 'performanceTesting642', '1', 'very-high', NULL, NULL, NULL),
	(644, '1', 'performanceTesting643', '1', 'very-high', NULL, NULL, NULL),
	(645, '1', 'performanceTesting644', '1', 'very-high', NULL, NULL, NULL),
	(646, '1', 'performanceTesting645', '1', 'very-high', NULL, NULL, NULL),
	(647, '1', 'performanceTesting646', '1', 'very-high', NULL, NULL, NULL),
	(648, '1', 'performanceTesting647', '1', 'very-high', NULL, NULL, NULL),
	(649, '1', 'performanceTesting648', '1', 'very-high', NULL, NULL, NULL),
	(650, '1', 'performanceTesting649', '1', 'very-high', NULL, NULL, NULL),
	(651, '1', 'performanceTesting650', '1', 'very-high', NULL, NULL, NULL),
	(652, '1', 'performanceTesting651', '1', 'very-high', NULL, NULL, NULL),
	(653, '1', 'performanceTesting652', '1', 'very-high', NULL, NULL, NULL),
	(654, '1', 'performanceTesting653', '1', 'very-high', NULL, NULL, NULL),
	(655, '1', 'performanceTesting654', '1', 'very-high', NULL, NULL, NULL),
	(656, '1', 'performanceTesting655', '1', 'very-high', NULL, NULL, NULL),
	(657, '1', 'performanceTesting656', '1', 'very-high', NULL, NULL, NULL),
	(658, '1', 'performanceTesting657', '1', 'very-high', NULL, NULL, NULL),
	(659, '1', 'performanceTesting658', '1', 'very-high', NULL, NULL, NULL),
	(660, '1', 'performanceTesting659', '1', 'very-high', NULL, NULL, NULL),
	(661, '1', 'performanceTesting660', '1', 'very-high', NULL, NULL, NULL),
	(662, '1', 'performanceTesting661', '1', 'very-high', NULL, NULL, NULL),
	(663, '1', 'performanceTesting662', '1', 'very-high', NULL, NULL, NULL),
	(664, '1', 'performanceTesting663', '1', 'very-high', NULL, NULL, NULL),
	(665, '1', 'performanceTesting664', '1', 'very-high', NULL, NULL, NULL),
	(666, '1', 'performanceTesting665', '1', 'very-high', NULL, NULL, NULL),
	(667, '1', 'performanceTesting666', '1', 'very-high', NULL, NULL, NULL),
	(668, '1', 'performanceTesting667', '1', 'very-high', NULL, NULL, NULL),
	(669, '1', 'performanceTesting668', '1', 'very-high', NULL, NULL, NULL),
	(670, '1', 'performanceTesting669', '1', 'very-high', NULL, NULL, NULL),
	(671, '1', 'performanceTesting670', '1', 'very-high', NULL, NULL, NULL),
	(672, '1', 'performanceTesting671', '1', 'very-high', NULL, NULL, NULL),
	(673, '1', 'performanceTesting672', '1', 'very-high', NULL, NULL, NULL),
	(674, '1', 'performanceTesting673', '1', 'very-high', NULL, NULL, NULL),
	(675, '1', 'performanceTesting674', '1', 'very-high', NULL, NULL, NULL),
	(676, '1', 'performanceTesting675', '1', 'very-high', NULL, NULL, NULL),
	(677, '1', 'performanceTesting676', '1', 'very-high', NULL, NULL, NULL),
	(678, '1', 'performanceTesting677', '1', 'very-high', NULL, NULL, NULL),
	(679, '1', 'performanceTesting678', '1', 'very-high', NULL, NULL, NULL),
	(680, '1', 'performanceTesting679', '1', 'very-high', NULL, NULL, NULL),
	(681, '1', 'performanceTesting680', '1', 'very-high', NULL, NULL, NULL),
	(682, '1', 'performanceTesting681', '1', 'very-high', NULL, NULL, NULL),
	(683, '1', 'performanceTesting682', '1', 'very-high', NULL, NULL, NULL),
	(684, '1', 'performanceTesting683', '1', 'very-high', NULL, NULL, NULL),
	(685, '1', 'performanceTesting684', '1', 'very-high', NULL, NULL, NULL),
	(686, '1', 'performanceTesting685', '1', 'very-high', NULL, NULL, NULL),
	(687, '1', 'performanceTesting686', '1', 'very-high', NULL, NULL, NULL),
	(688, '1', 'performanceTesting687', '1', 'very-high', NULL, NULL, NULL),
	(689, '1', 'performanceTesting688', '1', 'very-high', NULL, NULL, NULL),
	(690, '1', 'performanceTesting689', '1', 'very-high', NULL, NULL, NULL),
	(691, '1', 'performanceTesting690', '1', 'very-high', NULL, NULL, NULL),
	(692, '1', 'performanceTesting691', '1', 'very-high', NULL, NULL, NULL),
	(693, '1', 'performanceTesting692', '1', 'very-high', NULL, NULL, NULL),
	(694, '1', 'performanceTesting693', '1', 'very-high', NULL, NULL, NULL),
	(695, '1', 'performanceTesting694', '1', 'very-high', NULL, NULL, NULL),
	(696, '1', 'performanceTesting695', '1', 'very-high', NULL, NULL, NULL),
	(697, '1', 'performanceTesting696', '1', 'very-high', NULL, NULL, NULL),
	(698, '1', 'performanceTesting697', '1', 'very-high', NULL, NULL, NULL),
	(699, '1', 'performanceTesting698', '1', 'very-high', NULL, NULL, NULL),
	(700, '1', 'performanceTesting699', '1', 'very-high', NULL, NULL, NULL),
	(701, '1', 'performanceTesting700', '1', 'very-high', NULL, NULL, NULL),
	(702, '1', 'performanceTesting701', '1', 'very-high', NULL, NULL, NULL),
	(703, '1', 'performanceTesting702', '1', 'very-high', NULL, NULL, NULL),
	(704, '1', 'performanceTesting703', '1', 'very-high', NULL, NULL, NULL),
	(705, '1', 'performanceTesting704', '1', 'very-high', NULL, NULL, NULL),
	(706, '1', 'performanceTesting705', '1', 'very-high', NULL, NULL, NULL),
	(707, '1', 'performanceTesting706', '1', 'very-high', NULL, NULL, NULL),
	(708, '1', 'performanceTesting707', '1', 'very-high', NULL, NULL, NULL),
	(709, '1', 'performanceTesting708', '1', 'very-high', NULL, NULL, NULL),
	(710, '1', 'performanceTesting709', '1', 'very-high', NULL, NULL, NULL),
	(711, '1', 'performanceTesting710', '1', 'very-high', NULL, NULL, NULL),
	(712, '1', 'performanceTesting711', '1', 'very-high', NULL, NULL, NULL),
	(713, '1', 'performanceTesting712', '1', 'very-high', NULL, NULL, NULL),
	(714, '1', 'performanceTesting713', '1', 'very-high', NULL, NULL, NULL),
	(715, '1', 'performanceTesting714', '1', 'very-high', NULL, NULL, NULL),
	(716, '1', 'performanceTesting715', '1', 'very-high', NULL, NULL, NULL),
	(717, '1', 'performanceTesting716', '1', 'very-high', NULL, NULL, NULL),
	(718, '1', 'performanceTesting717', '1', 'very-high', NULL, NULL, NULL),
	(719, '1', 'performanceTesting718', '1', 'very-high', NULL, NULL, NULL),
	(720, '1', 'performanceTesting719', '1', 'very-high', NULL, NULL, NULL),
	(721, '1', 'performanceTesting720', '1', 'very-high', NULL, NULL, NULL),
	(722, '1', 'performanceTesting721', '1', 'very-high', NULL, NULL, NULL),
	(723, '1', 'performanceTesting722', '1', 'very-high', NULL, NULL, NULL),
	(724, '1', 'performanceTesting723', '1', 'very-high', NULL, NULL, NULL),
	(725, '1', 'performanceTesting724', '1', 'very-high', NULL, NULL, NULL),
	(726, '1', 'performanceTesting725', '1', 'very-high', NULL, NULL, NULL),
	(727, '1', 'performanceTesting726', '1', 'very-high', NULL, NULL, NULL),
	(728, '1', 'performanceTesting727', '1', 'very-high', NULL, NULL, NULL),
	(729, '1', 'performanceTesting728', '1', 'very-high', NULL, NULL, NULL),
	(730, '1', 'performanceTesting729', '1', 'very-high', NULL, NULL, NULL),
	(731, '1', 'performanceTesting730', '1', 'very-high', NULL, NULL, NULL),
	(732, '1', 'performanceTesting731', '1', 'very-high', NULL, NULL, NULL),
	(733, '1', 'performanceTesting732', '1', 'very-high', NULL, NULL, NULL),
	(734, '1', 'performanceTesting733', '1', 'very-high', NULL, NULL, NULL),
	(735, '1', 'performanceTesting734', '1', 'very-high', NULL, NULL, NULL),
	(736, '1', 'performanceTesting735', '1', 'very-high', NULL, NULL, NULL),
	(737, '1', 'performanceTesting736', '1', 'very-high', NULL, NULL, NULL),
	(738, '1', 'performanceTesting737', '1', 'very-high', NULL, NULL, NULL),
	(739, '1', 'performanceTesting738', '1', 'very-high', NULL, NULL, NULL),
	(740, '1', 'performanceTesting739', '1', 'very-high', NULL, NULL, NULL),
	(741, '1', 'performanceTesting740', '1', 'very-high', NULL, NULL, NULL),
	(742, '1', 'performanceTesting741', '1', 'very-high', NULL, NULL, NULL),
	(743, '1', 'performanceTesting742', '1', 'very-high', NULL, NULL, NULL),
	(744, '1', 'performanceTesting743', '1', 'very-high', NULL, NULL, NULL),
	(745, '1', 'performanceTesting744', '1', 'very-high', NULL, NULL, NULL),
	(746, '1', 'performanceTesting745', '1', 'very-high', NULL, NULL, NULL),
	(747, '1', 'performanceTesting746', '1', 'very-high', NULL, NULL, NULL),
	(748, '1', 'performanceTesting747', '1', 'very-high', NULL, NULL, NULL),
	(749, '1', 'performanceTesting748', '1', 'very-high', NULL, NULL, NULL),
	(750, '1', 'performanceTesting749', '1', 'very-high', NULL, NULL, NULL),
	(751, '1', 'performanceTesting750', '1', 'very-high', NULL, NULL, NULL),
	(752, '1', 'performanceTesting751', '1', 'very-high', NULL, NULL, NULL),
	(753, '1', 'performanceTesting752', '1', 'very-high', NULL, NULL, NULL),
	(754, '1', 'performanceTesting753', '1', 'very-high', NULL, NULL, NULL),
	(755, '1', 'performanceTesting754', '1', 'very-high', NULL, NULL, NULL),
	(756, '1', 'performanceTesting755', '1', 'very-high', NULL, NULL, NULL),
	(757, '1', 'performanceTesting756', '1', 'very-high', NULL, NULL, NULL),
	(758, '1', 'performanceTesting757', '1', 'very-high', NULL, NULL, NULL),
	(759, '1', 'performanceTesting758', '1', 'very-high', NULL, NULL, NULL),
	(760, '1', 'performanceTesting759', '1', 'very-high', NULL, NULL, NULL),
	(761, '1', 'performanceTesting760', '1', 'very-high', NULL, NULL, NULL),
	(762, '1', 'performanceTesting761', '1', 'very-high', NULL, NULL, NULL),
	(763, '1', 'performanceTesting762', '1', 'very-high', NULL, NULL, NULL),
	(764, '1', 'performanceTesting763', '1', 'very-high', NULL, NULL, NULL),
	(765, '1', 'performanceTesting764', '1', 'very-high', NULL, NULL, NULL),
	(766, '1', 'performanceTesting765', '1', 'very-high', NULL, NULL, NULL),
	(767, '1', 'performanceTesting766', '1', 'very-high', NULL, NULL, NULL),
	(768, '1', 'performanceTesting767', '1', 'very-high', NULL, NULL, NULL),
	(769, '1', 'performanceTesting768', '1', 'very-high', NULL, NULL, NULL),
	(770, '1', 'performanceTesting769', '1', 'very-high', NULL, NULL, NULL),
	(771, '1', 'performanceTesting770', '1', 'very-high', NULL, NULL, NULL),
	(772, '1', 'performanceTesting771', '1', 'very-high', NULL, NULL, NULL),
	(773, '1', 'performanceTesting772', '1', 'very-high', NULL, NULL, NULL),
	(774, '1', 'performanceTesting773', '1', 'very-high', NULL, NULL, NULL),
	(775, '1', 'performanceTesting774', '1', 'very-high', NULL, NULL, NULL),
	(776, '1', 'performanceTesting775', '1', 'very-high', NULL, NULL, NULL),
	(777, '1', 'performanceTesting776', '1', 'very-high', NULL, NULL, NULL),
	(778, '1', 'performanceTesting777', '1', 'very-high', NULL, NULL, NULL),
	(779, '1', 'performanceTesting778', '1', 'very-high', NULL, NULL, NULL),
	(780, '1', 'performanceTesting779', '1', 'very-high', NULL, NULL, NULL),
	(781, '1', 'performanceTesting780', '1', 'very-high', NULL, NULL, NULL),
	(782, '1', 'performanceTesting781', '1', 'very-high', NULL, NULL, NULL),
	(783, '1', 'performanceTesting782', '1', 'very-high', NULL, NULL, NULL),
	(784, '1', 'performanceTesting783', '1', 'very-high', NULL, NULL, NULL),
	(785, '1', 'performanceTesting784', '1', 'very-high', NULL, NULL, NULL),
	(786, '1', 'performanceTesting785', '1', 'very-high', NULL, NULL, NULL),
	(787, '1', 'performanceTesting786', '1', 'very-high', NULL, NULL, NULL),
	(788, '1', 'performanceTesting787', '1', 'very-high', NULL, NULL, NULL),
	(789, '1', 'performanceTesting788', '1', 'very-high', NULL, NULL, NULL),
	(790, '1', 'performanceTesting789', '1', 'very-high', NULL, NULL, NULL),
	(791, '1', 'performanceTesting790', '1', 'very-high', NULL, NULL, NULL),
	(792, '1', 'performanceTesting791', '1', 'very-high', NULL, NULL, NULL),
	(793, '1', 'performanceTesting792', '1', 'very-high', NULL, NULL, NULL),
	(794, '1', 'performanceTesting793', '1', 'very-high', NULL, NULL, NULL),
	(795, '1', 'performanceTesting794', '1', 'very-high', NULL, NULL, NULL),
	(796, '1', 'performanceTesting795', '1', 'very-high', NULL, NULL, NULL),
	(797, '1', 'performanceTesting796', '1', 'very-high', NULL, NULL, NULL),
	(798, '1', 'performanceTesting797', '1', 'very-high', NULL, NULL, NULL),
	(799, '1', 'performanceTesting798', '1', 'very-high', NULL, NULL, NULL),
	(800, '1', 'performanceTesting799', '1', 'very-high', NULL, NULL, NULL),
	(801, '1', 'performanceTesting800', '1', 'very-high', NULL, NULL, NULL),
	(802, '1', 'performanceTesting801', '1', 'very-high', NULL, NULL, NULL),
	(803, '1', 'performanceTesting802', '1', 'very-high', NULL, NULL, NULL),
	(804, '1', 'performanceTesting803', '1', 'very-high', NULL, NULL, NULL),
	(805, '1', 'performanceTesting804', '1', 'very-high', NULL, NULL, NULL),
	(806, '1', 'performanceTesting805', '1', 'very-high', NULL, NULL, NULL),
	(807, '1', 'performanceTesting806', '1', 'very-high', NULL, NULL, NULL),
	(808, '1', 'performanceTesting807', '1', 'very-high', NULL, NULL, NULL),
	(809, '1', 'performanceTesting808', '1', 'very-high', NULL, NULL, NULL),
	(810, '1', 'performanceTesting809', '1', 'very-high', NULL, NULL, NULL),
	(811, '1', 'performanceTesting810', '1', 'very-high', NULL, NULL, NULL),
	(812, '1', 'performanceTesting811', '1', 'very-high', NULL, NULL, NULL),
	(813, '1', 'performanceTesting812', '1', 'very-high', NULL, NULL, NULL),
	(814, '1', 'performanceTesting813', '1', 'very-high', NULL, NULL, NULL),
	(815, '1', 'performanceTesting814', '1', 'very-high', NULL, NULL, NULL),
	(816, '1', 'performanceTesting815', '1', 'very-high', NULL, NULL, NULL),
	(817, '1', 'performanceTesting816', '1', 'very-high', NULL, NULL, NULL),
	(818, '1', 'performanceTesting817', '1', 'very-high', NULL, NULL, NULL),
	(819, '1', 'performanceTesting818', '1', 'very-high', NULL, NULL, NULL),
	(820, '1', 'performanceTesting819', '1', 'very-high', NULL, NULL, NULL),
	(821, '1', 'performanceTesting820', '1', 'very-high', NULL, NULL, NULL),
	(822, '1', 'performanceTesting821', '1', 'very-high', NULL, NULL, NULL),
	(823, '1', 'performanceTesting822', '1', 'very-high', NULL, NULL, NULL),
	(824, '1', 'performanceTesting823', '1', 'very-high', NULL, NULL, NULL),
	(825, '1', 'performanceTesting824', '1', 'very-high', NULL, NULL, NULL),
	(826, '1', 'performanceTesting825', '1', 'very-high', NULL, NULL, NULL),
	(827, '1', 'performanceTesting826', '1', 'very-high', NULL, NULL, NULL),
	(828, '1', 'performanceTesting827', '1', 'very-high', NULL, NULL, NULL),
	(829, '1', 'performanceTesting828', '1', 'very-high', NULL, NULL, NULL),
	(830, '1', 'performanceTesting829', '1', 'very-high', NULL, NULL, NULL),
	(831, '1', 'performanceTesting830', '1', 'very-high', NULL, NULL, NULL),
	(832, '1', 'performanceTesting831', '1', 'very-high', NULL, NULL, NULL),
	(833, '1', 'performanceTesting832', '1', 'very-high', NULL, NULL, NULL),
	(834, '1', 'performanceTesting833', '1', 'very-high', NULL, NULL, NULL),
	(835, '1', 'performanceTesting834', '1', 'very-high', NULL, NULL, NULL),
	(836, '1', 'performanceTesting835', '1', 'very-high', NULL, NULL, NULL),
	(837, '1', 'performanceTesting836', '1', 'very-high', NULL, NULL, NULL),
	(838, '1', 'performanceTesting837', '1', 'very-high', NULL, NULL, NULL),
	(839, '1', 'performanceTesting838', '1', 'very-high', NULL, NULL, NULL),
	(840, '1', 'performanceTesting839', '1', 'very-high', NULL, NULL, NULL),
	(841, '1', 'performanceTesting840', '1', 'very-high', NULL, NULL, NULL),
	(842, '1', 'performanceTesting841', '1', 'very-high', NULL, NULL, NULL),
	(843, '1', 'performanceTesting842', '1', 'very-high', NULL, NULL, NULL),
	(844, '1', 'performanceTesting843', '1', 'very-high', NULL, NULL, NULL),
	(845, '1', 'performanceTesting844', '1', 'very-high', NULL, NULL, NULL),
	(846, '1', 'performanceTesting845', '1', 'very-high', NULL, NULL, NULL),
	(847, '1', 'performanceTesting846', '1', 'very-high', NULL, NULL, NULL),
	(848, '1', 'performanceTesting847', '1', 'very-high', NULL, NULL, NULL),
	(849, '1', 'performanceTesting848', '1', 'very-high', NULL, NULL, NULL),
	(850, '1', 'performanceTesting849', '1', 'very-high', NULL, NULL, NULL),
	(851, '1', 'performanceTesting850', '1', 'very-high', NULL, NULL, NULL),
	(852, '1', 'performanceTesting851', '1', 'very-high', NULL, NULL, NULL),
	(853, '1', 'performanceTesting852', '1', 'very-high', NULL, NULL, NULL),
	(854, '1', 'performanceTesting853', '1', 'very-high', NULL, NULL, NULL),
	(855, '1', 'performanceTesting854', '1', 'very-high', NULL, NULL, NULL),
	(856, '1', 'performanceTesting855', '1', 'very-high', NULL, NULL, NULL),
	(857, '1', 'performanceTesting856', '1', 'very-high', NULL, NULL, NULL),
	(858, '1', 'performanceTesting857', '1', 'very-high', NULL, NULL, NULL),
	(859, '1', 'performanceTesting858', '1', 'very-high', NULL, NULL, NULL),
	(860, '1', 'performanceTesting859', '1', 'very-high', NULL, NULL, NULL),
	(861, '1', 'performanceTesting860', '1', 'very-high', NULL, NULL, NULL),
	(862, '1', 'performanceTesting861', '1', 'very-high', NULL, NULL, NULL),
	(863, '1', 'performanceTesting862', '1', 'very-high', NULL, NULL, NULL),
	(864, '1', 'performanceTesting863', '1', 'very-high', NULL, NULL, NULL),
	(865, '1', 'performanceTesting864', '1', 'very-high', NULL, NULL, NULL),
	(866, '1', 'performanceTesting865', '1', 'very-high', NULL, NULL, NULL),
	(867, '1', 'performanceTesting866', '1', 'very-high', NULL, NULL, NULL),
	(868, '1', 'performanceTesting867', '1', 'very-high', NULL, NULL, NULL),
	(869, '1', 'performanceTesting868', '1', 'very-high', NULL, NULL, NULL),
	(870, '1', 'performanceTesting869', '1', 'very-high', NULL, NULL, NULL),
	(871, '1', 'performanceTesting870', '1', 'very-high', NULL, NULL, NULL),
	(872, '1', 'performanceTesting871', '1', 'very-high', NULL, NULL, NULL),
	(873, '1', 'performanceTesting872', '1', 'very-high', NULL, NULL, NULL),
	(874, '1', 'performanceTesting873', '1', 'very-high', NULL, NULL, NULL),
	(875, '1', 'performanceTesting874', '1', 'very-high', NULL, NULL, NULL),
	(876, '1', 'performanceTesting875', '1', 'very-high', NULL, NULL, NULL),
	(877, '1', 'performanceTesting876', '1', 'very-high', NULL, NULL, NULL),
	(878, '1', 'performanceTesting877', '1', 'very-high', NULL, NULL, NULL),
	(879, '1', 'performanceTesting878', '1', 'very-high', NULL, NULL, NULL),
	(880, '1', 'performanceTesting879', '1', 'very-high', NULL, NULL, NULL),
	(881, '1', 'performanceTesting880', '1', 'very-high', NULL, NULL, NULL),
	(882, '1', 'performanceTesting881', '1', 'very-high', NULL, NULL, NULL),
	(883, '1', 'performanceTesting882', '1', 'very-high', NULL, NULL, NULL),
	(884, '1', 'performanceTesting883', '1', 'very-high', NULL, NULL, NULL),
	(885, '1', 'performanceTesting884', '1', 'very-high', NULL, NULL, NULL),
	(886, '1', 'performanceTesting885', '1', 'very-high', NULL, NULL, NULL),
	(887, '1', 'performanceTesting886', '1', 'very-high', NULL, NULL, NULL),
	(888, '1', 'performanceTesting887', '1', 'very-high', NULL, NULL, NULL),
	(889, '1', 'performanceTesting888', '1', 'very-high', NULL, NULL, NULL),
	(890, '1', 'performanceTesting889', '1', 'very-high', NULL, NULL, NULL),
	(891, '1', 'performanceTesting890', '1', 'very-high', NULL, NULL, NULL),
	(892, '1', 'performanceTesting891', '1', 'very-high', NULL, NULL, NULL),
	(893, '1', 'performanceTesting892', '1', 'very-high', NULL, NULL, NULL),
	(894, '1', 'performanceTesting893', '1', 'very-high', NULL, NULL, NULL),
	(895, '1', 'performanceTesting894', '1', 'very-high', NULL, NULL, NULL),
	(896, '1', 'performanceTesting895', '1', 'very-high', NULL, NULL, NULL),
	(897, '1', 'performanceTesting896', '1', 'very-high', NULL, NULL, NULL),
	(898, '1', 'performanceTesting897', '1', 'very-high', NULL, NULL, NULL),
	(899, '1', 'performanceTesting898', '1', 'very-high', NULL, NULL, NULL),
	(900, '1', 'performanceTesting899', '1', 'very-high', NULL, NULL, NULL),
	(901, '1', 'performanceTesting900', '1', 'very-high', NULL, NULL, NULL),
	(902, '1', 'performanceTesting901', '1', 'very-high', NULL, NULL, NULL),
	(903, '1', 'performanceTesting902', '1', 'very-high', NULL, NULL, NULL),
	(904, '1', 'performanceTesting903', '1', 'very-high', NULL, NULL, NULL),
	(905, '1', 'performanceTesting904', '1', 'very-high', NULL, NULL, NULL),
	(906, '1', 'performanceTesting905', '1', 'very-high', NULL, NULL, NULL),
	(907, '1', 'performanceTesting906', '1', 'very-high', NULL, NULL, NULL),
	(908, '1', 'performanceTesting907', '1', 'very-high', NULL, NULL, NULL),
	(909, '1', 'performanceTesting908', '1', 'very-high', NULL, NULL, NULL),
	(910, '1', 'performanceTesting909', '1', 'very-high', NULL, NULL, NULL),
	(911, '1', 'performanceTesting910', '1', 'very-high', NULL, NULL, NULL),
	(912, '1', 'performanceTesting911', '1', 'very-high', NULL, NULL, NULL),
	(913, '1', 'performanceTesting912', '1', 'very-high', NULL, NULL, NULL),
	(914, '1', 'performanceTesting913', '1', 'very-high', NULL, NULL, NULL),
	(915, '1', 'performanceTesting914', '1', 'very-high', NULL, NULL, NULL),
	(916, '1', 'performanceTesting915', '1', 'very-high', NULL, NULL, NULL),
	(917, '1', 'performanceTesting916', '1', 'very-high', NULL, NULL, NULL),
	(918, '1', 'performanceTesting917', '1', 'very-high', NULL, NULL, NULL),
	(919, '1', 'performanceTesting918', '1', 'very-high', NULL, NULL, NULL),
	(920, '1', 'performanceTesting919', '1', 'very-high', NULL, NULL, NULL),
	(921, '1', 'performanceTesting920', '1', 'very-high', NULL, NULL, NULL),
	(922, '1', 'performanceTesting921', '1', 'very-high', NULL, NULL, NULL),
	(923, '1', 'performanceTesting922', '1', 'very-high', NULL, NULL, NULL),
	(924, '1', 'performanceTesting923', '1', 'very-high', NULL, NULL, NULL),
	(925, '1', 'performanceTesting924', '1', 'very-high', NULL, NULL, NULL),
	(926, '1', 'performanceTesting925', '1', 'very-high', NULL, NULL, NULL),
	(927, '1', 'performanceTesting926', '1', 'very-high', NULL, NULL, NULL),
	(928, '1', 'performanceTesting927', '1', 'very-high', NULL, NULL, NULL),
	(929, '1', 'performanceTesting928', '1', 'very-high', NULL, NULL, NULL),
	(930, '1', 'performanceTesting929', '1', 'very-high', NULL, NULL, NULL),
	(931, '1', 'performanceTesting930', '1', 'very-high', NULL, NULL, NULL),
	(932, '1', 'performanceTesting931', '1', 'very-high', NULL, NULL, NULL),
	(933, '1', 'performanceTesting932', '1', 'very-high', NULL, NULL, NULL),
	(934, '1', 'performanceTesting933', '1', 'very-high', NULL, NULL, NULL),
	(935, '1', 'performanceTesting934', '1', 'very-high', NULL, NULL, NULL),
	(936, '1', 'performanceTesting935', '1', 'very-high', NULL, NULL, NULL),
	(937, '1', 'performanceTesting936', '1', 'very-high', NULL, NULL, NULL),
	(938, '1', 'performanceTesting937', '1', 'very-high', NULL, NULL, NULL),
	(939, '1', 'performanceTesting938', '1', 'very-high', NULL, NULL, NULL),
	(940, '1', 'performanceTesting939', '1', 'very-high', NULL, NULL, NULL),
	(941, '1', 'performanceTesting940', '1', 'very-high', NULL, NULL, NULL),
	(942, '1', 'performanceTesting941', '1', 'very-high', NULL, NULL, NULL),
	(943, '1', 'performanceTesting942', '1', 'very-high', NULL, NULL, NULL),
	(944, '1', 'performanceTesting943', '1', 'very-high', NULL, NULL, NULL),
	(945, '1', 'performanceTesting944', '1', 'very-high', NULL, NULL, NULL),
	(946, '1', 'performanceTesting945', '1', 'very-high', NULL, NULL, NULL),
	(947, '1', 'performanceTesting946', '1', 'very-high', NULL, NULL, NULL),
	(948, '1', 'performanceTesting947', '1', 'very-high', NULL, NULL, NULL),
	(949, '1', 'performanceTesting948', '1', 'very-high', NULL, NULL, NULL),
	(950, '1', 'performanceTesting949', '1', 'very-high', NULL, NULL, NULL),
	(951, '1', 'performanceTesting950', '1', 'very-high', NULL, NULL, NULL),
	(952, '1', 'performanceTesting951', '1', 'very-high', NULL, NULL, NULL),
	(953, '1', 'performanceTesting952', '1', 'very-high', NULL, NULL, NULL),
	(954, '1', 'performanceTesting953', '1', 'very-high', NULL, NULL, NULL),
	(955, '1', 'performanceTesting954', '1', 'very-high', NULL, NULL, NULL),
	(956, '1', 'performanceTesting955', '1', 'very-high', NULL, NULL, NULL),
	(957, '1', 'performanceTesting956', '1', 'very-high', NULL, NULL, NULL),
	(958, '1', 'performanceTesting957', '1', 'very-high', NULL, NULL, NULL),
	(959, '1', 'performanceTesting958', '1', 'very-high', NULL, NULL, NULL),
	(960, '1', 'performanceTesting959', '1', 'very-high', NULL, NULL, NULL),
	(961, '1', 'performanceTesting960', '1', 'very-high', NULL, NULL, NULL),
	(962, '1', 'performanceTesting961', '1', 'very-high', NULL, NULL, NULL),
	(963, '1', 'performanceTesting962', '1', 'very-high', NULL, NULL, NULL),
	(964, '1', 'performanceTesting963', '1', 'very-high', NULL, NULL, NULL),
	(965, '1', 'performanceTesting964', '1', 'very-high', NULL, NULL, NULL),
	(966, '1', 'performanceTesting965', '1', 'very-high', NULL, NULL, NULL),
	(967, '1', 'performanceTesting966', '1', 'very-high', NULL, NULL, NULL),
	(968, '1', 'performanceTesting967', '1', 'very-high', NULL, NULL, NULL),
	(969, '1', 'performanceTesting968', '1', 'very-high', NULL, NULL, NULL),
	(970, '1', 'performanceTesting969', '1', 'very-high', NULL, NULL, NULL),
	(971, '1', 'performanceTesting970', '1', 'very-high', NULL, NULL, NULL),
	(972, '1', 'performanceTesting971', '1', 'very-high', NULL, NULL, NULL),
	(973, '1', 'performanceTesting972', '1', 'very-high', NULL, NULL, NULL),
	(974, '1', 'performanceTesting973', '1', 'very-high', NULL, NULL, NULL),
	(975, '1', 'performanceTesting974', '1', 'very-high', NULL, NULL, NULL),
	(976, '1', 'performanceTesting975', '1', 'very-high', NULL, NULL, NULL),
	(977, '1', 'performanceTesting976', '1', 'very-high', NULL, NULL, NULL),
	(978, '1', 'performanceTesting977', '1', 'very-high', NULL, NULL, NULL),
	(979, '1', 'performanceTesting978', '1', 'very-high', NULL, NULL, NULL),
	(980, '1', 'performanceTesting979', '1', 'very-high', NULL, NULL, NULL),
	(981, '1', 'performanceTesting980', '1', 'very-high', NULL, NULL, NULL),
	(982, '1', 'performanceTesting981', '1', 'very-high', NULL, NULL, NULL),
	(983, '1', 'performanceTesting982', '1', 'very-high', NULL, NULL, NULL),
	(984, '1', 'performanceTesting983', '1', 'very-high', NULL, NULL, NULL),
	(985, '1', 'performanceTesting984', '1', 'very-high', NULL, NULL, NULL),
	(986, '1', 'performanceTesting985', '1', 'very-high', NULL, NULL, NULL),
	(987, '1', 'performanceTesting986', '1', 'very-high', NULL, NULL, NULL),
	(988, '1', 'performanceTesting987', '1', 'very-high', NULL, NULL, NULL),
	(989, '1', 'performanceTesting988', '1', 'very-high', NULL, NULL, NULL),
	(990, '1', 'performanceTesting989', '1', 'very-high', NULL, NULL, NULL),
	(991, '1', 'performanceTesting990', '1', 'very-high', NULL, NULL, NULL),
	(992, '1', 'performanceTesting991', '1', 'very-high', NULL, NULL, NULL),
	(993, '1', 'performanceTesting992', '1', 'very-high', NULL, NULL, NULL),
	(994, '1', 'performanceTesting993', '1', 'very-high', NULL, NULL, NULL),
	(995, '1', 'performanceTesting994', '1', 'very-high', NULL, NULL, NULL),
	(996, '1', 'performanceTesting995', '1', 'very-high', NULL, NULL, NULL),
	(997, '1', 'performanceTesting996', '1', 'very-high', NULL, NULL, NULL),
	(998, '1', 'performanceTesting997', '1', 'very-high', NULL, NULL, NULL),
	(999, '1', 'performanceTesting998', '1', 'very-high', NULL, NULL, NULL),
	(1000, '1', 'performanceTesting999', '1', 'very-high', NULL, NULL, NULL),
	(1001, '1', 'performanceTesting1000', '1', 'very-high', NULL, NULL, NULL);
	`)

	// currentActivity = 1
	// currentTodo = 1
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":3000", nil)
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
		activities = append(activities, activity)

		w.Write(jData)

	case "GET":

		resp.Status = "Success"
		resp.Message = "Success"

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
	switch r.Method {
	case "POST":

		decoder := json.NewDecoder(r.Body)
		var t Todo
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Fprint(w, "Test Error")
			return
		}

		if t.ActivityGroupId == nil {

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

		if t.Title == nil {

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

		t.Priority = "very-high"

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

func HandleParamActivity(w http.ResponseWriter, r *http.Request, ids string) {

	if ids == "999999999" {
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
			resp.Data = activities[0]
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

			activities[0].Title = activity.Title
			if activity.Email != "" {
				activities[0].Email = activity.Email
			}

			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = activities[0]
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

	if ids == "999999999" {
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
			resp.Data = todos[0]
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

			todos[0].IsActive = t.IsActive

			if t.Title != "" && t.Title != nil {
				todos[0].Title = t.Title
			}

			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = todos[0]
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
	var path = r.URL.Path

	if path == "/todo-items" {
		if getTodo > 5 {
			if r.Method == "POST" {
				w.WriteHeader(http.StatusCreated)
			}
			return
		}

		getTodo = getTodo + 1
		w.Header().Set("Content-Type", "application/json")

		TodoRest(w, r)
		return
	}

	if path == "/activity-groups" {
		w.Header().Set("Content-Type", "application/json")
		ActivityRest(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if path == "/todo-items/1" {
		HandleParamTodo(w, r, "1")
		return
	}

	if path == "/activity-groups/1" {
		HandleParamActivity(w, r, "1")
		return
	}

	if path == "/todo-items/999999999" {
		HandleParamTodo(w, r, "999999999")
		return
	}

	if path == "/activity-groups/999999999" {
		HandleParamActivity(w, r, "999999999")
		return
	}

	names := "Oke"
	jData, err := json.Marshal(names)
	if err != nil {
		fmt.Fprint(w, "Internal Server Error")
	} else {
		w.Write(jData)
	}

}
