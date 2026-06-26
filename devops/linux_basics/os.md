Operating system 

every application in our system needs to talk to the hardware in some sense
like using memory or cpu or i/o devices so whats in between that helps these application to use these hardwares present in the system 
so this translor that works in between our application and hardwares , telling to allocate memory or to use any resource or like isolating resource contents , so all these because of operating system 

os tasks 

process management - process -> small unit that executes on a computer 
1 cpu = 1 process at a time 
(when we do multiple processes the cpu switches so effectively that we dont notice it much in terms of efficiency , however more cpu = more processes you can run simultaneously )

memory management - is just allocating memory for differnt processes 
in case there are a lot of processes running using the major memory available then in that case memory swapping takes place which is slow thus it feels like computer is working slow

storage management - the hardrive the secondary memory where the data persist
in a structured way
in linux - tree structure (single root)
in windows - multiple roots (a: c:)

io devices interaction

security and networking - manage user permissions


Operating system components 

Kernel - heart of operating system , loads first , manage the hardware components (cpu , memory, ram ) , kernel also handles various io devices using device drivers
so ultimately its the kernel which is the layer that get interact between applications and hardware ( it does the resource allocation , starts the processes , cleans up resouces)

Application layer - based on the same kernel , these layers are made , they are just differnt ui , graphical interface, icons , but are based on same kernel
e.g linux kernel based applicaiton distributions - ubuntu , mint , android
