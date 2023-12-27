<?php
# 写一个获取广州市天河区最近37天的天气情况的程序
# 1. 获取天气数据
function getWeather($city, $district, $days) {
    $url = "http://wthrcdn.etouch.cn/weather_mini?city=$city$district";
    $json = file_get_contents($url);
    $data = json_decode($json, true);
    return $data;
}
# 2. 解析数据
function parseWeather($data, $days) {
    $weather = array();
    for ($i = 0; $i < $days; $i++) {
        $weather[$i] = array(
            'date' => $data['data']['forecast'][$i]['date'],
            'high' => $data['data']['forecast'][$i]['high'],
            'low' => $data['data']['forecast'][$i]['low'],
            'type' => $data['data']['forecast'][$i]['type'],
        );
    }
    return $weather;
}
# 3. 保存数据
function saveWeather($weather) {
    $file = fopen('weather.csv', 'w');
    foreach ($weather as $day) {
        fputcsv($file, $day);
    }
    fclose($file);
}
# 4. 显示数据
function showWeather($weather) {
    echo '<table border="1">';
    echo '<tr><th>日期</th><th>最高温度</th><th>最低温度</th><th>天气</th></tr>';
    foreach ($weather as $day) {
        echo '<tr>';
        echo '<td>' . $day['date'] . '</td>';
        echo '<td>' . $day['high'] . '</td>';
        echo '<td>' . $day['low'] . '</td>';
        echo '<td>' . $day['type'] . '</td>';
        echo '</tr>';
    }
    echo '</table>';
}
# 5. 生成图表
function drawChart($weather) {
    $data = array();
    foreach ($weather as $day) {
        $data[] = $day['high'];
    }
    $data = implode(',', $data);
    echo '<img src="http://chart.apis.google.com/chart?cht=lc&chs=500x300&chd=t:' . $data . '">';
}
# 6. 主程序
$city = '广州';
$district = '天河区';
$days = 37;
$data = getWeather($city, $district, $days);
$weather = parseWeather($data, $days);
saveWeather($weather);
showWeather($weather);
drawChart($weather);
?>