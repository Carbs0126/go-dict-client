#!/bin/bash

# 获取当前路径
current_path=$(pwd)
# go 可执行文件的相对路径文件名
linux_dict_client_bin_relative_path="go-dict-client-linux"
# go 可执行文件的绝对路径文件名
linux_dict_client_bin_absolute_path="$current_path/$linux_dict_client_bin_relative_path"

# 检查文件是否存在
if ! [ -e "$linux_dict_client_bin_absolute_path" ]; then
    echo "文件不存在"
    exit 1
fi

# 路径存在，上传到服务器
scp "$linux_dict_client_bin_absolute_path" "root@121.40.243.60:/root"
