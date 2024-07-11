package register

import (
	"google.golang.org/grpc/resolver"
	"log"
)

type CustomNameServBuilder struct {
}

func (c CustomNameServBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	cns := CustomNameResolver{
		name: target.Endpoint,
		cc:   cc,
	}
	go cns.watcher()
	return &cns, nil
}

type CustomNameResolver struct {
	name string
	cc   resolver.ClientConn
}

func (r *CustomNameResolver) ResolveNow(options resolver.ResolveNowOptions) {
	log.Println("CustomNameResolver ResolveNow")
}

func (r *CustomNameResolver) Close() {
}

func (r *CustomNameResolver) watcher() {
	// todo:
	state := resolver.State{
		Addresses: []resolver.Address{
			{
				Addr:       "",
				ServerName: "",
				Attributes: nil,
				Type:       0,
				Metadata:   nil,
			},
		},
	}
	r.cc.UpdateState(state)
}

func (c CustomNameServBuilder) Scheme() string {
	panic("cz")
}
