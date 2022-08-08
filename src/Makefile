CPP=g++

redforest: redforest.cpp
	${CPP} -c interface/Interface.cpp 
	${CPP} -c interface/commands/Command.cpp
	${CPP} -c interface/commands/Sock.cpp
	${CPP} -c interface/ascii.cpp
	${CPP} -c modules/TcpClient/TcpClient.cpp
	${CPP} -c lib/Target.cpp
	${CPP} -c redforest.cpp
	${CPP} -o redforest redforest.o Interface.o Command.o Target.o ascii.o Sock.o TcpClient.o

release:
	# run normal make
	make
	# strip debug symbols
	strip redforest

clean:
	rm -f redforest *.o 

install:
	make release
	cp redforest /usr/local/bin

run:
	make
	# Running ./redforest
	./redforest

no-obj:
	make
	# Cleaning Object Files
	rm *o

