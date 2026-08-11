### FFMPEG Go封装

### Install FFMPEG

```shell
# MacOS
brew install ffmpeg
# Centos
yum install ffmpeg
# Ubuntu
apt install ffmpeg
# Windows
https://ffmpeg.org/download.html
下载 & 配置环境变量
```

#### Build

```shell
# 编译
./build.sh
# Local
./VideoTranscode
# MacOS
./VideoTranscode_darwin_arm64
# Windows
./VideoTranscode_windows_amd64.exe
# Linux
./VideoTranscode_linux_amd64
```

#### FFMPEG版本

```shell
./VideoTranscode version
```

#### 视频/音频 信息

```shell
./VideoTranscode info 123.mp4 345.mp4 ...
```

#### 视频格式转换

```shell
# mp4 -> avi
./VideoTranscode transcode 1234.mp4 1234.avi
# mp4 -> avi 无损
./VideoTranscode transcode --lossless 1234.mp4 1234_1.avi
# mp4 -> mpeg 无损
./VideoTranscode transcode --lossless 1234.mp4 1234_1.mpeg
# mp4 -> m3u8 无损 ts切片20秒
./VideoTranscode transcode --lossless --hls-time 20 1234.mp4 1234_1.m3u8
```

```shell
# video -> audio
./VideoTranscode extract-audio 1234.mp4 1234.mp3
```

```shell
# mp4 -> mp4缩放
./VideoTranscode transcode --scale 1280:720 1234.mp4 1234_1280.mp4
```

```shell
# mp4 -> image
./VideoTranscode snapshot 1234.mp4
```

```shell
# mp4 视频剪切
./VideoTranscode --crop start=00:00:00,duration=00:01:00 1234.mp4 1234_crop.mp4
```

```shell
# 生成ts文件列表input.txt
# input.txt
file '001.ts'
file '002.ts'
file '003.ts'
...
./VideoTranscode ts list /dir/ts input.txt
# ts -> mp4合并
./VideoTranscode ts merge input.txt 1234.mp4
```
