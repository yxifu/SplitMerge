# SplitMerge File splitting and merging tool

## Introduction

SplitMerge is a simple and easy-to-use command-line tool designed to help users effortlessly split large files into smaller parts and recombine these parts back into the original file. Whether you need to send large files via email or wish to split data for backup purposes, SplitMerge provides a quick and effective solution.

### Features and Capabilities:

- Customizable split size (supports KB/MB/GB)
- Cross-platform compatible (for example, files can be split on Windows and merged on Linux)

## Quick start

``` shell
    # Split file
    SplitMerge -s=10k hello.txt
    #-s: Size of the split files（10M 10K）
```

``` shell
    # Merge file
    SplitMerge hello.txt.s001
```

-----

# SplitMerge 文件分割与合并工具

## 项目简介

SplitMerge 是一个简单易用的命令行工具，旨在帮助用户轻松地将大文件分割成较小的部分以及将这些部分重新合并为原始文件。无论您是因为需要通过电子邮件发送大文件，还是为了备份目的而希望将数据分割存储，SplitMerge 都能提供快速有效的解决方案。

### 功能特点:

- 可定制分割大小（支持 KB/MB/GB）
- 可以跨系统（如在windows系统上分割，在linux系统上合并）

## 快速开始


``` shell
    # 分割文件
    SplitMerge -s=10k hello.txt
    #-s: 每块大小，例如：10k
```

``` shell
    # 合并文件
    SplitMerge hello.txt.s001
```