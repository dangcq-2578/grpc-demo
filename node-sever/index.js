const path = require("path");
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

const proto = protoLoader.loadSync(path.join(__dirname, "..", "proto/posts_service.proto"));
const definition = grpc.loadPackageDefinition(proto);

const postList = [
    { id: 1, title: 'Mình đã kiếm 1 triệu đô trên Viblo như thế nào?', content: 'This is content'},
    { id: 2, title:'Điều mà người làm dev nào cũng cần biết - SunCI', content: 'This is Sun* CI'},
];

const getPosts = (call, callback) => {
    callback(null, {posts: postList })
}

const severURL = 'localhost:9123';
const server = new grpc.Server();

server.addService(definition.PostService.service, { getPosts });
server.bindAsync(severURL, grpc.ServerCredentials.createInsecure(), port => {
    console.log(`Sever running on ${severURL}`);
    server.start();
});
