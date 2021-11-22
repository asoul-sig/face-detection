# face-detection
Face detection for asoul.video image cover.

## 简介
asoul.video 上的视频封面，展示时会将原本抖音上的封面图片居中裁剪为正方形。因此常常出现封面中的人物被截掉一半的情况。

因此，我们基于 https://github.com/nagadomi/lbpcascade_animeface 该项目实现的动漫人物面部捕捉，对每张封面图中人物面部进行定位并存储坐标。前端展示图片时根据人脸坐标对图片做相应的偏移。从而保证女孩们的脸始终居中。

