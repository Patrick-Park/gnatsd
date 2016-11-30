// Copyright 2012-2016 Apcera Inc. All rights reserved.

package test

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func createNonFatalClientConn(host string, port int) (net.Conn, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	return net.DialTimeout("tcp", addr, 1*time.Second)
}

func TestMaxUsers(t *testing.T) {

	srv, opts := RunServerWithConfig("./configs/max_users.conf")
	defer srv.Shutdown()

	c1, err := createNonFatalClientConn(opts.Host, opts.Port)
	if err != nil {
		t.Fatalf("Could not connect to server: %v\n", err)
	}
	defer c1.Close()
	c2, err := createNonFatalClientConn(opts.Host, opts.Port)
	if err == nil {
		t.Fatal("max_connections was not honored.\n")
	}
	defer c2.Close()
}
