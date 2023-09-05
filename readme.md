# 函数一览

函数名|说明|文件
---|---|---
main|程序入口|main.go
checkError|错误检查与处理|main.go
executeOutput|输出生成文本（Java Class）|main.go
getPropertyMap|从JSON文件获取字段名和数据类型|main.go
snakeToCamel|Snake Case转Camel Case|string.go
capitalizeFirst|首字符大写|string.go
getJavaType|Java类型转换|main.go
NewConfig|初始化配置|config.go
getConfig|获取配置|main.go

## main.go
主文件，存放主要流程相关函数。
### main
程序主入口。
### checkError
检查调用函数时是否返回错误，如果返回显示错误信息并中止程序运行。
### executeOutput
将生成的文本（预想为Java类定义）根据配置输出到命令行或文件中。
### getPropertyMap
读取JSON文件，并根据该JSON的一级属性获取【Java属性名->Java类型】形式的键值对。
### getJavaType
将通过JSON获取到的GO类型
### getConfig
根据Config的默认值和config.yml的配置生成全局配置信息。

## config.go
存放Config结构体定义以及构造函数。
### newConfig
Config结构体的构造函数，自带各配置属性默认值。

## string.go
存放字符串处理相关函数。
### snakeToCamel
将Snake Case字符串转换成Camel Case字符串，_后面的字母大写。
### capitalizeFirst
将字符串的首字母转换成大写。