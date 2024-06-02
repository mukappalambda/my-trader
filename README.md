# My Trader

A simple program written in Go to monitor data

```bash
::::    :::: :::   :::            ::::::::::::::::::::     :::    ::::::::: :::::::::::::::::::  
+:+:+: :+:+:+:+:   :+:                :+:    :+:    :+:  :+: :+:  :+:    :+::+:       :+:    :+: 
+:+ +:+:+ +:+ +:+ +:+                 +:+    +:+    +:+ +:+   +:+ +:+    +:++:+       +:+    +:+ 
+#+  +:+  +#+  +#++:  +#++:++#++:++   +#+    +#++:++#: +#++:++#++:+#+    +:++#++:++#  +#++:++#:  
+#+       +#+   +#+                   +#+    +#+    +#++#+     +#++#+    +#++#+       +#+    +#+ 
#+#       #+#   #+#                   #+#    #+#    #+##+#     #+##+#    #+##+#       #+#    #+# 
###       ###   ###                   ###    ###    ######     ############ #############    ### 
```

## Quick Start

```bash
git clone https://github.com/mukappalambda/my-trader.git
cd my-trader

## spin up the database container
bash scripts/run.sh

## run grpc server
export DATABASE_URL=postgresql://postgres:password@localhost:5432/demo?sslmode=disable
go run server/main.go

## run grpc client
go run client/main.go
```

You can also use [grpcurl](https://github.com/fullstorydev/grpcurl) to interact with the server.

**List the available services**

```bash
grpcurl -plaintext localhost:50051 list
```

Example output:

```console
grpc.health.v1.Health
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
message.MessageService
```

**List a service's methods**

```bash
grpcurl -plaintext localhost:50051 list grpc.health.v1.Health
```

Example output:

```console
grpc.health.v1.Health.Check
grpc.health.v1.Health.Watch
```

**Make a grpc request**

```bash
grpcurl -plaintext localhost:50051 grpc.health.v1.Health/Check
```

Example output:

```console
{
  "status": "SERVING"
}
```
