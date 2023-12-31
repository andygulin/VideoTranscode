### FFMPEG Go封装

#### Build

```shell
go build
```

#### FFMPEG版本

```shell
./VideoTranscode version
```

#### 视频/音频 信息

```shell
./VideoTranscode info 123.mp4
```

#### 视频格式转换

```shell
# mp4 -> avi
./VideoTranscode convert video 1234.mp4 1234.avi
# mp4 -> avi 无损
./VideoTranscode convert video 1234.mp4 1234_1.avi true
# mp4 -> mpeg 无损
./VideoTranscode convert video 1234.mp4 1234_1.mpeg true
# mp4 -> m3u8 无损 ts切片20秒
./VideoTranscode convert video 1234.mp4 1234_1.m3u8 20
```

```shell
# mp4 -> mp3
./VideoTranscode convert mp3 1234.mp4 1234.mp3
```

```shell
# mp4 -> mp4缩放
./VideoTranscode convert scale 1234.mp4 1234_1280.mp3 1280 720
```

```shell
# mp4 -> image
./VideoTranscode convert image 1234.mp4
```

```shell
# mp4 视频剪切
./VideoTranscode convert crop 1234.mp4 00:00:00 00:01:00
```