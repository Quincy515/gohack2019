package config

const (
	// 百度语音技术
	TTSApiKek    = "cxbNtyoceqNOSbSLe2YBua6u"
	TTSSecretKey = "byaeIte1RwCyzweuV2yTsYNoVtGXGmUQ"
	// 发音人选择, 基础音库：0为度小美，1为度小宇，3为度逍遥，4为度丫丫，
	// 精品音库：5为度小娇，103为度米朵，106为度博文，110为度小童，111为度小萌，默认为度小美
	PER = 4
	// 语速，取值0-15，默认为5中语速
	SPD = 5
	// 音调，取值0-15，默认为5中语调
	PIT = 5
	// 音量，取值0-9，默认为5中音量
	VOL = 5
	// 下载的文件格式, 3：mp3(default) 4： pcm-16k 5： pcm-8k 6. wav
	AUE     = 3
	CUID    = "123456PYTHON"
	TTS_URL = "http://tsn.baidu.com/text2audio"
	Host    = "0.0.0.0:8080" // 服务监听的地址
)
