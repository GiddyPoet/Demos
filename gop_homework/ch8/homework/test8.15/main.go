//练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。
// 修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好写。
// 或者为每一个客户端的消息发出channel建立缓冲区，这样大部分的消息便不会被丢掉；
// broadcaster应该用一个非阻塞的send向这个channel中发消息。
package main

func broadcaster() {
	clients := make(map[client]bool,30) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
