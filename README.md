## **gRPC Demo: Unary, Client Streaming, Server Streaming & Bi-Directional Streaming**  

This project is a **simple gRPC demonstration** that implements **all four types of gRPC communication** using **Go**. It covers:  
✅ **Unary RPC** – Simple request-response interaction.  
✅ **Client Streaming** – Client sends multiple requests, receives a single response.  
✅ **Server Streaming** – Client sends a request, server streams multiple responses.  
✅ **Bi-Directional Streaming** – Both client and server exchange multiple messages over a single connection.  

### **🛠 Tech Stack**  
- **Go (Golang)**  
- **gRPC**  
- **Protocol Buffers (protobuf)**  

### **📂 Project Structure**  
```  
/proto  
 ├── greet.proto              # Protocol buffer definitions  
 ├── greet.pb.go              # Generated Go code for message types  
 ├── greet_grpc.pb.go         # Generated Go code for gRPC service  
 
/server  
 ├── unary.go                 # Unary RPC implementation  
 ├── client_stream.go         # Client streaming implementation  
 ├── server_stream.go         # Server streaming implementation  
 ├── bi_stream.go             # Bi-directional streaming implementation  
 ├── main.go                  # starting point  

/client  
 ├── unary.go                 # Unary client  
 ├── client_stream.go         # Client streaming client  
 ├── server_stream.go         # Server streaming client  
 ├── bi_stream.go             # Bi-directional streaming client
 ├── main.go                  # starting point  
```  

### **🚀 How to Run the Project**  

#### **1️⃣ Install gRPC & Protobuf Compiler**  
```sh  
brew install protobuf  
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest  
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  
export PATH="$PATH:$(go env GOPATH)/bin"  
```

#### **2️⃣ Generate gRPC Code**  
```sh  
protoc --go_out=. --go-grpc_out=. proto/greet.proto  
```

#### **3️⃣ Run the Server**  
```sh  
go run server/main.go  
```

#### **4️⃣ Run the Client**  
```sh  
go run client/main.go  
```

_(i have commented the other services(unary, client and server streaming. use whichever required))_  

### **📌 Features**  
✅ Implements all gRPC communication types.  
✅ Uses **Protocol Buffers** for defining services.  
✅ Follows **best practices** for structuring a gRPC project.  
✅ Simple and beginner-friendly!  
