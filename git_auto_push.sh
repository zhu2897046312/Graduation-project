#!/bin/bash

# 检查是否提供了提交信息
if [ -z "$1" ]; then
    echo "错误：请提供提交信息。"
    echo "用法: ./git_auto_push.sh \"您的提交信息\""
    exit 1
fi

# 1. 添加所有更改
echo "正在执行 'git add .'..."
git add .

# 2. 提交更改
echo "正在提交更改，提交信息: \"$1\"..."
git commit -m "$1"

# 3. 推送到远程仓库（origin master）
echo "正在推送到 origin master..."
git push origin master

# 检查是否推送成功
if [ $? -eq 0 ]; then
    echo "成功推送到 origin master！"
else
    echo "推送失败，请检查错误信息。"
    exit 1
fi
