package utils

const RESPONSEOK = 0 //正常http为200的情况

const RESPONSEUNLOGIN = 10000 //用户未登陆

//**********不存在开始***************//

const RESPONSENOUSER = 11000 //没有这个用户

const RESPONSENOARTICLE = 11100 //文章不存在

//**********不存在结束***************//


const RESPONSEUSEREXSIST = 11001 //此用户已存在

const RESPONSEPARAMERROR = 12000 // 请求参数错误

const RESPONSESERVERERROR = 13000 //服务端出错

const RESPONSEUPDATEERROR = 14000 // 数据更新失败


