package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

func Init() {
	resolver.Register(&ConsulRegister{})
}

type ConsulCli struct {
	client *api.Client
}

func NewConsulCli() *ConsulCli {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ConsulCli{client: client}
}

func (cli *ConsulCli) RegisterService(name string, ip string, port int) {
	agent := cli.client.Agent()
	service := &api.AgentServiceRegistration{
		//Tags:    []string{name},                          // tag，可以为空
		ID:      fmt.Sprintf("%v-%v-%v", name, ip, port), // 服务节点的名称
		Name:    name,                                    // 服务名称
		Port:    port,                                    // 服务端口
		Address: ip,                                      // 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       (10 * time.Second).String(),               // 健康检查间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", ip, port, name),   // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: (time.Duration(1) * time.Minute).String(), // 注销时间，相当于过期时间
		},
	}

	if rErr := agent.ServiceRegister(service); rErr != nil {
		fmt.Printf("Service Register error\n%v", rErr)
		return
	}
}

func (cli *ConsulCli) DiscoveryAddrLst(service string, tag string, lastIndex uint64) (map[string][]string, uint64, error) {
	se, meta, err := cli.client.Health().Service(service, tag, true, &api.QueryOptions{WaitIndex: lastIndex})
	if err != nil {
		return nil, 0, err
	}
	svcAddrMap := make(map[string][]string)
	for _, e := range se {
		addr := fmt.Sprintf("%v:%v", e.Service.Address, e.Service.Port)
		svcAddrMap[e.Service.Service] = append(svcAddrMap[e.Service.Service], addr)
	}
	return svcAddrMap, meta.LastIndex, nil
}

type ConsulRegister struct {
}

func (l ConsulRegister) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	lr := ConsulResolver{
		name: target.Endpoint,
		cc:   cc,
	}
	go lr.watcher()
	return &lr, nil
}

func (l ConsulRegister) Scheme() string {
	return "consul"
}

type ConsulResolver struct {
	name      string
	tag       string
	cc        resolver.ClientConn
	lastIndex uint64
}

func (cr *ConsulResolver) ResolveNow(options resolver.ResolveNowOptions) {
	fmt.Println("ConsulResolver ResolveNow")
}

func (cr *ConsulResolver) Close() {
}

func (cr *ConsulResolver) watcher() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}
		log.Fatal(err)
	}()
	consulCli := NewConsulCli()
	for {
		svcAddrMap, lastIndex, checkErr := consulCli.DiscoveryAddrLst(cr.name, cr.tag, cr.lastIndex)
		if checkErr != nil {
			log.Fatal(checkErr)
			continue
		}
		cr.lastIndex = lastIndex

		updateAddresses := make([]resolver.Address, 0)
		for svc, addrLst := range svcAddrMap {
			for _, addr := range addrLst {
				updateAddresses = append(updateAddresses, resolver.Address{
					Addr:       addr,
					ServerName: svc,
					//Attributes: nil,
				})
			}
		}

		state := resolver.State{Addresses: updateAddresses}
		_ = cr.cc.UpdateState(state)
		log.Println("ConsulResolver UpdateState")
	}

}
