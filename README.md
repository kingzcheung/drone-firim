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

上传类型：ios 或者 android，必填

`bundle_id`(string):

App 的 bundleId，必填

`api_token`(string):

长度为 32, 用户在 fir 的 api_token

`file`(string):

要上传的文件（安装包）地址：一般是当前目录下开始，必填

`name`(string):

应用名称，必填

`version`(string):

版本号，必填

`build`(string):

Build号，必填

`release_type`(string):

打包类型，只针对 iOS,可选

`changelog`(string):

更新日志，可选