<?php

namespace Applications\Common\Helper;

class Curl {

    private static $options = array(
        CURLOPT_CONNECTTIMEOUT => 15,
        CURLOPT_TIMEOUT => 20, //超时
        CURLOPT_RETURNTRANSFER => 1, //获取内容不输出
        CURLOPT_HEADER => false, //设定不包含头
        CURLOPT_USERAGENT => "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.101 Safari/537.36",
        CURLOPT_ENCODING => "gzip,deflate",
        CURLOPT_FOLLOWLOCATION => 1,
        CURLOPT_MAXREDIRS => 3,
    );

    public static function getCookieFile($CookieFileContent) {
        self::$options[CURLOPT_COOKIE] = $CookieFileContent;
    }

    public static function run($url, $type = 'get', $fields = array()) {
        $ch = curl_init();
        if (!empty($fields) && $type != 'post')
            $url = $url . '?' . http_build_query($fields);

        self::$options[CURLOPT_URL] = $url;
        curl_setopt_array($ch, self::$options);

        if ($type == 'post') {
            curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($fields));
            curl_setopt($ch, CURLOPT_POST, 1);
        }

        $return = curl_exec($ch);

        curl_close($ch);
        return $return;
    }

    // 获取网页内容函数
    public static function https_request($url, $post = TRUE, $data = []) {
        $ch = curl_init($url);

        curl_setopt($ch, CURLOPT_POST, $post); // 是否POST方式
        curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($data)); // 数据
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE); // 为TRUE时 return内容，为FALSE时 直接在页面echo 且返回bool结果
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, FALSE);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, FALSE);

        $result = curl_exec($ch);

        curl_close($ch);
        return $result;
    }

}
