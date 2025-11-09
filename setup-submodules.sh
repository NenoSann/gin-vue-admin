#!/bin/bash

# Submodule 配置脚本
# 用于将已推送的 server 和 web 转换为 submodules

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}   Git Submodule 配置助手${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# 检查是否在项目根目录
if [ ! -d "server" ] || [ ! -d "web" ]; then
    echo -e "${RED}错误: 请在项目根目录运行此脚本${NC}"
    exit 1
fi

# 获取 server 和 web 的远程仓库地址
echo -e "${BLUE}Step 1: 检测远程仓库地址...${NC}"
echo ""

cd server
SERVER_REMOTE=$(git remote get-url origin 2>/dev/null || echo "")
cd ..

cd web
WEB_REMOTE=$(git remote get-url origin 2>/dev/null || echo "")
cd ..

if [ -z "$SERVER_REMOTE" ] || [ -z "$WEB_REMOTE" ]; then
    echo -e "${RED}错误: 未检测到 server 或 web 的远程仓库${NC}"
    echo ""
    echo "请确保："
    echo "1. server 目录已推送到远程: cd server && git remote -v"
    echo "2. web 目录已推送到远程: cd web && git remote -v"
    exit 1
fi

echo -e "${GREEN}✓ 检测到远程仓库:${NC}"
echo "  Server: $SERVER_REMOTE"
echo "  Web: $WEB_REMOTE"
echo ""

# 确认
read -p "是否使用以上仓库地址创建 submodules? [y/N] " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "已取消"
    exit 0
fi

# 创建备份
echo ""
echo -e "${BLUE}Step 2: 创建备份...${NC}"
BACKUP_DIR="../gin-vue-admin-backup-$(date +%Y%m%d-%H%M%S)"
cp -r . "$BACKUP_DIR"
echo -e "${GREEN}✓ 备份完成: $BACKUP_DIR${NC}"
echo ""

# 检查 server 和 web 是否已经被 Git 跟踪
echo -e "${BLUE}Step 3: 检查当前 Git 状态...${NC}"
if git ls-files --error-unmatch server/. >/dev/null 2>&1 || git ls-files --error-unmatch web/. >/dev/null 2>&1; then
    echo -e "${YELLOW}检测到 server/web 在主仓库中被跟踪${NC}"
    echo ""
    
    # 保存当前更改
    echo -e "${BLUE}Step 4: 从主仓库中移除 server 和 web...${NC}"
    
    # 先提交当前的更改（如果有）
    if ! git diff-index --quiet HEAD --; then
        echo -e "${YELLOW}检测到未提交的更改，正在提交...${NC}"
        git add .
        git commit -m "Save changes before submodule conversion" || true
    fi
    
    # 从 Git 中移除但保留文件
    git rm -r --cached server web 2>/dev/null || true
    
    # 提交删除
    git commit -m "Remove server and web directories for submodule conversion" || true
    echo -e "${GREEN}✓ 已从主仓库移除${NC}"
else
    echo -e "${GREEN}✓ server 和 web 未在主仓库中跟踪${NC}"
fi
echo ""

# 删除现有的 .git 目录（但保留文件）
echo -e "${BLUE}Step 5: 清理子目录的 .git...${NC}"
if [ -d "server/.git" ]; then
    rm -rf server/.git
    echo -e "${GREEN}✓ 已删除 server/.git${NC}"
fi

if [ -d "web/.git" ]; then
    rm -rf web/.git
    echo -e "${GREEN}✓ 已删除 web/.git${NC}"
fi

# 临时移动目录
echo -e "${BLUE}Step 6: 临时移动目录...${NC}"
mv server ../server-temp
mv web ../web-temp
echo -e "${GREEN}✓ 目录已临时移动${NC}"
echo ""

# 添加 submodules
echo -e "${BLUE}Step 7: 添加 submodules...${NC}"
echo ""

echo "添加 server submodule..."
if git submodule add "$SERVER_REMOTE" server; then
    echo -e "${GREEN}✓ Server submodule 添加成功${NC}"
else
    echo -e "${RED}✗ Server submodule 添加失败${NC}"
    # 恢复目录
    mv ../server-temp server 2>/dev/null || true
    mv ../web-temp web 2>/dev/null || true
    exit 1
fi

echo ""
echo "添加 web submodule..."
if git submodule add "$WEB_REMOTE" web; then
    echo -e "${GREEN}✓ Web submodule 添加成功${NC}"
else
    echo -e "${RED}✗ Web submodule 添加失败${NC}"
    # 恢复目录
    mv ../web-temp web 2>/dev/null || true
    exit 1
fi

echo ""

# 删除临时目录
echo -e "${BLUE}Step 8: 清理临时文件...${NC}"
rm -rf ../server-temp
rm -rf ../web-temp
echo -e "${GREEN}✓ 清理完成${NC}"
echo ""

# 初始化 submodules
echo -e "${BLUE}Step 9: 初始化 submodules...${NC}"
git submodule init
git submodule update
echo -e "${GREEN}✓ Submodules 初始化完成${NC}"
echo ""

# 提交更改
echo -e "${BLUE}Step 10: 提交更改...${NC}"
git add .gitmodules server web
if git commit -m "Add server and web as submodules"; then
    echo -e "${GREEN}✓ 更改已提交${NC}"
else
    echo -e "${YELLOW}⚠ 没有需要提交的更改${NC}"
fi
echo ""

# 显示当前状态
echo -e "${BLUE}Step 11: 验证配置...${NC}"
echo ""
echo "Submodules 状态:"
git submodule status
echo ""

# 显示 .gitmodules 内容
if [ -f ".gitmodules" ]; then
    echo ".gitmodules 内容:"
    cat .gitmodules
    echo ""
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}   配置完成！${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo -e "${YELLOW}接下来的步骤:${NC}"
echo ""
echo "1. 推送到远程仓库:"
echo -e "   ${BLUE}git push${NC}"
echo ""
echo "2. 团队成员克隆项目:"
echo -e "   ${BLUE}git clone --recursive <主仓库地址>${NC}"
echo ""
echo "3. 或者已克隆的仓库初始化 submodules:"
echo -e "   ${BLUE}git submodule init${NC}"
echo -e "   ${BLUE}git submodule update${NC}"
echo ""
echo "4. 更新 submodules:"
echo -e "   ${BLUE}git submodule update --remote${NC}"
echo ""

# 询问是否立即推送
read -p "是否立即推送到远程仓库? [y/N] " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo ""
    echo -e "${BLUE}正在推送...${NC}"
    if git push; then
        echo -e "${GREEN}✓ 推送成功${NC}"
    else
        echo -e "${RED}✗ 推送失败${NC}"
        echo "您可以稍后手动推送: git push"
    fi
fi

echo ""
echo -e "${GREEN}🎉 所有操作完成！${NC}"
