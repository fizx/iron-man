namespace go reddit.iron_man

include "baseplate.thrift"

struct EchoRequest { 
  1: string content;
} 

struct EchoResponse { 
  1: string content;
}

service IronManService extends baseplate.BaseplateService {
    EchoResponse echo(
      1: EchoRequest request,
    ) throws ();
}