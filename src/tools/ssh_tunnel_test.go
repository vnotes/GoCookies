package tools

import (
	_ "github.com/go-sql-driver/mysql"
)

//func TestSSHTunnel(t *testing.T) {
//	localEndpoint := &Endpoint{"", 9999}
//	serverEndpoint := &Endpoint{"", 22}
//	remoteEndpoint := &Endpoint{"", 3306}
//	sshConfig := &ssh.ClientConfig{
//		User: "xx",
//		Auth: []ssh.AuthMethod{
//			PrivateKeyFile("*.pem"),
//		},
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//	}
//	ch := make(chan bool)
//	sshTunnel := SSHTunnel{localEndpoint, serverEndpoint, remoteEndpoint, sshConfig}
//	go sshTunnel.Start(&ch)
//	<-ch
//	sourceName := fmt.Sprintf("root:pw@tcp(127.0.0.1:%d)/db", localEndpoint.Port)
//	t.Log(sourceName)
//	db, err := sql.Open("mysql", sourceName)
//	assert.Nil(t, err)
//
//	rows, err := db.Query(`select count(*) from xxx`)
//	assert.Nil(t, err)
//	t.Log(rows)
//}
