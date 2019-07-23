# drone-firim - Drone fir.im Plugin

`drone-firim` 是一个上传文件（安装包）到 fir.im 服务的插件

`.drone.yml` (1.*版本)配置示例

```yaml
- name: publish firim
  image: bbking/drone-firim
  settings:
      type: android
      bundle_id: com.google.android.....
      api_token: your api token
      file: drone output file
      name: app name
      version: app version
      build: build number

```

### Plugin Parameter Reference

`type`（string）:

上传类型：ios 或者 android

`bundle_id`(string):

App 的 bundleId

`api_token`(string):

长度为 32, 用户在 fir 的 api_token

`file`(string):

要上传的文件（安装包）地址：一般是当前目录下开始

`name`(string):

应用名称

`version`(string):

版本号

`build`(string):

Build号