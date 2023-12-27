<?php

// 写一个php的web框架
// 1. 定义一个函数，用来处理请求
function handleRequest($request) {
    $response = array(
        'status' => 200,
        'headers' => array(
            'Content-Type' => 'text/html',
        ),
        'body' => 'Hello World',
    );
    return $response;
}

// 2. 定义一个函数，用来发送响应
function sendResponse($response) {
    header('HTTP/1.1 ' . $response['status']);
    foreach ($response['headers'] as $key => $value) {
        header($key . ': ' . $value);
    }
    echo $response['body'];
}

// 3. 定义一个函数，用来处理异常
function handleException($e) {
    $response = array(
        'status' => 500,
        'headers' => array(
            'Content-Type' => 'text/html',
        ),
        'body' => 'Internal Server Error',
    );
    sendResponse($response);
}

// 4. 定义一个函数，用来处理错误
function handleError($errno, $errstr, $errfile, $errline) {
    $response = array(
        'status' => 500,
        'headers' => array(
            'Content-Type' => 'text/html',
        ),
        'body' => 'Internal Server Error',
    );
    sendResponse($response);
}

// 5. 定义一个函数，用来处理终止
function handleShutdown() {
    $error = error_get_last();
    if ($error) {
        $response = array(
            'status' => 500,
            'headers' => array(
                'Content-Type' => 'text/html',
            ),
            'body' => 'Internal Server Error',
        );
        sendResponse($response);
    }
}

// 合并数组函数
function array_merge_recursive_distinct ( array &$array1, array &$array2 )
{
    $merged = $array1;

    foreach ( $array2 as $key => &$value )
    {
        if ( is_array ( $value ) && isset ( $merged [$key] ) && is_array ( $merged [$key] ) )
        {
            $merged [$key] = array_merge_recursive_distinct ( $merged [$key], $value );
        }
        else
        {
            $merged [$key] = $value;
        }
    }

    return $merged;
}

// 解析xml代码
function parseXml($xml) {
    $data = array();
    $xml = simplexml_load_string($xml);
    $data['status'] = (string)$xml->status;
    $data['message'] = (string)$xml->message;
    $data['request'] = (string)$xml->request;
    $data['city'] = (string)$xml->city;
    $data['forecast'] = array();
    foreach ($xml->forecast->children() as $child) {
        $data['forecast'][] = array(
            'date' => (string)$child->date,
            'high' => (string)$child->high,
            'low' => (string)$child->low,
            'type' => (string)$child->type,
        );
    }
    return $data;
}

// 写一个parseXml测试用例
function testParseXml() {
    $xml = <<<EOF
<?xml version="1.0" encoding="UTF-8"?>
<resp>
    <status>0</status>
    <message>ok</message>
    <request>/v1/weather/query</request>
    <city>北京</city>
    <forecast>
        <yesterday date="3日星期二" high="高温 15℃" low="低温 8℃" type="晴"/>
        <today date="4日星期三" high="高温 16℃" low="低温 8℃" type="晴"/>
        <tomorrow date="5日星期四" high="高温 16℃" low="低温 8℃" type="晴"/>
    </forecast>
</resp>
EOF;
    $data = parseXml($xml);
    var_dump($data);
}






// 6. 定义一个函数，用来启动框架
function run() {
    set_exception_handler('handleException');
    set_error_handler('handleError');
    register_shutdown_function('handleShutdown');
    $request = $_SERVER['REQUEST_URI'];
    $response = handleRequest($request);
    sendResponse($response);
}

// 7. 主程序
// run();


testParseXml();

               