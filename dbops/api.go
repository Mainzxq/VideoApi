package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	defs "../defs"
	utils "../utils"
	"time"
)


// 基础要自己学习

func AddUserCredential(loginName string, pwd string) error {
	// Prepare 可以预处理数据库查询字串，避免出现字串攻击
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	//站退出才执行
	defer stmtIns.Close()
	return nil

}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "",err
	}
	stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error  {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

//返回整个OBJ会比较好，
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error)  {
	// create uuid
	vid, err := utils.NewUUID()
	if  err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // 固定的格式化字串
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info
(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)

	res := &defs.VideoInfo{Name: name, Id: vid, AuthorId: aid, DisplayCtime:ctime}

	defer stmtIns.Close()
	return res, nil
}

func GetVideo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT id FROM video_info WHERE id=? ")
	if err != nil {
		return nil, err
	}

	stmtOut.Exec(vid)
}