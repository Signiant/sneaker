package sneaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/kms"
)

type FakeKMS struct {
	GenerateInputs  []kms.GenerateDataKeyInput
	GenerateOutputs []kms.GenerateDataKeyOutput

	DecryptInputs  []kms.DecryptInput
	DecryptOutputs []kms.DecryptOutput
}

func (f *FakeKMS) GenerateDataKey(context context.Context, req *kms.GenerateDataKeyInput, optFns ...func(*kms.Options)) (*kms.GenerateDataKeyOutput, error) {
	f.GenerateInputs = append(f.GenerateInputs, *req)
	resp := f.GenerateOutputs[0]
	f.GenerateOutputs = f.GenerateOutputs[1:]
	return &resp, nil
}

func (f *FakeKMS) Decrypt(context context.Context, req *kms.DecryptInput, optFns ...func(*kms.Options)) (*kms.DecryptOutput, error) {
	f.DecryptInputs = append(f.DecryptInputs, *req)
	resp := f.DecryptOutputs[0]
	f.DecryptOutputs = f.DecryptOutputs[1:]
	return &resp, nil
}
