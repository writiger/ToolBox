package utils

import "fmt"

func Hello() {
	openWords := `
	-------------------------------------------------------------
	github:"https://github.com/writiger/ToolBox"
	一个简单的练手项目
	主要目的
	1. 实践代码整洁之道
	2. 体会设计模式
	3. 接触redis
	4. 熟悉gin+xorm开发
	5. 了解RESTFUL风格接口开发
	==================================================================
	| ___________           .__ __________               __      __  |
	| \__    ___/___   ____ |  |\______   \ _______  ___/  \    /  \ |
	|   |    | /  _ \ /  _ \|  | |    |  _//  _ \  \/  /\   \/\/   / |
	|   |    |(  <_> |  <_> )  |_|    |   (  <_> >    <  \        /  |
	|   |____| \____/ \____/|____/______  /\____/__/\_ \  \__/\  /   |
	|                                   \/            \/       \/    |
	==================================================================


	`
	fmt.Printf("%s\n", openWords)
}