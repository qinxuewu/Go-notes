/**
 *
 * @author qinxuewu
 * @create 19/6/25下午8:28
 * @since 1.0.0
 */
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Users struct {
	ID        uint `gorm:"primary_key"`
	Name string
	Age int
}




func main()  {
	dbs,_:=gorm.Open("mysql", "root:870439570@tcp(39.108.144.143:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer dbs.Close()
	// 检测指定表是否存在
	falg:=dbs.HasTable(&Users{})
	fmt.Println("指定表表名是否存在: ",falg)

	// 自动迁移模式
	dbs.AutoMigrate(&Users{})

	user := Users{Name: "qxw", Age: 18}
	dbs.Create(&user)

	// 获取第一条记录，按主键排序
	dbs.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	fmt.Println("获取第一条记录，按主键排序:",user)

	// 获取最后一条记录，按主键排序
	dbs.Last(&user)
	fmt.Println("获取最后一条记录，按主键排序:",user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	users := []Users{}
	// 获取所有记录
	dbs.Find(&users)
	fmt.Println("获取所有记录: ",users)
	//// SELECT * FROM users;

	// 使用主键获取记录
	dbs.First(&user, 1)
	fmt.Println("使用主键获取记录:",user)

	// SELECT * FROM users WHERE id = 1;

	//Where(dbs)
	//WhereStruct(dbs)
	//WherNot(dbs)
	//WhereParms(dbs)
	//WhereOr(dbs)
	//WHereLinkd(dbs)
	//WhereSet(dbs)
	WhereFirstOrInit(dbs)
}

func Where(db *gorm.DB)  {
	user := Users{}
	users := []Users{}

	// 获取第一个匹配记录
	db.Where("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

	// 获取所有匹配记录
	db.Where("name = ?", "jinzhu").Find(&users)
	for i, v := range users {
		fmt.Println("index :", i, v.Name)
	}
	// SELECT * FROM users WHERE name = 'jinzhu';

	db.Where("name <> ?", "jinzhu").Find(&users)

	// IN
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

	// LIKE
	db.Where("name LIKE ?", "%jin%").Find(&users)

	// AND
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	// Time
	db.Where("age > ?").Find(&users)

}


// 当使用struct查询时，GORM将只查询那些具有值的字段
func WhereStruct(db *gorm.DB)  {
	user := Users{}
	users := []Users{}

	// 使用Struct查询
	db.Where(&Users{Name:"jinzhu",Age:20}).First(&user)
    // SELECT * FROM users WHERE name = 'jinzhu' AND age=20;
	fmt.Println("使用Struct查询 :",user)

	// 使用Map 查询
	db.Where(map[string]interface{}{"name":"jinzhu","age":20}).First(&user)

	fmt.Println("使用Map 查询 :",user)

	// 主键的Slice

	db.Where([]int64{1,2,3}).Find(&users)
	//	SELECT * FROM users WHERE id IN (1 2, 3);

	fmt.Println("使用主键的Slice查询 :",users)
}


// Not条件查询
func WherNot(db *gorm.DB)  {
	user := Users{}
	users := []Users{}
	db.Not("name","jinzhu").First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1

	fmt.Println("使用Not查询 :",user)

	// Not In
	db.Not("name",[]string{"jinzhu","qxw"}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "qxw");

	// 使用主键查询模式的 not in
	db.Not([]int64{1,2,3}).Find(&users)
	// SELECT * FROM users WHERE id NOT IN (1,2,3);

	db.Not([]int64{}).First(&user) // SELECT * FROM users;

	//not和sql拼接查询
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT(name = "jinzhu");

	// Struct not查询
	db.Not(Users{Name: "jinzhu"}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu";
}

// 带内联条件的查询  占位符 参数
// 使用主键查询时，应仔细检查所传递的值是否为有效主键，以避免SQL注入
func WhereParms(db *gorm.DB)  {
	user := Users{}
	users := []Users{}

	// 主键查询
	db.First(&user,23) //  SELECT * FROM users WHERE id = 23 LIMIT 1;

	// 简单sql查询 占位符
	db.Find(&users,"name = ?","jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	db.Find(&users,Users{Age:20})
	// SELECT * FROM users WHERE age = 20;

	// map多条件查询
	db.Find(&users,map[string]interface{}{"age":20})
	// SELECT * FROM users WHERE age = 20;
}
// Or条件查询
func WhereOr(db *gorm.DB)  {

	users := []Users{}
	db.Where("name = ?","jinzhu").Or("name = ?","qxw").Find(&users)
	// select *  from users where name='jinzhu' or name='qxw';

	// Struct结构体查询
	db.Where("name = jinzhu").Or(Users{Name:"qxw"}).Find(&users)
	// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'qxw';

	// Map多条件查询
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name":"qxw"}).Find(&users)
}

//  Gorm有一个可链接的API，你可以这样使用它
func WHereLinkd(db *gorm.DB)  {
	users := []Users{}

	db.Where("name <> ?","jinzhu").Where("age >= ? ",20).Find(&users)

//	SELECT * FROM users WHERE name <> 'jinzhu' AND age >= 20 ;

    db.Where("name = ?","jinzhu").Or("age =?",20).Not("name = ?","qxw")
//	 select * from name='jinzhu' or age=20 and name <> 'qxw';
}

// 扩展查询选项
func WhereSet(db *gorm.DB)  {
	user := Users{}

	// 为Select语句添加扩展SQL选项  排它锁
	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)
	//// SELECT * FROM users WHERE id = 10 FOR UPDATE;
}

// 获取第一个匹配的记录，或者使用给定的条件初始化一个新的记录（仅适用于struct，map条件）
func WhereFirstOrInit(db *gorm.DB)  {
	user := Users{}
	// 未查询导数据 则返回指定的初始化数据
	db.FirstOrInit(&user,Users{Name:"non_existing"})
	// user -> User{Name: "non_existing"}

   //  查询有返回数据 则直接返回
	db.Where(Users{Name:"jinzhu"}).FirstOrInit(&user)
	// user -> User{Id: 111, Name: "Jinzhu", Age: 20}

	db.FirstOrInit(&user, map[string]interface{}{"name": "jinzhu"})
	// user -> User{Id: 111, Name: "Jinzhu", Age: 20}
}

// 如果未找到记录，则使用参数初始化结构
func WhereAttrs(db *gorm.DB)  {
	user := Users{}

	// 未查询到数据 则返回参数初始化结构
	db.Where(Users{Name:"non_qxwsss"}).Attrs(Users{Age:20}).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = 'non_qxwsss';
	// user -> User{Name: "non_qxwsss", Age: 20}

	//或
	db.Where(Users{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = 'non_existing';
	// user -> User{Name: "non_existing", Age: 20}

	// 查询到数据。 则直接返回查询的数据 忽略参数初始化的数据
	db.Where(Users{Name: "Jinzhu"}).Attrs(Users{Age: 30}).FirstOrInit(&user)

//	如果未找到记录，则为参数分配结构 也就是创建新记录

   // 未查询到 先添加 然后返回
   db.Where(Users{Name:"non_extisting"}).Attrs(Users{Age:20}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'non_existing';
	// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
	// 最后返回数据   user -> User{Id: 112, Name: "non_existing", Age: 20}

	//  查询到数据则直接返回查询的数据
	db.Where(Users{Name: "jinzhu"}).Attrs(Users{Age: 30}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'jinzhu';
	// user -> User{Id: 111, Name: "jinzhu", Age: 20}
}

//  将参数分配给结果，不管它是否被找到
func WhereAssign(db *gorm.DB)  {
	user := Users{}
	db.Where(Users{Name: "non_existing"}).Assign(Users{Age: 20}).FirstOrInit(&user)
	// user -> User{Name: "non_existing", Age: 20}

	db.Where(Users{Name: "Jinzhu"}).Assign(Users{Age: 30}).FirstOrInit(&user)
	// SELECT * FROM USERS WHERE name = jinzhu';
	// user -> User{Id: 111, Name: "Jinzhu", Age: 30}


	//  将其分配给记录，而不管它是否被找到，并保存回数据库。

	// 未查询到数据 就添加 然后返回
	db.Where(Users{Name: "non_existing"}).Assign(Users{Age: 20}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'non_existing';
	// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
	// user -> User{Id: 112, Name: "non_existing", Age: 20}

	// 查询到数据 就更新 并返回更新后的数据
	db.Where(Users{Name: "jinzhu"}).Assign(Users{Age: 30}).FirstOrCreate(&user)
	// SELECT * FROM users WHERE name = 'jinzhu';
	// UPDATE users SET age=30 WHERE id = 111;
	// user -> User{Id: 111, Name: "jinzhu", Age: 30}

}
// 获取第一个匹配的记录，或创建一个具有给定条件的新记录（仅适用于struct, map条件）
func WhereFirstOrCreate(db *gorm.DB)  {
	user := Users{}

	// 未查询到数据 则添加并返回
	db.FirstOrCreate(&user,Users{Name:"non_extisting"})
	// INSERT INTO "users" (name) VALUES ("non_existing");
	// user -> User{Id: 112, Name: "non_existing"}

	//  查询到数据则直接返回
	db.Where(Users{Name: "Jinzhu"}).FirstOrCreate(&user)
	// user -> User{Id: 111, Name: "Jinzhu"}
}

// 指定要从数据库检索的字段，默认情况下，将选择所有字段;
func WhereSelect(db *gorm.DB)  {
	users := []Users{}
	db.Select("name","age").Find(&users);
    //   SELECT name, age FROM users;

    db.Select([]string{"name","age"}).Find(&users)
	//   SELECT name, age FROM users;

	db.Table("users").Select("COALESCE(age,?)",42).Rows()
    //  SELECT COALESCE(age,'42') FROM users;
}

// 在从数据库检索记录时指定顺序，将重排序设置为true以覆盖定义的条件
func WhereOrder(db *gorm.DB)  {
//	 SELECT * FROM users ORDER BY age desc, name;

	users := []Users{}
	db.Order("age desc").Order("name").Find(&users)

	users1 := []Users{}
	users2 := []Users{}
	db.Order("age desc").Find(&users1).Order("age",true).Find(&users2)
	// SELECT * FROM users ORDER BY age desc; (users1)
	// SELECT * FROM users ORDER BY age; (users2)
}
// 指定要检索的记录数
func WhereLimit(db *gorm.DB) {

	users := []Users{}
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	users1 := []Users{}
	users2 := []Users{}
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)
}

// 指定在开始返回记录之前要跳过的记录数
func WhereOffset(db *gorm.DB)  {
	users := []Users{}
	db.Offset(3).Find(&users)
	//  SELECT * FROM users OFFSET 3;

	users1 := []Users{}
	users2 := []Users{}
	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)
}

func WhereGroupHaving(db *gorm.DB)  {
	rows, _ :=db.Table("users").Select("name,sum(age) as totalAge").Group("name").Rows()
	for rows.Next() {
		var id uint
		var name string
		var age int
		rows.Scan(&id, &name, &age)
		fmt.Println(id,name,age)
	}

	rows2,_:=db.Table("users").Select("name").Group("name").Having("count(name) > ?", 1).Rows()
	for  rows2.Next()  {
		var id uint
		var name string
		var age int
		rows2.Scan(&id, &name, &age)
		fmt.Println(id,name,age)
	}
}

type Result struct {
	Name string
	Email string
}

//  指定连接条件
func WhereJoin(db *gorm.DB)  {
	rows, _ := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	for rows.Next() {

	}
	results:=[]Result{};
	db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	
}

//  Save将包括执行更新SQL时的所有字段，即使它没有更改
// 如果只想更新更改的字段，可以使用Update, Updates
func SaveUpdate(db *gorm.DB)  {
	user := Users{}
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)

	// UPDATE users SET name='jinzhu 2', age=100 WHERE id=111;


	// 更新单个属性（如果更改）
	db.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello'  WHERE id=111;

	// 使用组合条件更新单个属性
	db.Model(&user).Where("age = ?", 20).Update("name", "hello")

	// UPDATE users SET name='hello' WHERE id=111 AND age=20;

	// 使用`map`更新多个属性，只会更新这些更改的字段
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18,})
	//// UPDATE users SET name='hello', age=18   WHERE id=111;

	// 使用`struct`更新多个属性，只会更新这些更改的和非空白字段
	db.Model(&user).Updates(Users{Name: "hello", Age: 18})
	// UPDATE users SET name='hello', age=18,  WHERE id = 111;

	// 警告:当使用struct更新时，FORM将仅更新具有非空值的字段
	// 对于下面的更新，什么都不会更新为""，0，false是其类型的空白值
	db.Model(&user).Updates(Users{Name: "", Age: 0})

//	 如果您只想在更新时更新或忽略某些字段，可以使用Select, Omit

	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET name='hello'  WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET age=18  WHERE id=111;

}
