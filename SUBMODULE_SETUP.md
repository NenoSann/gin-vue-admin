# Git Submodule 配置完成指南

## 🎯 当前状态

您已经完成：
- ✅ 创建了两个 private repository (server 和 web)
- ✅ 将代码推送到 GitHub

## 🚀 下一步：配置 Submodules

### 方式 1：使用自动化脚本（推荐）

```bash
cd /Users/liangheng/Code/Web/gin-vue-admin
./setup-submodules.sh
```

脚本会自动：
1. 检测 server 和 web 的远程仓库地址
2. 创建备份
3. 从主仓库中移除 server 和 web
4. 将它们添加为 submodules
5. 提交更改
6. 可选：立即推送到远程

### 方式 2：手动配置

如果您想手动操作，按以下步骤：

#### Step 1: 备份

```bash
cd /Users/liangheng/Code/Web/gin-vue-admin
cp -r . ../gin-vue-admin-backup
```

#### Step 2: 获取远程仓库地址

```bash
cd server
SERVER_REPO=$(git remote get-url origin)
echo "Server: $SERVER_REPO"
cd ..

cd web
WEB_REPO=$(git remote get-url origin)
echo "Web: $WEB_REPO"
cd ..
```

#### Step 3: 从主仓库移除 server 和 web

```bash
# 如果 server 和 web 在主仓库中被跟踪，需要移除
git rm -r --cached server web
git commit -m "Remove server and web for submodule conversion"
```

#### Step 4: 删除子目录的 .git

```bash
rm -rf server/.git
rm -rf web/.git
```

#### Step 5: 临时移动目录

```bash
mv server ../server-temp
mv web ../web-temp
```

#### Step 6: 添加 submodules

```bash
git submodule add <server-repo-url> server
git submodule add <web-repo-url> web
```

例如：
```bash
git submodule add git@github.com:NenoSann/gin-vue-admin-server.git server
git submodule add git@github.com:NenoSann/gin-vue-admin-web.git web
```

#### Step 7: 删除临时目录

```bash
rm -rf ../server-temp
rm -rf ../web-temp
```

#### Step 8: 提交并推送

```bash
git add .gitmodules server web
git commit -m "Add server and web as submodules"
git push
```

## 📋 验证配置

配置完成后，验证：

```bash
# 查看 submodules 状态
git submodule status

# 查看 .gitmodules 文件
cat .gitmodules
```

`.gitmodules` 应该包含：

```ini
[submodule "server"]
	path = server
	url = git@github.com:NenoSann/gin-vue-admin-server.git
[submodule "web"]
	path = web
	url = git@github.com:NenoSann/gin-vue-admin-web.git
```

## 👥 团队成员使用

### 新成员克隆项目

```bash
# 方式 1: 同时克隆主仓库和 submodules
git clone --recursive git@github.com:NenoSann/gin-vue-admin.git

# 方式 2: 先克隆主仓库，再初始化 submodules
git clone git@github.com:NenoSann/gin-vue-admin.git
cd gin-vue-admin
git submodule init
git submodule update
```

### 前端开发者

```bash
# 方式 1: 只克隆 web 仓库
git clone git@github.com:NenoSann/gin-vue-admin-web.git

# 方式 2: 克隆主仓库，只更新 web submodule
git clone git@github.com:NenoSann/gin-vue-admin.git
cd gin-vue-admin
git submodule init web
git submodule update web
```

### 后端开发者

```bash
# 方式 1: 只克隆 server 仓库
git clone git@github.com:NenoSann/gin-vue-admin-server.git

# 方式 2: 克隆主仓库，只更新 server submodule
git clone git@github.com:NenoSann/gin-vue-admin.git
cd gin-vue-admin
git submodule init server
git submodule update server
```

## 🔄 日常工作流程

### 在 server 中开发

```bash
cd server
# 进行开发
git add .
git commit -m "feat: add new feature"
git push

# 更新主仓库的 submodule 引用
cd ..
git add server
git commit -m "Update server submodule"
git push
```

### 在 web 中开发

```bash
cd web
# 进行开发
git add .
git commit -m "feat: update UI"
git push

# 更新主仓库的 submodule 引用
cd ..
git add web
git commit -m "Update web submodule"
git push
```

### 拉取最新的 submodule 更新

```bash
# 方式 1: 拉取所有 submodules 的最新更改
git submodule update --remote

# 方式 2: 只更新特定 submodule
git submodule update --remote server
git submodule update --remote web

# 方式 3: 进入 submodule 目录手动拉取
cd server
git pull origin main
cd ..

cd web
git pull origin main
cd ..
```

## 🔧 常用命令

```bash
# 查看所有 submodules
git submodule

# 查看 submodule 状态
git submodule status

# 初始化 submodules
git submodule init

# 更新 submodules
git submodule update

# 更新到最新版本
git submodule update --remote

# 递归初始化和更新
git submodule update --init --recursive

# 克隆时包含 submodules
git clone --recursive <repo-url>
```

## ⚠️ 注意事项

1. **提交顺序**
   - 先提交 submodule 的更改
   - 再提交主仓库的 submodule 引用更新

2. **私有仓库权限**
   - 确保团队成员有 server 和 web 仓库的访问权限
   - 配置 SSH 密钥或 HTTPS 凭证

3. **分支管理**
   - Submodule 默认处于 detached HEAD 状态
   - 开发时切换到对应分支：`cd server && git checkout main`

4. **更新冲突**
   - 如果 submodule 有未提交的更改，`git submodule update` 会失败
   - 需要先提交或 stash 更改

## 🎉 完成

配置完成后，您的项目结构：

```
gin-vue-admin/          (主仓库)
├── .gitmodules         (submodule 配置文件)
├── server/             (submodule -> gin-vue-admin-server)
│   └── .git/          (指向 server 仓库)
├── web/               (submodule -> gin-vue-admin-web)
│   └── .git/          (指向 web 仓库)
└── (其他文件)
```

现在：
- ✅ 前端修改只影响 web 仓库
- ✅ 后端修改只影响 server 仓库
- ✅ 主仓库只记录 submodule 的版本引用
- ✅ 团队协作清晰明了
