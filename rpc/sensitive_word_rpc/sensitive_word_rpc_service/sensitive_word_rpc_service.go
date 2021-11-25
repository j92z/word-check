package sensitive_word_rpc_service

import (
	"context"
	"sensitive_words_check/rpc/sensitive_word_rpc"
	"sensitive_words_check/service/sensitive_word_service"
)

type SensitiveWordRpcService struct {
	sensitive_word_rpc.UnimplementedSensitiveWordServer
}

func NewService() *SensitiveWordRpcService {
	return &SensitiveWordRpcService{}
}

func (s *SensitiveWordRpcService) AddSensitiveWord(_ context.Context, in *sensitive_word_rpc.OperateSensitiveWordRequest) (*sensitive_word_rpc.OperateSensitiveWordResponse, error) {
	word := in.Word
	response := new(sensitive_word_rpc.OperateSensitiveWordResponse)
	if err := sensitive_word_service.AddSensitiveWord(word); err != nil {
		response.Success = false
	} else {
		response.Success = true
	}
	return response, nil
}

func (s *SensitiveWordRpcService) RemoveSensitiveWord(_ context.Context, in *sensitive_word_rpc.OperateSensitiveWordRequest) (*sensitive_word_rpc.OperateSensitiveWordResponse, error) {
	word := in.Word
	response := new(sensitive_word_rpc.OperateSensitiveWordResponse)
	if err := sensitive_word_service.RemoveSensitiveWord(word); err != nil {
		response.Success = false
	} else {
		response.Success = true
	}
	return response, nil
}

func (s *SensitiveWordRpcService) CheckSensitiveWord(_ context.Context, in *sensitive_word_rpc.CheckSensitiveWordRequest) (*sensitive_word_rpc.CheckSensitiveWordResponse, error) {
	words := in.Words
	response := new(sensitive_word_rpc.CheckSensitiveWordResponse)
	result := sensitive_word_service.CheckSensitiveWord(words)
	for _, v := range result {
		response.Results = append(response.Results, &sensitive_word_rpc.CheckSensitiveWordResponse_Result{
			Text:      v.Text,
			Words:     v.Words,
			HitWords:  v.HitWords,
			Sensitive: v.Sensitive,
		})
	}
	return response, nil
}