# RedForest

This is a project under development

## Building

`make install`

### Dependencies

The project uses the [boost](https://www.boost.org/) and [ASIO](https://www.boost.org/doc/libs/1_75_0/doc/html/boost_asio.html) libraries.
You might need to install these if you want to build the project from source.

### Building structure

```
src/
├── (Makefile)
├── interface
│   ├── ascii.cpp
│   ├── ascii.hpp
│   ├── commands
│   │   ├── Command.cpp
│   │   ├── Command.hpp
│   │   ├── Curl.cpp
│   │   ├── Curl.hpp
│   │   ├── Sock.cpp
│   │   └── Sock.hpp
│   ├── Interface.cpp
│   └── Interface.hpp
├── lib
│   ├── Target.cpp
│   └── Target.hpp
├── modules
│   ├── HTTPClient
│   │   ├── HttpClient.cpp
│   │   └── HttpClient.hpp
│   └── TcpClient
│       ├── TcpClient.cpp
│       └── TcpClient.hpp
└── redforest.cpp -> main
```

![error displaying image -> media/footer.svg](media/footer.svg?raw=true "Title")


