# 功能实现



用户：登陆，注册，修改密码

问题：提出问题，修改问题，删除问题，获取所有问题，获取自己的问题

回答：写回答，修改回答，删除回答，获得自己的所有回答



## 接口

### 用户层   /user

​         登录 post      /login

​         注册 post    /register

​         修改密码 put     /changepass



### 问题层  /question

​         提出问题    post     /create

​         修改问题  put    /change

​         删除问题 delete    /delete

​         获取所有问题 get   /getall

​         获取自己的所有问题 get    /getmine



### 回答层 /answer

​         提出问题 post   /create

​         删除问题 delete  /delete

​         修改问题 put   /change

​          获取自己的回答   /getmine





## 数据库

user

| UID (主键)       USERNAME     PASSWORD |      |      |
| -------------------------------------- | ---- | ---- |
|                                        |      |      |

question

| QID  (主键)       uid         question |      |      |
| -------------------------------------- | ---- | ---- |
|                                        |      |      |

answer

| aid       qid      uid       answer |      |      |
| ----------------------------------- | ---- | ---- |
|                                     |      |      |

