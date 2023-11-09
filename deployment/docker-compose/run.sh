#!/bin/bash

tool_pre_check() {
    dependent_tools=(
        "docker-compose"
    )
    for item in "${dependent_tools[@]}"; do
        if ! command -v "$item" >/dev/null 2>&1;then
            echo "$item not found !!"
            exit 1
        fi
    done
}

dir_pre_check() {
    readonly _data_dir="$HOME/data/blog"
    readonly _example_dir="./deployment/config"
    [[ ! -d "$_data_dir/data/mongo" ]] || { mkdir -p "$_data_dir/data/mongo"; }
    [[ ! -d "$_data_dir/log/mongo" ]] || { mkdir -p "$_data_dir/log/mongo"; }
    [[ ! -d "$_data_dir/config" ]] || { mkdir -p "$_data_dir/config"; }
    [[ ! -f "$_data_dir/config/mongod.conf" ]] || { cp "$_example_dir/mongod.conf" "$_data_dir/config/mongod.conf"; }
}


tool_pre_check
dir_pre_check