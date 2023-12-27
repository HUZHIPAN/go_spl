<?php

// 写一个web框架
// 1. 定义一个函数，用来处理请求
function handleRequest($request) {
    $response = array(
        'status' => 200,
        'headers' => array(
            'Content-Type' => '