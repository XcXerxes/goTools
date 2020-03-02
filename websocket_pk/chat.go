package main

import (
	"bufio"
	"container/list"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

var (
	connid int
	conns *list.List
)

func ChatroomServer(ws *websocket.Conn)  {
	// 异常退出之前关闭
	defer ws.Close()
	// 进入一个 id + 1
	connid++
	id := connid
	// 提示连接成功
	fmt.Println("connection id:\n", id)
	// 将当前的 websocket 实例塞入到一个list 中
	item := conns.PushBack(ws)
	// 如果异常退出，移除掉当前的 ws 实例
	defer conns.Remove(item)
	// user + id
	name := fmt.Sprintf("user%d", id)
	// 发送消息， welcome user1 join
	SendMessage(nil, fmt.Sprintf("welcome %s join\n", name))
	// 创建一个缓存的读取对象
	r := bufio.NewReader(ws)
	// 死循环
	for{
		// 读取缓存数据
		data, err := r.ReadBytes('\n')
		if err != nil {
			// 读取失败，表示连接退出了， 同时发送消息，退出了
			fmt.Printf("disconnected id: %d\n", id)
			SendMessage(item, fmt.Sprintf("%s offline\n", name))
			break
		}
		// 正常连接就打赢 输入的数据和 名称
		fmt.Printf("%s:%s", name, data)
		SendMessage(item, fmt.Sprintf("%s\t> %s", name, data))
	}
}

func SendMessage(self *list.Element, data string)  {
	// 如果 当前的 ws 实例存在
	for item := conns.Front(); item != nil; item = item.Next() {
		// 获取到实例
		ws, ok := item.Value.(*websocket.Conn)
		if !ok {
			panic("item not *websocket.Conn")
		}
		if item == self {
			continue
		}
		// 写入实例和数据
		io.WriteString(ws, data)
	}
}

// 网页客户端
func Client(w http.ResponseWriter, r *http.Request) {
	html := `<!doctype html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>golang websocket chatroom</title>
    <script src="http://ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
    <script>
        var ws = new WebSocket("ws://127.0.0.1:6611/chatroom");
        ws.onopen = function(e){
            console.log("onopen");
            console.dir(e);
        };
        ws.onmessage = function(e){
            console.log("onmessage");
            console.dir(e);
            $('#log').append('<p>'+e.data+'<p>');
            $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
        };
        ws.onclose = function(e){
            console.log("onclose");
            console.dir(e);
        };
        ws.onerror = function(e){
            console.log("onerror");
            console.dir(e);
        };
        $(function(){
            $('#msgform').submit(function(){
                ws.send($('#msg').val()+"\n");
                $('#log').append('<p style="color:red;">My > '+$('#msg').val()+'<p>');
                $('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
                $('#msg').val('');
                return false;
            });
        });
    </script>
</head>
<body>
    <div id="log" style="height: 300px;overflow-y: scroll;border: 1px solid #CCC;">
    </div>
    <div>
        <form id="msgform">
            <input type="text" id="msg" size="60" />
        </form>
    </div>
</body>
</html>`
	io.WriteString(w, html)
}

func main()  {
	fmt.Printf(`Welcome chatroom server!
author: dotcoo zhao
url: http://www.dotcoo.com/golang-websocket-chatroom
`)
connid = 0
conns = list.New()
http.Handle("/chatroom", websocket.Handler(ChatroomServer))
http.HandleFunc("/", Client)
err := http.ListenAndServe(":6611", nil)
if err != nil {
	panic("ListenAndServe:" + err.Error())
}
}

