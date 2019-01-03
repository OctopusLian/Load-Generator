## Load Generator / 载荷发生器  
编译环境 Ubuntu 18.04 + liteide + VSCode  
软件性能评测的辅助工具，可向被测软件发送指定量的载荷，并记录被测软件处理载荷的结果。  
#### 测试运行  
- TestStart  
`go test -v -run=TestStart`  
- TestStop  
`go test -v -run=TestStop`  
#### 目录树  
```
.
├── gen.go
├── gen_test.go
├── lib
│   ├── base.go
│   ├── caller.go
│   └── gotickets.go
├── parameter.go
├── README.md
└── testhelper
    ├── comm.go
    ├── log
    │   ├── base
    │   │   ├── base.go
    │   │   ├── format.go
    │   │   ├── level.go
    │   │   ├── location.go
    │   │   └── type.go
    │   ├── field
    │   │   ├── cmd
    │   │   │   └── log_xfields_generator.go
    │   │   ├── field.go
    │   │   └── xfields.go
    │   ├── logger.go
    │   ├── logger_test.go
    │   └── logrus
    │       ├── logrus.go
    │       └── logrus_test.go
    └── server.go

7 directories, 21 files

```