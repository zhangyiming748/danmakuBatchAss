#!/bin/bash

# 指定文件夹路径和文件扩展名
folder_path="/sdcard/Movies/bili"
file_extension=".xml"

# 使用find命令查找文件
# shellcheck disable=SC2162
find "$folder_path" -type f -name "*$file_extension" | while read file; do
    # 去除扩展名的文件名
    base_name=$(basename "$file" "$file_extension")
    # 逐行打印文件名
    echo "$base_name"
    ass_name="$base_name".ass
    cmd=/Users/zen/Github/danmakuBatchAss/danmaku2ass.py "$file" -s 1280x720 -dm 15 -o "$ass_name"
    echo "$cmd"
done