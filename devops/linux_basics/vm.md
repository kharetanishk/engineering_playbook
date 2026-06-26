What is Virtualization?

Virtualization allows you to run multiple operating systems on the same physical machine by creating virtual computers called Virtual Machines (VMs).

Example:

Your Laptop (Windows)
│
├── Ubuntu VM
├── CentOS VM
└── Kali Linux VM

Even though you only have one physical laptop, it behaves like multiple computers.

What is a Hypervisor?

A Hypervisor is the software layer that creates and manages those Virtual Machines.

Think of it as a manager:

Physical Hardware
│
└── Hypervisor
    │
    ├── VM 1 (Ubuntu)
    ├── VM 2 (CentOS)
    └── VM 3 (Windows)

Without a hypervisor, virtualization wouldn't be possible.

Examples:

VirtualBox
VMware
Hyper-V
KVM

Hardware resources can be shared through hypervisor 

type 2 hypervisor can be used in personal computer to acces any vm 
type 1 hypervisor is implemented over a hardware (bare metal) e.g vmware,esxi

and what these hypervisor is for - they let us create os (guest os ) in our host (type 2) or from a hardware(bare meta type 1)

advantage of using a type 1 hypervisor and access those virtual machines is that user can choose custom resource combination(ram,memory,cpu) and yes the user have access to resources of big performant servers

with virtualization you have all the os stuff and applications all in one virtual machine image that we can create copies of (backups , portable , not dependent on physical server )