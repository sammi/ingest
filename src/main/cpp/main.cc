#include <iostream>
#include <memory>
#include <string>
#include <cstdlib>

#include <grpcpp/grpcpp.h>

#include "contract.grpc.pb.h"

using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using contract::HelloRequest;
using contract::HelloReply;
using contract::Greeter;

class GreeterServiceImpl final : public Greeter::Service {
    Status SayHello(ServerContext* context, const HelloRequest* request, HelloReply* reply) override {
        std::string prefix("Hi ");
        reply->set_message(prefix + request->name());
        return Status::OK;
    }
};

void RunServer(std::string port) {

    std::string server_address("0.0.0.0:" + port);

    GreeterServiceImpl service;

    ServerBuilder builder;
    builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
    builder.RegisterService(&service);
    std::unique_ptr<Server> server(builder.BuildAndStart());
    std::cout << "Server listening on " << server_address << std::endl;

    server->Wait();
}

int main(int argc, char** argv) {
    if (argc > 1) {
        std::string port(argv[1]);
        RunServer(port);
        return 0;
    } else {
        std::cout << "usage: ingest <port_number> \n, e.g:  ingest 50051\n";
        return 1;
    }

}