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
