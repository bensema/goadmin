export function formatDate(d) {
    var date = new Date(d);
    var YY = date.getFullYear() + '-';
    var MM = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
    var DD = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate());
    var hh = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
    var mm = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
    var ss = (date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds());
    return YY + MM + DD + " " + hh + mm + ss;
}

// yyyy-MM-dd hh:mm:ss  =>  1561953956
export function parseDateTime(date) {
    var d = new Date(date);
    return d.getTime()/1000;
}

export function DateTimeStart() {
    var t = new Date();//获取Date对象
    t.setHours(0);//设置小时
    t.setMinutes(0);//设置分钟
    t.setSeconds(0);//设置秒
    t.setMilliseconds(0);//设置毫妙
    return t
}

export function DateTimeEnd() {
    var t = new Date();//获取Date对象
    t.setHours(23);//设置小时
    t.setMinutes(59);//设置分钟
    t.setSeconds(59);//设置秒
    t.setMilliseconds(0);//设置毫妙
    return t
}
