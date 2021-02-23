package service

import (
	"gin-vue-admin/utils"
)

//@function: EmailTest
//@description: 发送邮件测试
//@return: err error

func EmailTest() (err error) {
<<<<<<< HEAD
	subject := "test"
	body := "test"
=======
	subject := "test111"
	body := "test111"
>>>>>>> develop
	err = utils.EmailTest(subject, body)
	return err
}
