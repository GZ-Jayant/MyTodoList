package db

import (
	"github.com/gin-gonic/gin"
)

func TaskCreate(c *gin.Context) (int, error) {
	tid := <-maxTaskID + 1

	// insert task
	_, err := dbConn.Exec(`insert into todo_task_info(
		task_id, task_name, task_desc, task_status, task_start_date, task_end_date
		)values(?, ?, ?, ?, ?, ?)`,
		tid, c.PostForm("task_name"), c.PostForm("task_desc"), c.PostForm("task_status"),
		c.PostForm("task_start_date"), c.PostForm("task_end_date"))
	if err != nil {
		return -1, err
	}

	maxTaskID <- tid
	return tid, nil
}

func TaskUpdate(c *gin.Context) error {
	_, err := dbConn.Exec(`update todo_task_info 
		set task_name = ?, task_desc = ?, task_status = ?,
		task_start_date = ?, task_end_date = ?
		where task_id = ?`,
		c.PostForm("task_name"), c.PostForm("task_desc"), c.PostForm("task_status"),
		c.PostForm("task_start_date"), c.PostForm("task_end_date"), c.Param("taskID"))
	if err != nil {
		return err
	}
	return nil
}

func TaskDelete(c *gin.Context) error {
	_, err := dbConn.Exec(`delete from todo_task_info where task_id = ?`, c.Param("taskID"))
	if err != nil {
		return err
	}
	return nil
}

func TaskSelect(c *gin.Context) ([]Task, error) {
	var (
		taskList []Task
		condi    string = ""
	)

	if c.PostForm("task_status") != "all" {
		condi = `where task_status = ` + c.PostForm("task_status")
	}

	dql := `select task_id, task_name, task_desc, task_status, task_start_date, task_end_date 
		from todo_task_info ` + condi + `order by task_end_date desc`
	rows, err := dbConn.Query(dql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		rows.Scan(&task.TaskID, &task.TaskName, &task.TaskDesc, &task.TaskStatus,
			&task.TaskStartDate, &task.TaskEndDate)
		taskList = append(taskList, task)
	}

	return taskList, nil

}
