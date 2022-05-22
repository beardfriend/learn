# 1. 비디오 프레임 생성

```bash
sudo apt update
sudo apt install ffmpeg

ffmpeg -i Video.mp4 -vf fps=1/10,scale=120:-1 assets/previewImgs/preview%d.jpg
```

비디오 프레임을 생성해준다.

## 2. 아이콘

구글 폰트 머테리얼 디자인 아이콘 방문해서 다운로드 할 수 있다.
