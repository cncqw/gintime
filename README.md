# Gin Time

基于GO和Vue.js的Gin框架项目

## 安装Gin
```
go get github.com/gin-gonic/gin
```

## 使用JWT
```
go get github.com/dgrijalva/jwt-go
```

## 文件读取配置
```
go get github.com/spf13/viper
```

## 安装Vue

1. 下载 nvm_setup.zip

https://github.com/coreybutler/nvm

2. 创建环境变量
```
NVM_NODEJS_ORG_MIRROR=https://npm.taobao.org/mirrors/node
```
   
3. 修改node源
```
nvm node_mirror https://npm.taobao.org/mirrors/node/
nvm npm_mirror https://npm.taobao.org/mirrors/npm/
```

4. 安装node.js
```
nvm install v14.0.0
nvm use 14.0.0
```

5. 安装yarn
```
npm install -g yarn
yarn config set registry https://registry.npm.taobao.org
```

6. 安装vue-cli
```
yarn global add @vue/cli
```

Windows下添加到Path环境变量
```
C:\Users\UserName\AppData\Local\Yarn\Data\global\node_modules\.bin
```

