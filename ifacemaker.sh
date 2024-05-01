#!/usr/bin/env bash

set -xe

go run -mod=mod github.com/vburenin/ifacemaker@latest \
  --file=/home/kurtis/go/pkg/mod/github.com/!azure/azure-sdk-for-go/sdk/ai/azopenai@v0.5.1/client.go \
  --struct=Client \
  --iface=OpenAI \
  --pkg=demo \
  --iface-comment='<iface> ...' \
  --doc=false \
  --output=/home/kurtis/wrkspc/github.com/angelokurtis/go-talk/pkg/demo/openai.go

