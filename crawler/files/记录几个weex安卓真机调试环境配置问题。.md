# 记录几个weex安卓真机调试环境配置问题。 

问题1：UnhandledPromiseRejectionWarning: TypeError: Cannot read property 'stack' of undefined

解决方法：没有用adb连接安卓设备问题，可以在命令行输入adb devices验证，然后下载安装adb直到adb devices能展示正确的结果。

问题2：Error: Command failed: call gradlew.bat assembleDebug

解决方案：升级jdk版本到8以上。

问题3： you have not accepted the license agreements of the following SDK components: [Android SDK Build-Tools 24, Android SDK Platform 24]. Before building your project, you need to accept the license agreements and complete the installation of the missing components using the Android Studio SDK Manager。

解决方案：命令行执行：android update sdk --no-ui --filter build-tools-24.0.0,android-24,extra-android-m2repository。

附上在真机上运行成功的图片：

