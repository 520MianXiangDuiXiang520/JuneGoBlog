package utils

import (
	"JuneGoBlog/src/dao"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"testing"
)

func TestUseTransaction(t *testing.T) {
	res, _ := UseTransaction(func(a int, b string) (string, error) {
		as, err := strconv.Atoi(b)
		if err != nil {
			return "", err
		}
		fmt.Println(a, b)
		if as != a {
			return "UnEqual", nil
		} else {
			s := a / as
			fmt.Println(s)
			return "Equal", nil
		}
	}, []interface{}{0, "0"})
	fmt.Println(res)
}

func TestUseTransaction2(t *testing.T) {
	_, _ = UseTransaction(func(tx *gorm.DB, id int) error {
		var err error
		// 删除 article 表中的数据
		err = tx.Where("id = ?", id).Delete(&dao.Article{}).Error
		if err != nil {
			return err
		}
		// 删除 article_tag 表中的数据
		err = tx.Where("article_id = ?", id).Delete(&dao.ArticleTags{}).Error
		return err

	}, []interface{}{&gorm.DB{}, 70})

}
