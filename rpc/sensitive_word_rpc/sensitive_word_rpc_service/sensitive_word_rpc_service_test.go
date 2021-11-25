package sensitive_word_rpc_service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"sensitive_words_check/config"
	"sensitive_words_check/pkg/dictionary"
	"sensitive_words_check/rpc/sensitive_word_rpc"
	"testing"
)


const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	sensitive_word_rpc.RegisterSensitiveWordServer(s, NewService())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(_ context.Context, s string) (net.Conn, error) {
	return lis.Dial()
}


func TestSensitiveWordRpcService_AddSensitiveWord(t *testing.T) {
	word := "测试数据"
	defer func() {
		dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
		_ = dty.RemoveWord(word)
	}()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(), grpc.WithBlock())
	assert.Nil(t, err, fmt.Sprintf("Failed to dial bufnet: %v", err))
	defer conn.Close()
	client := sensitive_word_rpc.NewSensitiveWordClient(conn)
	resp, err := client.AddSensitiveWord(ctx, &sensitive_word_rpc.OperateSensitiveWordRequest{Word: word})
	assert.Nil(t, err, fmt.Sprintf("AddSensitiveWord failed: %v", err))
	assert.Equal(t, true, resp.GetSuccess())
	// Test for output here.
}

func TestSensitiveWordRpcService_CheckSensitiveWord(t *testing.T) {
	word := "测试数据"
	dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
	_ = dty.AddWord(word)
	defer func() {
		_ = dty.RemoveWord(word)
	}()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(), grpc.WithBlock())
	assert.Nil(t, err, fmt.Sprintf("Failed to dial bufnet: %v", err))
	defer conn.Close()
	client := sensitive_word_rpc.NewSensitiveWordClient(conn)
	resp, err := client.CheckSensitiveWord(ctx, &sensitive_word_rpc.CheckSensitiveWordRequest{Words: []string{word}})
	assert.Nil(t, err, fmt.Sprintf("CheckSensitiveWord failed: %v", err))
	assert.Equal(t, true, resp.GetResults()[0].Sensitive)
}

func TestSensitiveWordRpcService_RemoveSensitiveWord(t *testing.T) {
	word := "测试数据"
	dty := dictionary.OpenDictionary(config.DictionaryConfig.StoreType)
	_ = dty.AddWord(word)
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure(), grpc.WithBlock())
	assert.Nil(t, err, fmt.Sprintf("Failed to dial bufnet: %v", err))
	defer conn.Close()
	client := sensitive_word_rpc.NewSensitiveWordClient(conn)
	resp, err := client.RemoveSensitiveWord(ctx, &sensitive_word_rpc.OperateSensitiveWordRequest{Word: word})
	assert.Nil(t, err, fmt.Sprintf("RemoveSensitiveWord failed: %v", err))
	assert.Equal(t, true, resp.GetSuccess())
}