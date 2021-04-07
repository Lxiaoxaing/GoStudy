package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/**
1、更新全部字段
*/
func Update_saveUser(user User) {
	//1、save:保存，id不存在时为新增
	//       更新，id存在时为更新，默认是更新所有的字段.
	user.Name = "小北"
	user.Age = 100
	user.Id = 1
	save := db.Save(&user)
	fmt.Println(save)
}

/**
2、更新修改字段
警告：如果id=0，则更新所有的数据
*/
func Update_updateUser(user User) {
	//更新单个属性。
	user.Id = 1
	db.Model(&user).Update("name", "hello")

	//更新指定条件
	db.Model(&user).Where("age=?", 100).Update("name", "hello")

	//Updates更新多个字段。map方式
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": "18"})

	//Updates更新多个字段。struct方式.  警告：使用struct更新时，只更新其中有变化且为非零的属性
	db.Model(&user).Updates(User{Name: "小黑黑", Age: 33})
}

/**
3、更新选定字段
警告：如果id=0，则更新所有的数据
*/
func Update_updateSelect(user User) {
	//更新单个属性。
	user.Id = 111
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": "18"})

	//UPDATE users SET age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Omit().Updates(map[string]interface{}{"name": "hello", "age": "18"})
}

/**
4、无 Hooks 更新
上面的更新操作会自动运行 model 的 BeforeUpdate, AfterUpdate 方法，更新 UpdatedAt 时间戳, 在更新时保存其 Associations, 如果你不想调用这些方法，你可以使用 UpdateColumn， UpdateColumns
*/
func Update_NoHooks(user User) {
	//更新单个属性。
	user.Id = 111
	db.Model(&user).UpdateColumn("name", "hello")

	//更新过个属性
	db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
}

/**
5、批量更新
*/
func Update_bulk(user User) {
	db.Table("user").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	//db.Where("id IN (?)",[]int{10,11}).Updates(map[string]interface{}{"name": "hello", "age": 18 })

	//使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}
	db.Model(User{}).Updates(User{Name: "hello", Age: 18})

	// 使用 `RowsAffected` 获取更新记录总数
	i := db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected
	fmt.Println("更新条数：", i)
}

/**
5、SQL表达式更新
*/
func Update_sql(user User) {
	//UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2'
	db.Model(&user).Update("age", gorm.Expr("age * ?+?", 1, 1))

	db.Model(&user).Updates(map[string]interface{}{"age": gorm.Expr("age * ?+?", 1, 1)})

	db.Model(&user).UpdateColumn("age", gorm.Expr("age * ?+?", 1, 1))

	db.Model(&user).Where("name=?", "小北").UpdateColumn("age", gorm.Expr("age * ?+?", 1, 1))
}

/**
6、修改Hooks中的值
*/
func Update_updateHooks(user User) {
}

/**
7、其他更新选项
 */
func Update_other(user User)  {
}